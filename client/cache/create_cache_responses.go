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

package cache

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateCacheReader is a Reader for the CreateCache structure.
type CreateCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCacheCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateCacheAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCacheBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateCacheConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCacheCreated creates a CreateCacheCreated with default headers values
func NewCreateCacheCreated() *CreateCacheCreated {
	return &CreateCacheCreated{}
}

/* CreateCacheCreated describes a response with status code 201, with default header values.

Cache created
*/
type CreateCacheCreated struct {
	Payload *models.Cache
}

func (o *CreateCacheCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheCreated  %+v", 201, o.Payload)
}
func (o *CreateCacheCreated) GetPayload() *models.Cache {
	return o.Payload
}

func (o *CreateCacheCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cache)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCacheAccepted creates a CreateCacheAccepted with default headers values
func NewCreateCacheAccepted() *CreateCacheAccepted {
	return &CreateCacheAccepted{}
}

/* CreateCacheAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateCacheAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Cache
}

func (o *CreateCacheAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheAccepted  %+v", 202, o.Payload)
}
func (o *CreateCacheAccepted) GetPayload() *models.Cache {
	return o.Payload
}

func (o *CreateCacheAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Cache)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCacheBadRequest creates a CreateCacheBadRequest with default headers values
func NewCreateCacheBadRequest() *CreateCacheBadRequest {
	return &CreateCacheBadRequest{}
}

/* CreateCacheBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateCacheBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateCacheBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheBadRequest  %+v", 400, o.Payload)
}
func (o *CreateCacheBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCacheConflict creates a CreateCacheConflict with default headers values
func NewCreateCacheConflict() *CreateCacheConflict {
	return &CreateCacheConflict{}
}

/* CreateCacheConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateCacheConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateCacheConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheConflict  %+v", 409, o.Payload)
}
func (o *CreateCacheConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCacheConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCacheDefault creates a CreateCacheDefault with default headers values
func NewCreateCacheDefault(code int) *CreateCacheDefault {
	return &CreateCacheDefault{
		_statusCode: code,
	}
}

/* CreateCacheDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateCacheDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create cache default response
func (o *CreateCacheDefault) Code() int {
	return o._statusCode
}

func (o *CreateCacheDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCache default  %+v", o._statusCode, o.Payload)
}
func (o *CreateCacheDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
