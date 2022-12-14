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

// NewDeleteTransactionParams creates a new DeleteTransactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTransactionParams() *DeleteTransactionParams {
	return &DeleteTransactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTransactionParamsWithTimeout creates a new DeleteTransactionParams object
// with the ability to set a timeout on a request.
func NewDeleteTransactionParamsWithTimeout(timeout time.Duration) *DeleteTransactionParams {
	return &DeleteTransactionParams{
		timeout: timeout,
	}
}

// NewDeleteTransactionParamsWithContext creates a new DeleteTransactionParams object
// with the ability to set a context for a request.
func NewDeleteTransactionParamsWithContext(ctx context.Context) *DeleteTransactionParams {
	return &DeleteTransactionParams{
		Context: ctx,
	}
}

// NewDeleteTransactionParamsWithHTTPClient creates a new DeleteTransactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTransactionParamsWithHTTPClient(client *http.Client) *DeleteTransactionParams {
	return &DeleteTransactionParams{
		HTTPClient: client,
	}
}

/* DeleteTransactionParams contains all the parameters to send to the API endpoint
   for the delete transaction operation.

   Typically these are written to a http.Request.
*/
type DeleteTransactionParams struct {

	/* ID.

	   Transaction id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTransactionParams) WithDefaults() *DeleteTransactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTransactionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete transaction params
func (o *DeleteTransactionParams) WithTimeout(timeout time.Duration) *DeleteTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete transaction params
func (o *DeleteTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete transaction params
func (o *DeleteTransactionParams) WithContext(ctx context.Context) *DeleteTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete transaction params
func (o *DeleteTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete transaction params
func (o *DeleteTransactionParams) WithHTTPClient(client *http.Client) *DeleteTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete transaction params
func (o *DeleteTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete transaction params
func (o *DeleteTransactionParams) WithID(id string) *DeleteTransactionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete transaction params
func (o *DeleteTransactionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
