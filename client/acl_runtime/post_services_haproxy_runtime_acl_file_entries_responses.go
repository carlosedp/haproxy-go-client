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

// PostServicesHaproxyRuntimeACLFileEntriesReader is a Reader for the PostServicesHaproxyRuntimeACLFileEntries structure.
type PostServicesHaproxyRuntimeACLFileEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostServicesHaproxyRuntimeACLFileEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostServicesHaproxyRuntimeACLFileEntriesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostServicesHaproxyRuntimeACLFileEntriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostServicesHaproxyRuntimeACLFileEntriesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostServicesHaproxyRuntimeACLFileEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostServicesHaproxyRuntimeACLFileEntriesCreated creates a PostServicesHaproxyRuntimeACLFileEntriesCreated with default headers values
func NewPostServicesHaproxyRuntimeACLFileEntriesCreated() *PostServicesHaproxyRuntimeACLFileEntriesCreated {
	return &PostServicesHaproxyRuntimeACLFileEntriesCreated{}
}

/* PostServicesHaproxyRuntimeACLFileEntriesCreated describes a response with status code 201, with default header values.

ACL entry created
*/
type PostServicesHaproxyRuntimeACLFileEntriesCreated struct {
	Payload *models.ACLFileEntry
}

func (o *PostServicesHaproxyRuntimeACLFileEntriesCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/acl_file_entries][%d] postServicesHaproxyRuntimeAclFileEntriesCreated  %+v", 201, o.Payload)
}
func (o *PostServicesHaproxyRuntimeACLFileEntriesCreated) GetPayload() *models.ACLFileEntry {
	return o.Payload
}

func (o *PostServicesHaproxyRuntimeACLFileEntriesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ACLFileEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServicesHaproxyRuntimeACLFileEntriesBadRequest creates a PostServicesHaproxyRuntimeACLFileEntriesBadRequest with default headers values
func NewPostServicesHaproxyRuntimeACLFileEntriesBadRequest() *PostServicesHaproxyRuntimeACLFileEntriesBadRequest {
	return &PostServicesHaproxyRuntimeACLFileEntriesBadRequest{}
}

/* PostServicesHaproxyRuntimeACLFileEntriesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostServicesHaproxyRuntimeACLFileEntriesBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *PostServicesHaproxyRuntimeACLFileEntriesBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/acl_file_entries][%d] postServicesHaproxyRuntimeAclFileEntriesBadRequest  %+v", 400, o.Payload)
}
func (o *PostServicesHaproxyRuntimeACLFileEntriesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServicesHaproxyRuntimeACLFileEntriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostServicesHaproxyRuntimeACLFileEntriesConflict creates a PostServicesHaproxyRuntimeACLFileEntriesConflict with default headers values
func NewPostServicesHaproxyRuntimeACLFileEntriesConflict() *PostServicesHaproxyRuntimeACLFileEntriesConflict {
	return &PostServicesHaproxyRuntimeACLFileEntriesConflict{}
}

/* PostServicesHaproxyRuntimeACLFileEntriesConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type PostServicesHaproxyRuntimeACLFileEntriesConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *PostServicesHaproxyRuntimeACLFileEntriesConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/acl_file_entries][%d] postServicesHaproxyRuntimeAclFileEntriesConflict  %+v", 409, o.Payload)
}
func (o *PostServicesHaproxyRuntimeACLFileEntriesConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServicesHaproxyRuntimeACLFileEntriesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostServicesHaproxyRuntimeACLFileEntriesDefault creates a PostServicesHaproxyRuntimeACLFileEntriesDefault with default headers values
func NewPostServicesHaproxyRuntimeACLFileEntriesDefault(code int) *PostServicesHaproxyRuntimeACLFileEntriesDefault {
	return &PostServicesHaproxyRuntimeACLFileEntriesDefault{
		_statusCode: code,
	}
}

/* PostServicesHaproxyRuntimeACLFileEntriesDefault describes a response with status code -1, with default header values.

General Error
*/
type PostServicesHaproxyRuntimeACLFileEntriesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the post services haproxy runtime ACL file entries default response
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) Code() int {
	return o._statusCode
}

func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/acl_file_entries][%d] PostServicesHaproxyRuntimeACLFileEntries default  %+v", o._statusCode, o.Payload)
}
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
