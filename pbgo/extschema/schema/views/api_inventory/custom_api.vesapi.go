// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package api_inventory

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
)

var (
	_ = fmt.Sprintf("dummy for fmt import use")
)

// Create CustomAPIPrivate GRPC Client satisfying server.CustomClient
type CustomAPIPrivateGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomAPIPrivateClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomAPIPrivateGrpcClient) doRPCGetPathSuggestions(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &GetPathSuggestionsReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.api_inventory.GetPathSuggestionsReq", yamlReq)
	}
	rsp, err := c.grpcClient.GetPathSuggestions(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIPrivateGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}
	if cco.YAMLReq == "" {
		return nil, fmt.Errorf("Error, empty request body")
	}
	ctx = client.AddHdrsToCtx(cco.Headers, ctx)

	rsp, err := rpcFn(ctx, cco.YAMLReq, cco.GrpcCallOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using GRPC")
	}
	if cco.OutCallResponse != nil {
		cco.OutCallResponse.ProtoMsg = rsp
	}
	return rsp, nil
}

func NewCustomAPIPrivateGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomAPIPrivateGrpcClient{
		conn:       cc,
		grpcClient: NewCustomAPIPrivateClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["GetPathSuggestions"] = ccl.doRPCGetPathSuggestions

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPIPrivate REST Client satisfying server.CustomClient
type CustomAPIPrivateRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomAPIPrivateRestClient) doRPCGetPathSuggestions(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &GetPathSuggestionsReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.api_inventory.GetPathSuggestionsReq: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := codec.ToJSON(req, codec.ToWithUseProtoFieldName())
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		var op string
		if hm == "post" {
			op = http.MethodPost
		} else {
			op = http.MethodPut
		}
		newReq, err := http.NewRequest(op, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrapf(err, "Creating new HTTP %s request for custom API", op)
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("match_value", fmt.Sprintf("%v", req.MatchValue))
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	// checking whether the status code is a successful status code (2xx series)
	if rsp.StatusCode < 200 || rsp.StatusCode > 299 {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &ves_io_schema.SuggestValuesResp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.SuggestValuesResp", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIPrivateRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}

	rsp, err := rpcFn(ctx, cco)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using Rest")
	}
	return rsp, nil
}

func NewCustomAPIPrivateRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomAPIPrivateRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["GetPathSuggestions"] = ccl.doRPCGetPathSuggestions

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customAPIPrivateInprocClient

// INPROC Client (satisfying CustomAPIPrivateClient interface)
type customAPIPrivateInprocClient struct {
	CustomAPIPrivateServer
}

func (c *customAPIPrivateInprocClient) GetPathSuggestions(ctx context.Context, in *GetPathSuggestionsReq, opts ...grpc.CallOption) (*ves_io_schema.SuggestValuesResp, error) {
	return c.CustomAPIPrivateServer.GetPathSuggestions(ctx, in)
}

func NewCustomAPIPrivateInprocClient(svc svcfw.Service) CustomAPIPrivateClient {
	return &customAPIPrivateInprocClient{CustomAPIPrivateServer: NewCustomAPIPrivateServer(svc)}
}

// RegisterGwCustomAPIPrivateHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomAPIPrivateHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomAPIPrivateHandlerClient(ctx, mux, NewCustomAPIPrivateInprocClient(s))
}

// Create customAPIPrivateSrv

// SERVER (satisfying CustomAPIPrivateServer interface)
type customAPIPrivateSrv struct {
	svc svcfw.Service
}

func (s *customAPIPrivateSrv) GetPathSuggestions(ctx context.Context, in *GetPathSuggestionsReq) (*ves_io_schema.SuggestValuesResp, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.views.api_inventory.CustomAPIPrivate")
	cah, ok := ah.(CustomAPIPrivateServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIPrivateServer", ah)
	}

	var (
		rsp *ves_io_schema.SuggestValuesResp
		err error
	)

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.views.api_inventory.CustomAPIPrivate.GetPathSuggestions"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.GetPathSuggestions(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}

func NewCustomAPIPrivateServer(svc svcfw.Service) CustomAPIPrivateServer {
	return &customAPIPrivateSrv{svc: svc}
}

var CustomAPIPrivateSwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "api definition object",
        "version": "version not set"
    },
    "schemes": [
        "http",
        "https"
    ],
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "tags": [],
    "paths": {
        "/private/namespaces/{namespace}/api_inventories/{name}/api_endpoint_path/suggestion": {
            "post": {
                "summary": "Get Path Suggestions",
                "description": "Get the suggestions for API Endpoint paths",
                "operationId": "ves.io.schema.views.api_inventory.CustomAPIPrivate.GetPathSuggestions",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/schemaSuggestValuesResp"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "description": "Namespace\n\nx-example: \"default\"\nNamespace for the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"all-api\"\nName of the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_inventoryGetPathSuggestionsReq"
                        }
                    }
                ],
                "tags": [
                    "CustomAPIPrivate"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-api_inventory-customapiprivate-getpathsuggestions"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.api_inventory.CustomAPIPrivate.GetPathSuggestions"
            },
            "x-displayname": "API Inventory Private APIs",
            "x-ves-proto-service": "ves.io.schema.views.api_inventory.CustomAPIPrivate",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        }
    },
    "definitions": {
        "api_inventoryGetPathSuggestionsReq": {
            "type": "object",
            "description": "This is the input message of the 'GetPathSuggestions' RPC",
            "title": "GetPathSuggestionsReq is used to get suggestions of API Endpoint paths",
            "x-displayname": "Get Paths Suggestions Request",
            "x-ves-proto-message": "ves.io.schema.views.api_inventory.GetPathSuggestionsReq",
            "properties": {
                "match_value": {
                    "type": "string",
                    "description": " A substring that must be present in either the value or description of each SuggestedItem in the response.\n\nExample: - \"some-substring\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n",
                    "title": "match_value",
                    "maxLength": 256,
                    "x-displayname": "Match Value",
                    "x-ves-example": "some-substring",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256"
                    }
                },
                "name": {
                    "type": "string",
                    "description": " Name of the object to be configured\n\nExample: - \"all-api\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "all-api"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                }
            }
        },
        "schemaSuggestValuesResp": {
            "type": "object",
            "description": "Response body of SuggestValues request",
            "title": "SuggestValuesResp",
            "x-displayname": "Response for SuggestValues",
            "x-ves-proto-message": "ves.io.schema.SuggestValuesResp",
            "properties": {
                "items": {
                    "type": "array",
                    "description": " List of suggested items.",
                    "title": "item_lists",
                    "items": {
                        "$ref": "#/definitions/schemaSuggestedItem"
                    },
                    "x-displayname": "Suggested Items"
                }
            }
        },
        "schemaSuggestedItem": {
            "type": "object",
            "description": "A tuple with a suggested value and it's description.",
            "title": "SuggestedItem",
            "x-displayname": "Suggested Item",
            "x-ves-proto-message": "ves.io.schema.SuggestedItem",
            "properties": {
                "description": {
                    "type": "string",
                    "description": " Optional description for the suggested value.",
                    "title": "description",
                    "x-displayname": "Description"
                },
                "value": {
                    "type": "string",
                    "description": " Suggested value for the field.",
                    "title": "value",
                    "x-displayname": "Value"
                }
            }
        }
    },
    "x-displayname": "",
    "x-ves-proto-file": "ves.io/schema/views/api_inventory/custom_api.proto"
}`