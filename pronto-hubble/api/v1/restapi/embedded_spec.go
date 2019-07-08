// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Hubble service for Pronto",
    "title": "Pronto-Hubble",
    "version": "0.1.0"
  },
  "basePath": "/pronto/v1",
  "paths": {
    "/users": {
      "$ref": "./user.yml#/paths/~1users"
    },
    "/users/{id}": {
      "$ref": "./user.yml#/paths/~1users~1%7Bid%7D"
    }
  },
  "definitions": {
    "user": {
      "$ref": "./user.yml#/definitions/user"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Hubble service for Pronto",
    "title": "Pronto-Hubble",
    "version": "0.1.0"
  },
  "basePath": "/pronto/v1",
  "paths": {
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "list",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "users"
        ],
        "operationId": "create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/id"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "get",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "users"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "users"
        ],
        "operationId": "deleteOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "id": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "created": {
          "type": "integer",
          "format": "int32"
        },
        "id": {
          "type": "string",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "role": {
          "type": "string",
          "default": "user"
        }
      }
    }
  }
}`))
}
