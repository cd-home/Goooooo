// Package admin GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package admin

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "God Yao",
            "email": "liyaoo1995@163.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/db": {
            "get": {
                "description": "Sys DBStats",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sys"
                ],
                "summary": "Sys DBStats",
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "description": "User Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Login",
                "responses": {}
            }
        },
        "/register": {
            "post": {
                "description": "User Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Register",
                "parameters": [
                    {
                        "description": "register",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RegisterParam"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "code\":0,\"data\":null,\"msg\":\"Error\"}",
                        "schema": {
                            "$ref": "#/definitions/types.CommonResponse"
                        }
                    },
                    "1": {
                        "description": "code\":1,\"data\":null,\"msg\":\"Success\"}",
                        "schema": {
                            "$ref": "#/definitions/types.CommonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.CommonResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "types.RegisterParam": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 4
                },
                "password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 6
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "this is Goooooo-Admin Sys.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
