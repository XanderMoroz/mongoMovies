// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/recipes": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get all recipes of authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "get all my recipes",
                "operationId": "get-all-recipes-of-current-user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RecipeModel"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Creating Recipe in DB with given request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "create new recipe",
                "operationId": "create-new-recipe",
                "parameters": [
                    {
                        "description": "Enter recipe data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RecipeCreateBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
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
        },
        "/api/recipes/{id}": {
            "get": {
                "description": "Get a recipe by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "get a recipe by ID",
                "operationId": "get-recipe-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "RecipeID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RecipeModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/users/current_user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Token check method for authentication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Check validity of token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TokenCheckResult"
                        }
                    }
                }
            }
        },
        "/api/users/login": {
            "post": {
                "description": "Login with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login to your account",
                "parameters": [
                    {
                        "description": "UserLogin",
                        "name": "userModelArgs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginArgs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginResult"
                        }
                    }
                }
            }
        },
        "/api/users/register": {
            "post": {
                "description": "Register and create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create a account",
                "parameters": [
                    {
                        "description": "UserRegister",
                        "name": "userModelArgs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegisterArgs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserRegisterResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.IngredientModel": {
            "type": "object",
            "properties": {
                "condition": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                }
            }
        },
        "models.RecipeCreateBody": {
            "type": "object",
            "properties": {
                "stages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.StageModel"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.RecipeModel": {
            "type": "object",
            "properties": {
                "author_email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "stages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.StageModel"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Result": {
            "type": "object",
            "properties": {
                "error_code": {
                    "$ref": "#/definitions/utils.Error"
                },
                "error_description": {
                    "type": "string"
                },
                "error_exception": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.StageModel": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.IngredientModel"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.TokenCheckResult": {
            "type": "object",
            "properties": {
                "expired": {
                    "type": "string"
                },
                "result": {
                    "$ref": "#/definitions/models.Result"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.UserLoginArgs": {
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
        "models.UserLoginResult": {
            "type": "object",
            "properties": {
                "authentication_token": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "result": {
                    "$ref": "#/definitions/models.Result"
                },
                "user_infos": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "models.UserRegisterArgs": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "validate_password": {
                    "type": "string"
                }
            }
        },
        "models.UserRegisterResult": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/models.Result"
                }
            }
        },
        "utils.Error": {
            "type": "string",
            "enum": [
                "ERR0303",
                "ERR0304",
                "ERR0401",
                "ERR0402",
                "ERR0403",
                "ERR0404",
                "ERR0405",
                "ERR0406",
                "ERR0407"
            ],
            "x-enum-varnames": [
                "ERR0303",
                "ERR0304",
                "ERR0401",
                "ERR0402",
                "ERR0403",
                "ERR0404",
                "ERR0405",
                "ERR0406",
                "ERR0407"
            ]
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
	Version:          "1.0",
	Host:             "127.0.0.1:8080/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "MonGorilla Project",
	Description:      "This is a sample server on Gorrilla Mux + MongoDB.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
