// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/ghodss/yaml"
)

type openAPIVersion struct {
	OpenAPI string `json:"openapi"`
	Swagger string `json:"swagger"`
}

const openAPI2 = "OpenAPI2"
const openAPI3 = "OpenAPI3"

func getInwayAPISpecsType(httpClient *http.Client, inwayAddress, serviceName string) (string, error) {
	data, err := getInwayAPISpecs(httpClient, inwayAddress, serviceName)
	if err != nil {
		return "", err
	}

	jsonBytes, err := yaml.YAMLToJSON(data)
	if err != nil {
		return "", err
	}

	versionCheck := &openAPIVersion{}
	err = json.Unmarshal(jsonBytes, versionCheck)

	if err != nil {
		return "", err
	}

	var versionValue string
	if versionCheck.OpenAPI != "" {
		versionValue = versionCheck.OpenAPI
	} else if versionCheck.Swagger != "" {
		versionValue = versionCheck.Swagger
	}

	switch versionValue {
	case "2.0":
		return openAPI2, nil
	case "3.0.0", "3.0.1", "3.0.2":
		return openAPI3, nil
	}

	return "", fmt.Errorf("documentation format is neither openAPI2 or openAPI3")
}

func getInwayAPISpecs(h *http.Client, inwayAddress, serviceName string) ([]byte, error) {
	inwayURL := url.URL{
		Scheme: "https",
		Host:   inwayAddress,
		Path:   path.Join("/.nlx/api-spec-doc/", serviceName),
	}

	res, err := h.Get(inwayURL.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}