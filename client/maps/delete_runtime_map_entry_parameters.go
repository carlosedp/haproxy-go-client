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

package maps

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
)

// NewDeleteRuntimeMapEntryParams creates a new DeleteRuntimeMapEntryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRuntimeMapEntryParams() *DeleteRuntimeMapEntryParams {
	return &DeleteRuntimeMapEntryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRuntimeMapEntryParamsWithTimeout creates a new DeleteRuntimeMapEntryParams object
// with the ability to set a timeout on a request.
func NewDeleteRuntimeMapEntryParamsWithTimeout(timeout time.Duration) *DeleteRuntimeMapEntryParams {
	return &DeleteRuntimeMapEntryParams{
		timeout: timeout,
	}
}

// NewDeleteRuntimeMapEntryParamsWithContext creates a new DeleteRuntimeMapEntryParams object
// with the ability to set a context for a request.
func NewDeleteRuntimeMapEntryParamsWithContext(ctx context.Context) *DeleteRuntimeMapEntryParams {
	return &DeleteRuntimeMapEntryParams{
		Context: ctx,
	}
}

// NewDeleteRuntimeMapEntryParamsWithHTTPClient creates a new DeleteRuntimeMapEntryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRuntimeMapEntryParamsWithHTTPClient(client *http.Client) *DeleteRuntimeMapEntryParams {
	return &DeleteRuntimeMapEntryParams{
		HTTPClient: client,
	}
}

/* DeleteRuntimeMapEntryParams contains all the parameters to send to the API endpoint
   for the delete runtime map entry operation.

   Typically these are written to a http.Request.
*/
type DeleteRuntimeMapEntryParams struct {

	/* ForceSync.

	   If true, immediately syncs changes to disk
	*/
	ForceSync *bool

	/* ID.

	   Map id
	*/
	ID string

	/* Map.

	   Mapfile attribute storage_name
	*/
	Map string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete runtime map entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRuntimeMapEntryParams) WithDefaults() *DeleteRuntimeMapEntryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete runtime map entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRuntimeMapEntryParams) SetDefaults() {
	var (
		forceSyncDefault = bool(false)
	)

	val := DeleteRuntimeMapEntryParams{
		ForceSync: &forceSyncDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithTimeout(timeout time.Duration) *DeleteRuntimeMapEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithContext(ctx context.Context) *DeleteRuntimeMapEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithHTTPClient(client *http.Client) *DeleteRuntimeMapEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceSync adds the forceSync to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithForceSync(forceSync *bool) *DeleteRuntimeMapEntryParams {
	o.SetForceSync(forceSync)
	return o
}

// SetForceSync adds the forceSync to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetForceSync(forceSync *bool) {
	o.ForceSync = forceSync
}

// WithID adds the id to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithID(id string) *DeleteRuntimeMapEntryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetID(id string) {
	o.ID = id
}

// WithMap adds the mapVar to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithMap(mapVar string) *DeleteRuntimeMapEntryParams {
	o.SetMap(mapVar)
	return o
}

// SetMap adds the map to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetMap(mapVar string) {
	o.Map = mapVar
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRuntimeMapEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceSync != nil {

		// query param force_sync
		var qrForceSync bool

		if o.ForceSync != nil {
			qrForceSync = *o.ForceSync
		}
		qForceSync := swag.FormatBool(qrForceSync)
		if qForceSync != "" {

			if err := r.SetQueryParam("force_sync", qForceSync); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param map
	qrMap := o.Map
	qMap := qrMap
	if qMap != "" {

		if err := r.SetQueryParam("map", qMap); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
