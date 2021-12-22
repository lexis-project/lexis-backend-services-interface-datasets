// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// DeleteCreatedCode is the HTTP code returned for type DeleteCreated
const DeleteCreatedCode int = 201

/*DeleteCreated The response code means that the data deletion has been initiated. Status of the deletion can be checked by querying the status.

swagger:response deleteCreated
*/
type DeleteCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RequestID `json:"body,omitempty"`
}

// NewDeleteCreated creates DeleteCreated with default headers values
func NewDeleteCreated() *DeleteCreated {

	return &DeleteCreated{}
}

// WithPayload adds the payload to the delete created response
func (o *DeleteCreated) WithPayload(payload *models.RequestID) *DeleteCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete created response
func (o *DeleteCreated) SetPayload(payload *models.RequestID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBadRequestCode is the HTTP code returned for type DeleteBadRequest
const DeleteBadRequestCode int = 400

/*DeleteBadRequest This means that there's something wrong in the input parameters and the server couldn't understand the request.

swagger:response deleteBadRequest
*/
type DeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteBadRequest creates DeleteBadRequest with default headers values
func NewDeleteBadRequest() *DeleteBadRequest {

	return &DeleteBadRequest{}
}

// WithPayload adds the payload to the delete bad request response
func (o *DeleteBadRequest) WithPayload(payload *models.ErrorResponse) *DeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete bad request response
func (o *DeleteBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUnauthorizedCode is the HTTP code returned for type DeleteUnauthorized
const DeleteUnauthorizedCode int = 401

/*DeleteUnauthorized This means that the user is not authenticated with keycloak and data deletion can't be triggered unless the user first log in with a valid user

swagger:response deleteUnauthorized
*/
type DeleteUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUnauthorized creates DeleteUnauthorized with default headers values
func NewDeleteUnauthorized() *DeleteUnauthorized {

	return &DeleteUnauthorized{}
}

// WithPayload adds the payload to the delete unauthorized response
func (o *DeleteUnauthorized) WithPayload(payload *models.ErrorResponse) *DeleteUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete unauthorized response
func (o *DeleteUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteForbiddenCode is the HTTP code returned for type DeleteForbidden
const DeleteForbiddenCode int = 403

/*DeleteForbidden This means that the resource the user is trying to delete from or to is not readable or writable by the user. User doesn't have the correct rights to delete the data on the target system location.

swagger:response deleteForbidden
*/
type DeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteForbidden creates DeleteForbidden with default headers values
func NewDeleteForbidden() *DeleteForbidden {

	return &DeleteForbidden{}
}

// WithPayload adds the payload to the delete forbidden response
func (o *DeleteForbidden) WithPayload(payload *models.ErrorResponse) *DeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete forbidden response
func (o *DeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNotFoundCode is the HTTP code returned for type DeleteNotFound
const DeleteNotFoundCode int = 404

/*DeleteNotFound This means that the the target path on the system doesn't exist.

swagger:response deleteNotFound
*/
type DeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteNotFound creates DeleteNotFound with default headers values
func NewDeleteNotFound() *DeleteNotFound {

	return &DeleteNotFound{}
}

// WithPayload adds the payload to the delete not found response
func (o *DeleteNotFound) WithPayload(payload *models.ErrorResponse) *DeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete not found response
func (o *DeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRequestURITooLongCode is the HTTP code returned for type DeleteRequestURITooLong
const DeleteRequestURITooLongCode int = 414

/*DeleteRequestURITooLong This means that the the target path is longer than the server is willing to interpret.

swagger:response deleteRequestUriTooLong
*/
type DeleteRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteRequestURITooLong creates DeleteRequestURITooLong with default headers values
func NewDeleteRequestURITooLong() *DeleteRequestURITooLong {

	return &DeleteRequestURITooLong{}
}

// WithPayload adds the payload to the delete request Uri too long response
func (o *DeleteRequestURITooLong) WithPayload(payload *models.ErrorResponse) *DeleteRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete request Uri too long response
func (o *DeleteRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTooManyRequestsCode is the HTTP code returned for type DeleteTooManyRequests
const DeleteTooManyRequestsCode int = 429

/*DeleteTooManyRequests This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.

swagger:response deleteTooManyRequests
*/
type DeleteTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteTooManyRequests creates DeleteTooManyRequests with default headers values
func NewDeleteTooManyRequests() *DeleteTooManyRequests {

	return &DeleteTooManyRequests{}
}

// WithPayload adds the payload to the delete too many requests response
func (o *DeleteTooManyRequests) WithPayload(payload *models.ErrorResponse) *DeleteTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete too many requests response
func (o *DeleteTooManyRequests) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteInternalServerErrorCode is the HTTP code returned for type DeleteInternalServerError
const DeleteInternalServerErrorCode int = 500

/*DeleteInternalServerError This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.

swagger:response deleteInternalServerError
*/
type DeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteInternalServerError creates DeleteInternalServerError with default headers values
func NewDeleteInternalServerError() *DeleteInternalServerError {

	return &DeleteInternalServerError{}
}

// WithPayload adds the payload to the delete internal server error response
func (o *DeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *DeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete internal server error response
func (o *DeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
