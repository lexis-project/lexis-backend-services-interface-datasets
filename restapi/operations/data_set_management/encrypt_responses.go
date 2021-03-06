// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// EncryptCreatedCode is the HTTP code returned for type EncryptCreated
const EncryptCreatedCode int = 201

/*EncryptCreated The response code means that the encryption has been initiated. Status of the operation can be checked by querying the status.

swagger:response encryptCreated
*/
type EncryptCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RequestID `json:"body,omitempty"`
}

// NewEncryptCreated creates EncryptCreated with default headers values
func NewEncryptCreated() *EncryptCreated {

	return &EncryptCreated{}
}

// WithPayload adds the payload to the encrypt created response
func (o *EncryptCreated) WithPayload(payload *models.RequestID) *EncryptCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the encrypt created response
func (o *EncryptCreated) SetPayload(payload *models.RequestID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EncryptCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EncryptBadRequestCode is the HTTP code returned for type EncryptBadRequest
const EncryptBadRequestCode int = 400

/*EncryptBadRequest This means that there's something wrong in the input parameters and the server couldn't understand the request.

swagger:response encryptBadRequest
*/
type EncryptBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewEncryptBadRequest creates EncryptBadRequest with default headers values
func NewEncryptBadRequest() *EncryptBadRequest {

	return &EncryptBadRequest{}
}

// WithPayload adds the payload to the encrypt bad request response
func (o *EncryptBadRequest) WithPayload(payload *models.ErrorResponse) *EncryptBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the encrypt bad request response
func (o *EncryptBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EncryptBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EncryptUnauthorizedCode is the HTTP code returned for type EncryptUnauthorized
const EncryptUnauthorizedCode int = 401

/*EncryptUnauthorized This means that the user is not authenticated with keycloak and encryption can't be triggered unless the user first log in with a valid user

swagger:response encryptUnauthorized
*/
type EncryptUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewEncryptUnauthorized creates EncryptUnauthorized with default headers values
func NewEncryptUnauthorized() *EncryptUnauthorized {

	return &EncryptUnauthorized{}
}

// WithPayload adds the payload to the encrypt unauthorized response
func (o *EncryptUnauthorized) WithPayload(payload *models.ErrorResponse) *EncryptUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the encrypt unauthorized response
func (o *EncryptUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EncryptUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EncryptForbiddenCode is the HTTP code returned for type EncryptForbidden
const EncryptForbiddenCode int = 403

/*EncryptForbidden This means that the resource the user is trying to encrypt from or to is not readable. User doesn't have the correct rights to read the source file.

swagger:response encryptForbidden
*/
type EncryptForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewEncryptForbidden creates EncryptForbidden with default headers values
func NewEncryptForbidden() *EncryptForbidden {

	return &EncryptForbidden{}
}

// WithPayload adds the payload to the encrypt forbidden response
func (o *EncryptForbidden) WithPayload(payload *models.ErrorResponse) *EncryptForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the encrypt forbidden response
func (o *EncryptForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EncryptForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EncryptNotFoundCode is the HTTP code returned for type EncryptNotFound
const EncryptNotFoundCode int = 404

/*EncryptNotFound This means that the source path on the system doesn't exist.

swagger:response encryptNotFound
*/
type EncryptNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewEncryptNotFound creates EncryptNotFound with default headers values
func NewEncryptNotFound() *EncryptNotFound {

	return &EncryptNotFound{}
}

// WithPayload adds the payload to the encrypt not found response
func (o *EncryptNotFound) WithPayload(payload *models.ErrorResponse) *EncryptNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the encrypt not found response
func (o *EncryptNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EncryptNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EncryptRequestURITooLongCode is the HTTP code returned for type EncryptRequestURITooLong
const EncryptRequestURITooLongCode int = 414

/*EncryptRequestURITooLong This means that the source path is longer than the server is willing to interpret.

swagger:response encryptRequestUriTooLong
*/
type EncryptRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewEncryptRequestURITooLong creates EncryptRequestURITooLong with default headers values
func NewEncryptRequestURITooLong() *EncryptRequestURITooLong {

	return &EncryptRequestURITooLong{}
}

// WithPayload adds the payload to the encrypt request Uri too long response
func (o *EncryptRequestURITooLong) WithPayload(payload *models.ErrorResponse) *EncryptRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the encrypt request Uri too long response
func (o *EncryptRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EncryptRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EncryptTooManyRequestsCode is the HTTP code returned for type EncryptTooManyRequests
const EncryptTooManyRequestsCode int = 429

/*EncryptTooManyRequests This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.

swagger:response encryptTooManyRequests
*/
type EncryptTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewEncryptTooManyRequests creates EncryptTooManyRequests with default headers values
func NewEncryptTooManyRequests() *EncryptTooManyRequests {

	return &EncryptTooManyRequests{}
}

// WithPayload adds the payload to the encrypt too many requests response
func (o *EncryptTooManyRequests) WithPayload(payload *models.ErrorResponse) *EncryptTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the encrypt too many requests response
func (o *EncryptTooManyRequests) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EncryptTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EncryptInternalServerErrorCode is the HTTP code returned for type EncryptInternalServerError
const EncryptInternalServerErrorCode int = 500

/*EncryptInternalServerError This means that something has gone wrong on the burst buffer. The user is advised to wait and send the request again.

swagger:response encryptInternalServerError
*/
type EncryptInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewEncryptInternalServerError creates EncryptInternalServerError with default headers values
func NewEncryptInternalServerError() *EncryptInternalServerError {

	return &EncryptInternalServerError{}
}

// WithPayload adds the payload to the encrypt internal server error response
func (o *EncryptInternalServerError) WithPayload(payload *models.ErrorResponse) *EncryptInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the encrypt internal server error response
func (o *EncryptInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EncryptInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
