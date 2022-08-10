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

package transactions

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

// NewGetTransactionParams creates a new GetTransactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTransactionParams() *GetTransactionParams {
	return &GetTransactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransactionParamsWithTimeout creates a new GetTransactionParams object
// with the ability to set a timeout on a request.
func NewGetTransactionParamsWithTimeout(timeout time.Duration) *GetTransactionParams {
	return &GetTransactionParams{
		timeout: timeout,
	}
}

// NewGetTransactionParamsWithContext creates a new GetTransactionParams object
// with the ability to set a context for a request.
func NewGetTransactionParamsWithContext(ctx context.Context) *GetTransactionParams {
	return &GetTransactionParams{
		Context: ctx,
	}
}

// NewGetTransactionParamsWithHTTPClient creates a new GetTransactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTransactionParamsWithHTTPClient(client *http.Client) *GetTransactionParams {
	return &GetTransactionParams{
		HTTPClient: client,
	}
}

/* GetTransactionParams contains all the parameters to send to the API endpoint
   for the get transaction operation.

   Typically these are written to a http.Request.
*/
type GetTransactionParams struct {

	/* ID.

	   Transaction id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionParams) WithDefaults() *GetTransactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get transaction params
func (o *GetTransactionParams) WithTimeout(timeout time.Duration) *GetTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transaction params
func (o *GetTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transaction params
func (o *GetTransactionParams) WithContext(ctx context.Context) *GetTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transaction params
func (o *GetTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transaction params
func (o *GetTransactionParams) WithHTTPClient(client *http.Client) *GetTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transaction params
func (o *GetTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get transaction params
func (o *GetTransactionParams) WithID(id string) *GetTransactionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get transaction params
func (o *GetTransactionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}