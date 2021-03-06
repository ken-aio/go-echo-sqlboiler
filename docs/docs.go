// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-03-07 13:36:27.516581272 +0900 JST m=+0.056707945

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server.",
        "title": "Go Echo SQLBoiler sample project",
        "contact": {
            "name": "API Support",
            "url": "dummy",
            "email": "hoge"
        },
        "license": {
            "name": "ken-aio",
            "url": "dummmy"
        },
        "version": "1.0"
    },
    "host": "localhost:1314",
    "basePath": "/",
    "paths": {
        "/api/v1/groups/{group_id}/members": {
            "get": {
                "description": "list group mmebers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GroupMemberList GroupMember list API",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "group id",
                        "name": "group_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.GroupWithMemberList"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "description": "list users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User list API",
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.UserList"
                        }
                    }
                }
            },
            "post": {
                "description": "create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User create API",
                "parameters": [
                    {
                        "description": "user create parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/v1.userCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.UserCreate"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{user_id}": {
            "put": {
                "description": "update users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User update API",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id parameter",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User delete API",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id parameter",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.GroupWithMemberList": {
            "type": "object",
            "properties": {
                "group_id": {
                    "type": "integer"
                },
                "group_name": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "type": "\u0026{%!s(token.Pos=275) groupMemberList}"
                    }
                }
            }
        },
        "models.UserCreate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.UserList": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "v1.userCreateReq": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
