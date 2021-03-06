// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// CheckStageStatusOKCode is the HTTP code returned for type CheckStageStatusOK
const CheckStageStatusOKCode int = 200

/*CheckStageStatusOK This means that the status has been returned to the user in the response body.

swagger:response checkStageStatusOK
*/
type CheckStageStatusOK struct {

	/*
	  In: Body
	*/
	Payload *CheckStageStatusOKBody `json:"body,omitempty"`
}

// NewCheckStageStatusOK creates CheckStageStatusOK with default headers values
func NewCheckStageStatusOK() *CheckStageStatusOK {

	return &CheckStageStatusOK{}
}

// WithPayload adds the payload to the check stage status o k response
func (o *CheckStageStatusOK) WithPayload(payload *CheckStageStatusOKBody) *CheckStageStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check stage status o k response
func (o *CheckStageStatusOK) SetPayload(payload *CheckStageStatusOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckStageStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckStageStatusBadRequestCode is the HTTP code returned for type CheckStageStatusBadRequest
const CheckStageStatusBadRequestCode int = 400

/*CheckStageStatusBadRequest This means that the request ID given by the user is incorrect.

swagger:response checkStageStatusBadRequest
*/
type CheckStageStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckStageStatusBadRequest creates CheckStageStatusBadRequest with default headers values
func NewCheckStageStatusBadRequest() *CheckStageStatusBadRequest {

	return &CheckStageStatusBadRequest{}
}

// WithPayload adds the payload to the check stage status bad request response
func (o *CheckStageStatusBadRequest) WithPayload(payload *models.ErrorResponse) *CheckStageStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check stage status bad request response
func (o *CheckStageStatusBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckStageStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckStageStatusUnauthorizedCode is the HTTP code returned for type CheckStageStatusUnauthorized
const CheckStageStatusUnauthorizedCode int = 401

/*CheckStageStatusUnauthorized This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user

swagger:response checkStageStatusUnauthorized
*/
type CheckStageStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckStageStatusUnauthorized creates CheckStageStatusUnauthorized with default headers values
func NewCheckStageStatusUnauthorized() *CheckStageStatusUnauthorized {

	return &CheckStageStatusUnauthorized{}
}

// WithPayload adds the payload to the check stage status unauthorized response
func (o *CheckStageStatusUnauthorized) WithPayload(payload *models.ErrorResponse) *CheckStageStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check stage status unauthorized response
func (o *CheckStageStatusUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckStageStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckStageStatusNotFoundCode is the HTTP code returned for type CheckStageStatusNotFound
const CheckStageStatusNotFoundCode int = 404

/*CheckStageStatusNotFound This means that the ID doesn't exist and thus a status can't be returned.

swagger:response checkStageStatusNotFound
*/
type CheckStageStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckStageStatusNotFound creates CheckStageStatusNotFound with default headers values
func NewCheckStageStatusNotFound() *CheckStageStatusNotFound {

	return &CheckStageStatusNotFound{}
}

// WithPayload adds the payload to the check stage status not found response
func (o *CheckStageStatusNotFound) WithPayload(payload *models.ErrorResponse) *CheckStageStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check stage status not found response
func (o *CheckStageStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckStageStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckStageStatusRequestURITooLongCode is the HTTP code returned for type CheckStageStatusRequestURITooLong
const CheckStageStatusRequestURITooLongCode int = 414

/*CheckStageStatusRequestURITooLong This means that the the request ID is longer than the server is willing to interpret.

swagger:response checkStageStatusRequestUriTooLong
*/
type CheckStageStatusRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckStageStatusRequestURITooLong creates CheckStageStatusRequestURITooLong with default headers values
func NewCheckStageStatusRequestURITooLong() *CheckStageStatusRequestURITooLong {

	return &CheckStageStatusRequestURITooLong{}
}

// WithPayload adds the payload to the check stage status request Uri too long response
func (o *CheckStageStatusRequestURITooLong) WithPayload(payload *models.ErrorResponse) *CheckStageStatusRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check stage status request Uri too long response
func (o *CheckStageStatusRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckStageStatusRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckStageStatusInternalServerErrorCode is the HTTP code returned for type CheckStageStatusInternalServerError
const CheckStageStatusInternalServerErrorCode int = 500

/*CheckStageStatusInternalServerError This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.

swagger:response checkStageStatusInternalServerError
*/
type CheckStageStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckStageStatusInternalServerError creates CheckStageStatusInternalServerError with default headers values
func NewCheckStageStatusInternalServerError() *CheckStageStatusInternalServerError {

	return &CheckStageStatusInternalServerError{}
}

// WithPayload adds the payload to the check stage status internal server error response
func (o *CheckStageStatusInternalServerError) WithPayload(payload *models.ErrorResponse) *CheckStageStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check stage status internal server error response
func (o *CheckStageStatusInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckStageStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
