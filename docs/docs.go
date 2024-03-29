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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.OkResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v1/account": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.OkResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/db.Account"
                                        }
                                    }
                                }
                            ]
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.OkResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/db.Account"
                                        }
                                    }
                                }
                            ]
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.OkResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/db.Account"
                                            }
                                        }
                                    }
                                }
                            ]
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
        "/v1/login": {
            "post": {
                "description": "returns accessToken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1/login"
                ],
                "summary": "logs in user",
                "parameters": [
                    {
                        "description": "loginUserRequest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.loginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.OkResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.loginUserResponse"
                                        }
                                    }
                                }
                            ]
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
        "/v1/transfer": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "takes input and transfers money from -\u003e to",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1/TransferMoney"
                ],
                "summary": "Transfer's money from Acc1 to Acc2",
                "parameters": [
                    {
                        "description": "TransferRequest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.TransferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.OkResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/db.TransferTxResult"
                                        }
                                    }
                                }
                            ]
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
        "/v1/user": {
            "post": {
                "description": "returns user newly created user profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1/User"
                ],
                "summary": "Creates user profile",
                "parameters": [
                    {
                        "description": "CreateUserRequest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.OkResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.userResponse"
                                        }
                                    }
                                }
                            ]
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
        "api.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "full_name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
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
        "api.TransferRequest": {
            "type": "object",
            "required": [
                "amount",
                "currency",
                "from_account_id",
                "to_account_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "currency": {
                    "type": "string",
                    "enum": [
                        "USD",
                        "EUR",
                        "CAD"
                    ]
                },
                "from_account_id": {
                    "type": "integer",
                    "minimum": 1
                },
                "to_account_id": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "api.loginUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.loginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.userResponse"
                }
            }
        },
        "api.userResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password_changed_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "db.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "db.Entry": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "amount": {
                    "description": "can be negative or positive",
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "db.Transfer": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "must be positive",
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "from_account_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "to_account_id": {
                    "type": "integer"
                }
            }
        },
        "db.TransferTxResult": {
            "type": "object",
            "properties": {
                "from_account": {
                    "$ref": "#/definitions/db.Account"
                },
                "from_entry": {
                    "$ref": "#/definitions/db.Entry"
                },
                "to_account": {
                    "$ref": "#/definitions/db.Account"
                },
                "to_entry": {
                    "$ref": "#/definitions/db.Entry"
                },
                "transfer": {
                    "$ref": "#/definitions/db.Transfer"
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
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
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
