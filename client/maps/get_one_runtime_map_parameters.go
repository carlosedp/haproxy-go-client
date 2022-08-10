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

package maps

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
)

// NewGetOneRuntimeMapParams creates a new GetOneRuntimeMapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOneRuntimeMapParams() *GetOneRuntimeMapParams {
	return &GetOneRuntimeMapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOneRuntimeMapParamsWithTimeout creates a new GetOneRuntimeMapParams object
// with the ability to set a timeout on a request.
func NewGetOneRuntimeMapParamsWithTimeout(timeout time.Duration) *GetOneRuntimeMapParams {
	return &GetOneRuntimeMapParams{
		timeout: timeout,
	}
}

// NewGetOneRuntimeMapParamsWithContext creates a new GetOneRuntimeMapParams object
// with the ability to set a context for a request.
func NewGetOneRuntimeMapParamsWithContext(ctx context.Context) *GetOneRuntimeMapParams {
	return &GetOneRuntimeMapParams{
		Context: ctx,
	}
}

// NewGetOneRuntimeMapParamsWithHTTPClient creates a new GetOneRuntimeMapParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOneRuntimeMapParamsWithHTTPClient(client *http.Client) *GetOneRuntimeMapParams {
	return &GetOneRuntimeMapParams{
		HTTPClient: client,
	}
}

/* GetOneRuntimeMapParams contains all the parameters to send to the API endpoint
   for the get one runtime map operation.

   Typically these are written to a http.Request.
*/
type GetOneRuntimeMapParams struct {

	/* Name.

	   Map file name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get one runtime map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOneRuntimeMapParams) WithDefaults() *GetOneRuntimeMapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get one runtime map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOneRuntimeMapParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get one runtime map params
func (o *GetOneRuntimeMapParams) WithTimeout(timeout time.Duration) *GetOneRuntimeMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get one runtime map params
func (o *GetOneRuntimeMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get one runtime map params
func (o *GetOneRuntimeMapParams) WithContext(ctx context.Context) *GetOneRuntimeMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get one runtime map params
func (o *GetOneRuntimeMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get one runtime map params
func (o *GetOneRuntimeMapParams) WithHTTPClient(client *http.Client) *GetOneRuntimeMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get one runtime map params
func (o *GetOneRuntimeMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get one runtime map params
func (o *GetOneRuntimeMapParams) WithName(name string) *GetOneRuntimeMapParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get one runtime map params
func (o *GetOneRuntimeMapParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetOneRuntimeMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}