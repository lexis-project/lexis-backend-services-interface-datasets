// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// DeleteProjectUserNoContentCode is the HTTP code returned for type DeleteProjectUserNoContent
const DeleteProjectUserNoContentCode int = 204

/*DeleteProjectUserNoContent User was removed from project

swagger:response deleteProjectUserNoContent
*/
type DeleteProjectUserNoContent struct {
}

// NewDeleteProjectUserNoContent creates DeleteProjectUserNoContent with default headers values
func NewDeleteProjectUserNoContent() *DeleteProjectUserNoContent {

	return &DeleteProjectUserNoContent{}
}

// WriteResponse to the client
func (o *DeleteProjectUserNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteProjectUserBadRequestCode is the HTTP code returned for type DeleteProjectUserBadRequest
const DeleteProjectUserBadRequestCode int = 400

/*DeleteProjectUserBadRequest Malformed request

swagger:response deleteProjectUserBadRequest
*/
type DeleteProjectUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteProjectUserBadRequest creates DeleteProjectUserBadRequest with default headers values
func NewDeleteProjectUserBadRequest() *DeleteProjectUserBadRequest {

	return &DeleteProjectUserBadRequest{}
}

// WithPayload adds the payload to the delete project user bad request response
func (o *DeleteProjectUserBadRequest) WithPayload(payload *models.ErrorResponse) *DeleteProjectUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project user bad request response
func (o *DeleteProjectUserBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectUserUnauthorizedCode is the HTTP code returned for type DeleteProjectUserUnauthorized
const DeleteProjectUserUnauthorizedCode int = 401

/*DeleteProjectUserUnauthorized Unauthorized

swagger:response deleteProjectUserUnauthorized
*/
type DeleteProjectUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteProjectUserUnauthorized creates DeleteProjectUserUnauthorized with default headers values
func NewDeleteProjectUserUnauthorized() *DeleteProjectUserUnauthorized {

	return &DeleteProjectUserUnauthorized{}
}

// WithPayload adds the payload to the delete project user unauthorized response
func (o *DeleteProjectUserUnauthorized) WithPayload(payload *models.ErrorResponse) *DeleteProjectUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project user unauthorized response
func (o *DeleteProjectUserUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectUserConflictCode is the HTTP code returned for type DeleteProjectUserConflict
const DeleteProjectUserConflictCode int = 409

/*DeleteProjectUserConflict Conflict, action could not be performed

swagger:response deleteProjectUserConflict
*/
type DeleteProjectUserConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteProjectUserConflict creates DeleteProjectUserConflict with default headers values
func NewDeleteProjectUserConflict() *DeleteProjectUserConflict {

	return &DeleteProjectUserConflict{}
}

// WithPayload adds the payload to the delete project user conflict response
func (o *DeleteProjectUserConflict) WithPayload(payload *models.ErrorResponse) *DeleteProjectUserConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project user conflict response
func (o *DeleteProjectUserConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectUserInternalServerErrorCode is the HTTP code returned for type DeleteProjectUserInternalServerError
const DeleteProjectUserInternalServerErrorCode int = 500

/*DeleteProjectUserInternalServerError Internal error processing request

swagger:response deleteProjectUserInternalServerError
*/
type DeleteProjectUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteProjectUserInternalServerError creates DeleteProjectUserInternalServerError with default headers values
func NewDeleteProjectUserInternalServerError() *DeleteProjectUserInternalServerError {

	return &DeleteProjectUserInternalServerError{}
}

// WithPayload adds the payload to the delete project user internal server error response
func (o *DeleteProjectUserInternalServerError) WithPayload(payload *models.ErrorResponse) *DeleteProjectUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project user internal server error response
func (o *DeleteProjectUserInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectUserBadGatewayCode is the HTTP code returned for type DeleteProjectUserBadGateway
const DeleteProjectUserBadGatewayCode int = 502

/*DeleteProjectUserBadGateway Bad Gateway

swagger:response deleteProjectUserBadGateway
*/
type DeleteProjectUserBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteProjectUserBadGateway creates DeleteProjectUserBadGateway with default headers values
func NewDeleteProjectUserBadGateway() *DeleteProjectUserBadGateway {

	return &DeleteProjectUserBadGateway{}
}

// WithPayload adds the payload to the delete project user bad gateway response
func (o *DeleteProjectUserBadGateway) WithPayload(payload *models.ErrorResponse) *DeleteProjectUserBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project user bad gateway response
func (o *DeleteProjectUserBadGateway) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectUserBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}