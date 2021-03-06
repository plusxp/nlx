package inspectionapi

const (
	SwaggerJSONDirectoryInspection = `
{
  "swagger": "2.0",
  "info": {
    "title": "inspectionapi.proto",
    "description": "Package inspectionapi defines the directory api.",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/directory/list-organizations": {
      "get": {
        "summary": "ListOrganizations lists all organizations and their details.",
        "operationId": "DirectoryInspection_ListOrganizations",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/inspectionapiListOrganizationsResponse"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "tags": [
          "DirectoryInspection"
        ]
      }
    },
    "/api/directory/list-services": {
      "get": {
        "summary": "ListServices lists all services and their gateways.",
        "operationId": "DirectoryInspection_ListServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/inspectionapiListServicesResponse"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "tags": [
          "DirectoryInspection"
        ]
      }
    }
  },
  "definitions": {
    "InwayState": {
      "type": "string",
      "enum": [
        "UNKNOWN",
        "UP",
        "DOWN"
      ],
      "default": "UNKNOWN"
    },
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
        },
        "internal": {
          "type": "boolean",
          "format": "boolean"
        },
        "public_support_contact": {
          "type": "string"
        },
        "healthy_states": {
          "type": "array",
          "items": {
            "type": "boolean",
            "format": "boolean"
          }
        },
        "inways": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/inspectionapiInway"
          }
        },
        "oneTimeCosts": {
          "type": "integer",
          "format": "int32"
        },
        "monthlyCosts": {
          "type": "integer",
          "format": "int32"
        },
        "requestCosts": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "inspectionapiGetOrganizationInwayResponse": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        }
      }
    },
    "inspectionapiInway": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "state": {
          "$ref": "#/definitions/InwayState"
        }
      }
    },
    "inspectionapiListOrganizationsResponse": {
      "type": "object",
      "properties": {
        "organizations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListOrganizationsResponseOrganization"
          }
        }
      }
    },
    "inspectionapiListServicesResponse": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListServicesResponseService"
          }
        }
      }
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "runtimeError": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    }
  }
}
`
)
