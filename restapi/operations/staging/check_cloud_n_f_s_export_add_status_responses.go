// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// CheckCloudNFSExportAddStatusOKCode is the HTTP code returned for type CheckCloudNFSExportAddStatusOK
const CheckCloudNFSExportAddStatusOKCode int = 200

/*CheckCloudNFSExportAddStatusOK This means that the status has been returned to the user in the response body.

swagger:response checkCloudNFSExportAddStatusOK
*/
type CheckCloudNFSExportAddStatusOK struct {

	/*
	  In: Body
	*/
	Payload *CheckCloudNFSExportAddStatusOKBody `json:"body,omitempty"`
}

// NewCheckCloudNFSExportAddStatusOK creates CheckCloudNFSExportAddStatusOK with default headers values
func NewCheckCloudNFSExportAddStatusOK() *CheckCloudNFSExportAddStatusOK {

	return &CheckCloudNFSExportAddStatusOK{}
}

// WithPayload adds the payload to the check cloud n f s export add status o k response
func (o *CheckCloudNFSExportAddStatusOK) WithPayload(payload *CheckCloudNFSExportAddStatusOKBody) *CheckCloudNFSExportAddStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check cloud n f s export add status o k response
func (o *CheckCloudNFSExportAddStatusOK) SetPayload(payload *CheckCloudNFSExportAddStatusOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCloudNFSExportAddStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCloudNFSExportAddStatusBadRequestCode is the HTTP code returned for type CheckCloudNFSExportAddStatusBadRequest
const CheckCloudNFSExportAddStatusBadRequestCode int = 400

/*CheckCloudNFSExportAddStatusBadRequest This means that the request ID given by the user is incorrect.

swagger:response checkCloudNFSExportAddStatusBadRequest
*/
type CheckCloudNFSExportAddStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCloudNFSExportAddStatusBadRequest creates CheckCloudNFSExportAddStatusBadRequest with default headers values
func NewCheckCloudNFSExportAddStatusBadRequest() *CheckCloudNFSExportAddStatusBadRequest {

	return &CheckCloudNFSExportAddStatusBadRequest{}
}

// WithPayload adds the payload to the check cloud n f s export add status bad request response
func (o *CheckCloudNFSExportAddStatusBadRequest) WithPayload(payload *models.ErrorResponse) *CheckCloudNFSExportAddStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check cloud n f s export add status bad request response
func (o *CheckCloudNFSExportAddStatusBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCloudNFSExportAddStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCloudNFSExportAddStatusUnauthorizedCode is the HTTP code returned for type CheckCloudNFSExportAddStatusUnauthorized
const CheckCloudNFSExportAddStatusUnauthorizedCode int = 401

/*CheckCloudNFSExportAddStatusUnauthorized This means that the user is not authenticated with keycloak and NFS export can't be triggered unless the user first log in with a valid user

swagger:response checkCloudNFSExportAddStatusUnauthorized
*/
type CheckCloudNFSExportAddStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCloudNFSExportAddStatusUnauthorized creates CheckCloudNFSExportAddStatusUnauthorized with default headers values
func NewCheckCloudNFSExportAddStatusUnauthorized() *CheckCloudNFSExportAddStatusUnauthorized {

	return &CheckCloudNFSExportAddStatusUnauthorized{}
}

// WithPayload adds the payload to the check cloud n f s export add status unauthorized response
func (o *CheckCloudNFSExportAddStatusUnauthorized) WithPayload(payload *models.ErrorResponse) *CheckCloudNFSExportAddStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check cloud n f s export add status unauthorized response
func (o *CheckCloudNFSExportAddStatusUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCloudNFSExportAddStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCloudNFSExportAddStatusForbiddenCode is the HTTP code returned for type CheckCloudNFSExportAddStatusForbidden
const CheckCloudNFSExportAddStatusForbiddenCode int = 403

/*CheckCloudNFSExportAddStatusForbidden The IP is outside of the allowed range

swagger:response checkCloudNFSExportAddStatusForbidden
*/
type CheckCloudNFSExportAddStatusForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCloudNFSExportAddStatusForbidden creates CheckCloudNFSExportAddStatusForbidden with default headers values
func NewCheckCloudNFSExportAddStatusForbidden() *CheckCloudNFSExportAddStatusForbidden {

	return &CheckCloudNFSExportAddStatusForbidden{}
}

// WithPayload adds the payload to the check cloud n f s export add status forbidden response
func (o *CheckCloudNFSExportAddStatusForbidden) WithPayload(payload *models.ErrorResponse) *CheckCloudNFSExportAddStatusForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check cloud n f s export add status forbidden response
func (o *CheckCloudNFSExportAddStatusForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCloudNFSExportAddStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCloudNFSExportAddStatusNotFoundCode is the HTTP code returned for type CheckCloudNFSExportAddStatusNotFound
const CheckCloudNFSExportAddStatusNotFoundCode int = 404

/*CheckCloudNFSExportAddStatusNotFound This means that the ID doesn't exist and thus a status can't be returned.

swagger:response checkCloudNFSExportAddStatusNotFound
*/
type CheckCloudNFSExportAddStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCloudNFSExportAddStatusNotFound creates CheckCloudNFSExportAddStatusNotFound with default headers values
func NewCheckCloudNFSExportAddStatusNotFound() *CheckCloudNFSExportAddStatusNotFound {

	return &CheckCloudNFSExportAddStatusNotFound{}
}

// WithPayload adds the payload to the check cloud n f s export add status not found response
func (o *CheckCloudNFSExportAddStatusNotFound) WithPayload(payload *models.ErrorResponse) *CheckCloudNFSExportAddStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check cloud n f s export add status not found response
func (o *CheckCloudNFSExportAddStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCloudNFSExportAddStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCloudNFSExportAddStatusRequestURITooLongCode is the HTTP code returned for type CheckCloudNFSExportAddStatusRequestURITooLong
const CheckCloudNFSExportAddStatusRequestURITooLongCode int = 414

/*CheckCloudNFSExportAddStatusRequestURITooLong This means that the the request ID is longer than the server is willing to interpret.

swagger:response checkCloudNFSExportAddStatusRequestUriTooLong
*/
type CheckCloudNFSExportAddStatusRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCloudNFSExportAddStatusRequestURITooLong creates CheckCloudNFSExportAddStatusRequestURITooLong with default headers values
func NewCheckCloudNFSExportAddStatusRequestURITooLong() *CheckCloudNFSExportAddStatusRequestURITooLong {

	return &CheckCloudNFSExportAddStatusRequestURITooLong{}
}

// WithPayload adds the payload to the check cloud n f s export add status request Uri too long response
func (o *CheckCloudNFSExportAddStatusRequestURITooLong) WithPayload(payload *models.ErrorResponse) *CheckCloudNFSExportAddStatusRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check cloud n f s export add status request Uri too long response
func (o *CheckCloudNFSExportAddStatusRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCloudNFSExportAddStatusRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCloudNFSExportAddStatusInternalServerErrorCode is the HTTP code returned for type CheckCloudNFSExportAddStatusInternalServerError
const CheckCloudNFSExportAddStatusInternalServerErrorCode int = 500

/*CheckCloudNFSExportAddStatusInternalServerError This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.

swagger:response checkCloudNFSExportAddStatusInternalServerError
*/
type CheckCloudNFSExportAddStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCloudNFSExportAddStatusInternalServerError creates CheckCloudNFSExportAddStatusInternalServerError with default headers values
func NewCheckCloudNFSExportAddStatusInternalServerError() *CheckCloudNFSExportAddStatusInternalServerError {

	return &CheckCloudNFSExportAddStatusInternalServerError{}
}

// WithPayload adds the payload to the check cloud n f s export add status internal server error response
func (o *CheckCloudNFSExportAddStatusInternalServerError) WithPayload(payload *models.ErrorResponse) *CheckCloudNFSExportAddStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check cloud n f s export add status internal server error response
func (o *CheckCloudNFSExportAddStatusInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCloudNFSExportAddStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}