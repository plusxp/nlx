// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway/config"
)

func newTestEnv(t *testing.T, tlsOptions orgtls.TLSOptions) (proxy, mock *httptest.Server) {

	// Mock endpoint (service)
	mockEndPoint := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	// defer is missing do this in you test!!
	// defer mockEndPoint.Close()

	serviceConfig := &config.ServiceConfig{}
	serviceConfig.Services = make(map[string]config.ServiceDetails)
	serviceConfig.Services["mock-service-whitelist"] = config.ServiceDetails{
		EndpointURL:            mockEndPoint.URL,
		AuthorizationWhitelist: []string{"nlx-test"},
		AuthorizationModel:     "whitelist",
	}
	serviceConfig.Services["mock-service-whitelist-unauthorized"] = config.ServiceDetails{
		EndpointURL:            mockEndPoint.URL,
		AuthorizationWhitelist: []string{"nlx-forbidden"},
		AuthorizationModel:     "whitelist",
	}
	serviceConfig.Services["mock-service-public"] = config.ServiceDetails{
		EndpointURL:        mockEndPoint.URL,
		AuthorizationModel: "none",
	}

	logger := zap.NewNop()
	p := process.NewProcess(logger)

	iw, err := NewInway(logger, nil, "localhost:1812", tlsOptions,
		"localhost:1815", serviceConfig)
	assert.Nil(t, err)

	// Add service endpoints
	for serviceName, serviceDetails := range serviceConfig.Services { //nolint
		endpoint, errr := iw.NewHTTPServiceEndpoint(logger, serviceName, serviceDetails.EndpointURL, nil)
		if errr != nil {
			t.Fatal("failed to create service endpoint", err)
		}

		switch serviceDetails.AuthorizationModel {
		case "none", "":
			endpoint.SetAuthorizationPublic()
		case "whitelist":
			endpoint.SetAuthorizationWhitelist(serviceDetails.AuthorizationWhitelist)
		default:
			logger.Fatal(fmt.Sprintf(`invalid authorization model "%s" for service "%s"`, serviceDetails.AuthorizationModel, serviceName))
		}

		err = iw.AddServiceEndpoint(p, endpoint, serviceDetails)
		assert.Nil(t, err)
	}

	assert.Nil(t, err)

	proxyRequestMockServer := httptest.NewUnstartedServer(http.HandlerFunc(iw.handleProxyRequest))
	proxyRequestMockServer.TLS = &tls.Config{
		ClientCAs:  iw.roots,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	return proxyRequestMockServer, mockEndPoint

}

func TestInwayProxyRequest(t *testing.T) {

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "root.crt"),
		OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
		OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
	}

	proxyRequestMockServer, mockEndPoint := newTestEnv(t, tlsOptions)
	proxyRequestMockServer.StartTLS()
	defer proxyRequestMockServer.Close()
	defer mockEndPoint.Close()

	client := setupClient(t, tlsOptions)

	//nolint
	tests := []struct {
		url          string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusOK, ""},
		{fmt.Sprintf("%s/mock-service-whitelist/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusOK, ""},
		{fmt.Sprintf("%s/mock-service-whitelist-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-outway: could not handle your request, organization \"nlx-test\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx-inway: invalid path in url\n"},
		{fmt.Sprintf("%s/mock-service/fictive", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx-inway: no endpoint for service\n"},
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "", http.StatusBadRequest, "nlx-outway: missing logrecord id\n"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		assert.Nil(t, err)

		req.Header.Add("X-NLX-Logrecord-Id", test.logRecordID)
		resp, err := client.Do(req)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, test.statusCode, resp.StatusCode)
		if resp.StatusCode != test.statusCode {
			t.Fatalf(`result: "%d" for url "%s", expected http status code : "%d"`, resp.StatusCode, test.url, test.statusCode)
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error parsing result.body", err)
		}
		assert.Equal(t, test.errorMessage, string(bytes))
	}
}

func TestInwayNoOrgProxyRequest(t *testing.T) {

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "root.crt"),
		OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
		OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
	}

	tlsNoOrgOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "root.crt"),
		OrgCertFile: filepath.Join("..", "testing", "no-org-nlx-test.crt"),
		OrgKeyFile:  filepath.Join("..", "testing", "no-org-nlx-test.key"),
	}

	// Clients with no organization specified in the certificate
	// should not be allowed on the nlx network.
	proxyRequestMockServer, mockEndPoint := newTestEnv(t, tlsOptions)
	proxyRequestMockServer.StartTLS()
	defer proxyRequestMockServer.Close()
	defer mockEndPoint.Close()

	//nolint
	tests := []struct {
		url          string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, ""},
		{fmt.Sprintf("%s/mock-service-whitelist/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, ""},
		{fmt.Sprintf("%s/mock-service-whitelist-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-outway: could not handle your request, organization \"nlx-test\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx inway error: invalid path in url\n"},
		{fmt.Sprintf("%s/mock-service/fictive", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx inway error: no endpoint for service\n"},
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "", http.StatusBadRequest, "nlx-outway: missing logrecord id\n"},
	}

	noOrgClient := setupClient(t, tlsNoOrgOptions)

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		assert.Nil(t, err)

		req.Header.Add("X-NLX-Logrecord-Id", test.logRecordID)
		resp, err := noOrgClient.Do(req)
		assert.Nil(t, err)

		if resp.StatusCode != 400 {
			t.Fatalf(
				`result: "%d" for url "%s", expected http status code : "%d"`,
				resp.StatusCode, test.url, 400)
		}
		resp.Body.Close()
	}
}
