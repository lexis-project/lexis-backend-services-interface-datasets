// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// CloudNFSExportRemoveCreatedCode is the HTTP code returned for type CloudNFSExportRemoveCreated
const CloudNFSExportRemoveCreatedCode int = 201

/*CloudNFSExportRemoveCreated The response code means that the request for removal has been initiated. Status of the removal process can be checked by querying the status.

swagger:response cloudNFSExportRemoveCreated
*/
type CloudNFSExportRemoveCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RequestID `json:"body,omitempty"`
}

// NewCloudNFSExportRemoveCreated creates CloudNFSExportRemoveCreated with default headers values
func NewCloudNFSExportRemoveCreated() *CloudNFSExportRemoveCreated {

	return &CloudNFSExportRemoveCreated{}
}

// WithPayload adds the payload to the cloud n f s export remove created response
func (o *CloudNFSExportRemoveCreated) WithPayload(payload *models.RequestID) *CloudNFSExportRemoveCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cloud n f s export remove created response
func (o *CloudNFSExportRemoveCreated) SetPayload(payload *models.RequestID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloudNFSExportRemoveCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CloudNFSExportRemoveBadRequestCode is the HTTP code returned for type CloudNFSExportRemoveBadRequest
const CloudNFSExportRemoveBadRequestCode int = 400

/*CloudNFSExportRemoveBadRequest This means that the request ID given by the user is incorrect.

swagger:response cloudNFSExportRemoveBadRequest
*/
type CloudNFSExportRemoveBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCloudNFSExportRemoveBadRequest creates CloudNFSExportRemoveBadRequest with default headers values
func NewCloudNFSExportRemoveBadRequest() *CloudNFSExportRemoveBadRequest {

	return &CloudNFSExportRemoveBadRequest{}
}

// WithPayload adds the payload to the cloud n f s export remove bad request response
func (o *CloudNFSExportRemoveBadRequest) WithPayload(payload *models.ErrorResponse) *CloudNFSExportRemoveBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cloud n f s export remove bad request response
func (o *CloudNFSExportRemoveBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloudNFSExportRemoveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CloudNFSExportRemoveUnauthorizedCode is the HTTP code returned for type CloudNFSExportRemoveUnauthorized
const CloudNFSExportRemoveUnauthorizedCode int = 401

/*CloudNFSExportRemoveUnauthorized This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user

swagger:response cloudNFSExportRemoveUnauthorized
*/
type CloudNFSExportRemoveUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCloudNFSExportRemoveUnauthorized creates CloudNFSExportRemoveUnauthorized with default headers values
func NewCloudNFSExportRemoveUnauthorized() *CloudNFSExportRemoveUnauthorized {

	return &CloudNFSExportRemoveUnauthorized{}
}

// WithPayload adds the payload to the cloud n f s export remove unauthorized response
func (o *CloudNFSExportRemoveUnauthorized) WithPayload(payload *models.ErrorResponse) *CloudNFSExportRemoveUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cloud n f s export remove unauthorized response
func (o *CloudNFSExportRemoveUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloudNFSExportRemoveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CloudNFSExportRemoveRequestURITooLongCode is the HTTP code returned for type CloudNFSExportRemoveRequestURITooLong
const CloudNFSExportRemoveRequestURITooLongCode int = 414

/*CloudNFSExportRemoveRequestURITooLong This means that the the request ID is longer than the server is willing to interpret.

swagger:response cloudNFSExportRemoveRequestUriTooLong
*/
type CloudNFSExportRemoveRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCloudNFSExportRemoveRequestURITooLong creates CloudNFSExportRemoveRequestURITooLong with default headers values
func NewCloudNFSExportRemoveRequestURITooLong() *CloudNFSExportRemoveRequestURITooLong {

	return &CloudNFSExportRemoveRequestURITooLong{}
}

// WithPayload adds the payload to the cloud n f s export remove request Uri too long response
func (o *CloudNFSExportRemoveRequestURITooLong) WithPayload(payload *models.ErrorResponse) *CloudNFSExportRemoveRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cloud n f s export remove request Uri too long response
func (o *CloudNFSExportRemoveRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloudNFSExportRemoveRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CloudNFSExportRemoveInternalServerErrorCode is the HTTP code returned for type CloudNFSExportRemoveInternalServerError
const CloudNFSExportRemoveInternalServerErrorCode int = 500

/*CloudNFSExportRemoveInternalServerError This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.

swagger:response cloudNFSExportRemoveInternalServerError
*/
type CloudNFSExportRemoveInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCloudNFSExportRemoveInternalServerError creates CloudNFSExportRemoveInternalServerError with default headers values
func NewCloudNFSExportRemoveInternalServerError() *CloudNFSExportRemoveInternalServerError {

	return &CloudNFSExportRemoveInternalServerError{}
}

// WithPayload adds the payload to the cloud n f s export remove internal server error response
func (o *CloudNFSExportRemoveInternalServerError) WithPayload(payload *models.ErrorResponse) *CloudNFSExportRemoveInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cloud n f s export remove internal server error response
func (o *CloudNFSExportRemoveInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloudNFSExportRemoveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
