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

package log_forward

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

// NewGetLogForwardParams creates a new GetLogForwardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogForwardParams() *GetLogForwardParams {
	return &GetLogForwardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogForwardParamsWithTimeout creates a new GetLogForwardParams object
// with the ability to set a timeout on a request.
func NewGetLogForwardParamsWithTimeout(timeout time.Duration) *GetLogForwardParams {
	return &GetLogForwardParams{
		timeout: timeout,
	}
}

// NewGetLogForwardParamsWithContext creates a new GetLogForwardParams object
// with the ability to set a context for a request.
func NewGetLogForwardParamsWithContext(ctx context.Context) *GetLogForwardParams {
	return &GetLogForwardParams{
		Context: ctx,
	}
}

// NewGetLogForwardParamsWithHTTPClient creates a new GetLogForwardParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogForwardParamsWithHTTPClient(client *http.Client) *GetLogForwardParams {
	return &GetLogForwardParams{
		HTTPClient: client,
	}
}

/* GetLogForwardParams contains all the parameters to send to the API endpoint
   for the get log forward operation.

   Typically these are written to a http.Request.
*/
type GetLogForwardParams struct {

	/* Name.

	   Log Forward name
	*/
	Name string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogForwardParams) WithDefaults() *GetLogForwardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogForwardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log forward params
func (o *GetLogForwardParams) WithTimeout(timeout time.Duration) *GetLogForwardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log forward params
func (o *GetLogForwardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log forward params
func (o *GetLogForwardParams) WithContext(ctx context.Context) *GetLogForwardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log forward params
func (o *GetLogForwardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log forward params
func (o *GetLogForwardParams) WithHTTPClient(client *http.Client) *GetLogForwardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log forward params
func (o *GetLogForwardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get log forward params
func (o *GetLogForwardParams) WithName(name string) *GetLogForwardParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get log forward params
func (o *GetLogForwardParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the get log forward params
func (o *GetLogForwardParams) WithTransactionID(transactionID *string) *GetLogForwardParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get log forward params
func (o *GetLogForwardParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogForwardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string

		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {

			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
