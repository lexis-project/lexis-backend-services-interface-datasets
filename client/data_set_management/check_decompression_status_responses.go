// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

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

// CheckDecompressionStatusReader is a Reader for the CheckDecompressionStatus structure.
type CheckDecompressionStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckDecompressionStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckDecompressionStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckDecompressionStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckDecompressionStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckDecompressionStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewCheckDecompressionStatusRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckDecompressionStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckDecompressionStatusOK creates a CheckDecompressionStatusOK with default headers values
func NewCheckDecompressionStatusOK() *CheckDecompressionStatusOK {
	return &CheckDecompressionStatusOK{}
}

/*CheckDecompressionStatusOK handles this case with default header values.

This means that the status has been returned to the user in the response body.
*/
type CheckDecompressionStatusOK struct {
	Payload *CheckDecompressionStatusOKBody
}

func (o *CheckDecompressionStatusOK) Error() string {
	return fmt.Sprintf("[GET /encryption/decompress/{request_id}][%d] checkDecompressionStatusOK  %+v", 200, o.Payload)
}

func (o *CheckDecompressionStatusOK) GetPayload() *CheckDecompressionStatusOKBody {
	return o.Payload
}

func (o *CheckDecompressionStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckDecompressionStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDecompressionStatusBadRequest creates a CheckDecompressionStatusBadRequest with default headers values
func NewCheckDecompressionStatusBadRequest() *CheckDecompressionStatusBadRequest {
	return &CheckDecompressionStatusBadRequest{}
}

/*CheckDecompressionStatusBadRequest handles this case with default header values.

This means that the request ID given by the user is incorrect.
*/
type CheckDecompressionStatusBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CheckDecompressionStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /encryption/decompress/{request_id}][%d] checkDecompressionStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckDecompressionStatusBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDecompressionStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDecompressionStatusUnauthorized creates a CheckDecompressionStatusUnauthorized with default headers values
func NewCheckDecompressionStatusUnauthorized() *CheckDecompressionStatusUnauthorized {
	return &CheckDecompressionStatusUnauthorized{}
}

/*CheckDecompressionStatusUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and decompression can't be triggered unless the user first log in with a valid user
*/
type CheckDecompressionStatusUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckDecompressionStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /encryption/decompress/{request_id}][%d] checkDecompressionStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckDecompressionStatusUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDecompressionStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDecompressionStatusNotFound creates a CheckDecompressionStatusNotFound with default headers values
func NewCheckDecompressionStatusNotFound() *CheckDecompressionStatusNotFound {
	return &CheckDecompressionStatusNotFound{}
}

/*CheckDecompressionStatusNotFound handles this case with default header values.

This means that the ID doesn't exist and thus a status can't be returned.
*/
type CheckDecompressionStatusNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CheckDecompressionStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /encryption/decompress/{request_id}][%d] checkDecompressionStatusNotFound  %+v", 404, o.Payload)
}

func (o *CheckDecompressionStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDecompressionStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDecompressionStatusRequestURITooLong creates a CheckDecompressionStatusRequestURITooLong with default headers values
func NewCheckDecompressionStatusRequestURITooLong() *CheckDecompressionStatusRequestURITooLong {
	return &CheckDecompressionStatusRequestURITooLong{}
}

/*CheckDecompressionStatusRequestURITooLong handles this case with default header values.

This means that the the request ID is longer than the server is willing to interpret.
*/
type CheckDecompressionStatusRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *CheckDecompressionStatusRequestURITooLong) Error() string {
	return fmt.Sprintf("[GET /encryption/decompress/{request_id}][%d] checkDecompressionStatusRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *CheckDecompressionStatusRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDecompressionStatusRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDecompressionStatusInternalServerError creates a CheckDecompressionStatusInternalServerError with default headers values
func NewCheckDecompressionStatusInternalServerError() *CheckDecompressionStatusInternalServerError {
	return &CheckDecompressionStatusInternalServerError{}
}

/*CheckDecompressionStatusInternalServerError handles this case with default header values.

This means that something has gone wrong on the burst buffer. The user is advised to wait and send the request again.
*/
type CheckDecompressionStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CheckDecompressionStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /encryption/decompress/{request_id}][%d] checkDecompressionStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckDecompressionStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDecompressionStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckDecompressionStatusOKBody check decompression status o k body
swagger:model CheckDecompressionStatusOKBody
*/
type CheckDecompressionStatusOKBody struct {

	// status
	// Required: true
	Status *string `json:"status"`

	// Single path for encryption endpoints
	TargetPath string `json:"target_path,omitempty"`
}

// Validate validates this check decompression status o k body
func (o *CheckDecompressionStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckDecompressionStatusOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("checkDecompressionStatusOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckDecompressionStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckDecompressionStatusOKBody) UnmarshalBinary(b []byte) error {
	var res CheckDecompressionStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
