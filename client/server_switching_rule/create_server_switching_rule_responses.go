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

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateServerSwitchingRuleReader is a Reader for the CreateServerSwitchingRule structure.
type CreateServerSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServerSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServerSwitchingRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateServerSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServerSwitchingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateServerSwitchingRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateServerSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServerSwitchingRuleCreated creates a CreateServerSwitchingRuleCreated with default headers values
func NewCreateServerSwitchingRuleCreated() *CreateServerSwitchingRuleCreated {
	return &CreateServerSwitchingRuleCreated{}
}

/* CreateServerSwitchingRuleCreated describes a response with status code 201, with default header values.

Server Switching Rule created
*/
type CreateServerSwitchingRuleCreated struct {
	Payload *models.ServerSwitchingRule
}

func (o *CreateServerSwitchingRuleCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_switching_rules][%d] createServerSwitchingRuleCreated  %+v", 201, o.Payload)
}
func (o *CreateServerSwitchingRuleCreated) GetPayload() *models.ServerSwitchingRule {
	return o.Payload
}

func (o *CreateServerSwitchingRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerSwitchingRuleAccepted creates a CreateServerSwitchingRuleAccepted with default headers values
func NewCreateServerSwitchingRuleAccepted() *CreateServerSwitchingRuleAccepted {
	return &CreateServerSwitchingRuleAccepted{}
}

/* CreateServerSwitchingRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateServerSwitchingRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.ServerSwitchingRule
}

func (o *CreateServerSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_switching_rules][%d] createServerSwitchingRuleAccepted  %+v", 202, o.Payload)
}
func (o *CreateServerSwitchingRuleAccepted) GetPayload() *models.ServerSwitchingRule {
	return o.Payload
}

func (o *CreateServerSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.ServerSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerSwitchingRuleBadRequest creates a CreateServerSwitchingRuleBadRequest with default headers values
func NewCreateServerSwitchingRuleBadRequest() *CreateServerSwitchingRuleBadRequest {
	return &CreateServerSwitchingRuleBadRequest{}
}

/* CreateServerSwitchingRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateServerSwitchingRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateServerSwitchingRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_switching_rules][%d] createServerSwitchingRuleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateServerSwitchingRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerSwitchingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerSwitchingRuleConflict creates a CreateServerSwitchingRuleConflict with default headers values
func NewCreateServerSwitchingRuleConflict() *CreateServerSwitchingRuleConflict {
	return &CreateServerSwitchingRuleConflict{}
}

/* CreateServerSwitchingRuleConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateServerSwitchingRuleConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateServerSwitchingRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_switching_rules][%d] createServerSwitchingRuleConflict  %+v", 409, o.Payload)
}
func (o *CreateServerSwitchingRuleConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerSwitchingRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerSwitchingRuleDefault creates a CreateServerSwitchingRuleDefault with default headers values
func NewCreateServerSwitchingRuleDefault(code int) *CreateServerSwitchingRuleDefault {
	return &CreateServerSwitchingRuleDefault{
		_statusCode: code,
	}
}

/* CreateServerSwitchingRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateServerSwitchingRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create server switching rule default response
func (o *CreateServerSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateServerSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_switching_rules][%d] createServerSwitchingRule default  %+v", o._statusCode, o.Payload)
}
func (o *CreateServerSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
