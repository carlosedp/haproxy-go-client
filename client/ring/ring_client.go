// Code generated by go-swagger; DO NOT EDIT.

// /*
// MIT License
//
// Copyright (c) 2022 Carlos Eduardo de Paula
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
// */
//
//

package ring

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ring API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ring API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRing(params *CreateRingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRingCreated, *CreateRingAccepted, error)

	DeleteRing(params *DeleteRingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRingAccepted, *DeleteRingNoContent, error)

	GetRing(params *GetRingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRingOK, error)

	GetRings(params *GetRingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRingsOK, error)

	ReplaceRing(params *ReplaceRingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceRingOK, *ReplaceRingAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRing adds a ring

  Adds a new ring to the configuration file.
*/
func (a *Client) CreateRing(params *CreateRingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRingCreated, *CreateRingAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRing",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/rings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateRingCreated:
		return value, nil, nil
	case *CreateRingAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRingDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteRing deletes a ring

  Deletes a ring from the configuration by it's name.
*/
func (a *Client) DeleteRing(params *DeleteRingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRingAccepted, *DeleteRingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRing",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/rings/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteRingAccepted:
		return value, nil, nil
	case *DeleteRingNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRingDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetRing returns a ring

  Returns one ring configuration by it's name.
*/
func (a *Client) GetRing(params *GetRingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRing",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/rings/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetRings returns an array of rings

  Returns an array of all configured rings.
*/
func (a *Client) GetRings(params *GetRingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRings",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/rings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceRing replaces a ring

  Replaces a ring configuration by it's name.
*/
func (a *Client) ReplaceRing(params *ReplaceRingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceRingOK, *ReplaceRingAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceRingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceRing",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/rings/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceRingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceRingOK:
		return value, nil, nil
	case *ReplaceRingAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceRingDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
