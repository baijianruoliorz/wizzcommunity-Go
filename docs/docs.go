// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-10-20 21:15:34.7853667 +0800 CST m=+0.052842601

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "yangxiangrui",
            "url": "https://www.yangxiangrui.site"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/posts2": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可按社区按时间或分数排序查询帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "升级版帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "可以为空",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数据量",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序依据",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/signUp": {
            "post": {
                "description": "用户注册用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "简单注册接口",
                "parameters": [
                    {
                        "description": "注册请求参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        }
    },
    "definitions": {
        "controller._ResponsePostList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ApiPostDetail"
                    }
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.ApiPostDetail": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "description": "作者id",
                    "type": "integer"
                },
                "author_name": {
                    "description": "作者",
                    "type": "string"
                },
                "community_id": {
                    "description": "社区id",
                    "type": "integer"
                },
                "content": {
                    "description": "帖子内容",
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "description": "帖子状态",
                    "type": "integer"
                },
                "title": {
                    "description": "帖子标题",
                    "type": "string"
                },
                "vote_num": {
                    "description": "投票数",
                    "type": "integer"
                }
            }
        },
        "models.ParamPostList": {
            "type": "object",
            "properties": {
                "community_id": {
                    "description": "可以为空",
                    "type": "integer"
                },
                "order": {
                    "description": "排序依据",
                    "type": "string",
                    "example": "score"
                },
                "page": {
                    "description": "页码",
                    "type": "integer",
                    "example": 1
                },
                "size": {
                    "description": "每页数据量",
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "models.ParamSignUp": {
            "type": "object",
            "required": [
                "password",
                "re_password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "re_password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:8083",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "wizzCommunity项目接口文档",
	Description: "wizz成员交流社区",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
