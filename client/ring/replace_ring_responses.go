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

// ReplaceRingReader is a Reader for the ReplaceRing structure.
type ReplaceRingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceRingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceRingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceRingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceRingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceRingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceRingOK creates a ReplaceRingOK with default headers values
func NewReplaceRingOK() *ReplaceRingOK {
	return &ReplaceRingOK{}
}

/* ReplaceRingOK describes a response with status code 200, with default header values.

Ring replaced
*/
type ReplaceRingOK struct {
	Payload *models.Ring
}

func (o *ReplaceRingOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/rings/{name}][%d] replaceRingOK  %+v", 200, o.Payload)
}
func (o *ReplaceRingOK) GetPayload() *models.Ring {
	return o.Payload
}

func (o *ReplaceRingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ring)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRingAccepted creates a ReplaceRingAccepted with default headers values
func NewReplaceRingAccepted() *ReplaceRingAccepted {
	return &ReplaceRingAccepted{}
}

/* ReplaceRingAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceRingAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Ring
}

func (o *ReplaceRingAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/rings/{name}][%d] replaceRingAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceRingAccepted) GetPayload() *models.Ring {
	return o.Payload
}

func (o *ReplaceRingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceRingBadRequest creates a ReplaceRingBadRequest with default headers values
func NewReplaceRingBadRequest() *ReplaceRingBadRequest {
	return &ReplaceRingBadRequest{}
}

/* ReplaceRingBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceRingBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceRingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/rings/{name}][%d] replaceRingBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceRingBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceRingNotFound creates a ReplaceRingNotFound with default headers values
func NewReplaceRingNotFound() *ReplaceRingNotFound {
	return &ReplaceRingNotFound{}
}

/* ReplaceRingNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceRingNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceRingNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/rings/{name}][%d] replaceRingNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceRingNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceRingDefault creates a ReplaceRingDefault with default headers values
func NewReplaceRingDefault(code int) *ReplaceRingDefault {
	return &ReplaceRingDefault{
		_statusCode: code,
	}
}

/* ReplaceRingDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceRingDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace ring default response
func (o *ReplaceRingDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceRingDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/rings/{name}][%d] replaceRing default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceRingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
