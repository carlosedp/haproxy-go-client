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

package backend_switching_rule

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

// NewReplaceBackendSwitchingRuleParams creates a new ReplaceBackendSwitchingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceBackendSwitchingRuleParams() *ReplaceBackendSwitchingRuleParams {
	return &ReplaceBackendSwitchingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceBackendSwitchingRuleParamsWithTimeout creates a new ReplaceBackendSwitchingRuleParams object
// with the ability to set a timeout on a request.
func NewReplaceBackendSwitchingRuleParamsWithTimeout(timeout time.Duration) *ReplaceBackendSwitchingRuleParams {
	return &ReplaceBackendSwitchingRuleParams{
		timeout: timeout,
	}
}

// NewReplaceBackendSwitchingRuleParamsWithContext creates a new ReplaceBackendSwitchingRuleParams object
// with the ability to set a context for a request.
func NewReplaceBackendSwitchingRuleParamsWithContext(ctx context.Context) *ReplaceBackendSwitchingRuleParams {
	return &ReplaceBackendSwitchingRuleParams{
		Context: ctx,
	}
}

// NewReplaceBackendSwitchingRuleParamsWithHTTPClient creates a new ReplaceBackendSwitchingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceBackendSwitchingRuleParamsWithHTTPClient(client *http.Client) *ReplaceBackendSwitchingRuleParams {
	return &ReplaceBackendSwitchingRuleParams{
		HTTPClient: client,
	}
}

/* ReplaceBackendSwitchingRuleParams contains all the parameters to send to the API endpoint
   for the replace backend switching rule operation.

   Typically these are written to a http.Request.
*/
type ReplaceBackendSwitchingRuleParams struct {

	// Data.
	Data *models.BackendSwitchingRule

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Frontend.

	   Frontend name
	*/
	Frontend string

	/* Index.

	   Switching Rule Index
	*/
	Index int64

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

// WithDefaults hydrates default values in the replace backend switching rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceBackendSwitchingRuleParams) WithDefaults() *ReplaceBackendSwitchingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace backend switching rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceBackendSwitchingRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceBackendSwitchingRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithTimeout(timeout time.Duration) *ReplaceBackendSwitchingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithContext(ctx context.Context) *ReplaceBackendSwitchingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithHTTPClient(client *http.Client) *ReplaceBackendSwitchingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithData(data *models.BackendSwitchingRule) *ReplaceBackendSwitchingRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetData(data *models.BackendSwitchingRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithForceReload(forceReload *bool) *ReplaceBackendSwitchingRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithFrontend adds the frontend to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithFrontend(frontend string) *ReplaceBackendSwitchingRuleParams {
	o.SetFrontend(frontend)
	return o
}

// SetFrontend adds the frontend to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetFrontend(frontend string) {
	o.Frontend = frontend
}

// WithIndex adds the index to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithIndex(index int64) *ReplaceBackendSwitchingRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithTransactionID adds the transactionID to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithTransactionID(transactionID *string) *ReplaceBackendSwitchingRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) WithVersion(version *int64) *ReplaceBackendSwitchingRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace backend switching rule params
func (o *ReplaceBackendSwitchingRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceBackendSwitchingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param frontend
	qrFrontend := o.Frontend
	qFrontend := qrFrontend
	if qFrontend != "" {

		if err := r.SetQueryParam("frontend", qFrontend); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
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
