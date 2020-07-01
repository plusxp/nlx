// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"encoding/json"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

const PREFIX = "nlx"

// ETCDConfigDatabase is the etcd implementation of ConfigDatabase
type ETCDConfigDatabase struct {
	pathPrefix string
	etcdCli    *clientv3.Client
	logger     *zap.Logger
	clock      clock.Clock
}

// NewEtcdConfigDatabase constructs a new EtcdConfigDatabase
func NewEtcdConfigDatabase(logger *zap.Logger, p *process.Process, connectionStrings []string, c clock.Clock) (ConfigDatabase, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   connectionStrings,
		DialTimeout: time.Second,
	})
	if err != nil {
		return nil, err
	}

	pathPrefix := PREFIX
	if !strings.HasPrefix(PREFIX, "/") {
		pathPrefix = "/" + pathPrefix
	}

	p.CloseGracefully(cli.Close)

	return &ETCDConfigDatabase{
		pathPrefix: pathPrefix,
		etcdCli:    cli,
		logger:     logger,
		clock:      c,
	}, nil
}

func (db ETCDConfigDatabase) put(ctx context.Context, key string, value interface{}) error {
	key = path.Join(db.pathPrefix, key)

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = db.etcdCli.Put(ctx, key, string(data))
	if err != nil {
		return err
	}

	return nil
}

// list retrieves all the values under the key prefix and stores the result in the value pointed to by dest.
// The value pointed by dest must be a slice.
func (db ETCDConfigDatabase) list(ctx context.Context, key string, dest interface{}) error {
	destValue := reflect.ValueOf(dest).Elem()
	destValueType := destValue.Type()
	destElementType := destValueType.Elem().Elem()

	key = path.Join(db.pathPrefix, key)

	response, err := db.etcdCli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	if response.Count == 0 {
		return nil
	}

	sliceLen := int(response.Count)
	destValue.Set(reflect.MakeSlice(destValueType, sliceLen, sliceLen))

	for i, kv := range response.Kvs {
		value := reflect.New(destElementType)
		valueElement := value.Elem()

		err := json.Unmarshal(kv.Value, valueElement.Addr().Interface())
		if err != nil {
			return err
		}

		destValue.Index(i).Set(value)
	}

	return nil
}