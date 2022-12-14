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

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceDefaultsReader is a Reader for the ReplaceDefaults structure.
type ReplaceDefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceDefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceDefaultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceDefaultsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceDefaultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceDefaultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceDefaultsOK creates a ReplaceDefaultsOK with default headers values
func NewReplaceDefaultsOK() *ReplaceDefaultsOK {
	return &ReplaceDefaultsOK{}
}

/* ReplaceDefaultsOK describes a response with status code 200, with default header values.

Defaults replaced
*/
type ReplaceDefaultsOK struct {
	Payload *models.Defaults
}

func (o *ReplaceDefaultsOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/defaults][%d] replaceDefaultsOK  %+v", 200, o.Payload)
}
func (o *ReplaceDefaultsOK) GetPayload() *models.Defaults {
	return o.Payload
}

func (o *ReplaceDefaultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Defaults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDefaultsAccepted creates a ReplaceDefaultsAccepted with default headers values
func NewReplaceDefaultsAccepted() *ReplaceDefaultsAccepted {
	return &ReplaceDefaultsAccepted{}
}

/* ReplaceDefaultsAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceDefaultsAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Defaults
}

func (o *ReplaceDefaultsAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/defaults][%d] replaceDefaultsAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceDefaultsAccepted) GetPayload() *models.Defaults {
	return o.Payload
}

func (o *ReplaceDefaultsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Defaults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDefaultsBadRequest creates a ReplaceDefaultsBadRequest with default headers values
func NewReplaceDefaultsBadRequest() *ReplaceDefaultsBadRequest {
	return &ReplaceDefaultsBadRequest{}
}

/* ReplaceDefaultsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceDefaultsBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceDefaultsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/defaults][%d] replaceDefaultsBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceDefaultsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceDefaultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceDefaultsDefault creates a ReplaceDefaultsDefault with default headers values
func NewReplaceDefaultsDefault(code int) *ReplaceDefaultsDefault {
	return &ReplaceDefaultsDefault{
		_statusCode: code,
	}
}

/* ReplaceDefaultsDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceDefaultsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace defaults default response
func (o *ReplaceDefaultsDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceDefaultsDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/defaults][%d] replaceDefaults default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceDefaultsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceDefaultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
