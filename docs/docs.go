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
            "name": "rice408",
            "url": "http://47.108.226.33/",
            "email": "hongwei.wang408@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/info": {
            "get": {
                "description": "管理员信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminInfo"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "管理员登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "parameters": [
                    {
                        "description": "登录",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminLoginResp"
                        }
                    }
                }
            }
        },
        "/admin/register": {
            "post": {
                "description": "管理员注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "parameters": [
                    {
                        "description": "注册",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminRegisterResp"
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
        "/role/add": {
            "post": {
                "description": "添加角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "parameters": [
                    {
                        "description": "添加角色",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddRoleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "添加成功",
                        "schema": {
                            "$ref": "#/definitions/response.AddRoleResp"
                        }
                    }
                }
            }
        },
        "/role/delete": {
            "post": {
                "description": "删除角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "description": "角色id",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.DeleteRoleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功",
                        "schema": {
                            "$ref": "#/definitions/response.DeleteRoleResp"
                        }
                    }
                }
            }
        },
        "/role/list": {
            "get": {
                "description": "查询角色列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "查询角色列表",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "page_num",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询成功",
                        "schema": {
                            "$ref": "#/definitions/response.GetRoleListResp"
                        }
                    }
                }
            }
        },
        "/role/update": {
            "post": {
                "description": "更新角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "parameters": [
                    {
                        "description": "更新角色",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateRoleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新成功",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateRoleResp"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "parameters": [
                    {
                        "description": "登录",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "$ref": "#/definitions/response.UserLoginResp"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "parameters": [
                    {
                        "description": "注册",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "$ref": "#/definitions/response.UserRegisterResp"
                        }
                    },
                    "500": {
                        "description": "参数错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "role_name": {
                    "type": "string"
                }
            }
        },
        "request.AddRoleReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "request.AdminLoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.AdminRegisterReq": {
            "type": "object",
            "required": [
                "email",
                "mobile",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.DeleteRoleReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "request.UpdateRoleReq": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.UserLoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.UserRegisterReq": {
            "type": "object",
            "required": [
                "email",
                "mobile",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.AddRoleResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "response.AdminInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mobile": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.AdminLoginResp": {
            "type": "object",
            "properties": {
                "expire": {
                    "description": "过期时间",
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "response.AdminRegisterResp": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mobile": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.DeleteRoleResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "response.GetRoleListResp": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Role"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.UpdateRoleResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "response.UserLoginResp": {
            "type": "object",
            "properties": {
                "expire": {
                    "description": "过期时间",
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "response.UserRegisterResp": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mobile": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "商城系统",
	Description:      "商城系统接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
