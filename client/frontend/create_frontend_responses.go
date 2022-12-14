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

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateFrontendReader is a Reader for the CreateFrontend structure.
type CreateFrontendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFrontendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFrontendCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateFrontendAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFrontendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateFrontendConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateFrontendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateFrontendCreated creates a CreateFrontendCreated with default headers values
func NewCreateFrontendCreated() *CreateFrontendCreated {
	return &CreateFrontendCreated{}
}

/* CreateFrontendCreated describes a response with status code 201, with default header values.

Frontend created
*/
type CreateFrontendCreated struct {
	Payload *models.Frontend
}

func (o *CreateFrontendCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/frontends][%d] createFrontendCreated  %+v", 201, o.Payload)
}
func (o *CreateFrontendCreated) GetPayload() *models.Frontend {
	return o.Payload
}

func (o *CreateFrontendCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Frontend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFrontendAccepted creates a CreateFrontendAccepted with default headers values
func NewCreateFrontendAccepted() *CreateFrontendAccepted {
	return &CreateFrontendAccepted{}
}

/* CreateFrontendAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateFrontendAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Frontend
}

func (o *CreateFrontendAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/frontends][%d] createFrontendAccepted  %+v", 202, o.Payload)
}
func (o *CreateFrontendAccepted) GetPayload() *models.Frontend {
	return o.Payload
}

func (o *CreateFrontendAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Frontend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFrontendBadRequest creates a CreateFrontendBadRequest with default headers values
func NewCreateFrontendBadRequest() *CreateFrontendBadRequest {
	return &CreateFrontendBadRequest{}
}

/* CreateFrontendBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateFrontendBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateFrontendBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/frontends][%d] createFrontendBadRequest  %+v", 400, o.Payload)
}
func (o *CreateFrontendBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFrontendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFrontendConflict creates a CreateFrontendConflict with default headers values
func NewCreateFrontendConflict() *CreateFrontendConflict {
	return &CreateFrontendConflict{}
}

/* CreateFrontendConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateFrontendConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateFrontendConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/frontends][%d] createFrontendConflict  %+v", 409, o.Payload)
}
func (o *CreateFrontendConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFrontendConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFrontendDefault creates a CreateFrontendDefault with default headers values
func NewCreateFrontendDefault(code int) *CreateFrontendDefault {
	return &CreateFrontendDefault{
		_statusCode: code,
	}
}

/* CreateFrontendDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateFrontendDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create frontend default response
func (o *CreateFrontendDefault) Code() int {
	return o._statusCode
}

func (o *CreateFrontendDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/frontends][%d] createFrontend default  %+v", o._statusCode, o.Payload)
}
func (o *CreateFrontendDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFrontendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
