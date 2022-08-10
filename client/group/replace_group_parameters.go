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

package group

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

	"models"
)

// NewReplaceGroupParams creates a new ReplaceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceGroupParams() *ReplaceGroupParams {
	return &ReplaceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceGroupParamsWithTimeout creates a new ReplaceGroupParams object
// with the ability to set a timeout on a request.
func NewReplaceGroupParamsWithTimeout(timeout time.Duration) *ReplaceGroupParams {
	return &ReplaceGroupParams{
		timeout: timeout,
	}
}

// NewReplaceGroupParamsWithContext creates a new ReplaceGroupParams object
// with the ability to set a context for a request.
func NewReplaceGroupParamsWithContext(ctx context.Context) *ReplaceGroupParams {
	return &ReplaceGroupParams{
		Context: ctx,
	}
}

// NewReplaceGroupParamsWithHTTPClient creates a new ReplaceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceGroupParamsWithHTTPClient(client *http.Client) *ReplaceGroupParams {
	return &ReplaceGroupParams{
		HTTPClient: client,
	}
}

/* ReplaceGroupParams contains all the parameters to send to the API endpoint
   for the replace group operation.

   Typically these are written to a http.Request.
*/
type ReplaceGroupParams struct {

	// Data.
	Data *models.Group

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   group name
	*/
	Name string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	/* Userlist.

	   Parent userlist name
	*/
	Userlist string

	/* Version.

	   Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceGroupParams) WithDefaults() *ReplaceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceGroupParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceGroupParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace group params
func (o *ReplaceGroupParams) WithTimeout(timeout time.Duration) *ReplaceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace group params
func (o *ReplaceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace group params
func (o *ReplaceGroupParams) WithContext(ctx context.Context) *ReplaceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace group params
func (o *ReplaceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace group params
func (o *ReplaceGroupParams) WithHTTPClient(client *http.Client) *ReplaceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace group params
func (o *ReplaceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace group params
func (o *ReplaceGroupParams) WithData(data *models.Group) *ReplaceGroupParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace group params
func (o *ReplaceGroupParams) SetData(data *models.Group) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace group params
func (o *ReplaceGroupParams) WithForceReload(forceReload *bool) *ReplaceGroupParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace group params
func (o *ReplaceGroupParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the replace group params
func (o *ReplaceGroupParams) WithName(name string) *ReplaceGroupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace group params
func (o *ReplaceGroupParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the replace group params
func (o *ReplaceGroupParams) WithTransactionID(transactionID *string) *ReplaceGroupParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace group params
func (o *ReplaceGroupParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithUserlist adds the userlist to the replace group params
func (o *ReplaceGroupParams) WithUserlist(userlist string) *ReplaceGroupParams {
	o.SetUserlist(userlist)
	return o
}

// SetUserlist adds the userlist to the replace group params
func (o *ReplaceGroupParams) SetUserlist(userlist string) {
	o.Userlist = userlist
}

// WithVersion adds the version to the replace group params
func (o *ReplaceGroupParams) WithVersion(version *int64) *ReplaceGroupParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace group params
func (o *ReplaceGroupParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool

		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {

			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}
	}

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

	// query param userlist
	qrUserlist := o.Userlist
	qUserlist := qrUserlist
	if qUserlist != "" {

		if err := r.SetQueryParam("userlist", qUserlist); err != nil {
			return err
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
