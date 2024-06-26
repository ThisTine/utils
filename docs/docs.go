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
            "name": "Thistine",
            "url": "https://thistine.com",
            "email": "tine@thistine.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/qr/v1/make": {
            "get": {
                "description": "Generate QRCode",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "QRCodeAPI"
                ],
                "summary": "Generate QRCode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Text to generate QRCode",
                        "name": "text",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Size of QRCode (in pixel)",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Border of QRCode",
                        "name": "disableBorder",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
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
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Thistine's Utils API",
	Description:      "This contains a set of utilities API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
