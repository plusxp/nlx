package directoryservice

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

	"go.nlx.io/nlx/directory/directoryapi"
)

// compile-time interface implementation verification
var _ directoryapi.DirectoryServer = &DirectoryService{}

// DirectoryService handles all requests for a directory api
type DirectoryService struct {
	*registerInwayHandler
	*listServicesHandler
	*listOrganizationsHandler
	*getServiceAPISpecHandler
}

// New sets up a new DirectoryService and returns an error when something failed during set.
func New(logger *zap.Logger, db *sqlx.DB, rootCA *x509.CertPool, certKeyPair tls.Certificate, demoEnv string, demoDomain string) (*DirectoryService, error) {
	s := &DirectoryService{}

	var err error

	s.registerInwayHandler, err = newRegisterInwayHandler(db, logger, rootCA, certKeyPair)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup RegisterInway handler")
	}
	s.listServicesHandler, err = newListServicesHandler(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ListServices handler")
	}
	s.listOrganizationsHandler, err = newListOrganizationsHandler(db, logger, demoEnv, demoDomain)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ListOrganizations handler")
	}
	s.getServiceAPISpecHandler, err = newGetServiceAPISpecHandler(db, logger, rootCA, certKeyPair)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup GetServiceAPISpecHandler handler")
	}

	return s, nil
}

func getOrganisationNameFromRequest(ctx context.Context) (string, error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("failed to obtain peer from context")
	}
	tlsInfo := peer.AuthInfo.(credentials.TLSInfo)
	return tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0], nil
}

func newHTTPClient(rootCA *x509.CertPool, certKeyPair tls.Certificate) *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			RootCAs:      rootCA,
			Certificates: []tls.Certificate{certKeyPair},
		},
	}
	return &http.Client{
		Transport: transport,
	}
}
