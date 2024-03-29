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
    "/clusters": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "listClusters",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "clusters"
        ],
        "operationId": "createCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "./common.yml#/id"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
            }
          }
        }
      }
    },
    "/clusters/{id}": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "getCluster",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "clusters"
        ],
        "operationId": "updateCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "clusters"
        ],
        "operationId": "deleteCluster",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
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
    },
    "/clusters/{id}/join": {
      "put": {
        "tags": [
          "clusters"
        ],
        "operationId": "joinCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/deviceIds"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
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
    },
    "/clusters/{id}/leave": {
      "put": {
        "tags": [
          "clusters"
        ],
        "operationId": "leaveCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/deviceIds"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
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
    },
    "/devices": {
      "get": {
        "tags": [
          "devices"
        ],
        "operationId": "listDevices",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/device"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "devices"
        ],
        "operationId": "registerDevice",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/device"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "./common.yml#/id"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
            }
          }
        }
      }
    },
    "/devices/{id}": {
      "get": {
        "tags": [
          "devices"
        ],
        "operationId": "getDevice",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/device"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "devices"
        ],
        "operationId": "updateDevice",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/device"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/device"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "devices"
        ],
        "operationId": "deleteDevice",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
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
    },
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "listUsers",
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
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "users"
        ],
        "operationId": "createUser",
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
              "$ref": "./common.yml#/id"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
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
        "operationId": "getUser",
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
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "users"
        ],
        "operationId": "updateUser",
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
              "$ref": "./common.yml#/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "users"
        ],
        "operationId": "deleteUser",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./common.yml#/error"
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
    "cluster": {
      "$ref": "./cluster.yml#/definitions/cluster"
    },
    "device": {
      "$ref": "./device.yml#/definitions/device"
    },
    "deviceIds": {
      "$ref": "./cluster.yml#/definitions/deviceIds"
    },
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
    "/clusters": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "listClusters",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
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
          "clusters"
        ],
        "operationId": "createCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cluster"
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
    "/clusters/{id}": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "getCluster",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/cluster"
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
          "clusters"
        ],
        "operationId": "updateCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/cluster"
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
          "clusters"
        ],
        "operationId": "deleteCluster",
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
    },
    "/clusters/{id}/join": {
      "put": {
        "tags": [
          "clusters"
        ],
        "operationId": "joinCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/deviceIds"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/cluster"
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
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/clusters/{id}/leave": {
      "put": {
        "tags": [
          "clusters"
        ],
        "operationId": "leaveCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/deviceIds"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/cluster"
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
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/devices": {
      "get": {
        "tags": [
          "devices"
        ],
        "operationId": "listDevices",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/device"
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
          "devices"
        ],
        "operationId": "registerDevice",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/device"
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
    "/devices/{id}": {
      "get": {
        "tags": [
          "devices"
        ],
        "operationId": "getDevice",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/device"
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
          "devices"
        ],
        "operationId": "updateDevice",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/device"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/device"
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
          "devices"
        ],
        "operationId": "deleteDevice",
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
    },
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "listUsers",
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
        "operationId": "createUser",
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
        "operationId": "getUser",
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
        "operationId": "updateUser",
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
        "operationId": "deleteUser",
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
    "cluster": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "created": {
          "type": "integer",
          "format": "int32"
        },
        "deviceIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "string",
          "minLength": 3
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "device": {
      "type": "object",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "created": {
          "type": "integer",
          "format": "int32"
        },
        "id": {
          "type": "string",
          "minLength": 3
        },
        "metaData": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "deviceIds": {
      "type": "object",
      "properties": {
        "deviceIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
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
