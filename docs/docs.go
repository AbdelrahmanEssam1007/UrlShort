// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/create-short-url": {
            "post": {
                "description": "Accepts a long URL and user ID and returns a shortened version",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortener"
                ],
                "summary": "Create a new short URL",
                "parameters": [
                    {
                        "description": "Original URL and User ID",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UrlCreationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/{shortUrl}": {
            "get": {
                "description": "Redirects the user from a short URL to the original long URL",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Shortener"
                ],
                "summary": "Redirect to original URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL code",
                        "name": "shortUrl",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.UrlCreationRequest": {
            "type": "object",
            "required": [
                "long_url",
                "user_id"
            ],
            "properties": {
                "long_url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9808",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "URL Shortener API",
	Description:      "A simple URL shortener built with Go, Gin, and Redis.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
