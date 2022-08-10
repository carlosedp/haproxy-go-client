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
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"models"
)

// NewCreateConsulParams creates a new CreateConsulParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateConsulParams() *CreateConsulParams {
	return &CreateConsulParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateConsulParamsWithTimeout creates a new CreateConsulParams object
// with the ability to set a timeout on a request.
func NewCreateConsulParamsWithTimeout(timeout time.Duration) *CreateConsulParams {
	return &CreateConsulParams{
		timeout: timeout,
	}
}

// NewCreateConsulParamsWithContext creates a new CreateConsulParams object
// with the ability to set a context for a request.
func NewCreateConsulParamsWithContext(ctx context.Context) *CreateConsulParams {
	return &CreateConsulParams{
		Context: ctx,
	}
}

// NewCreateConsulParamsWithHTTPClient creates a new CreateConsulParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateConsulParamsWithHTTPClient(client *http.Client) *CreateConsulParams {
	return &CreateConsulParams{
		HTTPClient: client,
	}
}

/* CreateConsulParams contains all the parameters to send to the API endpoint
   for the create consul operation.

   Typically these are written to a http.Request.
*/
type CreateConsulParams struct {

	// Data.
	Data *models.Consul

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create consul params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateConsulParams) WithDefaults() *CreateConsulParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create consul params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateConsulParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create consul params
func (o *CreateConsulParams) WithTimeout(timeout time.Duration) *CreateConsulParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create consul params
func (o *CreateConsulParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create consul params
func (o *CreateConsulParams) WithContext(ctx context.Context) *CreateConsulParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create consul params
func (o *CreateConsulParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create consul params
func (o *CreateConsulParams) WithHTTPClient(client *http.Client) *CreateConsulParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create consul params
func (o *CreateConsulParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create consul params
func (o *CreateConsulParams) WithData(data *models.Consul) *CreateConsulParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create consul params
func (o *CreateConsulParams) SetData(data *models.Consul) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CreateConsulParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}