package directoryapi
const (
SwaggerJSONDirectory = `
{
  "swagger": "2.0",
  "info": {
    "title": "directory.proto",
    "description": "Package directoryapi defines the directory api.",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/directory/get-service-api-spec/{organization_name}/{service_name}": {
      "get": {
        "operationId": "GetServiceAPISpec",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/directoryapiGetServiceAPISpecResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "organization_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "service_name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Directory"
        ]
      }
    },
    "/directory/list-organizations": {
      "get": {
        "summary": "ListOrganizations lists all organizations and their details.",
        "operationId": "ListOrganizations",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/directoryapiListOrganizationsResponse"
            }
          }
        },
        "tags": [
          "Directory"
        ]
      }
    },
    "/directory/list-services": {
      "get": {
        "summary": "ListServices lists all services and their gateways.",
        "operationId": "ListServices",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/directoryapiListServicesResponse"
            }
          }
        },
        "tags": [
          "Directory"
        ]
      }
    }
  },
  "definitions": {
    "ListOrganizationsResponseOrganization": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "insight_irma_endpoint": {
          "type": "string"
        },
        "insight_log_endpoint": {
          "type": "string"
        }
      }
    },
    "ListServicesResponseService": {
      "type": "object",
      "properties": {
        "organization_name": {
          "type": "string"
        },
        "service_name": {
          "type": "string"
        },
        "inway_addresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "documentation_url": {
          "type": "string"
        },
        "api_specification_type": {
          "type": "string"
        }
      }
    },
    "RegisterInwayRequestRegisterService": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "documentation_url": {
          "type": "string"
        },
        "api_specification_type": {
          "type": "string"
        }
      }
    },
    "directoryapiGetServiceAPISpecResponse": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "document": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "directoryapiListOrganizationsResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "organizations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListOrganizationsResponseOrganization"
          }
        }
      }
    },
    "directoryapiListServicesResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListServicesResponseService"
          }
        }
      }
    },
    "directoryapiRegisterInwayResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        }
      }
    }
  }
}
`)
