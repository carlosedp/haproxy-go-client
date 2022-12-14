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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetAllStorageGeneralFilesReader is a Reader for the GetAllStorageGeneralFiles structure.
type GetAllStorageGeneralFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllStorageGeneralFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllStorageGeneralFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllStorageGeneralFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAllStorageGeneralFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllStorageGeneralFilesOK creates a GetAllStorageGeneralFilesOK with default headers values
func NewGetAllStorageGeneralFilesOK() *GetAllStorageGeneralFilesOK {
	return &GetAllStorageGeneralFilesOK{}
}

/* GetAllStorageGeneralFilesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetAllStorageGeneralFilesOK struct {
	Payload models.GeneralFiles
}

func (o *GetAllStorageGeneralFilesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFilesOK  %+v", 200, o.Payload)
}
func (o *GetAllStorageGeneralFilesOK) GetPayload() models.GeneralFiles {
	return o.Payload
}

func (o *GetAllStorageGeneralFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllStorageGeneralFilesNotFound creates a GetAllStorageGeneralFilesNotFound with default headers values
func NewGetAllStorageGeneralFilesNotFound() *GetAllStorageGeneralFilesNotFound {
	return &GetAllStorageGeneralFilesNotFound{}
}

/* GetAllStorageGeneralFilesNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetAllStorageGeneralFilesNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetAllStorageGeneralFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFilesNotFound  %+v", 404, o.Payload)
}
func (o *GetAllStorageGeneralFilesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllStorageGeneralFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAllStorageGeneralFilesDefault creates a GetAllStorageGeneralFilesDefault with default headers values
func NewGetAllStorageGeneralFilesDefault(code int) *GetAllStorageGeneralFilesDefault {
	return &GetAllStorageGeneralFilesDefault{
		_statusCode: code,
	}
}

/* GetAllStorageGeneralFilesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetAllStorageGeneralFilesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get all storage general files default response
func (o *GetAllStorageGeneralFilesDefault) Code() int {
	return o._statusCode
}

func (o *GetAllStorageGeneralFilesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFiles default  %+v", o._statusCode, o.Payload)
}
func (o *GetAllStorageGeneralFilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllStorageGeneralFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
