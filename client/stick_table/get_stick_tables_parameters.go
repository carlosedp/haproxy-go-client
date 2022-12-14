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

package stick_table

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
	"github.com/go-openapi/swag"
)

// NewGetStickTablesParams creates a new GetStickTablesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStickTablesParams() *GetStickTablesParams {
	return &GetStickTablesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStickTablesParamsWithTimeout creates a new GetStickTablesParams object
// with the ability to set a timeout on a request.
func NewGetStickTablesParamsWithTimeout(timeout time.Duration) *GetStickTablesParams {
	return &GetStickTablesParams{
		timeout: timeout,
	}
}

// NewGetStickTablesParamsWithContext creates a new GetStickTablesParams object
// with the ability to set a context for a request.
func NewGetStickTablesParamsWithContext(ctx context.Context) *GetStickTablesParams {
	return &GetStickTablesParams{
		Context: ctx,
	}
}

// NewGetStickTablesParamsWithHTTPClient creates a new GetStickTablesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStickTablesParamsWithHTTPClient(client *http.Client) *GetStickTablesParams {
	return &GetStickTablesParams{
		HTTPClient: client,
	}
}

/* GetStickTablesParams contains all the parameters to send to the API endpoint
   for the get stick tables operation.

   Typically these are written to a http.Request.
*/
type GetStickTablesParams struct {

	/* Process.

	   Process number if master-worker mode, if not all processes are returned
	*/
	Process *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get stick tables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStickTablesParams) WithDefaults() *GetStickTablesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get stick tables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStickTablesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get stick tables params
func (o *GetStickTablesParams) WithTimeout(timeout time.Duration) *GetStickTablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stick tables params
func (o *GetStickTablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stick tables params
func (o *GetStickTablesParams) WithContext(ctx context.Context) *GetStickTablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stick tables params
func (o *GetStickTablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stick tables params
func (o *GetStickTablesParams) WithHTTPClient(client *http.Client) *GetStickTablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stick tables params
func (o *GetStickTablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProcess adds the process to the get stick tables params
func (o *GetStickTablesParams) WithProcess(process *int64) *GetStickTablesParams {
	o.SetProcess(process)
	return o
}

// SetProcess adds the process to the get stick tables params
func (o *GetStickTablesParams) SetProcess(process *int64) {
	o.Process = process
}

// WriteToRequest writes these params to a swagger request
func (o *GetStickTablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Process != nil {

		// query param process
		var qrProcess int64

		if o.Process != nil {
			qrProcess = *o.Process
		}
		qProcess := swag.FormatInt64(qrProcess)
		if qProcess != "" {

			if err := r.SetQueryParam("process", qProcess); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
