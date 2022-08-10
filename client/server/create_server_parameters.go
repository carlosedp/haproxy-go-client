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

package server

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

// NewCreateServerParams creates a new CreateServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServerParams() *CreateServerParams {
	return &CreateServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServerParamsWithTimeout creates a new CreateServerParams object
// with the ability to set a timeout on a request.
func NewCreateServerParamsWithTimeout(timeout time.Duration) *CreateServerParams {
	return &CreateServerParams{
		timeout: timeout,
	}
}

// NewCreateServerParamsWithContext creates a new CreateServerParams object
// with the ability to set a context for a request.
func NewCreateServerParamsWithContext(ctx context.Context) *CreateServerParams {
	return &CreateServerParams{
		Context: ctx,
	}
}

// NewCreateServerParamsWithHTTPClient creates a new CreateServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServerParamsWithHTTPClient(client *http.Client) *CreateServerParams {
	return &CreateServerParams{
		HTTPClient: client,
	}
}

/* CreateServerParams contains all the parameters to send to the API endpoint
   for the create server operation.

   Typically these are written to a http.Request.
*/
type CreateServerParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend *string

	// Data.
	Data *models.Server

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

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

// WithDefaults hydrates default values in the create server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerParams) WithDefaults() *CreateServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateServerParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create server params
func (o *CreateServerParams) WithTimeout(timeout time.Duration) *CreateServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create server params
func (o *CreateServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create server params
func (o *CreateServerParams) WithContext(ctx context.Context) *CreateServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create server params
func (o *CreateServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create server params
func (o *CreateServerParams) WithHTTPClient(client *http.Client) *CreateServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create server params
func (o *CreateServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the create server params
func (o *CreateServerParams) WithBackend(backend *string) *CreateServerParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the create server params
func (o *CreateServerParams) SetBackend(backend *string) {
	o.Backend = backend
}

// WithData adds the data to the create server params
func (o *CreateServerParams) WithData(data *models.Server) *CreateServerParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create server params
func (o *CreateServerParams) SetData(data *models.Server) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create server params
func (o *CreateServerParams) WithForceReload(forceReload *bool) *CreateServerParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create server params
func (o *CreateServerParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithParentName adds the parentName to the create server params
func (o *CreateServerParams) WithParentName(parentName *string) *CreateServerParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the create server params
func (o *CreateServerParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the create server params
func (o *CreateServerParams) WithParentType(parentType *string) *CreateServerParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the create server params
func (o *CreateServerParams) SetParentType(parentType *string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the create server params
func (o *CreateServerParams) WithTransactionID(transactionID *string) *CreateServerParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create server params
func (o *CreateServerParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create server params
func (o *CreateServerParams) WithVersion(version *int64) *CreateServerParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create server params
func (o *CreateServerParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Backend != nil {

		// query param backend
		var qrBackend string

		if o.Backend != nil {
			qrBackend = *o.Backend
		}
		qBackend := qrBackend
		if qBackend != "" {

			if err := r.SetQueryParam("backend", qBackend); err != nil {
				return err
			}
		}
	}
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
