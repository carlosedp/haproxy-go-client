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

package stick_rule

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

// GetStickRuleReader is a Reader for the GetStickRule structure.
type GetStickRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStickRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStickRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStickRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStickRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStickRuleOK creates a GetStickRuleOK with default headers values
func NewGetStickRuleOK() *GetStickRuleOK {
	return &GetStickRuleOK{}
}

/* GetStickRuleOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetStickRuleOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetStickRuleOKBody
}

func (o *GetStickRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/stick_rules/{index}][%d] getStickRuleOK  %+v", 200, o.Payload)
}
func (o *GetStickRuleOK) GetPayload() *GetStickRuleOKBody {
	return o.Payload
}

func (o *GetStickRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetStickRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickRuleNotFound creates a GetStickRuleNotFound with default headers values
func NewGetStickRuleNotFound() *GetStickRuleNotFound {
	return &GetStickRuleNotFound{}
}

/* GetStickRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetStickRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetStickRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/stick_rules/{index}][%d] getStickRuleNotFound  %+v", 404, o.Payload)
}
func (o *GetStickRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStickRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStickRuleDefault creates a GetStickRuleDefault with default headers values
func NewGetStickRuleDefault(code int) *GetStickRuleDefault {
	return &GetStickRuleDefault{
		_statusCode: code,
	}
}

/* GetStickRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type GetStickRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get stick rule default response
func (o *GetStickRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetStickRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/stick_rules/{index}][%d] getStickRule default  %+v", o._statusCode, o.Payload)
}
func (o *GetStickRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStickRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetStickRuleOKBody get stick rule o k body
swagger:model GetStickRuleOKBody
*/
type GetStickRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.StickRule `json:"data,omitempty"`
}

// Validate validates this get stick rule o k body
func (o *GetStickRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStickRuleOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStickRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStickRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get stick rule o k body based on the context it is used
func (o *GetStickRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStickRuleOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStickRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStickRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStickRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStickRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetStickRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
