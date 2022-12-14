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

package nameserver

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

// NewGetNameserversParams creates a new GetNameserversParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNameserversParams() *GetNameserversParams {
	return &GetNameserversParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNameserversParamsWithTimeout creates a new GetNameserversParams object
// with the ability to set a timeout on a request.
func NewGetNameserversParamsWithTimeout(timeout time.Duration) *GetNameserversParams {
	return &GetNameserversParams{
		timeout: timeout,
	}
}

// NewGetNameserversParamsWithContext creates a new GetNameserversParams object
// with the ability to set a context for a request.
func NewGetNameserversParamsWithContext(ctx context.Context) *GetNameserversParams {
	return &GetNameserversParams{
		Context: ctx,
	}
}

// NewGetNameserversParamsWithHTTPClient creates a new GetNameserversParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNameserversParamsWithHTTPClient(client *http.Client) *GetNameserversParams {
	return &GetNameserversParams{
		HTTPClient: client,
	}
}

/* GetNameserversParams contains all the parameters to send to the API endpoint
   for the get nameservers operation.

   Typically these are written to a http.Request.
*/
type GetNameserversParams struct {

	/* Resolver.

	   Parent resolver name
	*/
	Resolver string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nameservers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNameserversParams) WithDefaults() *GetNameserversParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nameservers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNameserversParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nameservers params
func (o *GetNameserversParams) WithTimeout(timeout time.Duration) *GetNameserversParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nameservers params
func (o *GetNameserversParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nameservers params
func (o *GetNameserversParams) WithContext(ctx context.Context) *GetNameserversParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nameservers params
func (o *GetNameserversParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nameservers params
func (o *GetNameserversParams) WithHTTPClient(client *http.Client) *GetNameserversParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nameservers params
func (o *GetNameserversParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResolver adds the resolver to the get nameservers params
func (o *GetNameserversParams) WithResolver(resolver string) *GetNameserversParams {
	o.SetResolver(resolver)
	return o
}

// SetResolver adds the resolver to the get nameservers params
func (o *GetNameserversParams) SetResolver(resolver string) {
	o.Resolver = resolver
}

// WithTransactionID adds the transactionID to the get nameservers params
func (o *GetNameserversParams) WithTransactionID(transactionID *string) *GetNameserversParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get nameservers params
func (o *GetNameserversParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNameserversParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param resolver
	qrResolver := o.Resolver
	qResolver := qrResolver
	if qResolver != "" {

		if err := r.SetQueryParam("resolver", qResolver); err != nil {
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
