///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Service Manager\n",
    "title": "Service Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/serviceclass": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceClass"
        ],
        "summary": "List all existing service classes",
        "operationId": "getServiceClasses",
        "parameters": [
          {
            "type": "string",
            "description": "Broker name",
            "name": "broker",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter on service class tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ServiceClass"
              }
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/serviceclass/{serviceClassName}": {
      "get": {
        "description": "Returns a single service class",
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceClass"
        ],
        "summary": "Find service class by name",
        "operationId": "getServiceClassByName",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ServiceClass"
            }
          },
          "400": {
            "description": "Invalid name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Service class not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of service class to return",
          "name": "serviceClassName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/serviceinstance": {
      "get": {
        "description": "List all service instances",
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceInstance"
        ],
        "summary": "Get all service instances",
        "operationId": "getServiceInstances",
        "parameters": [
          {
            "type": "string",
            "description": "service class name",
            "name": "serviceclass",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter on service instance tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ServiceInstance"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceInstance"
        ],
        "summary": "Add a new service instance",
        "operationId": "addServiceInstance",
        "parameters": [
          {
            "description": "Service instance object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ServiceInstance"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/ServiceInstance"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/serviceinstance/{serviceInstanceName}": {
      "get": {
        "description": "Returns a single service instance",
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceInstance"
        ],
        "summary": "Find service instance by name",
        "operationId": "getServiceInstanceByName",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ServiceInstance"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Service instance not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceInstance"
        ],
        "summary": "Deletes a service instance",
        "operationId": "deleteServiceInstanceByName",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ServiceInstance"
            }
          },
          "400": {
            "description": "Invalid name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Service instance not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of service instance to return",
          "name": "serviceInstanceName",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ServiceBinding": {
      "type": "object",
      "properties": {
        "bindingSecret": {
          "type": "string"
        },
        "createdTime": {
          "type": "integer"
        },
        "parameters": {
          "type": "object"
        },
        "reason": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "secretParameters": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "$ref": "#/definitions/Status"
        }
      }
    },
    "ServiceClass": {
      "type": "object",
      "required": [
        "name",
        "broker"
      ],
      "properties": {
        "bindable": {
          "type": "boolean"
        },
        "broker": {
          "type": "string"
        },
        "createdTime": {
          "type": "integer"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "kind": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "plans": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ServicePlan"
          }
        },
        "reason": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "$ref": "#/definitions/Status"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "ServiceInstance": {
      "type": "object",
      "required": [
        "name",
        "serviceClass",
        "servicePlan"
      ],
      "properties": {
        "binding": {
          "$ref": "#/definitions/ServiceBinding"
        },
        "createdTime": {
          "type": "integer"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "kind": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "parameters": {
          "type": "object"
        },
        "reason": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "secretParameters": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "serviceClass": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "servicePlan": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "status": {
          "$ref": "#/definitions/Status"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "ServicePlan": {
      "type": "object",
      "properties": {
        "bindable": {
          "type": "boolean"
        },
        "description": {
          "type": "string"
        },
        "free": {
          "type": "boolean"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "kind": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "readOnly": true
        },
        "metadata": {
          "type": "object"
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "schema": {
          "$ref": "#/definitions/ServicePlanSchema"
        }
      }
    },
    "ServicePlanSchema": {
      "type": "object",
      "properties": {
        "bind": {
          "type": "object"
        },
        "create": {
          "type": "object"
        },
        "update": {
          "type": "object"
        }
      }
    },
    "Status": {
      "type": "string",
      "enum": [
        "INITIALIZED",
        "CREATING",
        "READY",
        "ERROR",
        "DELETED"
      ]
    },
    "Tag": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    },
    {
      "bearer": []
    }
  ],
  "tags": [
    {
      "description": "Operations on service classes",
      "name": "serviceClass"
    },
    {
      "description": "Operations on service instances",
      "name": "serviceInstance"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Service Manager\n",
    "title": "Service Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/serviceclass": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceClass"
        ],
        "summary": "List all existing service classes",
        "operationId": "getServiceClasses",
        "parameters": [
          {
            "type": "string",
            "description": "Broker name",
            "name": "broker",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter on service class tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ServiceClass"
              }
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/serviceclass/{serviceClassName}": {
      "get": {
        "description": "Returns a single service class",
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceClass"
        ],
        "summary": "Find service class by name",
        "operationId": "getServiceClassByName",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ServiceClass"
            }
          },
          "400": {
            "description": "Invalid name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Service class not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of service class to return",
          "name": "serviceClassName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/serviceinstance": {
      "get": {
        "description": "List all service instances",
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceInstance"
        ],
        "summary": "Get all service instances",
        "operationId": "getServiceInstances",
        "parameters": [
          {
            "type": "string",
            "description": "service class name",
            "name": "serviceclass",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter on service instance tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ServiceInstance"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceInstance"
        ],
        "summary": "Add a new service instance",
        "operationId": "addServiceInstance",
        "parameters": [
          {
            "description": "Service instance object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ServiceInstance"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/ServiceInstance"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/serviceinstance/{serviceInstanceName}": {
      "get": {
        "description": "Returns a single service instance",
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceInstance"
        ],
        "summary": "Find service instance by name",
        "operationId": "getServiceInstanceByName",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ServiceInstance"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Service instance not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceInstance"
        ],
        "summary": "Deletes a service instance",
        "operationId": "deleteServiceInstanceByName",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ServiceInstance"
            }
          },
          "400": {
            "description": "Invalid name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Service instance not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of service instance to return",
          "name": "serviceInstanceName",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ServiceBinding": {
      "type": "object",
      "properties": {
        "bindingSecret": {
          "type": "string"
        },
        "createdTime": {
          "type": "integer"
        },
        "parameters": {
          "type": "object"
        },
        "reason": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "secretParameters": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "$ref": "#/definitions/Status"
        }
      }
    },
    "ServiceClass": {
      "type": "object",
      "required": [
        "name",
        "broker"
      ],
      "properties": {
        "bindable": {
          "type": "boolean"
        },
        "broker": {
          "type": "string"
        },
        "createdTime": {
          "type": "integer"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "kind": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "plans": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ServicePlan"
          }
        },
        "reason": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "$ref": "#/definitions/Status"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "ServiceInstance": {
      "type": "object",
      "required": [
        "name",
        "serviceClass",
        "servicePlan"
      ],
      "properties": {
        "binding": {
          "$ref": "#/definitions/ServiceBinding"
        },
        "createdTime": {
          "type": "integer"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "kind": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "parameters": {
          "type": "object"
        },
        "reason": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "secretParameters": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "serviceClass": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "servicePlan": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "status": {
          "$ref": "#/definitions/Status"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "ServicePlan": {
      "type": "object",
      "properties": {
        "bindable": {
          "type": "boolean"
        },
        "description": {
          "type": "string"
        },
        "free": {
          "type": "boolean"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "kind": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "readOnly": true
        },
        "metadata": {
          "type": "object"
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "schema": {
          "$ref": "#/definitions/ServicePlanSchema"
        }
      }
    },
    "ServicePlanSchema": {
      "type": "object",
      "properties": {
        "bind": {
          "type": "object"
        },
        "create": {
          "type": "object"
        },
        "update": {
          "type": "object"
        }
      }
    },
    "Status": {
      "type": "string",
      "enum": [
        "INITIALIZED",
        "CREATING",
        "READY",
        "ERROR",
        "DELETED"
      ]
    },
    "Tag": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    },
    {
      "bearer": []
    }
  ],
  "tags": [
    {
      "description": "Operations on service classes",
      "name": "serviceClass"
    },
    {
      "description": "Operations on service instances",
      "name": "serviceInstance"
    }
  ]
}`))
}
