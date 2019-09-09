package configapi_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/config-api/configapi"
)

func TestServiceValidate(t *testing.T) {
	testService := configapi.Service{
		Name: "my-service",
	}

	err := testService.Validate()
	assert.Equal(t, errors.New("invalid endpoint URL for service my-service"), err)

	testService = configapi.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	err = testService.Validate()
	assert.Equal(t, nil, err)
}

func TestInwayValidate(t *testing.T) {
	testInway := configapi.Inway{}

	err := testInway.Validate()
	assert.Equal(t, errors.New("invalid inway name"), err)

	testInway = configapi.Inway{
		Name: "inway42.test",
	}

	err = testInway.Validate()
	assert.Equal(t, nil, err)
}