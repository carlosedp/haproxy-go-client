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

package bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteBindReader is a Reader for the DeleteBind structure.
type DeleteBindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteBindAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteBindNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteBindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBindDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBindAccepted creates a DeleteBindAccepted with default headers values
func NewDeleteBindAccepted() *DeleteBindAccepted {
	return &DeleteBindAccepted{}
}

/* DeleteBindAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteBindAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteBindAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/binds/{name}][%d] deleteBindAccepted ", 202)
}

func (o *DeleteBindAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteBindNoContent creates a DeleteBindNoContent with default headers values
func NewDeleteBindNoContent() *DeleteBindNoContent {
	return &DeleteBindNoContent{}
}

/* DeleteBindNoContent describes a response with status code 204, with default header values.

Bind deleted
*/
type DeleteBindNoContent struct {
}

func (o *DeleteBindNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/binds/{name}][%d] deleteBindNoContent ", 204)
}

func (o *DeleteBindNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBindNotFound creates a DeleteBindNotFound with default headers values
func NewDeleteBindNotFound() *DeleteBindNotFound {
	return &DeleteBindNotFound{}
}

/* DeleteBindNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteBindNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteBindNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/binds/{name}][%d] deleteBindNotFound  %+v", 404, o.Payload)
}
func (o *DeleteBindNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteBindDefault creates a DeleteBindDefault with default headers values
func NewDeleteBindDefault(code int) *DeleteBindDefault {
	return &DeleteBindDefault{
		_statusCode: code,
	}
}

/* DeleteBindDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteBindDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete bind default response
func (o *DeleteBindDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBindDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/binds/{name}][%d] deleteBind default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteBindDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBindDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
