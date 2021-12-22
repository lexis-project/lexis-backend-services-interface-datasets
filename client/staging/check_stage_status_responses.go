// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// CheckStageStatusReader is a Reader for the CheckStageStatus structure.
type CheckStageStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckStageStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckStageStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckStageStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckStageStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckStageStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewCheckStageStatusRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckStageStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckStageStatusOK creates a CheckStageStatusOK with default headers values
func NewCheckStageStatusOK() *CheckStageStatusOK {
	return &CheckStageStatusOK{}
}

/*CheckStageStatusOK handles this case with default header values.

This means that the status has been returned to the user in the response body.
*/
type CheckStageStatusOK struct {
	Payload *CheckStageStatusOKBody
}

func (o *CheckStageStatusOK) Error() string {
	return fmt.Sprintf("[GET /staging/stage/{request_id}][%d] checkStageStatusOK  %+v", 200, o.Payload)
}

func (o *CheckStageStatusOK) GetPayload() *CheckStageStatusOKBody {
	return o.Payload
}

func (o *CheckStageStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckStageStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckStageStatusBadRequest creates a CheckStageStatusBadRequest with default headers values
func NewCheckStageStatusBadRequest() *CheckStageStatusBadRequest {
	return &CheckStageStatusBadRequest{}
}

/*CheckStageStatusBadRequest handles this case with default header values.

This means that the request ID given by the user is incorrect.
*/
type CheckStageStatusBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CheckStageStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /staging/stage/{request_id}][%d] checkStageStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckStageStatusBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckStageStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckStageStatusUnauthorized creates a CheckStageStatusUnauthorized with default headers values
func NewCheckStageStatusUnauthorized() *CheckStageStatusUnauthorized {
	return &CheckStageStatusUnauthorized{}
}

/*CheckStageStatusUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user
*/
type CheckStageStatusUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckStageStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /staging/stage/{request_id}][%d] checkStageStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckStageStatusUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckStageStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckStageStatusNotFound creates a CheckStageStatusNotFound with default headers values
func NewCheckStageStatusNotFound() *CheckStageStatusNotFound {
	return &CheckStageStatusNotFound{}
}

/*CheckStageStatusNotFound handles this case with default header values.

This means that the ID doesn't exist and thus a status can't be returned.
*/
type CheckStageStatusNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CheckStageStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /staging/stage/{request_id}][%d] checkStageStatusNotFound  %+v", 404, o.Payload)
}

func (o *CheckStageStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckStageStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckStageStatusRequestURITooLong creates a CheckStageStatusRequestURITooLong with default headers values
func NewCheckStageStatusRequestURITooLong() *CheckStageStatusRequestURITooLong {
	return &CheckStageStatusRequestURITooLong{}
}

/*CheckStageStatusRequestURITooLong handles this case with default header values.

This means that the the request ID is longer than the server is willing to interpret.
*/
type CheckStageStatusRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *CheckStageStatusRequestURITooLong) Error() string {
	return fmt.Sprintf("[GET /staging/stage/{request_id}][%d] checkStageStatusRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *CheckStageStatusRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckStageStatusRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckStageStatusInternalServerError creates a CheckStageStatusInternalServerError with default headers values
func NewCheckStageStatusInternalServerError() *CheckStageStatusInternalServerError {
	return &CheckStageStatusInternalServerError{}
}

/*CheckStageStatusInternalServerError handles this case with default header values.

This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.
*/
type CheckStageStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CheckStageStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staging/stage/{request_id}][%d] checkStageStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckStageStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckStageStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckStageStatusOKBody check stage status o k body
swagger:model CheckStageStatusOKBody
*/
type CheckStageStatusOKBody struct {

	// status
	// Required: true
	Status *string `json:"status"`

	// Single path for staging endpoints
	TargetPath string `json:"target_path,omitempty"`
}

// Validate validates this check stage status o k body
func (o *CheckStageStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckStageStatusOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("checkStageStatusOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckStageStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckStageStatusOKBody) UnmarshalBinary(b []byte) error {
	var res CheckStageStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
