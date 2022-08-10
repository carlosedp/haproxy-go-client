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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateRingReader is a Reader for the CreateRing structure.
type CreateRingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateRingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateRingConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateRingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRingCreated creates a CreateRingCreated with default headers values
func NewCreateRingCreated() *CreateRingCreated {
	return &CreateRingCreated{}
}

/* CreateRingCreated describes a response with status code 201, with default header values.

Ring created
*/
type CreateRingCreated struct {
	Payload *models.Ring
}

func (o *CreateRingCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/rings][%d] createRingCreated  %+v", 201, o.Payload)
}
func (o *CreateRingCreated) GetPayload() *models.Ring {
	return o.Payload
}

func (o *CreateRingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ring)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRingAccepted creates a CreateRingAccepted with default headers values
func NewCreateRingAccepted() *CreateRingAccepted {
	return &CreateRingAccepted{}
}

/* CreateRingAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateRingAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Ring
}

func (o *CreateRingAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/rings][%d] createRingAccepted  %+v", 202, o.Payload)
}
func (o *CreateRingAccepted) GetPayload() *models.Ring {
	return o.Payload
}

func (o *CreateRingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Ring)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRingBadRequest creates a CreateRingBadRequest with default headers values
func NewCreateRingBadRequest() *CreateRingBadRequest {
	return &CreateRingBadRequest{}
}

/* CreateRingBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateRingBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateRingBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/rings][%d] createRingBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRingBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRingConflict creates a CreateRingConflict with default headers values
func NewCreateRingConflict() *CreateRingConflict {
	return &CreateRingConflict{}
}

/* CreateRingConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateRingConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateRingConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/rings][%d] createRingConflict  %+v", 409, o.Payload)
}
func (o *CreateRingConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRingConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRingDefault creates a CreateRingDefault with default headers values
func NewCreateRingDefault(code int) *CreateRingDefault {
	return &CreateRingDefault{
		_statusCode: code,
	}
}

/* CreateRingDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateRingDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create ring default response
func (o *CreateRingDefault) Code() int {
	return o._statusCode
}

func (o *CreateRingDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/rings][%d] createRing default  %+v", o._statusCode, o.Payload)
}
func (o *CreateRingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
