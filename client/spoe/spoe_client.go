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

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new spoe API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for spoe API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSpoe(params *CreateSpoeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeCreated, error)

	CreateSpoeAgent(params *CreateSpoeAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeAgentCreated, error)

	CreateSpoeGroup(params *CreateSpoeGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeGroupCreated, error)

	CreateSpoeMessage(params *CreateSpoeMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeMessageCreated, error)

	CreateSpoeScope(params *CreateSpoeScopeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeScopeCreated, error)

	DeleteSpoeAgent(params *DeleteSpoeAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeAgentNoContent, error)

	DeleteSpoeFile(params *DeleteSpoeFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeFileNoContent, error)

	DeleteSpoeGroup(params *DeleteSpoeGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeGroupNoContent, error)

	DeleteSpoeMessage(params *DeleteSpoeMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeMessageNoContent, error)

	DeleteSpoeScope(params *DeleteSpoeScopeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeScopeNoContent, error)

	GetAllSpoeFiles(params *GetAllSpoeFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllSpoeFilesOK, error)

	GetOneSpoeFile(params *GetOneSpoeFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOneSpoeFileOK, error)

	GetSpoeAgent(params *GetSpoeAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeAgentOK, error)

	GetSpoeAgents(params *GetSpoeAgentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeAgentsOK, error)

	GetSpoeConfigurationVersion(params *GetSpoeConfigurationVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeConfigurationVersionOK, error)

	GetSpoeGroup(params *GetSpoeGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeGroupOK, error)

	GetSpoeGroups(params *GetSpoeGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeGroupsOK, error)

	GetSpoeMessage(params *GetSpoeMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeMessageOK, error)

	GetSpoeMessages(params *GetSpoeMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeMessagesOK, error)

	GetSpoeScope(params *GetSpoeScopeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeScopeOK, error)

	GetSpoeScopes(params *GetSpoeScopesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeScopesOK, error)

	ReplaceSpoeAgent(params *ReplaceSpoeAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceSpoeAgentOK, error)

	ReplaceSpoeGroup(params *ReplaceSpoeGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceSpoeGroupOK, error)

	ReplaceSpoeMessage(params *ReplaceSpoeMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceSpoeMessageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateSpoe creates s p o e file with its entries

  Creates SPOE file with its entries.
*/
func (a *Client) CreateSpoe(params *CreateSpoeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSpoeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSpoe",
		Method:             "POST",
		PathPattern:        "/services/haproxy/spoe/spoe_files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSpoeReader{formats: a.formats},
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
	success, ok := result.(*CreateSpoeCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSpoeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateSpoeAgent adds a new spoe agent to scope

  Adds a new spoe agent to the spoe scope.
*/
func (a *Client) CreateSpoeAgent(params *CreateSpoeAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeAgentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSpoeAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSpoeAgent",
		Method:             "POST",
		PathPattern:        "/services/haproxy/spoe/spoe_agents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSpoeAgentReader{formats: a.formats},
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
	success, ok := result.(*CreateSpoeAgentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSpoeAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateSpoeGroup adds a new s p o e groups

  Adds a new SPOE groups to the SPOE scope.
*/
func (a *Client) CreateSpoeGroup(params *CreateSpoeGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSpoeGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSpoeGroup",
		Method:             "POST",
		PathPattern:        "/services/haproxy/spoe/spoe_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSpoeGroupReader{formats: a.formats},
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
	success, ok := result.(*CreateSpoeGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSpoeGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateSpoeMessage adds a new spoe message to scope

  Adds a new spoe message to the spoe scope.
*/
func (a *Client) CreateSpoeMessage(params *CreateSpoeMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeMessageCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSpoeMessageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSpoeMessage",
		Method:             "POST",
		PathPattern:        "/services/haproxy/spoe/spoe_messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSpoeMessageReader{formats: a.formats},
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
	success, ok := result.(*CreateSpoeMessageCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSpoeMessageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateSpoeScope adds a new spoe scope

  Adds a new spoe scope.
*/
func (a *Client) CreateSpoeScope(params *CreateSpoeScopeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSpoeScopeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSpoeScopeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSpoeScope",
		Method:             "POST",
		PathPattern:        "/services/haproxy/spoe/spoe_scopes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSpoeScopeReader{formats: a.formats},
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
	success, ok := result.(*CreateSpoeScopeCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSpoeScopeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSpoeAgent deletes a s p o e agent

  Deletes a SPOE agent from the configuration in one SPOE scope.
*/
func (a *Client) DeleteSpoeAgent(params *DeleteSpoeAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeAgentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSpoeAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSpoeAgent",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/spoe/spoe_agents/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSpoeAgentReader{formats: a.formats},
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
	success, ok := result.(*DeleteSpoeAgentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSpoeAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSpoeFile deletes s p o e file

  Deletes SPOE file.
*/
func (a *Client) DeleteSpoeFile(params *DeleteSpoeFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeFileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSpoeFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSpoeFile",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/spoe/spoe_files/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSpoeFileReader{formats: a.formats},
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
	success, ok := result.(*DeleteSpoeFileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSpoeFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSpoeGroup deletes a s p o e groups

  Deletes a SPOE groups from the one SPOE scope.
*/
func (a *Client) DeleteSpoeGroup(params *DeleteSpoeGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeGroupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSpoeGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSpoeGroup",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/spoe/spoe_groups/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSpoeGroupReader{formats: a.formats},
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
	success, ok := result.(*DeleteSpoeGroupNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSpoeGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSpoeMessage deletes a spoe message

  Deletes a spoe message from the SPOE scope.
*/
func (a *Client) DeleteSpoeMessage(params *DeleteSpoeMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeMessageNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSpoeMessageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSpoeMessage",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/spoe/spoe_messages/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSpoeMessageReader{formats: a.formats},
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
	success, ok := result.(*DeleteSpoeMessageNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSpoeMessageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSpoeScope deletes a s p o e scope

  Deletes a SPOE scope from the configuration file.
*/
func (a *Client) DeleteSpoeScope(params *DeleteSpoeScopeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSpoeScopeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSpoeScopeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSpoeScope",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/spoe/spoe_scopes/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSpoeScopeReader{formats: a.formats},
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
	success, ok := result.(*DeleteSpoeScopeNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSpoeScopeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAllSpoeFiles returns all available s p o e files

  Returns all available SPOE files.
*/
func (a *Client) GetAllSpoeFiles(params *GetAllSpoeFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllSpoeFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllSpoeFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllSpoeFiles",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllSpoeFilesReader{formats: a.formats},
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
	success, ok := result.(*GetAllSpoeFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAllSpoeFilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetOneSpoeFile returns one s p o e file

  Returns one SPOE file.
*/
func (a *Client) GetOneSpoeFile(params *GetOneSpoeFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOneSpoeFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOneSpoeFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOneSpoeFile",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_files/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOneSpoeFileReader{formats: a.formats},
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
	success, ok := result.(*GetOneSpoeFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOneSpoeFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeAgent returns a spoe agent

  Returns one spoe agent configuration in one SPOE scope.
*/
func (a *Client) GetSpoeAgent(params *GetSpoeAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeAgent",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_agents/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeAgentReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeAgents returns an array of spoe agents in one scope

  Returns an array of all configured spoe agents in one scope.
*/
func (a *Client) GetSpoeAgents(params *GetSpoeAgentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeAgentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeAgents",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_agents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeAgentsReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeAgentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeConfigurationVersion returns a s p o e configuration version

  Returns SPOE configuration version.
*/
func (a *Client) GetSpoeConfigurationVersion(params *GetSpoeConfigurationVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeConfigurationVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeConfigurationVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeConfigurationVersion",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeConfigurationVersionReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeConfigurationVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeConfigurationVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeGroup returns a s p o e groups

  Returns one SPOE groups configuration in one SPOE scope.
*/
func (a *Client) GetSpoeGroup(params *GetSpoeGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeGroup",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_groups/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeGroupReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeGroups returns an array of s p o e groups

  Returns an array of all configured SPOE groups in one SPOE scope.
*/
func (a *Client) GetSpoeGroups(params *GetSpoeGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeGroups",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeGroupsReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeMessage returns a spoe message

  Returns one spoe message configuration in SPOE scope.
*/
func (a *Client) GetSpoeMessage(params *GetSpoeMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeMessageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeMessage",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_messages/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeMessageReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeMessageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeMessageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeMessages returns an array of spoe messages in one scope

  Returns an array of all configured spoe messages in one scope.
*/
func (a *Client) GetSpoeMessages(params *GetSpoeMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeMessagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeMessages",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeMessagesReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeMessagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeMessagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeScope returns one s p o e scope

  Returns one SPOE scope in one SPOE file.
*/
func (a *Client) GetSpoeScope(params *GetSpoeScopeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeScopeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeScopeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeScope",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_scopes/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeScopeReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeScopeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeScopeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSpoeScopes returns an array of spoe scopes

  Returns an array of all configured spoe scopes.
*/
func (a *Client) GetSpoeScopes(params *GetSpoeScopesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeScopesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeScopesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeScopes",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe/spoe_scopes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpoeScopesReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeScopesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeScopesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceSpoeAgent replaces a s p o e agent

  Replaces a SPOE agent configuration in one SPOE scope.
*/
func (a *Client) ReplaceSpoeAgent(params *ReplaceSpoeAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceSpoeAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceSpoeAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceSpoeAgent",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/spoe/spoe_agents/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceSpoeAgentReader{formats: a.formats},
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
	success, ok := result.(*ReplaceSpoeAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceSpoeAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceSpoeGroup replaces a s p o e groups

  Replaces a SPOE groups configuration in one SPOE scope.
*/
func (a *Client) ReplaceSpoeGroup(params *ReplaceSpoeGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceSpoeGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceSpoeGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceSpoeGroup",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/spoe/spoe_groups/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceSpoeGroupReader{formats: a.formats},
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
	success, ok := result.(*ReplaceSpoeGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceSpoeGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceSpoeMessage replaces a spoe message

  Replaces a spoe message configuration in one SPOE scope.
*/
func (a *Client) ReplaceSpoeMessage(params *ReplaceSpoeMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceSpoeMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceSpoeMessageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceSpoeMessage",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/spoe/spoe_messages/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceSpoeMessageReader{formats: a.formats},
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
	success, ok := result.(*ReplaceSpoeMessageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceSpoeMessageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}