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

package tcp_check

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
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v4/models"
)

// GetTCPChecksReader is a Reader for the GetTCPChecks structure.
type GetTCPChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTCPChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTCPChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTCPChecksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTCPChecksOK creates a GetTCPChecksOK with default headers values
func NewGetTCPChecksOK() *GetTCPChecksOK {
	return &GetTCPChecksOK{}
}

/* GetTCPChecksOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetTCPChecksOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetTCPChecksOKBody
}

func (o *GetTCPChecksOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_checks][%d] getTcpChecksOK  %+v", 200, o.Payload)
}
func (o *GetTCPChecksOK) GetPayload() *GetTCPChecksOKBody {
	return o.Payload
}

func (o *GetTCPChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetTCPChecksOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPChecksDefault creates a GetTCPChecksDefault with default headers values
func NewGetTCPChecksDefault(code int) *GetTCPChecksDefault {
	return &GetTCPChecksDefault{
		_statusCode: code,
	}
}

/* GetTCPChecksDefault describes a response with status code -1, with default header values.

General Error
*/
type GetTCPChecksDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get TCP checks default response
func (o *GetTCPChecksDefault) Code() int {
	return o._statusCode
}

func (o *GetTCPChecksDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_checks][%d] getTCPChecks default  %+v", o._statusCode, o.Payload)
}
func (o *GetTCPChecksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPChecksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetTCPChecksOKBody get TCP checks o k body
swagger:model GetTCPChecksOKBody
*/
type GetTCPChecksOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.TCPChecks `json:"data"`
}

// Validate validates this get TCP checks o k body
func (o *GetTCPChecksOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPChecksOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getTcpChecksOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getTcpChecksOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getTcpChecksOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get TCP checks o k body based on the context it is used
func (o *GetTCPChecksOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPChecksOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getTcpChecksOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getTcpChecksOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTCPChecksOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTCPChecksOKBody) UnmarshalBinary(b []byte) error {
	var res GetTCPChecksOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
