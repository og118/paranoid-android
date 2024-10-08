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
        "/export/issue/{issue_id}": {
            "get": {
                "description": "Export an issue to EPUB format",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exporters"
                ],
                "summary": "Export an issue",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address to send the EPUB file",
                        "name": "mail_to",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Issue ID",
                        "name": "issue_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Issue exported successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/export/issue/{issue_id}/article/{article_id}": {
            "get": {
                "description": "Export an issue to EPUB format",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exporters"
                ],
                "summary": "Export an issue article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address to send the EPUB file",
                        "name": "mail_to",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Issue ID",
                        "name": "issue_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "article_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Issue exported successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
