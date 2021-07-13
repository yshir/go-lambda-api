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
    "description": "A Example app for serverless api",
    "title": "Example app",
    "version": "1.0.0"
  },
  "host": "localhost:18888",
  "basePath": "/v1",
  "paths": {
    "/users": {
      "get": {
        "summary": "ユーザ一覧取得",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/users"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "summary": "ユーザ登録",
        "operationId": "postUsers",
        "parameters": [
          {
            "description": "登録するユーザ情報",
            "name": "postUsers",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "description": "エラー",
      "type": "object",
      "title": "Error",
      "properties": {
        "message": {
          "description": "エラーメッセージ",
          "type": "string"
        }
      }
    },
    "user": {
      "description": "ユーザ",
      "type": "object",
      "title": "User",
      "required": [
        "user_id",
        "name"
      ],
      "properties": {
        "name": {
          "description": "ユーザ名",
          "type": "string",
          "x-omitempty": false
        },
        "user_id": {
          "description": "ユーザID",
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "users": {
      "description": "ユーザ一覧",
      "type": "array",
      "title": "users",
      "items": {
        "$ref": "#/definitions/user"
      }
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
    "description": "A Example app for serverless api",
    "title": "Example app",
    "version": "1.0.0"
  },
  "host": "localhost:18888",
  "basePath": "/v1",
  "paths": {
    "/users": {
      "get": {
        "summary": "ユーザ一覧取得",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/users"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "summary": "ユーザ登録",
        "operationId": "postUsers",
        "parameters": [
          {
            "description": "登録するユーザ情報",
            "name": "postUsers",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "description": "エラー",
      "type": "object",
      "title": "Error",
      "properties": {
        "message": {
          "description": "エラーメッセージ",
          "type": "string"
        }
      }
    },
    "user": {
      "description": "ユーザ",
      "type": "object",
      "title": "User",
      "required": [
        "user_id",
        "name"
      ],
      "properties": {
        "name": {
          "description": "ユーザ名",
          "type": "string",
          "x-omitempty": false
        },
        "user_id": {
          "description": "ユーザID",
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "users": {
      "description": "ユーザ一覧",
      "type": "array",
      "title": "users",
      "items": {
        "$ref": "#/definitions/user"
      }
    }
  }
}`))
}
