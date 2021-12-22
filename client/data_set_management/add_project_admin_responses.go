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

// AddProjectAdminReader is a Reader for the AddProjectAdmin structure.
type AddProjectAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProjectAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddProjectAdminCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddProjectAdminBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddProjectAdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddProjectAdminConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddProjectAdminInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewAddProjectAdminBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddProjectAdminCreated creates a AddProjectAdminCreated with default headers values
func NewAddProjectAdminCreated() *AddProjectAdminCreated {
	return &AddProjectAdminCreated{}
}

/*AddProjectAdminCreated handles this case with default header values.

Admin was added to Project
*/
type AddProjectAdminCreated struct {
}

func (o *AddProjectAdminCreated) Error() string {
	return fmt.Sprintf("[POST /project/admin][%d] addProjectAdminCreated ", 201)
}

func (o *AddProjectAdminCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddProjectAdminBadRequest creates a AddProjectAdminBadRequest with default headers values
func NewAddProjectAdminBadRequest() *AddProjectAdminBadRequest {
	return &AddProjectAdminBadRequest{}
}

/*AddProjectAdminBadRequest handles this case with default header values.

Malformed request
*/
type AddProjectAdminBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *AddProjectAdminBadRequest) Error() string {
	return fmt.Sprintf("[POST /project/admin][%d] addProjectAdminBadRequest  %+v", 400, o.Payload)
}

func (o *AddProjectAdminBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddProjectAdminBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectAdminUnauthorized creates a AddProjectAdminUnauthorized with default headers values
func NewAddProjectAdminUnauthorized() *AddProjectAdminUnauthorized {
	return &AddProjectAdminUnauthorized{}
}

/*AddProjectAdminUnauthorized handles this case with default header values.

Unauthorized
*/
type AddProjectAdminUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *AddProjectAdminUnauthorized) Error() string {
	return fmt.Sprintf("[POST /project/admin][%d] addProjectAdminUnauthorized  %+v", 401, o.Payload)
}

func (o *AddProjectAdminUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddProjectAdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectAdminConflict creates a AddProjectAdminConflict with default headers values
func NewAddProjectAdminConflict() *AddProjectAdminConflict {
	return &AddProjectAdminConflict{}
}

/*AddProjectAdminConflict handles this case with default header values.

Conflict, action could not be performed
*/
type AddProjectAdminConflict struct {
	Payload *models.ErrorResponse
}

func (o *AddProjectAdminConflict) Error() string {
	return fmt.Sprintf("[POST /project/admin][%d] addProjectAdminConflict  %+v", 409, o.Payload)
}

func (o *AddProjectAdminConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddProjectAdminConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectAdminInternalServerError creates a AddProjectAdminInternalServerError with default headers values
func NewAddProjectAdminInternalServerError() *AddProjectAdminInternalServerError {
	return &AddProjectAdminInternalServerError{}
}

/*AddProjectAdminInternalServerError handles this case with default header values.

Internal error processing request
*/
type AddProjectAdminInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *AddProjectAdminInternalServerError) Error() string {
	return fmt.Sprintf("[POST /project/admin][%d] addProjectAdminInternalServerError  %+v", 500, o.Payload)
}

func (o *AddProjectAdminInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddProjectAdminInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectAdminBadGateway creates a AddProjectAdminBadGateway with default headers values
func NewAddProjectAdminBadGateway() *AddProjectAdminBadGateway {
	return &AddProjectAdminBadGateway{}
}

/*AddProjectAdminBadGateway handles this case with default header values.

Bad Gateway
*/
type AddProjectAdminBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *AddProjectAdminBadGateway) Error() string {
	return fmt.Sprintf("[POST /project/admin][%d] addProjectAdminBadGateway  %+v", 502, o.Payload)
}

func (o *AddProjectAdminBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddProjectAdminBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}