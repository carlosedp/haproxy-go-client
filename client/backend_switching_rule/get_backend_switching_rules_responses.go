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

package backend_switching_rule

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

// GetBackendSwitchingRulesReader is a Reader for the GetBackendSwitchingRules structure.
type GetBackendSwitchingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackendSwitchingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackendSwitchingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackendSwitchingRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackendSwitchingRulesOK creates a GetBackendSwitchingRulesOK with default headers values
func NewGetBackendSwitchingRulesOK() *GetBackendSwitchingRulesOK {
	return &GetBackendSwitchingRulesOK{}
}

/* GetBackendSwitchingRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetBackendSwitchingRulesOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetBackendSwitchingRulesOKBody
}

func (o *GetBackendSwitchingRulesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backend_switching_rules][%d] getBackendSwitchingRulesOK  %+v", 200, o.Payload)
}
func (o *GetBackendSwitchingRulesOK) GetPayload() *GetBackendSwitchingRulesOKBody {
	return o.Payload
}

func (o *GetBackendSwitchingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetBackendSwitchingRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackendSwitchingRulesDefault creates a GetBackendSwitchingRulesDefault with default headers values
func NewGetBackendSwitchingRulesDefault(code int) *GetBackendSwitchingRulesDefault {
	return &GetBackendSwitchingRulesDefault{
		_statusCode: code,
	}
}

/* GetBackendSwitchingRulesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetBackendSwitchingRulesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get backend switching rules default response
func (o *GetBackendSwitchingRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetBackendSwitchingRulesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backend_switching_rules][%d] getBackendSwitchingRules default  %+v", o._statusCode, o.Payload)
}
func (o *GetBackendSwitchingRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackendSwitchingRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetBackendSwitchingRulesOKBody get backend switching rules o k body
swagger:model GetBackendSwitchingRulesOKBody
*/
type GetBackendSwitchingRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.BackendSwitchingRules `json:"data"`
}

// Validate validates this get backend switching rules o k body
func (o *GetBackendSwitchingRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackendSwitchingRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBackendSwitchingRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBackendSwitchingRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getBackendSwitchingRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get backend switching rules o k body based on the context it is used
func (o *GetBackendSwitchingRulesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackendSwitchingRulesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBackendSwitchingRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getBackendSwitchingRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBackendSwitchingRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendSwitchingRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetBackendSwitchingRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
