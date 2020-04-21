//nolint
package configapi
const (
SwaggerJSONDirectory = `
{
  "swagger": "2.0",
  "info": {
    "title": "configapi.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v1/insight-configuration": {
      "get": {
        "operationId": "ConfigApi_GetInsightConfiguration",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInsightConfiguration"
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
          "ConfigApi"
        ]
      },
      "put": {
        "operationId": "ConfigApi_PutInsightConfiguration",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInsightConfiguration"
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
          "ConfigApi"
        ]
      }
    },
    "/api/v1/inways": {
      "get": {
        "operationId": "ConfigApi_ListInways",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiListInwaysResponse"
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
          "ConfigApi"
        ]
      },
      "post": {
        "operationId": "ConfigApi_CreateInway",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      }
    },
    "/api/v1/inways/{name}": {
      "get": {
        "operationId": "ConfigApi_GetInway",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "delete": {
        "operationId": "ConfigApi_DeleteInway",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiEmpty"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "put": {
        "operationId": "ConfigApi_UpdateInway",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      }
    },
    "/api/v1/services": {
      "get": {
        "operationId": "ConfigApi_ListServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiListServicesResponse"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "inwayName",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "post": {
        "operationId": "ConfigApi_CreateService",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      }
    },
    "/api/v1/services/{name}": {
      "get": {
        "operationId": "ConfigApi_GetService",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "delete": {
        "operationId": "ConfigApi_DeleteService",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiEmpty"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "put": {
        "operationId": "ConfigApi_UpdateService",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      }
    }
  },
  "definitions": {
    "AuthorizationSettingsAuthorization": {
      "type": "object",
      "properties": {
        "organizationName": {
          "type": "string"
        },
        "publicKeyHash": {
          "type": "string"
        }
      }
    },
    "ServiceAuthorizationSettings": {
      "type": "object",
      "properties": {
        "mode": {
          "type": "string"
        },
        "authorizations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/AuthorizationSettingsAuthorization"
          }
        }
      }
    },
    "configapiEmpty": {
      "type": "object"
    },
    "configapiInsightConfiguration": {
      "type": "object",
      "properties": {
        "irmaServerURL": {
          "type": "string"
        },
        "insightAPIURL": {
          "type": "string"
        }
      }
    },
    "configapiInway": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "hostname": {
          "type": "string"
        },
        "selfAddress": {
          "type": "string"
        },
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/configapiInwayService"
          }
        }
      }
    },
    "configapiInwayService": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "configapiListInwaysResponse": {
      "type": "object",
      "properties": {
        "inways": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/configapiInway"
          }
        }
      }
    },
    "configapiListServicesResponse": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/configapiService"
          }
        }
      }
    },
    "configapiService": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "endpointURL": {
          "type": "string"
        },
        "documentationURL": {
          "type": "string"
        },
        "apiSpecificationURL": {
          "type": "string"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean"
        },
        "techSupportContact": {
          "type": "string"
        },
        "publicSupportContact": {
          "type": "string"
        },
        "authorizationSettings": {
          "$ref": "#/definitions/ServiceAuthorizationSettings"
        },
        "inways": {
          "type": "array",
          "items": {
            "type": "string"
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
`)
