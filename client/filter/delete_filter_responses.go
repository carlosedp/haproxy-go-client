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

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// DeleteFilterReader is a Reader for the DeleteFilter structure.
type DeleteFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteFilterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteFilterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFilterAccepted creates a DeleteFilterAccepted with default headers values
func NewDeleteFilterAccepted() *DeleteFilterAccepted {
	return &DeleteFilterAccepted{}
}

/* DeleteFilterAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteFilterAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteFilterAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterAccepted ", 202)
}

func (o *DeleteFilterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteFilterNoContent creates a DeleteFilterNoContent with default headers values
func NewDeleteFilterNoContent() *DeleteFilterNoContent {
	return &DeleteFilterNoContent{}
}

/* DeleteFilterNoContent describes a response with status code 204, with default header values.

Filter deleted
*/
type DeleteFilterNoContent struct {
}

func (o *DeleteFilterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterNoContent ", 204)
}

func (o *DeleteFilterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFilterNotFound creates a DeleteFilterNotFound with default headers values
func NewDeleteFilterNotFound() *DeleteFilterNotFound {
	return &DeleteFilterNotFound{}
}

/* DeleteFilterNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteFilterNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteFilterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterNotFound  %+v", 404, o.Payload)
}
func (o *DeleteFilterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteFilterDefault creates a DeleteFilterDefault with default headers values
func NewDeleteFilterDefault(code int) *DeleteFilterDefault {
	return &DeleteFilterDefault{
		_statusCode: code,
	}
}

/* DeleteFilterDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteFilterDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete filter default response
func (o *DeleteFilterDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFilterDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilter default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteFilterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
