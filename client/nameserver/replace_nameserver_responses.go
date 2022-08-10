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

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceNameserverReader is a Reader for the ReplaceNameserver structure.
type ReplaceNameserverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceNameserverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceNameserverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceNameserverAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceNameserverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceNameserverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceNameserverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceNameserverOK creates a ReplaceNameserverOK with default headers values
func NewReplaceNameserverOK() *ReplaceNameserverOK {
	return &ReplaceNameserverOK{}
}

/* ReplaceNameserverOK describes a response with status code 200, with default header values.

Nameserver replaced
*/
type ReplaceNameserverOK struct {
	Payload *models.Nameserver
}

func (o *ReplaceNameserverOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserverOK  %+v", 200, o.Payload)
}
func (o *ReplaceNameserverOK) GetPayload() *models.Nameserver {
	return o.Payload
}

func (o *ReplaceNameserverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Nameserver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNameserverAccepted creates a ReplaceNameserverAccepted with default headers values
func NewReplaceNameserverAccepted() *ReplaceNameserverAccepted {
	return &ReplaceNameserverAccepted{}
}

/* ReplaceNameserverAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceNameserverAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Nameserver
}

func (o *ReplaceNameserverAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserverAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceNameserverAccepted) GetPayload() *models.Nameserver {
	return o.Payload
}

func (o *ReplaceNameserverAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Nameserver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNameserverBadRequest creates a ReplaceNameserverBadRequest with default headers values
func NewReplaceNameserverBadRequest() *ReplaceNameserverBadRequest {
	return &ReplaceNameserverBadRequest{}
}

/* ReplaceNameserverBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceNameserverBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceNameserverBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserverBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceNameserverBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceNameserverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceNameserverNotFound creates a ReplaceNameserverNotFound with default headers values
func NewReplaceNameserverNotFound() *ReplaceNameserverNotFound {
	return &ReplaceNameserverNotFound{}
}

/* ReplaceNameserverNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceNameserverNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceNameserverNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserverNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceNameserverNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceNameserverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceNameserverDefault creates a ReplaceNameserverDefault with default headers values
func NewReplaceNameserverDefault(code int) *ReplaceNameserverDefault {
	return &ReplaceNameserverDefault{
		_statusCode: code,
	}
}

/* ReplaceNameserverDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceNameserverDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace nameserver default response
func (o *ReplaceNameserverDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceNameserverDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserver default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceNameserverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceNameserverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
