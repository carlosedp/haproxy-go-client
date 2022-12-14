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

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceGroupReader is a Reader for the ReplaceGroup structure.
type ReplaceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceGroupAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceGroupOK creates a ReplaceGroupOK with default headers values
func NewReplaceGroupOK() *ReplaceGroupOK {
	return &ReplaceGroupOK{}
}

/* ReplaceGroupOK describes a response with status code 200, with default header values.

Group replaced
*/
type ReplaceGroupOK struct {
	Payload *models.Group
}

func (o *ReplaceGroupOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/groups/{name}][%d] replaceGroupOK  %+v", 200, o.Payload)
}
func (o *ReplaceGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *ReplaceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGroupAccepted creates a ReplaceGroupAccepted with default headers values
func NewReplaceGroupAccepted() *ReplaceGroupAccepted {
	return &ReplaceGroupAccepted{}
}

/* ReplaceGroupAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceGroupAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Group
}

func (o *ReplaceGroupAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/groups/{name}][%d] replaceGroupAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceGroupAccepted) GetPayload() *models.Group {
	return o.Payload
}

func (o *ReplaceGroupAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGroupBadRequest creates a ReplaceGroupBadRequest with default headers values
func NewReplaceGroupBadRequest() *ReplaceGroupBadRequest {
	return &ReplaceGroupBadRequest{}
}

/* ReplaceGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceGroupBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/groups/{name}][%d] replaceGroupBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceGroupNotFound creates a ReplaceGroupNotFound with default headers values
func NewReplaceGroupNotFound() *ReplaceGroupNotFound {
	return &ReplaceGroupNotFound{}
}

/* ReplaceGroupNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceGroupNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/groups/{name}][%d] replaceGroupNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceGroupDefault creates a ReplaceGroupDefault with default headers values
func NewReplaceGroupDefault(code int) *ReplaceGroupDefault {
	return &ReplaceGroupDefault{
		_statusCode: code,
	}
}

/* ReplaceGroupDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceGroupDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace group default response
func (o *ReplaceGroupDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceGroupDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/groups/{name}][%d] replaceGroup default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
