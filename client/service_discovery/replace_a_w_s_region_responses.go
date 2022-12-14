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

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceAWSRegionReader is a Reader for the ReplaceAWSRegion structure.
type ReplaceAWSRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAWSRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAWSRegionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceAWSRegionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceAWSRegionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceAWSRegionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceAWSRegionOK creates a ReplaceAWSRegionOK with default headers values
func NewReplaceAWSRegionOK() *ReplaceAWSRegionOK {
	return &ReplaceAWSRegionOK{}
}

/* ReplaceAWSRegionOK describes a response with status code 200, with default header values.

Resource updated
*/
type ReplaceAWSRegionOK struct {
	Payload *models.AwsRegion
}

func (o *ReplaceAWSRegionOK) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/aws/{id}][%d] replaceAWSRegionOK  %+v", 200, o.Payload)
}
func (o *ReplaceAWSRegionOK) GetPayload() *models.AwsRegion {
	return o.Payload
}

func (o *ReplaceAWSRegionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAWSRegionBadRequest creates a ReplaceAWSRegionBadRequest with default headers values
func NewReplaceAWSRegionBadRequest() *ReplaceAWSRegionBadRequest {
	return &ReplaceAWSRegionBadRequest{}
}

/* ReplaceAWSRegionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceAWSRegionBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceAWSRegionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/aws/{id}][%d] replaceAWSRegionBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceAWSRegionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceAWSRegionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceAWSRegionNotFound creates a ReplaceAWSRegionNotFound with default headers values
func NewReplaceAWSRegionNotFound() *ReplaceAWSRegionNotFound {
	return &ReplaceAWSRegionNotFound{}
}

/* ReplaceAWSRegionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceAWSRegionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceAWSRegionNotFound) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/aws/{id}][%d] replaceAWSRegionNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceAWSRegionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceAWSRegionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceAWSRegionDefault creates a ReplaceAWSRegionDefault with default headers values
func NewReplaceAWSRegionDefault(code int) *ReplaceAWSRegionDefault {
	return &ReplaceAWSRegionDefault{
		_statusCode: code,
	}
}

/* ReplaceAWSRegionDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceAWSRegionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace a w s region default response
func (o *ReplaceAWSRegionDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceAWSRegionDefault) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/aws/{id}][%d] replaceAWSRegion default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceAWSRegionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceAWSRegionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
