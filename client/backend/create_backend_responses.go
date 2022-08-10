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

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateBackendReader is a Reader for the CreateBackend structure.
type CreateBackendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBackendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateBackendCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateBackendAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBackendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateBackendConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateBackendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBackendCreated creates a CreateBackendCreated with default headers values
func NewCreateBackendCreated() *CreateBackendCreated {
	return &CreateBackendCreated{}
}

/* CreateBackendCreated describes a response with status code 201, with default header values.

Backend created
*/
type CreateBackendCreated struct {
	Payload *models.Backend
}

func (o *CreateBackendCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backends][%d] createBackendCreated  %+v", 201, o.Payload)
}
func (o *CreateBackendCreated) GetPayload() *models.Backend {
	return o.Payload
}

func (o *CreateBackendCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Backend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBackendAccepted creates a CreateBackendAccepted with default headers values
func NewCreateBackendAccepted() *CreateBackendAccepted {
	return &CreateBackendAccepted{}
}

/* CreateBackendAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateBackendAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Backend
}

func (o *CreateBackendAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backends][%d] createBackendAccepted  %+v", 202, o.Payload)
}
func (o *CreateBackendAccepted) GetPayload() *models.Backend {
	return o.Payload
}

func (o *CreateBackendAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Backend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBackendBadRequest creates a CreateBackendBadRequest with default headers values
func NewCreateBackendBadRequest() *CreateBackendBadRequest {
	return &CreateBackendBadRequest{}
}

/* CreateBackendBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateBackendBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateBackendBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backends][%d] createBackendBadRequest  %+v", 400, o.Payload)
}
func (o *CreateBackendBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBackendConflict creates a CreateBackendConflict with default headers values
func NewCreateBackendConflict() *CreateBackendConflict {
	return &CreateBackendConflict{}
}

/* CreateBackendConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateBackendConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateBackendConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backends][%d] createBackendConflict  %+v", 409, o.Payload)
}
func (o *CreateBackendConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBackendDefault creates a CreateBackendDefault with default headers values
func NewCreateBackendDefault(code int) *CreateBackendDefault {
	return &CreateBackendDefault{
		_statusCode: code,
	}
}

/* CreateBackendDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateBackendDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create backend default response
func (o *CreateBackendDefault) Code() int {
	return o._statusCode
}

func (o *CreateBackendDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backends][%d] createBackend default  %+v", o._statusCode, o.Payload)
}
func (o *CreateBackendDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
