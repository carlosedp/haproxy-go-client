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

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetServicesHaproxyRuntimeACLFileEntriesReader is a Reader for the GetServicesHaproxyRuntimeACLFileEntries structure.
type GetServicesHaproxyRuntimeACLFileEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesHaproxyRuntimeACLFileEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesHaproxyRuntimeACLFileEntriesOK creates a GetServicesHaproxyRuntimeACLFileEntriesOK with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesOK() *GetServicesHaproxyRuntimeACLFileEntriesOK {
	return &GetServicesHaproxyRuntimeACLFileEntriesOK{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServicesHaproxyRuntimeACLFileEntriesOK struct {
	Payload models.ACLFilesEntries
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries][%d] getServicesHaproxyRuntimeAclFileEntriesOK  %+v", 200, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesOK) GetPayload() models.ACLFilesEntries {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesHaproxyRuntimeACLFileEntriesBadRequest creates a GetServicesHaproxyRuntimeACLFileEntriesBadRequest with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesBadRequest() *GetServicesHaproxyRuntimeACLFileEntriesBadRequest {
	return &GetServicesHaproxyRuntimeACLFileEntriesBadRequest{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetServicesHaproxyRuntimeACLFileEntriesBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries][%d] getServicesHaproxyRuntimeAclFileEntriesBadRequest  %+v", 400, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServicesHaproxyRuntimeACLFileEntriesNotFound creates a GetServicesHaproxyRuntimeACLFileEntriesNotFound with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesNotFound() *GetServicesHaproxyRuntimeACLFileEntriesNotFound {
	return &GetServicesHaproxyRuntimeACLFileEntriesNotFound{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetServicesHaproxyRuntimeACLFileEntriesNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries][%d] getServicesHaproxyRuntimeAclFileEntriesNotFound  %+v", 404, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServicesHaproxyRuntimeACLFileEntriesDefault creates a GetServicesHaproxyRuntimeACLFileEntriesDefault with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesDefault(code int) *GetServicesHaproxyRuntimeACLFileEntriesDefault {
	return &GetServicesHaproxyRuntimeACLFileEntriesDefault{
		_statusCode: code,
	}
}

/* GetServicesHaproxyRuntimeACLFileEntriesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServicesHaproxyRuntimeACLFileEntriesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get services haproxy runtime ACL file entries default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries][%d] GetServicesHaproxyRuntimeACLFileEntries default  %+v", o._statusCode, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
