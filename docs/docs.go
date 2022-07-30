// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/:code": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authenticationService"
                ],
                "summary": "Authenticates a user by code grant",
                "operationId": "authUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Discord code grant",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apiModel.UserAuth"
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
        },
        "/user": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userService"
                ],
                "summary": "Gets a user given its authentication token",
                "operationId": "getUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Discord code grant",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apiModel.User"
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
        },
        "/user/guilds": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guildService"
                ],
                "summary": "Gets a user's guilds given its authentication token",
                "operationId": "getGuilds",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Discord authentication token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Get guilds after this id",
                        "name": "afterId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/discordgo.UserGuild"
                            }
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
        },
        "/user/guilds/:guildId": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guildService"
                ],
                "summary": "Gets a guild only if the user is inside it and we have a record of it",
                "operationId": "getGuild",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Discord authentication token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Guild id",
                        "name": "guildId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/processedModel.Guild"
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
    },
    "definitions": {
        "apiModel.User": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "accent_color": {
                    "type": "integer"
                },
                "avatar": {
                    "type": "string"
                },
                "banner": {
                    "type": "string"
                },
                "discriminator": {
                    "type": "string"
                },
                "locale": {
                    "type": "string"
                },
                "mfa_enabled": {
                    "type": "boolean"
                },
                "premium_type": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "apiModel.UserAuth": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "discordgo.UserGuild": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "boolean"
                },
                "permissions": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "processedModel.ChannelData": {
            "type": "object",
            "properties": {
                "channelName": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                }
            }
        },
        "processedModel.Guild": {
            "type": "object",
            "properties": {
                "guildID": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "topMessageUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/processedModel.UserScore"
                    }
                },
                "topUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/processedModel.UserScore"
                    }
                },
                "userData": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/processedModel.User"
                    }
                },
                "userMessageData": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/processedModel.User"
                    }
                }
            }
        },
        "processedModel.User": {
            "type": "object",
            "properties": {
                "channelData": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/processedModel.ChannelData"
                    }
                },
                "score": {
                    "type": "integer"
                }
            }
        },
        "processedModel.UserScore": {
            "type": "object",
            "properties": {
                "score": {
                    "type": "integer"
                },
                "username": {
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
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "DC-STATS user API",
	Description:      "This api helps user see their data using discord code grant authentication",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
