// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// ListingOKCode is the HTTP code returned for type ListingOK
const ListingOKCode int = 200

/*ListingOK Content of the dataset

swagger:response listingOK
*/
type ListingOK struct {

	/*
	  In: Body
	*/
	Payload *models.DatasetContent `json:"body,omitempty"`
}

// NewListingOK creates ListingOK with default headers values
func NewListingOK() *ListingOK {

	return &ListingOK{}
}

// WithPayload adds the payload to the listing o k response
func (o *ListingOK) WithPayload(payload *models.DatasetContent) *ListingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listing o k response
func (o *ListingOK) SetPayload(payload *models.DatasetContent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListingBadRequestCode is the HTTP code returned for type ListingBadRequest
const ListingBadRequestCode int = 400

/*ListingBadRequest Malformed request

swagger:response listingBadRequest
*/
type ListingBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListingBadRequest creates ListingBadRequest with default headers values
func NewListingBadRequest() *ListingBadRequest {

	return &ListingBadRequest{}
}

// WithPayload adds the payload to the listing bad request response
func (o *ListingBadRequest) WithPayload(payload *models.ErrorResponse) *ListingBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listing bad request response
func (o *ListingBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListingUnauthorizedCode is the HTTP code returned for type ListingUnauthorized
const ListingUnauthorizedCode int = 401

/*ListingUnauthorized Authorization failed

swagger:response listingUnauthorized
*/
type ListingUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListingUnauthorized creates ListingUnauthorized with default headers values
func NewListingUnauthorized() *ListingUnauthorized {

	return &ListingUnauthorized{}
}

// WithPayload adds the payload to the listing unauthorized response
func (o *ListingUnauthorized) WithPayload(payload *models.ErrorResponse) *ListingUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listing unauthorized response
func (o *ListingUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListingForbiddenCode is the HTTP code returned for type ListingForbidden
const ListingForbiddenCode int = 403

/*ListingForbidden User does not have permission

swagger:response listingForbidden
*/
type ListingForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListingForbidden creates ListingForbidden with default headers values
func NewListingForbidden() *ListingForbidden {

	return &ListingForbidden{}
}

// WithPayload adds the payload to the listing forbidden response
func (o *ListingForbidden) WithPayload(payload *models.ErrorResponse) *ListingForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listing forbidden response
func (o *ListingForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListingForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListingBadGatewayCode is the HTTP code returned for type ListingBadGateway
const ListingBadGatewayCode int = 502

/*ListingBadGateway Bad Gateway

swagger:response listingBadGateway
*/
type ListingBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListingBadGateway creates ListingBadGateway with default headers values
func NewListingBadGateway() *ListingBadGateway {

	return &ListingBadGateway{}
}

// WithPayload adds the payload to the listing bad gateway response
func (o *ListingBadGateway) WithPayload(payload *models.ErrorResponse) *ListingBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listing bad gateway response
func (o *ListingBadGateway) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListingBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListingServiceUnavailableCode is the HTTP code returned for type ListingServiceUnavailable
const ListingServiceUnavailableCode int = 503

/*ListingServiceUnavailable Error accessing backend service

swagger:response listingServiceUnavailable
*/
type ListingServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListingServiceUnavailable creates ListingServiceUnavailable with default headers values
func NewListingServiceUnavailable() *ListingServiceUnavailable {

	return &ListingServiceUnavailable{}
}

// WithPayload adds the payload to the listing service unavailable response
func (o *ListingServiceUnavailable) WithPayload(payload *models.ErrorResponse) *ListingServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listing service unavailable response
func (o *ListingServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListingServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
