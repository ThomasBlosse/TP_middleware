// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Thomas Blosse",
            "email": "Thomas.BLOSSE@etu.uca.fr"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/alerts": {
            "get": {
                "description": "Retrieves all the alerts in the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "alerts"
                ],
                "summary": "Get all alerts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Alerts"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new alert from the JSON payload and saves it in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "alerts"
                ],
                "summary": "Create a new alert",
                "parameters": [
                    {
                        "description": "Alert to create",
                        "name": "alert",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Alerts"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Alerts"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            }
        },
        "/alerts/{id}": {
            "get": {
                "description": "Retrieves alerts associated with a specific resource",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "alerts"
                ],
                "summary": "Get alerts by resource ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Resource ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Alerts"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates the targets of an existing alert for the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "alerts"
                ],
                "summary": "Update alert",
                "responses": {
                    "200": {
                        "description": "Alert updated successfully",
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
                            "$ref": "#/definitions/models.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an alert associated with the authenticated user",
                "tags": [
                    "alerts"
                ],
                "summary": "Delete alert",
                "responses": {
                    "200": {
                        "description": "Alert deleted successfully",
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
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            }
        },
        "/resources": {
            "get": {
                "description": "Retrieves all the resources",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resources"
                ],
                "summary": "Get all resources",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Resources"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new resource",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resources"
                ],
                "summary": "Create resource",
                "parameters": [
                    {
                        "description": "New resource",
                        "name": "resource",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Resources"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Resources"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            }
        },
        "/resources/{id}": {
            "get": {
                "description": "Retrieves a resource given its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resources"
                ],
                "summary": "Get resource by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Resources"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a resource by its ID",
                "tags": [
                    "resources"
                ],
                "summary": "Delete resource",
                "responses": {
                    "200": {
                        "description": "Resource deleted successfully",
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
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Alerts": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "targets": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.CustomError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "default": 200
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Resources": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
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
	Title:            "Config API",
	Description:      "APi to handle configurations",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
