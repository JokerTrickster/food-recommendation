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
        "/v0.1/auth/logout": {
            "post": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패\nPLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "로그아웃 하기",
                "parameters": [
                    {
                        "type": "string",
                        "description": "accessToken",
                        "name": "tkn",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v0.1/auth/password/request": {
            "post": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "비밀번호 변경 요청",
                "parameters": [
                    {
                        "description": "email",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqRequestPassword"
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
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v0.1/auth/password/validate": {
            "post": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "비밀번호 변경 검증",
                "parameters": [
                    {
                        "description": "email, password, code",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqValidatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v0.1/auth/signin": {
            "post": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "로그인",
                "parameters": [
                    {
                        "description": "이메일, 비밀번호",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqSignin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResSignin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v0.1/auth/signup": {
            "post": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "회원 가입",
                "parameters": [
                    {
                        "description": "이름, 이메일, 비밀번호",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqSignup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v0.1/auth/token/reissue": {
            "put": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "액세스 토큰 재발급",
                "parameters": [
                    {
                        "description": "액세스 토큰, 리프레시 토큰",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqReissue"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResReissue"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v0.1/foods/recommend": {
            "post": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패\nPLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "food"
                ],
                "summary": "음식 추천 받기",
                "parameters": [
                    {
                        "type": "string",
                        "description": "accessToken",
                        "name": "tkn",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "type",
                        "name": "type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqRecommendFood"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResRecommendFood"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v0.1/users/check": {
            "get": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패\nPLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "유저 프로필 가져오기",
                "parameters": [
                    {
                        "type": "string",
                        "description": "accessToken",
                        "name": "tkn",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResGetUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v0.1/users/profile": {
            "put": {
                "description": "■ errCode with 400\nPARAM_BAD : 파라미터 오류\nUSER_NOT_EXIST : 유저가 존재하지 않음\nUSER_ALREADY_EXISTED : 유저가 이미 존재\n\n■ errCode with 500\nINTERNAL_SERVER : 내부 로직 처리 실패\nINTERNAL_DB : DB 처리 실패\nPLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "유저 프로필 저장하기",
                "parameters": [
                    {
                        "type": "string",
                        "description": "accessToken",
                        "name": "tkn",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "유저 프로필",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqUpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "request.ReqRecommendFood": {
            "type": "object",
            "properties": {
                "scenario": {
                    "description": "전체 , 혼밥, 가족, 친구들",
                    "type": "string",
                    "example": "혼밥"
                },
                "time": {
                    "type": "string",
                    "example": "점심"
                },
                "type": {
                    "description": "전체, 양식, 한식, 중식 등",
                    "type": "string",
                    "example": "중식"
                }
            }
        },
        "request.ReqReissue": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "request.ReqRequestPassword": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "request.ReqSignin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.ReqSignup": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.ReqUpdateUser": {
            "type": "object",
            "required": [
                "birth",
                "sex"
            ],
            "properties": {
                "birth": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                }
            }
        },
        "request.ReqValidatePassword": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "response.ResGetUser": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                }
            }
        },
        "response.ResRecommendFood": {
            "type": "object",
            "properties": {
                "foodNames": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "response.ResReissue": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "response.ResSignin": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
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
