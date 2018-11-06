package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	logoptions.LogOptions

	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:80" description:"Adress for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`
}

func main() {

	// Parse options
	args, err := flags.Parse(&options)
	if err != nil {
		if et, ok := err.(*flags.Error); ok {
			if et.Type == flags.ErrHelp {
				return
			}
		}
		log.Fatalf("error parsing flags: %v", err)
	}
	if len(args) > 0 {
		log.Fatalf("unexpected arguments: %v", args)
	}

	// Setup new zap logger
	config := options.LogOptions.ZapConfig()
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	process := process.NewProcess(logger)
	logDB, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}
	logDB.MapperFunc(xstrings.ToSnakeCase)
	process.CloseGracefully(logDB.Close)

	dbversion.WaitUntilLatestTxlogDBVersion(logger, logDB.DB)

	r := chi.NewRouter()
	r.HandleFunc("/in", newLogFetcher(logger, logDB, transactionlog.DirectionIn))
	r.HandleFunc("/out", newLogFetcher(logger, logDB, transactionlog.DirectionOut))

	server := &http.Server{
		Addr:    options.ListenAddress,
		Handler: r,
	}

	process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		return server.Shutdown(localCtx)
	})

	err = server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			logger.Error("failed to ListenAndServe", zap.Error(err))
		}
		return
	}

}

type Record struct {
	*transactionlog.Record
	Created  time.Time      `json:"created"`
	DataJSON types.JSONText `json:"-"`
}

type Out struct {
	Records []*Record `json:"records"`
}

func newLogFetcher(logger *zap.Logger, logDB *sqlx.DB, direction transactionlog.Direction) http.HandlerFunc {
	stmtFetchLogs, err := logDB.Preparex(`
		SELECT
			created,
			src_organization,
			dest_organization,
			service_name,
			logrecord_id,
			data AS data_json
		FROM transactionlog.records
		WHERE direction = '` + string(direction) + `'::transactionlog.direction
		ORDER BY created
	`)
	if err != nil {
		logger.Fatal("failed to prepare query for fetching logs "+string(direction), zap.Error(err))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res, err := stmtFetchLogs.Queryx()
		if err != nil {
			logger.Error("failed to fetch transaction logs", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		var out = Out{
			Records: make([]*Record, 0),
		}
		for res.Next() {
			rec := &Record{}
			err = res.StructScan(rec)
			if err != nil {
				logger.Error("failed to scan transaction log into struct", zap.Error(err))
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			err = rec.DataJSON.Unmarshal(&rec.Record.Data)
			if err != nil {
				logger.Error("failed to unmarshal record data fields", zap.Error(err))
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			out.Records = append(out.Records, rec)
		}

		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			logger.Error("failed to output records", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return

		}
	}
}
