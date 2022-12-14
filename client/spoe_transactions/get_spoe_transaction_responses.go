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

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetSpoeTransactionReader is a Reader for the GetSpoeTransaction structure.
type GetSpoeTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSpoeTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSpoeTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeTransactionOK creates a GetSpoeTransactionOK with default headers values
func NewGetSpoeTransactionOK() *GetSpoeTransactionOK {
	return &GetSpoeTransactionOK{}
}

/* GetSpoeTransactionOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetSpoeTransactionOK struct {
	Payload *models.SpoeTransaction
}

func (o *GetSpoeTransactionOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe_transactions/{id}][%d] getSpoeTransactionOK  %+v", 200, o.Payload)
}
func (o *GetSpoeTransactionOK) GetPayload() *models.SpoeTransaction {
	return o.Payload
}

func (o *GetSpoeTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeTransactionNotFound creates a GetSpoeTransactionNotFound with default headers values
func NewGetSpoeTransactionNotFound() *GetSpoeTransactionNotFound {
	return &GetSpoeTransactionNotFound{}
}

/* GetSpoeTransactionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetSpoeTransactionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetSpoeTransactionNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe_transactions/{id}][%d] getSpoeTransactionNotFound  %+v", 404, o.Payload)
}
func (o *GetSpoeTransactionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSpoeTransactionDefault creates a GetSpoeTransactionDefault with default headers values
func NewGetSpoeTransactionDefault(code int) *GetSpoeTransactionDefault {
	return &GetSpoeTransactionDefault{
		_statusCode: code,
	}
}

/* GetSpoeTransactionDefault describes a response with status code -1, with default header values.

General Error
*/
type GetSpoeTransactionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get spoe transaction default response
func (o *GetSpoeTransactionDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeTransactionDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe_transactions/{id}][%d] getSpoeTransaction default  %+v", o._statusCode, o.Payload)
}
func (o *GetSpoeTransactionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
