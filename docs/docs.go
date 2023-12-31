// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/raffle/doDraw": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "抽奖管理"
                ],
                "summary": "抽奖",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "raffleKey",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/remark/list": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "说明管理"
                ],
                "summary": "说明",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "raffleKey",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/user/createUser": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "人员管理"
                ],
                "summary": "创建人员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "raffleKey",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号码",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/user/getToken": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取token"
                ],
                "summary": "获取token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号码",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Resp": {
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
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Raffle",
	Description:      "a toy of golang",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
