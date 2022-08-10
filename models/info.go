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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Info Information
//
// General API, OS and hardware information
// Example: {"api":{"build_date":"2019-08-21T17:31:56.000Z","version":"v1.2.1 45a3288.dev"},"system":{"cpu_info":{"model":"Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz","num_cpus":4},"hostname":"test","mem_info":{"dataplaneapi_memory":44755536,"free_memory":5790642176,"total_memory":16681517056},"os_string":"Linux 4.15.0-58-generic #64-Ubuntu SMP Tue Aug 6 11:12:41 UTC 2019","time":1566401525,"uptime":87340}}
//
// swagger:model info
type Info struct {

	// api
	API *InfoAPI `json:"api,omitempty"`

	// system
	System *InfoSystem `json:"system,omitempty"`
}

// Validate validates this info
func (m *Info) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Info) validateAPI(formats strfmt.Registry) error {
	if swag.IsZero(m.API) { // not required
		return nil
	}

	if m.API != nil {
		if err := m.API.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api")
			}
			return err
		}
	}

	return nil
}

func (m *Info) validateSystem(formats strfmt.Registry) error {
	if swag.IsZero(m.System) { // not required
		return nil
	}

	if m.System != nil {
		if err := m.System.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this info based on the context it is used
func (m *Info) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Info) contextValidateAPI(ctx context.Context, formats strfmt.Registry) error {

	if m.API != nil {
		if err := m.API.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api")
			}
			return err
		}
	}

	return nil
}

func (m *Info) contextValidateSystem(ctx context.Context, formats strfmt.Registry) error {

	if m.System != nil {
		if err := m.System.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Info) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Info) UnmarshalBinary(b []byte) error {
	var res Info
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoAPI info API
//
// swagger:model InfoAPI
type InfoAPI struct {

	// HAProxy Dataplane API build date
	// Format: date-time
	BuildDate strfmt.DateTime `json:"build_date,omitempty"`

	// HAProxy Dataplane API version string
	Version string `json:"version,omitempty"`
}

// Validate validates this info API
func (m *InfoAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfoAPI) validateBuildDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BuildDate) { // not required
		return nil
	}

	if err := validate.FormatOf("api"+"."+"build_date", "body", "date-time", m.BuildDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this info API based on context it is used
func (m *InfoAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfoAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoAPI) UnmarshalBinary(b []byte) error {
	var res InfoAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoSystem info system
//
// swagger:model InfoSystem
type InfoSystem struct {

	// cpu info
	CPUInfo *InfoSystemCPUInfo `json:"cpu_info,omitempty"`

	// Hostname where the HAProxy is running
	Hostname string `json:"hostname,omitempty"`

	// mem info
	MemInfo *InfoSystemMemInfo `json:"mem_info,omitempty"`

	// OS string
	OsString string `json:"os_string,omitempty"`

	// Current time in milliseconds since Epoch.
	Time int64 `json:"time,omitempty"`

	// System uptime
	Uptime *int64 `json:"uptime,omitempty"`
}

// Validate validates this info system
func (m *InfoSystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfoSystem) validateCPUInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUInfo) { // not required
		return nil
	}

	if m.CPUInfo != nil {
		if err := m.CPUInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system" + "." + "cpu_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system" + "." + "cpu_info")
			}
			return err
		}
	}

	return nil
}

func (m *InfoSystem) validateMemInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.MemInfo) { // not required
		return nil
	}

	if m.MemInfo != nil {
		if err := m.MemInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system" + "." + "mem_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system" + "." + "mem_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this info system based on the context it is used
func (m *InfoSystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPUInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfoSystem) contextValidateCPUInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CPUInfo != nil {
		if err := m.CPUInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system" + "." + "cpu_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system" + "." + "cpu_info")
			}
			return err
		}
	}

	return nil
}

func (m *InfoSystem) contextValidateMemInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MemInfo != nil {
		if err := m.MemInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system" + "." + "mem_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system" + "." + "mem_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfoSystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoSystem) UnmarshalBinary(b []byte) error {
	var res InfoSystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoSystemCPUInfo info system CPU info
//
// swagger:model InfoSystemCPUInfo
type InfoSystemCPUInfo struct {

	// model
	Model string `json:"model,omitempty"`

	// Number of logical CPUs
	NumCpus int64 `json:"num_cpus,omitempty"`
}

// Validate validates this info system CPU info
func (m *InfoSystemCPUInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info system CPU info based on context it is used
func (m *InfoSystemCPUInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfoSystemCPUInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoSystemCPUInfo) UnmarshalBinary(b []byte) error {
	var res InfoSystemCPUInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoSystemMemInfo info system mem info
//
// swagger:model InfoSystemMemInfo
type InfoSystemMemInfo struct {

	// dataplaneapi memory
	DataplaneapiMemory int64 `json:"dataplaneapi_memory,omitempty"`

	// free memory
	FreeMemory int64 `json:"free_memory,omitempty"`

	// total memory
	TotalMemory int64 `json:"total_memory,omitempty"`
}

// Validate validates this info system mem info
func (m *InfoSystemMemInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info system mem info based on context it is used
func (m *InfoSystemMemInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfoSystemMemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoSystemMemInfo) UnmarshalBinary(b []byte) error {
	var res InfoSystemMemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}