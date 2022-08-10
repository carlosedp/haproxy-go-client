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

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// DeletePeerReader is a Reader for the DeletePeer structure.
type DeletePeerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePeerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeletePeerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeletePeerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeletePeerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePeerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePeerAccepted creates a DeletePeerAccepted with default headers values
func NewDeletePeerAccepted() *DeletePeerAccepted {
	return &DeletePeerAccepted{}
}

/* DeletePeerAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeletePeerAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeletePeerAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerAccepted ", 202)
}

func (o *DeletePeerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeletePeerNoContent creates a DeletePeerNoContent with default headers values
func NewDeletePeerNoContent() *DeletePeerNoContent {
	return &DeletePeerNoContent{}
}

/* DeletePeerNoContent describes a response with status code 204, with default header values.

Peer deleted
*/
type DeletePeerNoContent struct {
}

func (o *DeletePeerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerNoContent ", 204)
}

func (o *DeletePeerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePeerNotFound creates a DeletePeerNotFound with default headers values
func NewDeletePeerNotFound() *DeletePeerNotFound {
	return &DeletePeerNotFound{}
}

/* DeletePeerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeletePeerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeletePeerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerNotFound  %+v", 404, o.Payload)
}
func (o *DeletePeerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePeerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePeerDefault creates a DeletePeerDefault with default headers values
func NewDeletePeerDefault(code int) *DeletePeerDefault {
	return &DeletePeerDefault{
		_statusCode: code,
	}
}

/* DeletePeerDefault describes a response with status code -1, with default header values.

General Error
*/
type DeletePeerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete peer default response
func (o *DeletePeerDefault) Code() int {
	return o._statusCode
}

func (o *DeletePeerDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeer default  %+v", o._statusCode, o.Payload)
}
func (o *DeletePeerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePeerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
