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

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// DeleteBackendSwitchingRuleReader is a Reader for the DeleteBackendSwitchingRule structure.
type DeleteBackendSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBackendSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteBackendSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteBackendSwitchingRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteBackendSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBackendSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBackendSwitchingRuleAccepted creates a DeleteBackendSwitchingRuleAccepted with default headers values
func NewDeleteBackendSwitchingRuleAccepted() *DeleteBackendSwitchingRuleAccepted {
	return &DeleteBackendSwitchingRuleAccepted{}
}

/* DeleteBackendSwitchingRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteBackendSwitchingRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteBackendSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backend_switching_rules/{index}][%d] deleteBackendSwitchingRuleAccepted ", 202)
}

func (o *DeleteBackendSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteBackendSwitchingRuleNoContent creates a DeleteBackendSwitchingRuleNoContent with default headers values
func NewDeleteBackendSwitchingRuleNoContent() *DeleteBackendSwitchingRuleNoContent {
	return &DeleteBackendSwitchingRuleNoContent{}
}

/* DeleteBackendSwitchingRuleNoContent describes a response with status code 204, with default header values.

Backend Switching Rule deleted
*/
type DeleteBackendSwitchingRuleNoContent struct {
}

func (o *DeleteBackendSwitchingRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backend_switching_rules/{index}][%d] deleteBackendSwitchingRuleNoContent ", 204)
}

func (o *DeleteBackendSwitchingRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBackendSwitchingRuleNotFound creates a DeleteBackendSwitchingRuleNotFound with default headers values
func NewDeleteBackendSwitchingRuleNotFound() *DeleteBackendSwitchingRuleNotFound {
	return &DeleteBackendSwitchingRuleNotFound{}
}

/* DeleteBackendSwitchingRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteBackendSwitchingRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteBackendSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backend_switching_rules/{index}][%d] deleteBackendSwitchingRuleNotFound  %+v", 404, o.Payload)
}
func (o *DeleteBackendSwitchingRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBackendSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteBackendSwitchingRuleDefault creates a DeleteBackendSwitchingRuleDefault with default headers values
func NewDeleteBackendSwitchingRuleDefault(code int) *DeleteBackendSwitchingRuleDefault {
	return &DeleteBackendSwitchingRuleDefault{
		_statusCode: code,
	}
}

/* DeleteBackendSwitchingRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteBackendSwitchingRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete backend switching rule default response
func (o *DeleteBackendSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBackendSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backend_switching_rules/{index}][%d] deleteBackendSwitchingRule default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteBackendSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBackendSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
