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

// NewGetOneStorageSSLCertificateParams creates a new GetOneStorageSSLCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOneStorageSSLCertificateParams() *GetOneStorageSSLCertificateParams {
	return &GetOneStorageSSLCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOneStorageSSLCertificateParamsWithTimeout creates a new GetOneStorageSSLCertificateParams object
// with the ability to set a timeout on a request.
func NewGetOneStorageSSLCertificateParamsWithTimeout(timeout time.Duration) *GetOneStorageSSLCertificateParams {
	return &GetOneStorageSSLCertificateParams{
		timeout: timeout,
	}
}

// NewGetOneStorageSSLCertificateParamsWithContext creates a new GetOneStorageSSLCertificateParams object
// with the ability to set a context for a request.
func NewGetOneStorageSSLCertificateParamsWithContext(ctx context.Context) *GetOneStorageSSLCertificateParams {
	return &GetOneStorageSSLCertificateParams{
		Context: ctx,
	}
}

// NewGetOneStorageSSLCertificateParamsWithHTTPClient creates a new GetOneStorageSSLCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOneStorageSSLCertificateParamsWithHTTPClient(client *http.Client) *GetOneStorageSSLCertificateParams {
	return &GetOneStorageSSLCertificateParams{
		HTTPClient: client,
	}
}

/* GetOneStorageSSLCertificateParams contains all the parameters to send to the API endpoint
   for the get one storage s s l certificate operation.

   Typically these are written to a http.Request.
*/
type GetOneStorageSSLCertificateParams struct {

	/* Name.

	   SSL certificate name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get one storage s s l certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOneStorageSSLCertificateParams) WithDefaults() *GetOneStorageSSLCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get one storage s s l certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOneStorageSSLCertificateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get one storage s s l certificate params
func (o *GetOneStorageSSLCertificateParams) WithTimeout(timeout time.Duration) *GetOneStorageSSLCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get one storage s s l certificate params
func (o *GetOneStorageSSLCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get one storage s s l certificate params
func (o *GetOneStorageSSLCertificateParams) WithContext(ctx context.Context) *GetOneStorageSSLCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get one storage s s l certificate params
func (o *GetOneStorageSSLCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get one storage s s l certificate params
func (o *GetOneStorageSSLCertificateParams) WithHTTPClient(client *http.Client) *GetOneStorageSSLCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get one storage s s l certificate params
func (o *GetOneStorageSSLCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get one storage s s l certificate params
func (o *GetOneStorageSSLCertificateParams) WithName(name string) *GetOneStorageSSLCertificateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get one storage s s l certificate params
func (o *GetOneStorageSSLCertificateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetOneStorageSSLCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
