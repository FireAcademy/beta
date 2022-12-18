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
        "termsOfService": "https://fireacademy.io/terms-and-conditions.txt",
        "contact": {
            "name": "yakuhito",
            "url": "https://twitter.com/yakuh1t0",
            "email": "y@kuhi.to"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/FireAcademy/beta/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/get_peak_synced_block": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Beta constantly syncs to the blockchain by fetching the latest blocks. This function returns information about the latest processed block (the one with the biggest 'height' value).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sync"
                ],
                "summary": "returns the peak synced block",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetPeakSyncedBlockResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/NoAPIKeyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Beta constantly syncs to the blockchain by fetching the latest blocks. This function returns information about the latest processed block (the one with the biggest 'height' value).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sync"
                ],
                "summary": "returns the peak synced block",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetPeakSyncedBlockResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/NoAPIKeyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get_puzzle": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Beta stores all revealed inner puzzles of singletons. Use this method to get it from the corresponding puzzle hash. The puzzle will be returned as a hex string.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Puzzles"
                ],
                "summary": "returns the stored puzzle for a given puzzle_hash",
                "parameters": [
                    {
                        "description": "The inner puzzle hash",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/GetPuzzleArgs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetPuzzleResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/NoAPIKeyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get_singleton_states": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Singletons are like souls - when a coin dies (gets spent), they might move on to a new coin or disappear (melt). Beta views these transitions as changes of state.\nThis endpoint takes multiple optional arguments and returns the states that match the criteria. Most body parameters are self-explanatory: if they match a field of the returned struct model (e.g., 'height' or 'launcher_id'), they act as filters. Only states with the specified values will be returned.\nThe first special parameter is 'limit' - by default, this function returns 100 results at most. Use this parameter to tweak this value.\n'order_by' can be 'coin_id', 'header_hash', 'height', 'parent_coin_id', 'puzzle_hash', 'amount', 'launcher_id', or 'inner_puzzle_hash'. The default ordering is ascending, but that can be changed by setting the 'order' parameter to 'desc'.\nIf your query returns more than 100 results and you need all of them for some reasons, you can also use the 'offset' parameter.\nNote: Two singleton states CAN have the same coin id but a different 'melted' value - the primary key is a composite one: (coin_id, melted)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Singleton States"
                ],
                "summary": "returns singleton states for given conditions",
                "parameters": [
                    {
                        "description": "The arguments (see description for more details)",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/GetSingletonStatesArgs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetSingletonStatesResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/NoAPIKeyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get_synced_block": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Beta constantly syncs to the blockchain by fetching the latest blocks. This function returns information about the block at the given height.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sync"
                ],
                "summary": "returns block for given height",
                "parameters": [
                    {
                        "description": "The height",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/GetSyncedBlockArgs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetSyncedBlockResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/NoAPIKeyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get_synced_blocks": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Beta constantly syncs to the blockchain by fetching the latest blocks. This function returns information about the blocks with height in [start, end).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sync"
                ],
                "summary": "returns blocks for given range",
                "parameters": [
                    {
                        "description": "The start and end heights",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/GetSyncedBlocksArgs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetSyncedBlocksResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/NoAPIKeyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "example error message"
                },
                "results": {
                    "type": "integer",
                    "example": 0
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "GetPeakSyncedBlockResponse": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "integer",
                    "example": 1
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "synced_block": {
                    "$ref": "#/definitions/SyncedBlock"
                }
            }
        },
        "GetPuzzleArgs": {
            "type": "object",
            "properties": {
                "puzzle_hash": {
                    "type": "string",
                    "example": "6b665c0e059050f71a1c3e8a7d5b58e4e1d7abbd02d937e9b5ab5abfd7f8eaba"
                }
            }
        },
        "GetPuzzleResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/Puzzle"
                },
                "results": {
                    "type": "integer",
                    "example": 1
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "GetSingletonStatesArgs": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "coin_id": {
                    "type": "string",
                    "format": "hex",
                    "example": "625799464319c8703ac2d0664af98cf45b9b306f7dcf717b1070d170bb5916a9"
                },
                "header_hash": {
                    "type": "string",
                    "format": "hex",
                    "example": "796c33c3905150e649211fdd9ed42c7c418758c30c321271973a7c792a5bd403"
                },
                "height": {
                    "type": "integer",
                    "format": "int64",
                    "example": 2174316
                },
                "inner_puzzle_hash": {
                    "type": "string",
                    "format": "hex",
                    "example": "6b665c0e059050f71a1c3e8a7d5b58e4e1d7abbd02d937e9b5ab5abfd7f8eaba"
                },
                "launcher_id": {
                    "type": "string",
                    "format": "hex",
                    "example": "f4dd6f4ec490974f7eb98223748f47340a9e9363b4c2dccc1932cdbbc54d03fd"
                },
                "limit": {
                    "type": "integer",
                    "example": 7
                },
                "offset": {
                    "type": "integer",
                    "example": 1
                },
                "order": {
                    "type": "string",
                    "example": "desc"
                },
                "order_by": {
                    "type": "string",
                    "example": "amount"
                },
                "parent_coin_id": {
                    "type": "string",
                    "format": "hex",
                    "example": "e9676e8ce096c5be27dee2fbf2120054d206e4df2de9ef59c24a651d3c558c95"
                },
                "puzzle_hash": {
                    "type": "string",
                    "format": "hex",
                    "example": "0a5a9c760970ebcc094c6f9faa3d9730f066c7a8f7450841a94fc4fd59229bc2"
                }
            }
        },
        "GetSingletonStatesResponse": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "integer",
                    "example": 1
                },
                "singleton_states": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/SingletonState"
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "GetSyncedBlockArgs": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer",
                    "example": 2000000
                }
            }
        },
        "GetSyncedBlockResponse": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "integer",
                    "example": 1
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "synced_block": {
                    "$ref": "#/definitions/SyncedBlock"
                }
            }
        },
        "GetSyncedBlocksArgs": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "integer",
                    "example": 2000001
                },
                "start": {
                    "type": "integer",
                    "example": 2000000
                }
            }
        },
        "GetSyncedBlocksResponse": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "integer",
                    "example": 1
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "synced_blocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/SyncedBlock"
                    }
                }
            }
        },
        "NoAPIKeyResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "No API key provided."
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "Puzzle": {
            "type": "object",
            "properties": {
                "puzzle": {
                    "type": "string",
                    "format": "hex",
                    "example": "ff02ffff01ff02ffff01ff02ff3effff04ff02ffff04ff05ffff04ffff02ff2fff5f80ffff04ff80ffff04ffff04ffff04ff0bffff04ff17ff808080ffff01ff808080ffff01ff8080808080808080ffff04ffff01ffffff0233ff04ff0101ffff02ff02ffff03ff05ffff01ff02ff1affff04ff02ffff04ff0dffff04ffff0bff12ffff0bff2cff1480ffff0bff12ffff0bff12ffff0bff2cff3c80ff0980ffff0bff12ff0bffff0bff2cff8080808080ff8080808080ffff010b80ff0180ffff0bff12ffff0bff2cff1080ffff0bff12ffff0bff12ffff0bff2cff3c80ff0580ffff0bff12ffff02ff1affff04ff02ffff04ff07ffff04ffff0bff2cff2c80ff8080808080ffff0bff2cff8080808080ffff02ffff03ffff07ff0580ffff01ff0bffff0102ffff02ff2effff04ff02ffff04ff09ff80808080ffff02ff2effff04ff02ffff04ff0dff8080808080ffff01ff0bffff0101ff058080ff0180ff02ffff03ff0bffff01ff02ffff03ffff09ff23ff1880ffff01ff02ffff03ffff18ff81b3ff2c80ffff01ff02ffff03ffff20ff1780ffff01ff02ff3effff04ff02ffff04ff05ffff04ff1bffff04ff33ffff04ff2fffff04ff5fff8080808080808080ffff01ff088080ff0180ffff01ff04ff13ffff02ff3effff04ff02ffff04ff05ffff04ff1bffff04ff17ffff04ff2fffff04ff5fff80808080808080808080ff0180ffff01ff02ffff03ffff09ff23ffff0181e880ffff01ff02ff3effff04ff02ffff04ff05ffff04ff1bffff04ff17ffff04ffff02ffff03ffff22ffff09ffff02ff2effff04ff02ffff04ff53ff80808080ff82014f80ffff20ff5f8080ffff01ff02ff53ffff04ff818fffff04ff82014fffff04ff81b3ff8080808080ffff01ff088080ff0180ffff04ff2cff8080808080808080ffff01ff04ff13ffff02ff3effff04ff02ffff04ff05ffff04ff1bffff04ff17ffff04ff2fffff04ff5fff80808080808080808080ff018080ff0180ffff01ff04ffff04ff18ffff04ffff02ff16ffff04ff02ffff04ff05ffff04ff27ffff04ffff0bff2cff82014f80ffff04ffff02ff2effff04ff02ffff04ff818fff80808080ffff04ffff0bff2cff0580ff8080808080808080ff378080ff81af8080ff0180ff018080ffff04ffff01a0a04d9f57764f54a43e4030befb4d80026e870519aaa66334aef8304f5d0393c2ffff04ffff01ffff75ffc05968747470733a2f2f6261666b726569657175656b6879786b34643575366e77653462733378366f6c7369616b646269753467353276736a7665666f6b78686f6d6b6d712e697066732e6e667473746f726167652e6c696e6b2f80ffff68a090a1147c5d5c1f69e6d89c0cb77f3972401430a29c37755926a42b9573b98a64ffff826d75ffc05968747470733a2f2f6261666b7265696468336273697933626835686870766e6d367572357176757474373566783277697772646f62736d70686e6c3237776f726372792e697066732e6e667473746f726167652e6c696e6b2f80ffff826c7580ffff82736e01ffff82737401ffff826d68a067d8648c6c27e9cefab59ea47b0ad273ff4b7d591688dc1931e76af5fb3a228e80ffff04ffff01a0fe8a4b4e27a2e29a4d3fc7ce9d527adbcaccbab6ada3903ccf3ba9a769d2d78bffff04ffff01ff02ffff01ff02ffff01ff02ffff03ff0bffff01ff02ffff03ffff09ff05ffff1dff0bffff1effff0bff0bffff02ff06ffff04ff02ffff04ff17ff8080808080808080ffff01ff02ff17ff2f80ffff01ff088080ff0180ffff01ff04ffff04ff04ffff04ff05ffff04ffff02ff06ffff04ff02ffff04ff17ff80808080ff80808080ffff02ff17ff2f808080ff0180ffff04ffff01ff32ff02ffff03ffff07ff0580ffff01ff0bffff0102ffff02ff06ffff04ff02ffff04ff09ff80808080ffff02ff06ffff04ff02ffff04ff0dff8080808080ffff01ff0bffff0101ff058080ff0180ff018080ffff04ffff01b093c8e9237b64d8e9282a7f8e73369c671bdf4e8fa631704c58ae8aa7a8912322994e22df1b01dd8c896e13cf1683ba2dff018080ff018080808080"
                },
                "puzzle_hash": {
                    "type": "string",
                    "format": "hex",
                    "example": "6b665c0e059050f71a1c3e8a7d5b58e4e1d7abbd02d937e9b5ab5abfd7f8eaba"
                }
            }
        },
        "SingletonState": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "coin_id": {
                    "type": "string",
                    "format": "hex",
                    "example": "625799464319c8703ac2d0664af98cf45b9b306f7dcf717b1070d170bb5916a9"
                },
                "header_hash": {
                    "type": "string",
                    "format": "hex",
                    "example": "796c33c3905150e649211fdd9ed42c7c418758c30c321271973a7c792a5bd403"
                },
                "height": {
                    "type": "integer",
                    "format": "int64",
                    "example": 2174316
                },
                "inner_puzzle_hash": {
                    "type": "string",
                    "format": "hex",
                    "example": "6b665c0e059050f71a1c3e8a7d5b58e4e1d7abbd02d937e9b5ab5abfd7f8eaba"
                },
                "launcher_id": {
                    "type": "string",
                    "format": "hex",
                    "example": "f4dd6f4ec490974f7eb98223748f47340a9e9363b4c2dccc1932cdbbc54d03fd"
                },
                "melted": {
                    "type": "boolean",
                    "example": false
                },
                "parent_coin_id": {
                    "type": "string",
                    "format": "hex",
                    "example": "e9676e8ce096c5be27dee2fbf2120054d206e4df2de9ef59c24a651d3c558c95"
                },
                "puzzle_hash": {
                    "type": "string",
                    "format": "hex",
                    "example": "0a5a9c760970ebcc094c6f9faa3d9730f066c7a8f7450841a94fc4fd59229bc2"
                }
            }
        },
        "SyncedBlock": {
            "type": "object",
            "properties": {
                "header_hash": {
                    "type": "string",
                    "format": "hex",
                    "example": "9bb7135ae2b4a207807be2661007a89c8e3d0de1a58c0670da07b6099b9fedf7"
                },
                "height": {
                    "type": "integer",
                    "format": "int64",
                    "example": 2000000
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "The API key can also be included in the URL or as a GET parameter - please see [https://docs.fireacademy.io/developers/using-api-keys](https://docs.fireacademy.io/developers/using-api-keys) for more details.",
            "type": "apiKey",
            "name": "X-API-Key",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "kraken.fireacademy.io",
	BasePath:         "/beta",
	Schemes:          []string{},
	Title:            "Beta API",
	Description:      "The Beta FireAcademy.io API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
