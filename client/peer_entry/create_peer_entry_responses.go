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

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreatePeerEntryReader is a Reader for the CreatePeerEntry structure.
type CreatePeerEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePeerEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePeerEntryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreatePeerEntryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePeerEntryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreatePeerEntryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePeerEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePeerEntryCreated creates a CreatePeerEntryCreated with default headers values
func NewCreatePeerEntryCreated() *CreatePeerEntryCreated {
	return &CreatePeerEntryCreated{}
}

/* CreatePeerEntryCreated describes a response with status code 201, with default header values.

PeerEntry created
*/
type CreatePeerEntryCreated struct {
	Payload *models.PeerEntry
}

func (o *CreatePeerEntryCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryCreated  %+v", 201, o.Payload)
}
func (o *CreatePeerEntryCreated) GetPayload() *models.PeerEntry {
	return o.Payload
}

func (o *CreatePeerEntryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerEntryAccepted creates a CreatePeerEntryAccepted with default headers values
func NewCreatePeerEntryAccepted() *CreatePeerEntryAccepted {
	return &CreatePeerEntryAccepted{}
}

/* CreatePeerEntryAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreatePeerEntryAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.PeerEntry
}

func (o *CreatePeerEntryAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryAccepted  %+v", 202, o.Payload)
}
func (o *CreatePeerEntryAccepted) GetPayload() *models.PeerEntry {
	return o.Payload
}

func (o *CreatePeerEntryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.PeerEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerEntryBadRequest creates a CreatePeerEntryBadRequest with default headers values
func NewCreatePeerEntryBadRequest() *CreatePeerEntryBadRequest {
	return &CreatePeerEntryBadRequest{}
}

/* CreatePeerEntryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreatePeerEntryBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreatePeerEntryBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryBadRequest  %+v", 400, o.Payload)
}
func (o *CreatePeerEntryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerEntryConflict creates a CreatePeerEntryConflict with default headers values
func NewCreatePeerEntryConflict() *CreatePeerEntryConflict {
	return &CreatePeerEntryConflict{}
}

/* CreatePeerEntryConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreatePeerEntryConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreatePeerEntryConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryConflict  %+v", 409, o.Payload)
}
func (o *CreatePeerEntryConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerEntryDefault creates a CreatePeerEntryDefault with default headers values
func NewCreatePeerEntryDefault(code int) *CreatePeerEntryDefault {
	return &CreatePeerEntryDefault{
		_statusCode: code,
	}
}

/* CreatePeerEntryDefault describes a response with status code -1, with default header values.

General Error
*/
type CreatePeerEntryDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create peer entry default response
func (o *CreatePeerEntryDefault) Code() int {
	return o._statusCode
}

func (o *CreatePeerEntryDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntry default  %+v", o._statusCode, o.Payload)
}
func (o *CreatePeerEntryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
