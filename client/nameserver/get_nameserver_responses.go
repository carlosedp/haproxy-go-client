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

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"models"
)

// GetNameserverReader is a Reader for the GetNameserver structure.
type GetNameserverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNameserverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNameserverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNameserverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNameserverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNameserverOK creates a GetNameserverOK with default headers values
func NewGetNameserverOK() *GetNameserverOK {
	return &GetNameserverOK{}
}

/* GetNameserverOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNameserverOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetNameserverOKBody
}

func (o *GetNameserverOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/nameservers/{name}][%d] getNameserverOK  %+v", 200, o.Payload)
}
func (o *GetNameserverOK) GetPayload() *GetNameserverOKBody {
	return o.Payload
}

func (o *GetNameserverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetNameserverOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNameserverNotFound creates a GetNameserverNotFound with default headers values
func NewGetNameserverNotFound() *GetNameserverNotFound {
	return &GetNameserverNotFound{}
}

/* GetNameserverNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetNameserverNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetNameserverNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/nameservers/{name}][%d] getNameserverNotFound  %+v", 404, o.Payload)
}
func (o *GetNameserverNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNameserverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetNameserverDefault creates a GetNameserverDefault with default headers values
func NewGetNameserverDefault(code int) *GetNameserverDefault {
	return &GetNameserverDefault{
		_statusCode: code,
	}
}

/* GetNameserverDefault describes a response with status code -1, with default header values.

General Error
*/
type GetNameserverDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get nameserver default response
func (o *GetNameserverDefault) Code() int {
	return o._statusCode
}

func (o *GetNameserverDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/nameservers/{name}][%d] getNameserver default  %+v", o._statusCode, o.Payload)
}
func (o *GetNameserverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNameserverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetNameserverOKBody get nameserver o k body
swagger:model GetNameserverOKBody
*/
type GetNameserverOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Nameserver `json:"data,omitempty"`
}

// Validate validates this get nameserver o k body
func (o *GetNameserverOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNameserverOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNameserverOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNameserverOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get nameserver o k body based on the context it is used
func (o *GetNameserverOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNameserverOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNameserverOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNameserverOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNameserverOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNameserverOKBody) UnmarshalBinary(b []byte) error {
	var res GetNameserverOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
