// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Tetrex"
        },
        "license": {
            "name": "MIT License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Gives us Server Time , To check health of server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "For health check, of server",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.OkResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v1/account/": {
            "post": {
                "description": "takes input of Owner,Currency , and creates account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1/Account"
                ],
                "summary": "Creates account Of user",
                "parameters": [
                    {
                        "description": "CreateAccountRequest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.OkResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/account/{id}": {
            "get": {
                "description": "takes id of user and returns user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1/Account"
                ],
                "summary": "Gets User Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.OkResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/accounts": {
            "post": {
                "description": "takes pages and pagesize",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1/Account"
                ],
                "summary": "Gets List Of User Account",
                "parameters": [
                    {
                        "description": "ListAccountRequest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ListAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.OkResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CreateAccountRequest": {
            "type": "object",
            "required": [
                "currency",
                "owner"
            ],
            "properties": {
                "currency": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "api.ListAccountRequest": {
            "type": "object",
            "required": [
                "page_id",
                "page_size"
            ],
            "properties": {
                "page_id": {
                    "type": "integer",
                    "minimum": 1
                },
                "page_size": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 5
                }
            }
        },
        "util.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "util.OkResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API",
	Description:      "This is a backend api for simple bank",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
