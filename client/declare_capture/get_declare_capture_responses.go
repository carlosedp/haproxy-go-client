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

package declare_capture

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

	"github.com/haproxytech/client-native/v4/models"
)

// GetDeclareCaptureReader is a Reader for the GetDeclareCapture structure.
type GetDeclareCaptureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeclareCaptureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeclareCaptureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeclareCaptureNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeclareCaptureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeclareCaptureOK creates a GetDeclareCaptureOK with default headers values
func NewGetDeclareCaptureOK() *GetDeclareCaptureOK {
	return &GetDeclareCaptureOK{}
}

/* GetDeclareCaptureOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeclareCaptureOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetDeclareCaptureOKBody
}

func (o *GetDeclareCaptureOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/captures/{index}][%d] getDeclareCaptureOK  %+v", 200, o.Payload)
}
func (o *GetDeclareCaptureOK) GetPayload() *GetDeclareCaptureOKBody {
	return o.Payload
}

func (o *GetDeclareCaptureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetDeclareCaptureOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeclareCaptureNotFound creates a GetDeclareCaptureNotFound with default headers values
func NewGetDeclareCaptureNotFound() *GetDeclareCaptureNotFound {
	return &GetDeclareCaptureNotFound{}
}

/* GetDeclareCaptureNotFound describes a response with status code 404, with default header values.

The specified resource already exists
*/
type GetDeclareCaptureNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetDeclareCaptureNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/captures/{index}][%d] getDeclareCaptureNotFound  %+v", 404, o.Payload)
}
func (o *GetDeclareCaptureNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeclareCaptureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeclareCaptureDefault creates a GetDeclareCaptureDefault with default headers values
func NewGetDeclareCaptureDefault(code int) *GetDeclareCaptureDefault {
	return &GetDeclareCaptureDefault{
		_statusCode: code,
	}
}

/* GetDeclareCaptureDefault describes a response with status code -1, with default header values.

General Error
*/
type GetDeclareCaptureDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get declare capture default response
func (o *GetDeclareCaptureDefault) Code() int {
	return o._statusCode
}

func (o *GetDeclareCaptureDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/captures/{index}][%d] getDeclareCapture default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeclareCaptureDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeclareCaptureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetDeclareCaptureOKBody get declare capture o k body
swagger:model GetDeclareCaptureOKBody
*/
type GetDeclareCaptureOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Capture `json:"data,omitempty"`
}

// Validate validates this get declare capture o k body
func (o *GetDeclareCaptureOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeclareCaptureOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDeclareCaptureOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getDeclareCaptureOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get declare capture o k body based on the context it is used
func (o *GetDeclareCaptureOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeclareCaptureOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDeclareCaptureOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getDeclareCaptureOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeclareCaptureOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeclareCaptureOKBody) UnmarshalBinary(b []byte) error {
	var res GetDeclareCaptureOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
