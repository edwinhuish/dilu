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
        "/api/v1/demo/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Demo"
                ],
                "summary": "创建",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dilu_modules_demo_service_dto.DemoDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dilu_modules_demo_models.Demo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/demo/del": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Demo"
                ],
                "summary": "删除",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/base.ReqIds"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dilu_modules_demo_models.Demo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/demo/get": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Demo"
                ],
                "summary": "获得",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/base.ReqIds"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dilu_modules_demo_models.Demo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/demo/page": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Demo"
                ],
                "summary": "Page接口",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dilu_modules_demo_service_dto.DemePageReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/base.PageResp"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/dilu_modules_demo_models.Demo"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/demo/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Demo"
                ],
                "summary": "更新",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dilu_modules_demo_service_dto.DemoDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dilu_modules_demo_models.Demo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/init": {
            "get": {
                "description": "init接口",
                "tags": [
                    "Default"
                ],
                "summary": "init接口",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/ping": {
            "get": {
                "description": "Ping接口",
                "tags": [
                    "Default"
                ],
                "summary": "Ping接口",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/utils.Server"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.PageResp": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "数据列表"
                },
                "page": {
                    "description": "当前第几页",
                    "type": "integer"
                },
                "size": {
                    "description": "分页大小",
                    "type": "integer"
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                }
            }
        },
        "base.ReqIds": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "多id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "base.Resp": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "返回码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "msg": {
                    "description": "消息",
                    "type": "string"
                },
                "reqId": {
                    "description": "` + "`" + `json:\"请求id\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "dilu_modules_demo_models.Demo": {
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
        "dilu_modules_demo_service_dto.DemePageReq": {
            "type": "object",
            "properties": {
                "page": {
                    "description": "页码",
                    "type": "integer"
                },
                "size": {
                    "description": "每页大小",
                    "type": "integer"
                }
            }
        },
        "dilu_modules_demo_service_dto.DemoDto": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 2
                }
            }
        },
        "utils.Cpu": {
            "type": "object",
            "properties": {
                "cores": {
                    "type": "integer"
                },
                "cpus": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        },
        "utils.Disk": {
            "type": "object",
            "properties": {
                "totalGb": {
                    "type": "integer"
                },
                "totalMb": {
                    "type": "integer"
                },
                "usedGb": {
                    "type": "integer"
                },
                "usedMb": {
                    "type": "integer"
                },
                "usedPercent": {
                    "type": "integer"
                }
            }
        },
        "utils.Os": {
            "type": "object",
            "properties": {
                "compiler": {
                    "type": "string"
                },
                "goVersion": {
                    "type": "string"
                },
                "goos": {
                    "type": "string"
                },
                "numCpu": {
                    "type": "integer"
                },
                "numGoroutine": {
                    "type": "integer"
                }
            }
        },
        "utils.Ram": {
            "type": "object",
            "properties": {
                "totalMb": {
                    "type": "integer"
                },
                "usedMb": {
                    "type": "integer"
                },
                "usedPercent": {
                    "type": "integer"
                }
            }
        },
        "utils.Server": {
            "type": "object",
            "properties": {
                "cpu": {
                    "$ref": "#/definitions/utils.Cpu"
                },
                "disk": {
                    "$ref": "#/definitions/utils.Disk"
                },
                "os": {
                    "$ref": "#/definitions/utils.Os"
                },
                "ram": {
                    "$ref": "#/definitions/utils.Ram"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "V0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "go-walker API",
	Description:      "一个简单的脚手",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
