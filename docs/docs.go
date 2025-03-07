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
        "/api/v1/diary/create": {
            "post": {
                "description": "Create a diary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diary service"
                ],
                "summary": "Create a diary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Request ID",
                        "name": "x-request-id",
                        "in": "header"
                    },
                    {
                        "description": "The command to create a diary",
                        "name": "CreateCommand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/diary.CreateCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/diary.CreateView"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/diary/pageList": {
            "post": {
                "description": "list diaries in page according to the last id and page size",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diary service"
                ],
                "summary": "list diaries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Request ID",
                        "name": "x-request-id",
                        "in": "header"
                    },
                    {
                        "description": "The command to page list",
                        "name": "Query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/diary.Query"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/diary.QueryView"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/stickynote/create": {
            "post": {
                "description": "Create a sticky note",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sticky note service"
                ],
                "summary": "Create a sticky note",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Request ID",
                        "name": "x-request-id",
                        "in": "header"
                    },
                    {
                        "description": "The command to create a sticky note",
                        "name": "CreateCommand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/stickynote.CreateCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/stickynote.CreateView"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/stickynote/generateDiaryContent": {
            "post": {
                "description": "Generate content for a diary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sticky note service"
                ],
                "summary": "Generate content for a diary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Request ID",
                        "name": "x-request-id",
                        "in": "header"
                    },
                    {
                        "description": "The command to generate content for a diary",
                        "name": "GenerateContentCommand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/stickynote.GenerateContentCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/stickynote.GenerateContentView"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/stickynote/modify": {
            "post": {
                "description": "Modify a sticky note",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sticky note service"
                ],
                "summary": "Modify a sticky note",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Request ID",
                        "name": "x-request-id",
                        "in": "header"
                    },
                    {
                        "description": "The command to modify a sticky note",
                        "name": "ModifyCommand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/stickynote.ModifyCommand"
                        }
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
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "diary.CreateCommand": {
            "type": "object",
            "required": [
                "diary_date",
                "uid"
            ],
            "properties": {
                "diary_date": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "diary.CreateView": {
            "type": "object",
            "properties": {
                "diary_id": {
                    "type": "string"
                }
            }
        },
        "diary.Query": {
            "type": "object",
            "required": [
                "page_size"
            ],
            "properties": {
                "last_id": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                }
            }
        },
        "diary.QueryView": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "diary_date_str": {
                    "type": "string"
                },
                "diary_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "stickynote.CreateCommand": {
            "type": "object",
            "required": [
                "content",
                "diary_id",
                "occurrence_time_str",
                "participants",
                "uid"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "diary_id": {
                    "type": "string"
                },
                "occurrence_time_str": {
                    "type": "string"
                },
                "participants": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "stickynote.CreateView": {
            "type": "object",
            "properties": {
                "sticky_note_id": {
                    "type": "string"
                }
            }
        },
        "stickynote.GenerateContentCommand": {
            "type": "object",
            "required": [
                "diary_id"
            ],
            "properties": {
                "diary_id": {
                    "type": "string"
                }
            }
        },
        "stickynote.GenerateContentView": {
            "type": "object",
            "properties": {
                "generated_content": {
                    "type": "string"
                }
            }
        },
        "stickynote.ModifyCommand": {
            "type": "object",
            "required": [
                "content",
                "participants",
                "sticky_note_id"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "participants": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "sticky_note_id": {
                    "type": "string"
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
