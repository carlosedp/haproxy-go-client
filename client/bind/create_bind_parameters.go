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

package bind

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

// NewCreateBindParams creates a new CreateBindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBindParams() *CreateBindParams {
	return &CreateBindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBindParamsWithTimeout creates a new CreateBindParams object
// with the ability to set a timeout on a request.
func NewCreateBindParamsWithTimeout(timeout time.Duration) *CreateBindParams {
	return &CreateBindParams{
		timeout: timeout,
	}
}

// NewCreateBindParamsWithContext creates a new CreateBindParams object
// with the ability to set a context for a request.
func NewCreateBindParamsWithContext(ctx context.Context) *CreateBindParams {
	return &CreateBindParams{
		Context: ctx,
	}
}

// NewCreateBindParamsWithHTTPClient creates a new CreateBindParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBindParamsWithHTTPClient(client *http.Client) *CreateBindParams {
	return &CreateBindParams{
		HTTPClient: client,
	}
}

/* CreateBindParams contains all the parameters to send to the API endpoint
   for the create bind operation.

   Typically these are written to a http.Request.
*/
type CreateBindParams struct {

	// Data.
	Data *models.Bind

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Frontend.

	   Parent frontend name
	*/
	Frontend *string

	/* ParentName.

	   Parent name
	*/
	ParentName *string

	/* ParentType.

	   Parent type
	*/
	ParentType *string

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

// WithDefaults hydrates default values in the create bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBindParams) WithDefaults() *CreateBindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBindParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateBindParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create bind params
func (o *CreateBindParams) WithTimeout(timeout time.Duration) *CreateBindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create bind params
func (o *CreateBindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create bind params
func (o *CreateBindParams) WithContext(ctx context.Context) *CreateBindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create bind params
func (o *CreateBindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create bind params
func (o *CreateBindParams) WithHTTPClient(client *http.Client) *CreateBindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create bind params
func (o *CreateBindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create bind params
func (o *CreateBindParams) WithData(data *models.Bind) *CreateBindParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create bind params
func (o *CreateBindParams) SetData(data *models.Bind) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create bind params
func (o *CreateBindParams) WithForceReload(forceReload *bool) *CreateBindParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create bind params
func (o *CreateBindParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithFrontend adds the frontend to the create bind params
func (o *CreateBindParams) WithFrontend(frontend *string) *CreateBindParams {
	o.SetFrontend(frontend)
	return o
}

// SetFrontend adds the frontend to the create bind params
func (o *CreateBindParams) SetFrontend(frontend *string) {
	o.Frontend = frontend
}

// WithParentName adds the parentName to the create bind params
func (o *CreateBindParams) WithParentName(parentName *string) *CreateBindParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the create bind params
func (o *CreateBindParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the create bind params
func (o *CreateBindParams) WithParentType(parentType *string) *CreateBindParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the create bind params
func (o *CreateBindParams) SetParentType(parentType *string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the create bind params
func (o *CreateBindParams) WithTransactionID(transactionID *string) *CreateBindParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create bind params
func (o *CreateBindParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create bind params
func (o *CreateBindParams) WithVersion(version *int64) *CreateBindParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create bind params
func (o *CreateBindParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Frontend != nil {

		// query param frontend
		var qrFrontend string

		if o.Frontend != nil {
			qrFrontend = *o.Frontend
		}
		qFrontend := qrFrontend
		if qFrontend != "" {

			if err := r.SetQueryParam("frontend", qFrontend); err != nil {
				return err
			}
		}
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

	if o.ParentType != nil {

		// query param parent_type
		var qrParentType string

		if o.ParentType != nil {
			qrParentType = *o.ParentType
		}
		qParentType := qrParentType
		if qParentType != "" {

			if err := r.SetQueryParam("parent_type", qParentType); err != nil {
				return err
			}
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
