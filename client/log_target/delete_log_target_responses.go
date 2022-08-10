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

	"models"
)

// DeleteLogTargetReader is a Reader for the DeleteLogTarget structure.
type DeleteLogTargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLogTargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteLogTargetAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteLogTargetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteLogTargetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteLogTargetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLogTargetAccepted creates a DeleteLogTargetAccepted with default headers values
func NewDeleteLogTargetAccepted() *DeleteLogTargetAccepted {
	return &DeleteLogTargetAccepted{}
}

/* DeleteLogTargetAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteLogTargetAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteLogTargetAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_targets/{index}][%d] deleteLogTargetAccepted ", 202)
}

func (o *DeleteLogTargetAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteLogTargetNoContent creates a DeleteLogTargetNoContent with default headers values
func NewDeleteLogTargetNoContent() *DeleteLogTargetNoContent {
	return &DeleteLogTargetNoContent{}
}

/* DeleteLogTargetNoContent describes a response with status code 204, with default header values.

Log Target deleted
*/
type DeleteLogTargetNoContent struct {
}

func (o *DeleteLogTargetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_targets/{index}][%d] deleteLogTargetNoContent ", 204)
}

func (o *DeleteLogTargetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLogTargetNotFound creates a DeleteLogTargetNotFound with default headers values
func NewDeleteLogTargetNotFound() *DeleteLogTargetNotFound {
	return &DeleteLogTargetNotFound{}
}

/* DeleteLogTargetNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteLogTargetNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteLogTargetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_targets/{index}][%d] deleteLogTargetNotFound  %+v", 404, o.Payload)
}
func (o *DeleteLogTargetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLogTargetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLogTargetDefault creates a DeleteLogTargetDefault with default headers values
func NewDeleteLogTargetDefault(code int) *DeleteLogTargetDefault {
	return &DeleteLogTargetDefault{
		_statusCode: code,
	}
}

/* DeleteLogTargetDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteLogTargetDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete log target default response
func (o *DeleteLogTargetDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLogTargetDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_targets/{index}][%d] deleteLogTarget default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteLogTargetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLogTargetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
