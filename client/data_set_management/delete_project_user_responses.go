// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// DeleteProjectUserReader is a Reader for the DeleteProjectUser structure.
type DeleteProjectUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProjectUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteProjectUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewDeleteProjectUserBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectUserNoContent creates a DeleteProjectUserNoContent with default headers values
func NewDeleteProjectUserNoContent() *DeleteProjectUserNoContent {
	return &DeleteProjectUserNoContent{}
}

/*DeleteProjectUserNoContent handles this case with default header values.

User was removed from project
*/
type DeleteProjectUserNoContent struct {
}

func (o *DeleteProjectUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /project/user][%d] deleteProjectUserNoContent ", 204)
}

func (o *DeleteProjectUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectUserBadRequest creates a DeleteProjectUserBadRequest with default headers values
func NewDeleteProjectUserBadRequest() *DeleteProjectUserBadRequest {
	return &DeleteProjectUserBadRequest{}
}

/*DeleteProjectUserBadRequest handles this case with default header values.

Malformed request
*/
type DeleteProjectUserBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /project/user][%d] deleteProjectUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteProjectUserBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectUserUnauthorized creates a DeleteProjectUserUnauthorized with default headers values
func NewDeleteProjectUserUnauthorized() *DeleteProjectUserUnauthorized {
	return &DeleteProjectUserUnauthorized{}
}

/*DeleteProjectUserUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteProjectUserUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /project/user][%d] deleteProjectUserUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteProjectUserUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectUserConflict creates a DeleteProjectUserConflict with default headers values
func NewDeleteProjectUserConflict() *DeleteProjectUserConflict {
	return &DeleteProjectUserConflict{}
}

/*DeleteProjectUserConflict handles this case with default header values.

Conflict, action could not be performed
*/
type DeleteProjectUserConflict struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectUserConflict) Error() string {
	return fmt.Sprintf("[DELETE /project/user][%d] deleteProjectUserConflict  %+v", 409, o.Payload)
}

func (o *DeleteProjectUserConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectUserInternalServerError creates a DeleteProjectUserInternalServerError with default headers values
func NewDeleteProjectUserInternalServerError() *DeleteProjectUserInternalServerError {
	return &DeleteProjectUserInternalServerError{}
}

/*DeleteProjectUserInternalServerError handles this case with default header values.

Internal error processing request
*/
type DeleteProjectUserInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /project/user][%d] deleteProjectUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteProjectUserInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectUserBadGateway creates a DeleteProjectUserBadGateway with default headers values
func NewDeleteProjectUserBadGateway() *DeleteProjectUserBadGateway {
	return &DeleteProjectUserBadGateway{}
}

/*DeleteProjectUserBadGateway handles this case with default header values.

Bad Gateway
*/
type DeleteProjectUserBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectUserBadGateway) Error() string {
	return fmt.Sprintf("[DELETE /project/user][%d] deleteProjectUserBadGateway  %+v", 502, o.Payload)
}

func (o *DeleteProjectUserBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectUserBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
