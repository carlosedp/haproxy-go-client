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

package global

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// ReplaceGlobalReader is a Reader for the ReplaceGlobal structure.
type ReplaceGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceGlobalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceGlobalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceGlobalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceGlobalOK creates a ReplaceGlobalOK with default headers values
func NewReplaceGlobalOK() *ReplaceGlobalOK {
	return &ReplaceGlobalOK{}
}

/* ReplaceGlobalOK describes a response with status code 200, with default header values.

Global replaced
*/
type ReplaceGlobalOK struct {
	Payload *models.Global
}

func (o *ReplaceGlobalOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalOK  %+v", 200, o.Payload)
}
func (o *ReplaceGlobalOK) GetPayload() *models.Global {
	return o.Payload
}

func (o *ReplaceGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Global)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGlobalAccepted creates a ReplaceGlobalAccepted with default headers values
func NewReplaceGlobalAccepted() *ReplaceGlobalAccepted {
	return &ReplaceGlobalAccepted{}
}

/* ReplaceGlobalAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceGlobalAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Global
}

func (o *ReplaceGlobalAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceGlobalAccepted) GetPayload() *models.Global {
	return o.Payload
}

func (o *ReplaceGlobalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Global)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGlobalBadRequest creates a ReplaceGlobalBadRequest with default headers values
func NewReplaceGlobalBadRequest() *ReplaceGlobalBadRequest {
	return &ReplaceGlobalBadRequest{}
}

/* ReplaceGlobalBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceGlobalBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceGlobalBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceGlobalBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceGlobalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceGlobalDefault creates a ReplaceGlobalDefault with default headers values
func NewReplaceGlobalDefault(code int) *ReplaceGlobalDefault {
	return &ReplaceGlobalDefault{
		_statusCode: code,
	}
}

/* ReplaceGlobalDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceGlobalDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace global default response
func (o *ReplaceGlobalDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceGlobalDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobal default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceGlobalDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceGlobalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
