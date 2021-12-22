// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// AddGridmapEntryCreatedCode is the HTTP code returned for type AddGridmapEntryCreated
const AddGridmapEntryCreatedCode int = 201

/*AddGridmapEntryCreated DN added

swagger:response addGridmapEntryCreated
*/
type AddGridmapEntryCreated struct {
}

// NewAddGridmapEntryCreated creates AddGridmapEntryCreated with default headers values
func NewAddGridmapEntryCreated() *AddGridmapEntryCreated {

	return &AddGridmapEntryCreated{}
}

// WriteResponse to the client
func (o *AddGridmapEntryCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddGridmapEntryUnauthorizedCode is the HTTP code returned for type AddGridmapEntryUnauthorized
const AddGridmapEntryUnauthorizedCode int = 401

/*AddGridmapEntryUnauthorized Unauthorized

swagger:response addGridmapEntryUnauthorized
*/
type AddGridmapEntryUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddGridmapEntryUnauthorized creates AddGridmapEntryUnauthorized with default headers values
func NewAddGridmapEntryUnauthorized() *AddGridmapEntryUnauthorized {

	return &AddGridmapEntryUnauthorized{}
}

// WithPayload adds the payload to the add gridmap entry unauthorized response
func (o *AddGridmapEntryUnauthorized) WithPayload(payload *models.ErrorResponse) *AddGridmapEntryUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add gridmap entry unauthorized response
func (o *AddGridmapEntryUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGridmapEntryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGridmapEntryForbiddenCode is the HTTP code returned for type AddGridmapEntryForbidden
const AddGridmapEntryForbiddenCode int = 403

/*AddGridmapEntryForbidden Forbidden

swagger:response addGridmapEntryForbidden
*/
type AddGridmapEntryForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddGridmapEntryForbidden creates AddGridmapEntryForbidden with default headers values
func NewAddGridmapEntryForbidden() *AddGridmapEntryForbidden {

	return &AddGridmapEntryForbidden{}
}

// WithPayload adds the payload to the add gridmap entry forbidden response
func (o *AddGridmapEntryForbidden) WithPayload(payload *models.ErrorResponse) *AddGridmapEntryForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add gridmap entry forbidden response
func (o *AddGridmapEntryForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGridmapEntryForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGridmapEntryBadGatewayCode is the HTTP code returned for type AddGridmapEntryBadGateway
const AddGridmapEntryBadGatewayCode int = 502

/*AddGridmapEntryBadGateway Bad Gateway

swagger:response addGridmapEntryBadGateway
*/
type AddGridmapEntryBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddGridmapEntryBadGateway creates AddGridmapEntryBadGateway with default headers values
func NewAddGridmapEntryBadGateway() *AddGridmapEntryBadGateway {

	return &AddGridmapEntryBadGateway{}
}

// WithPayload adds the payload to the add gridmap entry bad gateway response
func (o *AddGridmapEntryBadGateway) WithPayload(payload *models.ErrorResponse) *AddGridmapEntryBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add gridmap entry bad gateway response
func (o *AddGridmapEntryBadGateway) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGridmapEntryBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGridmapEntryServiceUnavailableCode is the HTTP code returned for type AddGridmapEntryServiceUnavailable
const AddGridmapEntryServiceUnavailableCode int = 503

/*AddGridmapEntryServiceUnavailable Server error

swagger:response addGridmapEntryServiceUnavailable
*/
type AddGridmapEntryServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddGridmapEntryServiceUnavailable creates AddGridmapEntryServiceUnavailable with default headers values
func NewAddGridmapEntryServiceUnavailable() *AddGridmapEntryServiceUnavailable {

	return &AddGridmapEntryServiceUnavailable{}
}

// WithPayload adds the payload to the add gridmap entry service unavailable response
func (o *AddGridmapEntryServiceUnavailable) WithPayload(payload *models.ErrorResponse) *AddGridmapEntryServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add gridmap entry service unavailable response
func (o *AddGridmapEntryServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGridmapEntryServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
