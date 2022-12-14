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

package spoe_transactions

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

// NewDeleteSpoeTransactionParams creates a new DeleteSpoeTransactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSpoeTransactionParams() *DeleteSpoeTransactionParams {
	return &DeleteSpoeTransactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSpoeTransactionParamsWithTimeout creates a new DeleteSpoeTransactionParams object
// with the ability to set a timeout on a request.
func NewDeleteSpoeTransactionParamsWithTimeout(timeout time.Duration) *DeleteSpoeTransactionParams {
	return &DeleteSpoeTransactionParams{
		timeout: timeout,
	}
}

// NewDeleteSpoeTransactionParamsWithContext creates a new DeleteSpoeTransactionParams object
// with the ability to set a context for a request.
func NewDeleteSpoeTransactionParamsWithContext(ctx context.Context) *DeleteSpoeTransactionParams {
	return &DeleteSpoeTransactionParams{
		Context: ctx,
	}
}

// NewDeleteSpoeTransactionParamsWithHTTPClient creates a new DeleteSpoeTransactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSpoeTransactionParamsWithHTTPClient(client *http.Client) *DeleteSpoeTransactionParams {
	return &DeleteSpoeTransactionParams{
		HTTPClient: client,
	}
}

/* DeleteSpoeTransactionParams contains all the parameters to send to the API endpoint
   for the delete spoe transaction operation.

   Typically these are written to a http.Request.
*/
type DeleteSpoeTransactionParams struct {

	/* ID.

	   Transaction id
	*/
	ID string

	/* Spoe.

	   Spoe file name
	*/
	Spoe string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete spoe transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSpoeTransactionParams) WithDefaults() *DeleteSpoeTransactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete spoe transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSpoeTransactionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithTimeout(timeout time.Duration) *DeleteSpoeTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithContext(ctx context.Context) *DeleteSpoeTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithHTTPClient(client *http.Client) *DeleteSpoeTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithID(id string) *DeleteSpoeTransactionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetID(id string) {
	o.ID = id
}

// WithSpoe adds the spoe to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithSpoe(spoe string) *DeleteSpoeTransactionParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSpoeTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {

		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
