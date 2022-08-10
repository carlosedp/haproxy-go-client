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

package stick_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetStickTableEntriesReader is a Reader for the GetStickTableEntries structure.
type GetStickTableEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStickTableEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStickTableEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStickTableEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStickTableEntriesOK creates a GetStickTableEntriesOK with default headers values
func NewGetStickTableEntriesOK() *GetStickTableEntriesOK {
	return &GetStickTableEntriesOK{}
}

/* GetStickTableEntriesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetStickTableEntriesOK struct {
	Payload models.StickTableEntries
}

func (o *GetStickTableEntriesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/stick_table_entries][%d] getStickTableEntriesOK  %+v", 200, o.Payload)
}
func (o *GetStickTableEntriesOK) GetPayload() models.StickTableEntries {
	return o.Payload
}

func (o *GetStickTableEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickTableEntriesDefault creates a GetStickTableEntriesDefault with default headers values
func NewGetStickTableEntriesDefault(code int) *GetStickTableEntriesDefault {
	return &GetStickTableEntriesDefault{
		_statusCode: code,
	}
}

/* GetStickTableEntriesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetStickTableEntriesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get stick table entries default response
func (o *GetStickTableEntriesDefault) Code() int {
	return o._statusCode
}

func (o *GetStickTableEntriesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/stick_table_entries][%d] getStickTableEntries default  %+v", o._statusCode, o.Payload)
}
func (o *GetStickTableEntriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStickTableEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
