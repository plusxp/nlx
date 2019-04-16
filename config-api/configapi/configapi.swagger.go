package configapi
const (
SwaggerJSONDirectory = `
{
  "swagger": "2.0",
  "info": {
    "title": "configapi.proto",
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
  "paths": {},
  "definitions": {
    "ListAnnouncementsResponseAnnouncement": {
      "type": "object",
      "properties": {
        "componentName": {
          "type": "string"
        },
        "componentType": {
          "type": "string"
        }
      }
    },
    "configapiAnnounceResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        }
      }
    },
    "configapiConfig": {
      "type": "object",
      "properties": {
        "kind": {
          "type": "string"
        },
        "config": {
          "type": "string"
        }
      }
    },
    "configapiListAnnouncementsResponse": {
      "type": "object",
      "properties": {
        "announcements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListAnnouncementsResponseAnnouncement"
          }
        }
      }
    },
    "configapiSetConfigResponse": {
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
