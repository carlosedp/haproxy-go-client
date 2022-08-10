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

package service_discovery

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

	"github.com/haproxytech/client-native/v4/models"
)

// NewCreateAWSRegionParams creates a new CreateAWSRegionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAWSRegionParams() *CreateAWSRegionParams {
	return &CreateAWSRegionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAWSRegionParamsWithTimeout creates a new CreateAWSRegionParams object
// with the ability to set a timeout on a request.
func NewCreateAWSRegionParamsWithTimeout(timeout time.Duration) *CreateAWSRegionParams {
	return &CreateAWSRegionParams{
		timeout: timeout,
	}
}

// NewCreateAWSRegionParamsWithContext creates a new CreateAWSRegionParams object
// with the ability to set a context for a request.
func NewCreateAWSRegionParamsWithContext(ctx context.Context) *CreateAWSRegionParams {
	return &CreateAWSRegionParams{
		Context: ctx,
	}
}

// NewCreateAWSRegionParamsWithHTTPClient creates a new CreateAWSRegionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAWSRegionParamsWithHTTPClient(client *http.Client) *CreateAWSRegionParams {
	return &CreateAWSRegionParams{
		HTTPClient: client,
	}
}

/* CreateAWSRegionParams contains all the parameters to send to the API endpoint
   for the create a w s region operation.

   Typically these are written to a http.Request.
*/
type CreateAWSRegionParams struct {

	// Data.
	Data *models.AwsRegion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create a w s region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAWSRegionParams) WithDefaults() *CreateAWSRegionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create a w s region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAWSRegionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create a w s region params
func (o *CreateAWSRegionParams) WithTimeout(timeout time.Duration) *CreateAWSRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create a w s region params
func (o *CreateAWSRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create a w s region params
func (o *CreateAWSRegionParams) WithContext(ctx context.Context) *CreateAWSRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create a w s region params
func (o *CreateAWSRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create a w s region params
func (o *CreateAWSRegionParams) WithHTTPClient(client *http.Client) *CreateAWSRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create a w s region params
func (o *CreateAWSRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create a w s region params
func (o *CreateAWSRegionParams) WithData(data *models.AwsRegion) *CreateAWSRegionParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create a w s region params
func (o *CreateAWSRegionParams) SetData(data *models.AwsRegion) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAWSRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
