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

package dgram_bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateDgramBindReader is a Reader for the CreateDgramBind structure.
type CreateDgramBindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDgramBindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDgramBindCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateDgramBindAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDgramBindBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDgramBindConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDgramBindDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDgramBindCreated creates a CreateDgramBindCreated with default headers values
func NewCreateDgramBindCreated() *CreateDgramBindCreated {
	return &CreateDgramBindCreated{}
}

/* CreateDgramBindCreated describes a response with status code 201, with default header values.

Bind created
*/
type CreateDgramBindCreated struct {
	Payload *models.DgramBind
}

func (o *CreateDgramBindCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/dgram_binds][%d] createDgramBindCreated  %+v", 201, o.Payload)
}
func (o *CreateDgramBindCreated) GetPayload() *models.DgramBind {
	return o.Payload
}

func (o *CreateDgramBindCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DgramBind)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDgramBindAccepted creates a CreateDgramBindAccepted with default headers values
func NewCreateDgramBindAccepted() *CreateDgramBindAccepted {
	return &CreateDgramBindAccepted{}
}

/* CreateDgramBindAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateDgramBindAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.DgramBind
}

func (o *CreateDgramBindAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/dgram_binds][%d] createDgramBindAccepted  %+v", 202, o.Payload)
}
func (o *CreateDgramBindAccepted) GetPayload() *models.DgramBind {
	return o.Payload
}

func (o *CreateDgramBindAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.DgramBind)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDgramBindBadRequest creates a CreateDgramBindBadRequest with default headers values
func NewCreateDgramBindBadRequest() *CreateDgramBindBadRequest {
	return &CreateDgramBindBadRequest{}
}

/* CreateDgramBindBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateDgramBindBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateDgramBindBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/dgram_binds][%d] createDgramBindBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDgramBindBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDgramBindBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDgramBindConflict creates a CreateDgramBindConflict with default headers values
func NewCreateDgramBindConflict() *CreateDgramBindConflict {
	return &CreateDgramBindConflict{}
}

/* CreateDgramBindConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateDgramBindConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateDgramBindConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/dgram_binds][%d] createDgramBindConflict  %+v", 409, o.Payload)
}
func (o *CreateDgramBindConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDgramBindConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDgramBindDefault creates a CreateDgramBindDefault with default headers values
func NewCreateDgramBindDefault(code int) *CreateDgramBindDefault {
	return &CreateDgramBindDefault{
		_statusCode: code,
	}
}

/* CreateDgramBindDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateDgramBindDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create dgram bind default response
func (o *CreateDgramBindDefault) Code() int {
	return o._statusCode
}

func (o *CreateDgramBindDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/dgram_binds][%d] createDgramBind default  %+v", o._statusCode, o.Payload)
}
func (o *CreateDgramBindDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDgramBindDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
