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

package http_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateHTTPCheckReader is a Reader for the CreateHTTPCheck structure.
type CreateHTTPCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHTTPCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateHTTPCheckCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateHTTPCheckAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHTTPCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateHTTPCheckConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateHTTPCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHTTPCheckCreated creates a CreateHTTPCheckCreated with default headers values
func NewCreateHTTPCheckCreated() *CreateHTTPCheckCreated {
	return &CreateHTTPCheckCreated{}
}

/* CreateHTTPCheckCreated describes a response with status code 201, with default header values.

HTTP check created
*/
type CreateHTTPCheckCreated struct {
	Payload *models.HTTPCheck
}

func (o *CreateHTTPCheckCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_checks][%d] createHttpCheckCreated  %+v", 201, o.Payload)
}
func (o *CreateHTTPCheckCreated) GetPayload() *models.HTTPCheck {
	return o.Payload
}

func (o *CreateHTTPCheckCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPCheck)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPCheckAccepted creates a CreateHTTPCheckAccepted with default headers values
func NewCreateHTTPCheckAccepted() *CreateHTTPCheckAccepted {
	return &CreateHTTPCheckAccepted{}
}

/* CreateHTTPCheckAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateHTTPCheckAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.HTTPCheck
}

func (o *CreateHTTPCheckAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_checks][%d] createHttpCheckAccepted  %+v", 202, o.Payload)
}
func (o *CreateHTTPCheckAccepted) GetPayload() *models.HTTPCheck {
	return o.Payload
}

func (o *CreateHTTPCheckAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.HTTPCheck)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPCheckBadRequest creates a CreateHTTPCheckBadRequest with default headers values
func NewCreateHTTPCheckBadRequest() *CreateHTTPCheckBadRequest {
	return &CreateHTTPCheckBadRequest{}
}

/* CreateHTTPCheckBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateHTTPCheckBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateHTTPCheckBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_checks][%d] createHttpCheckBadRequest  %+v", 400, o.Payload)
}
func (o *CreateHTTPCheckBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHTTPCheckConflict creates a CreateHTTPCheckConflict with default headers values
func NewCreateHTTPCheckConflict() *CreateHTTPCheckConflict {
	return &CreateHTTPCheckConflict{}
}

/* CreateHTTPCheckConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateHTTPCheckConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateHTTPCheckConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_checks][%d] createHttpCheckConflict  %+v", 409, o.Payload)
}
func (o *CreateHTTPCheckConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPCheckConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHTTPCheckDefault creates a CreateHTTPCheckDefault with default headers values
func NewCreateHTTPCheckDefault(code int) *CreateHTTPCheckDefault {
	return &CreateHTTPCheckDefault{
		_statusCode: code,
	}
}

/* CreateHTTPCheckDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateHTTPCheckDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create HTTP check default response
func (o *CreateHTTPCheckDefault) Code() int {
	return o._statusCode
}

func (o *CreateHTTPCheckDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_checks][%d] createHTTPCheck default  %+v", o._statusCode, o.Payload)
}
func (o *CreateHTTPCheckDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
