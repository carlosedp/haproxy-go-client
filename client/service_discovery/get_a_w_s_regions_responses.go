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

package service_discovery

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

// GetAWSRegionsReader is a Reader for the GetAWSRegions structure.
type GetAWSRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAWSRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAWSRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAWSRegionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAWSRegionsOK creates a GetAWSRegionsOK with default headers values
func NewGetAWSRegionsOK() *GetAWSRegionsOK {
	return &GetAWSRegionsOK{}
}

/* GetAWSRegionsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetAWSRegionsOK struct {
	Payload *GetAWSRegionsOKBody
}

func (o *GetAWSRegionsOK) Error() string {
	return fmt.Sprintf("[GET /service_discovery/aws][%d] getAWSRegionsOK  %+v", 200, o.Payload)
}
func (o *GetAWSRegionsOK) GetPayload() *GetAWSRegionsOKBody {
	return o.Payload
}

func (o *GetAWSRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAWSRegionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSRegionsDefault creates a GetAWSRegionsDefault with default headers values
func NewGetAWSRegionsDefault(code int) *GetAWSRegionsDefault {
	return &GetAWSRegionsDefault{
		_statusCode: code,
	}
}

/* GetAWSRegionsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetAWSRegionsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get a w s regions default response
func (o *GetAWSRegionsDefault) Code() int {
	return o._statusCode
}

func (o *GetAWSRegionsDefault) Error() string {
	return fmt.Sprintf("[GET /service_discovery/aws][%d] getAWSRegions default  %+v", o._statusCode, o.Payload)
}
func (o *GetAWSRegionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSRegionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetAWSRegionsOKBody get a w s regions o k body
swagger:model GetAWSRegionsOKBody
*/
type GetAWSRegionsOKBody struct {

	// data
	// Required: true
	Data models.AwsRegions `json:"data"`
}

// Validate validates this get a w s regions o k body
func (o *GetAWSRegionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAWSRegionsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getAWSRegionsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getAWSRegionsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getAWSRegionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get a w s regions o k body based on the context it is used
func (o *GetAWSRegionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAWSRegionsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getAWSRegionsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getAWSRegionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAWSRegionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAWSRegionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAWSRegionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
