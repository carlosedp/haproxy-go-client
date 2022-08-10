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

package http_check

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

// NewReplaceHTTPCheckParams creates a new ReplaceHTTPCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceHTTPCheckParams() *ReplaceHTTPCheckParams {
	return &ReplaceHTTPCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceHTTPCheckParamsWithTimeout creates a new ReplaceHTTPCheckParams object
// with the ability to set a timeout on a request.
func NewReplaceHTTPCheckParamsWithTimeout(timeout time.Duration) *ReplaceHTTPCheckParams {
	return &ReplaceHTTPCheckParams{
		timeout: timeout,
	}
}

// NewReplaceHTTPCheckParamsWithContext creates a new ReplaceHTTPCheckParams object
// with the ability to set a context for a request.
func NewReplaceHTTPCheckParamsWithContext(ctx context.Context) *ReplaceHTTPCheckParams {
	return &ReplaceHTTPCheckParams{
		Context: ctx,
	}
}

// NewReplaceHTTPCheckParamsWithHTTPClient creates a new ReplaceHTTPCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceHTTPCheckParamsWithHTTPClient(client *http.Client) *ReplaceHTTPCheckParams {
	return &ReplaceHTTPCheckParams{
		HTTPClient: client,
	}
}

/* ReplaceHTTPCheckParams contains all the parameters to send to the API endpoint
   for the replace HTTP check operation.

   Typically these are written to a http.Request.
*/
type ReplaceHTTPCheckParams struct {

	// Data.
	Data *models.HTTPCheck

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   HTTP Check Index
	*/
	Index int64

	/* ParentName.

	   Parent name
	*/
	ParentName *string

	/* ParentType.

	   Parent type
	*/
	ParentType string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	/* Version.

	   Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace HTTP check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceHTTPCheckParams) WithDefaults() *ReplaceHTTPCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace HTTP check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceHTTPCheckParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceHTTPCheckParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithTimeout(timeout time.Duration) *ReplaceHTTPCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithContext(ctx context.Context) *ReplaceHTTPCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithHTTPClient(client *http.Client) *ReplaceHTTPCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithData(data *models.HTTPCheck) *ReplaceHTTPCheckParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetData(data *models.HTTPCheck) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithForceReload(forceReload *bool) *ReplaceHTTPCheckParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithIndex(index int64) *ReplaceHTTPCheckParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithParentName(parentName *string) *ReplaceHTTPCheckParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithParentType(parentType string) *ReplaceHTTPCheckParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithTransactionID(transactionID *string) *ReplaceHTTPCheckParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) WithVersion(version *int64) *ReplaceHTTPCheckParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace HTTP check params
func (o *ReplaceHTTPCheckParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceHTTPCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	if o.ParentName != nil {

		// query param parent_name
		var qrParentName string

		if o.ParentName != nil {
			qrParentName = *o.ParentName
		}
		qParentName := qrParentName
		if qParentName != "" {

			if err := r.SetQueryParam("parent_name", qParentName); err != nil {
				return err
			}
		}
	}

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {

		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
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