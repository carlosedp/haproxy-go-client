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

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetTransactionReader is a Reader for the GetTransaction structure.
type GetTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTransactionOK creates a GetTransactionOK with default headers values
func NewGetTransactionOK() *GetTransactionOK {
	return &GetTransactionOK{}
}

/* GetTransactionOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetTransactionOK struct {
	Payload *models.Transaction
}

func (o *GetTransactionOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/transactions/{id}][%d] getTransactionOK  %+v", 200, o.Payload)
}
func (o *GetTransactionOK) GetPayload() *models.Transaction {
	return o.Payload
}

func (o *GetTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Transaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionNotFound creates a GetTransactionNotFound with default headers values
func NewGetTransactionNotFound() *GetTransactionNotFound {
	return &GetTransactionNotFound{}
}

/* GetTransactionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetTransactionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetTransactionNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/transactions/{id}][%d] getTransactionNotFound  %+v", 404, o.Payload)
}
func (o *GetTransactionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetTransactionDefault creates a GetTransactionDefault with default headers values
func NewGetTransactionDefault(code int) *GetTransactionDefault {
	return &GetTransactionDefault{
		_statusCode: code,
	}
}

/* GetTransactionDefault describes a response with status code -1, with default header values.

General Error
*/
type GetTransactionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get transaction default response
func (o *GetTransactionDefault) Code() int {
	return o._statusCode
}

func (o *GetTransactionDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/transactions/{id}][%d] getTransaction default  %+v", o._statusCode, o.Payload)
}
func (o *GetTransactionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
