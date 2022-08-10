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

package acl_runtime

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

	"models"
)

// NewAddPayloadRuntimeACLParams creates a new AddPayloadRuntimeACLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddPayloadRuntimeACLParams() *AddPayloadRuntimeACLParams {
	return &AddPayloadRuntimeACLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddPayloadRuntimeACLParamsWithTimeout creates a new AddPayloadRuntimeACLParams object
// with the ability to set a timeout on a request.
func NewAddPayloadRuntimeACLParamsWithTimeout(timeout time.Duration) *AddPayloadRuntimeACLParams {
	return &AddPayloadRuntimeACLParams{
		timeout: timeout,
	}
}

// NewAddPayloadRuntimeACLParamsWithContext creates a new AddPayloadRuntimeACLParams object
// with the ability to set a context for a request.
func NewAddPayloadRuntimeACLParamsWithContext(ctx context.Context) *AddPayloadRuntimeACLParams {
	return &AddPayloadRuntimeACLParams{
		Context: ctx,
	}
}

// NewAddPayloadRuntimeACLParamsWithHTTPClient creates a new AddPayloadRuntimeACLParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddPayloadRuntimeACLParamsWithHTTPClient(client *http.Client) *AddPayloadRuntimeACLParams {
	return &AddPayloadRuntimeACLParams{
		HTTPClient: client,
	}
}

/* AddPayloadRuntimeACLParams contains all the parameters to send to the API endpoint
   for the add payload runtime ACL operation.

   Typically these are written to a http.Request.
*/
type AddPayloadRuntimeACLParams struct {

	/* ACLID.

	   ACL ID
	*/
	ACLID string

	// Data.
	Data models.ACLFilesEntries

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add payload runtime ACL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPayloadRuntimeACLParams) WithDefaults() *AddPayloadRuntimeACLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add payload runtime ACL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPayloadRuntimeACLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithTimeout(timeout time.Duration) *AddPayloadRuntimeACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithContext(ctx context.Context) *AddPayloadRuntimeACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithHTTPClient(client *http.Client) *AddPayloadRuntimeACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLID adds the aCLID to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithACLID(aCLID string) *AddPayloadRuntimeACLParams {
	o.SetACLID(aCLID)
	return o
}

// SetACLID adds the aclId to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetACLID(aCLID string) {
	o.ACLID = aCLID
}

// WithData adds the data to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithData(data models.ACLFilesEntries) *AddPayloadRuntimeACLParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetData(data models.ACLFilesEntries) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *AddPayloadRuntimeACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param acl_id
	qrACLID := o.ACLID
	qACLID := qrACLID
	if qACLID != "" {

		if err := r.SetQueryParam("acl_id", qACLID); err != nil {
			return err
		}
	}
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
