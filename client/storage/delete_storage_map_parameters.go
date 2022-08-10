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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteStorageMapParams creates a new DeleteStorageMapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteStorageMapParams() *DeleteStorageMapParams {
	return &DeleteStorageMapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStorageMapParamsWithTimeout creates a new DeleteStorageMapParams object
// with the ability to set a timeout on a request.
func NewDeleteStorageMapParamsWithTimeout(timeout time.Duration) *DeleteStorageMapParams {
	return &DeleteStorageMapParams{
		timeout: timeout,
	}
}

// NewDeleteStorageMapParamsWithContext creates a new DeleteStorageMapParams object
// with the ability to set a context for a request.
func NewDeleteStorageMapParamsWithContext(ctx context.Context) *DeleteStorageMapParams {
	return &DeleteStorageMapParams{
		Context: ctx,
	}
}

// NewDeleteStorageMapParamsWithHTTPClient creates a new DeleteStorageMapParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteStorageMapParamsWithHTTPClient(client *http.Client) *DeleteStorageMapParams {
	return &DeleteStorageMapParams{
		HTTPClient: client,
	}
}

/* DeleteStorageMapParams contains all the parameters to send to the API endpoint
   for the delete storage map operation.

   Typically these are written to a http.Request.
*/
type DeleteStorageMapParams struct {

	/* Name.

	   Map file storage_name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete storage map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteStorageMapParams) WithDefaults() *DeleteStorageMapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete storage map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteStorageMapParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete storage map params
func (o *DeleteStorageMapParams) WithTimeout(timeout time.Duration) *DeleteStorageMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete storage map params
func (o *DeleteStorageMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete storage map params
func (o *DeleteStorageMapParams) WithContext(ctx context.Context) *DeleteStorageMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete storage map params
func (o *DeleteStorageMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete storage map params
func (o *DeleteStorageMapParams) WithHTTPClient(client *http.Client) *DeleteStorageMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete storage map params
func (o *DeleteStorageMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete storage map params
func (o *DeleteStorageMapParams) WithName(name string) *DeleteStorageMapParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete storage map params
func (o *DeleteStorageMapParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStorageMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
