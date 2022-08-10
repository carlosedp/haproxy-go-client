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

	"github.com/haproxytech/client-native/v4/models"
)

// AddPayloadRuntimeACLReader is a Reader for the AddPayloadRuntimeACL structure.
type AddPayloadRuntimeACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPayloadRuntimeACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddPayloadRuntimeACLCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddPayloadRuntimeACLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddPayloadRuntimeACLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPayloadRuntimeACLCreated creates a AddPayloadRuntimeACLCreated with default headers values
func NewAddPayloadRuntimeACLCreated() *AddPayloadRuntimeACLCreated {
	return &AddPayloadRuntimeACLCreated{}
}

/* AddPayloadRuntimeACLCreated describes a response with status code 201, with default header values.

ACL payload added
*/
type AddPayloadRuntimeACLCreated struct {
	Payload models.ACLFilesEntries
}

func (o *AddPayloadRuntimeACLCreated) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/acl_file_entries][%d] addPayloadRuntimeAclCreated  %+v", 201, o.Payload)
}
func (o *AddPayloadRuntimeACLCreated) GetPayload() models.ACLFilesEntries {
	return o.Payload
}

func (o *AddPayloadRuntimeACLCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPayloadRuntimeACLBadRequest creates a AddPayloadRuntimeACLBadRequest with default headers values
func NewAddPayloadRuntimeACLBadRequest() *AddPayloadRuntimeACLBadRequest {
	return &AddPayloadRuntimeACLBadRequest{}
}

/* AddPayloadRuntimeACLBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddPayloadRuntimeACLBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *AddPayloadRuntimeACLBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/acl_file_entries][%d] addPayloadRuntimeAclBadRequest  %+v", 400, o.Payload)
}
func (o *AddPayloadRuntimeACLBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddPayloadRuntimeACLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddPayloadRuntimeACLDefault creates a AddPayloadRuntimeACLDefault with default headers values
func NewAddPayloadRuntimeACLDefault(code int) *AddPayloadRuntimeACLDefault {
	return &AddPayloadRuntimeACLDefault{
		_statusCode: code,
	}
}

/* AddPayloadRuntimeACLDefault describes a response with status code -1, with default header values.

General Error
*/
type AddPayloadRuntimeACLDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the add payload runtime ACL default response
func (o *AddPayloadRuntimeACLDefault) Code() int {
	return o._statusCode
}

func (o *AddPayloadRuntimeACLDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/acl_file_entries][%d] addPayloadRuntimeACL default  %+v", o._statusCode, o.Payload)
}
func (o *AddPayloadRuntimeACLDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddPayloadRuntimeACLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
