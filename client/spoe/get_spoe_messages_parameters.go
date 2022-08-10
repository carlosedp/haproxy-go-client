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

package spoe

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

// NewGetSpoeMessagesParams creates a new GetSpoeMessagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSpoeMessagesParams() *GetSpoeMessagesParams {
	return &GetSpoeMessagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoeMessagesParamsWithTimeout creates a new GetSpoeMessagesParams object
// with the ability to set a timeout on a request.
func NewGetSpoeMessagesParamsWithTimeout(timeout time.Duration) *GetSpoeMessagesParams {
	return &GetSpoeMessagesParams{
		timeout: timeout,
	}
}

// NewGetSpoeMessagesParamsWithContext creates a new GetSpoeMessagesParams object
// with the ability to set a context for a request.
func NewGetSpoeMessagesParamsWithContext(ctx context.Context) *GetSpoeMessagesParams {
	return &GetSpoeMessagesParams{
		Context: ctx,
	}
}

// NewGetSpoeMessagesParamsWithHTTPClient creates a new GetSpoeMessagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSpoeMessagesParamsWithHTTPClient(client *http.Client) *GetSpoeMessagesParams {
	return &GetSpoeMessagesParams{
		HTTPClient: client,
	}
}

/* GetSpoeMessagesParams contains all the parameters to send to the API endpoint
   for the get spoe messages operation.

   Typically these are written to a http.Request.
*/
type GetSpoeMessagesParams struct {

	/* Scope.

	   Spoe scope
	*/
	Scope string

	/* Spoe.

	   Spoe file name
	*/
	Spoe string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get spoe messages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeMessagesParams) WithDefaults() *GetSpoeMessagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get spoe messages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeMessagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get spoe messages params
func (o *GetSpoeMessagesParams) WithTimeout(timeout time.Duration) *GetSpoeMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoe messages params
func (o *GetSpoeMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoe messages params
func (o *GetSpoeMessagesParams) WithContext(ctx context.Context) *GetSpoeMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoe messages params
func (o *GetSpoeMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoe messages params
func (o *GetSpoeMessagesParams) WithHTTPClient(client *http.Client) *GetSpoeMessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoe messages params
func (o *GetSpoeMessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the get spoe messages params
func (o *GetSpoeMessagesParams) WithScope(scope string) *GetSpoeMessagesParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get spoe messages params
func (o *GetSpoeMessagesParams) SetScope(scope string) {
	o.Scope = scope
}

// WithSpoe adds the spoe to the get spoe messages params
func (o *GetSpoeMessagesParams) WithSpoe(spoe string) *GetSpoeMessagesParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the get spoe messages params
func (o *GetSpoeMessagesParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the get spoe messages params
func (o *GetSpoeMessagesParams) WithTransactionID(transactionID *string) *GetSpoeMessagesParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get spoe messages params
func (o *GetSpoeMessagesParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoeMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param scope
	qrScope := o.Scope
	qScope := qrScope
	if qScope != "" {

		if err := r.SetQueryParam("scope", qScope); err != nil {
			return err
		}
	}

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {

		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
			return err
		}
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