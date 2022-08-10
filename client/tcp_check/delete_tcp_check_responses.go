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

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteTCPCheckReader is a Reader for the DeleteTCPCheck structure.
type DeleteTCPCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTCPCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteTCPCheckAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTCPCheckNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTCPCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTCPCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTCPCheckAccepted creates a DeleteTCPCheckAccepted with default headers values
func NewDeleteTCPCheckAccepted() *DeleteTCPCheckAccepted {
	return &DeleteTCPCheckAccepted{}
}

/* DeleteTCPCheckAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteTCPCheckAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteTCPCheckAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckAccepted ", 202)
}

func (o *DeleteTCPCheckAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteTCPCheckNoContent creates a DeleteTCPCheckNoContent with default headers values
func NewDeleteTCPCheckNoContent() *DeleteTCPCheckNoContent {
	return &DeleteTCPCheckNoContent{}
}

/* DeleteTCPCheckNoContent describes a response with status code 204, with default header values.

TCP check deleted
*/
type DeleteTCPCheckNoContent struct {
}

func (o *DeleteTCPCheckNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckNoContent ", 204)
}

func (o *DeleteTCPCheckNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTCPCheckNotFound creates a DeleteTCPCheckNotFound with default headers values
func NewDeleteTCPCheckNotFound() *DeleteTCPCheckNotFound {
	return &DeleteTCPCheckNotFound{}
}

/* DeleteTCPCheckNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteTCPCheckNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteTCPCheckNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckNotFound  %+v", 404, o.Payload)
}
func (o *DeleteTCPCheckNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTCPCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteTCPCheckDefault creates a DeleteTCPCheckDefault with default headers values
func NewDeleteTCPCheckDefault(code int) *DeleteTCPCheckDefault {
	return &DeleteTCPCheckDefault{
		_statusCode: code,
	}
}

/* DeleteTCPCheckDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteTCPCheckDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete TCP check default response
func (o *DeleteTCPCheckDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTCPCheckDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTCPCheck default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteTCPCheckDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTCPCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
