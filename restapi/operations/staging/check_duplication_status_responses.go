// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// CheckDuplicationStatusOKCode is the HTTP code returned for type CheckDuplicationStatusOK
const CheckDuplicationStatusOKCode int = 200

/*CheckDuplicationStatusOK This means that the status has been returned to the user in the response body.

swagger:response checkDuplicationStatusOK
*/
type CheckDuplicationStatusOK struct {

	/*
	  In: Body
	*/
	Payload *CheckDuplicationStatusOKBody `json:"body,omitempty"`
}

// NewCheckDuplicationStatusOK creates CheckDuplicationStatusOK with default headers values
func NewCheckDuplicationStatusOK() *CheckDuplicationStatusOK {

	return &CheckDuplicationStatusOK{}
}

// WithPayload adds the payload to the check duplication status o k response
func (o *CheckDuplicationStatusOK) WithPayload(payload *CheckDuplicationStatusOKBody) *CheckDuplicationStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check duplication status o k response
func (o *CheckDuplicationStatusOK) SetPayload(payload *CheckDuplicationStatusOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDuplicationStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDuplicationStatusBadRequestCode is the HTTP code returned for type CheckDuplicationStatusBadRequest
const CheckDuplicationStatusBadRequestCode int = 400

/*CheckDuplicationStatusBadRequest This means that the request ID given by the user is incorrect.

swagger:response checkDuplicationStatusBadRequest
*/
type CheckDuplicationStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDuplicationStatusBadRequest creates CheckDuplicationStatusBadRequest with default headers values
func NewCheckDuplicationStatusBadRequest() *CheckDuplicationStatusBadRequest {

	return &CheckDuplicationStatusBadRequest{}
}

// WithPayload adds the payload to the check duplication status bad request response
func (o *CheckDuplicationStatusBadRequest) WithPayload(payload *models.ErrorResponse) *CheckDuplicationStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check duplication status bad request response
func (o *CheckDuplicationStatusBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDuplicationStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDuplicationStatusUnauthorizedCode is the HTTP code returned for type CheckDuplicationStatusUnauthorized
const CheckDuplicationStatusUnauthorizedCode int = 401

/*CheckDuplicationStatusUnauthorized This means that the user is not authenticated with keycloak and duplication can't be triggered unless the user first log in with a valid user

swagger:response checkDuplicationStatusUnauthorized
*/
type CheckDuplicationStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDuplicationStatusUnauthorized creates CheckDuplicationStatusUnauthorized with default headers values
func NewCheckDuplicationStatusUnauthorized() *CheckDuplicationStatusUnauthorized {

	return &CheckDuplicationStatusUnauthorized{}
}

// WithPayload adds the payload to the check duplication status unauthorized response
func (o *CheckDuplicationStatusUnauthorized) WithPayload(payload *models.ErrorResponse) *CheckDuplicationStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check duplication status unauthorized response
func (o *CheckDuplicationStatusUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDuplicationStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDuplicationStatusNotFoundCode is the HTTP code returned for type CheckDuplicationStatusNotFound
const CheckDuplicationStatusNotFoundCode int = 404

/*CheckDuplicationStatusNotFound This means that the ID doesn't exist and thus a status can't be returned.

swagger:response checkDuplicationStatusNotFound
*/
type CheckDuplicationStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDuplicationStatusNotFound creates CheckDuplicationStatusNotFound with default headers values
func NewCheckDuplicationStatusNotFound() *CheckDuplicationStatusNotFound {

	return &CheckDuplicationStatusNotFound{}
}

// WithPayload adds the payload to the check duplication status not found response
func (o *CheckDuplicationStatusNotFound) WithPayload(payload *models.ErrorResponse) *CheckDuplicationStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check duplication status not found response
func (o *CheckDuplicationStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDuplicationStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDuplicationStatusRequestURITooLongCode is the HTTP code returned for type CheckDuplicationStatusRequestURITooLong
const CheckDuplicationStatusRequestURITooLongCode int = 414

/*CheckDuplicationStatusRequestURITooLong This means that the the request ID is longer than the server is willing to interpret.

swagger:response checkDuplicationStatusRequestUriTooLong
*/
type CheckDuplicationStatusRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDuplicationStatusRequestURITooLong creates CheckDuplicationStatusRequestURITooLong with default headers values
func NewCheckDuplicationStatusRequestURITooLong() *CheckDuplicationStatusRequestURITooLong {

	return &CheckDuplicationStatusRequestURITooLong{}
}

// WithPayload adds the payload to the check duplication status request Uri too long response
func (o *CheckDuplicationStatusRequestURITooLong) WithPayload(payload *models.ErrorResponse) *CheckDuplicationStatusRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check duplication status request Uri too long response
func (o *CheckDuplicationStatusRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDuplicationStatusRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDuplicationStatusInternalServerErrorCode is the HTTP code returned for type CheckDuplicationStatusInternalServerError
const CheckDuplicationStatusInternalServerErrorCode int = 500

/*CheckDuplicationStatusInternalServerError This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.

swagger:response checkDuplicationStatusInternalServerError
*/
type CheckDuplicationStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDuplicationStatusInternalServerError creates CheckDuplicationStatusInternalServerError with default headers values
func NewCheckDuplicationStatusInternalServerError() *CheckDuplicationStatusInternalServerError {

	return &CheckDuplicationStatusInternalServerError{}
}

// WithPayload adds the payload to the check duplication status internal server error response
func (o *CheckDuplicationStatusInternalServerError) WithPayload(payload *models.ErrorResponse) *CheckDuplicationStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check duplication status internal server error response
func (o *CheckDuplicationStatusInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDuplicationStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
