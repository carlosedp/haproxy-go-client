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

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteServerReader is a Reader for the DeleteServer structure.
type DeleteServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteServerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServerAccepted creates a DeleteServerAccepted with default headers values
func NewDeleteServerAccepted() *DeleteServerAccepted {
	return &DeleteServerAccepted{}
}

/* DeleteServerAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteServerAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteServerAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/servers/{name}][%d] deleteServerAccepted ", 202)
}

func (o *DeleteServerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteServerNoContent creates a DeleteServerNoContent with default headers values
func NewDeleteServerNoContent() *DeleteServerNoContent {
	return &DeleteServerNoContent{}
}

/* DeleteServerNoContent describes a response with status code 204, with default header values.

Server deleted
*/
type DeleteServerNoContent struct {
}

func (o *DeleteServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/servers/{name}][%d] deleteServerNoContent ", 204)
}

func (o *DeleteServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServerNotFound creates a DeleteServerNotFound with default headers values
func NewDeleteServerNotFound() *DeleteServerNotFound {
	return &DeleteServerNotFound{}
}

/* DeleteServerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteServerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteServerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/servers/{name}][%d] deleteServerNotFound  %+v", 404, o.Payload)
}
func (o *DeleteServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteServerDefault creates a DeleteServerDefault with default headers values
func NewDeleteServerDefault(code int) *DeleteServerDefault {
	return &DeleteServerDefault{
		_statusCode: code,
	}
}

/* DeleteServerDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete server default response
func (o *DeleteServerDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServerDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/servers/{name}][%d] deleteServer default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
