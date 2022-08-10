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

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateLogTargetReader is a Reader for the CreateLogTarget structure.
type CreateLogTargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLogTargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateLogTargetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateLogTargetAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLogTargetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateLogTargetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateLogTargetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateLogTargetCreated creates a CreateLogTargetCreated with default headers values
func NewCreateLogTargetCreated() *CreateLogTargetCreated {
	return &CreateLogTargetCreated{}
}

/* CreateLogTargetCreated describes a response with status code 201, with default header values.

Log Target created
*/
type CreateLogTargetCreated struct {
	Payload *models.LogTarget
}

func (o *CreateLogTargetCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/log_targets][%d] createLogTargetCreated  %+v", 201, o.Payload)
}
func (o *CreateLogTargetCreated) GetPayload() *models.LogTarget {
	return o.Payload
}

func (o *CreateLogTargetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogTargetAccepted creates a CreateLogTargetAccepted with default headers values
func NewCreateLogTargetAccepted() *CreateLogTargetAccepted {
	return &CreateLogTargetAccepted{}
}

/* CreateLogTargetAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateLogTargetAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.LogTarget
}

func (o *CreateLogTargetAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/log_targets][%d] createLogTargetAccepted  %+v", 202, o.Payload)
}
func (o *CreateLogTargetAccepted) GetPayload() *models.LogTarget {
	return o.Payload
}

func (o *CreateLogTargetAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.LogTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogTargetBadRequest creates a CreateLogTargetBadRequest with default headers values
func NewCreateLogTargetBadRequest() *CreateLogTargetBadRequest {
	return &CreateLogTargetBadRequest{}
}

/* CreateLogTargetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateLogTargetBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateLogTargetBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/log_targets][%d] createLogTargetBadRequest  %+v", 400, o.Payload)
}
func (o *CreateLogTargetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateLogTargetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateLogTargetConflict creates a CreateLogTargetConflict with default headers values
func NewCreateLogTargetConflict() *CreateLogTargetConflict {
	return &CreateLogTargetConflict{}
}

/* CreateLogTargetConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateLogTargetConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateLogTargetConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/log_targets][%d] createLogTargetConflict  %+v", 409, o.Payload)
}
func (o *CreateLogTargetConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateLogTargetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateLogTargetDefault creates a CreateLogTargetDefault with default headers values
func NewCreateLogTargetDefault(code int) *CreateLogTargetDefault {
	return &CreateLogTargetDefault{
		_statusCode: code,
	}
}

/* CreateLogTargetDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateLogTargetDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create log target default response
func (o *CreateLogTargetDefault) Code() int {
	return o._statusCode
}

func (o *CreateLogTargetDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/log_targets][%d] createLogTarget default  %+v", o._statusCode, o.Payload)
}
func (o *CreateLogTargetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateLogTargetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
