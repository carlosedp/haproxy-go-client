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

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetStatsEndpointsReader is a Reader for the GetStatsEndpoints structure.
type GetStatsEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatsEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatsEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStatsEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStatsEndpointsOK creates a GetStatsEndpointsOK with default headers values
func NewGetStatsEndpointsOK() *GetStatsEndpointsOK {
	return &GetStatsEndpointsOK{}
}

/* GetStatsEndpointsOK describes a response with status code 200, with default header values.

Success
*/
type GetStatsEndpointsOK struct {
	Payload models.Endpoints
}

func (o *GetStatsEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/stats][%d] getStatsEndpointsOK  %+v", 200, o.Payload)
}
func (o *GetStatsEndpointsOK) GetPayload() models.Endpoints {
	return o.Payload
}

func (o *GetStatsEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsEndpointsDefault creates a GetStatsEndpointsDefault with default headers values
func NewGetStatsEndpointsDefault(code int) *GetStatsEndpointsDefault {
	return &GetStatsEndpointsDefault{
		_statusCode: code,
	}
}

/* GetStatsEndpointsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetStatsEndpointsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get stats endpoints default response
func (o *GetStatsEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *GetStatsEndpointsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/stats][%d] getStatsEndpoints default  %+v", o._statusCode, o.Payload)
}
func (o *GetStatsEndpointsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStatsEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
