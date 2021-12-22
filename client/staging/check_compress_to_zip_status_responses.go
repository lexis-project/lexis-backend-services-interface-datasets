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

// CheckCompressToZipStatusReader is a Reader for the CheckCompressToZipStatus structure.
type CheckCompressToZipStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckCompressToZipStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckCompressToZipStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckCompressToZipStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckCompressToZipStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckCompressToZipStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewCheckCompressToZipStatusRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckCompressToZipStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckCompressToZipStatusOK creates a CheckCompressToZipStatusOK with default headers values
func NewCheckCompressToZipStatusOK() *CheckCompressToZipStatusOK {
	return &CheckCompressToZipStatusOK{}
}

/*CheckCompressToZipStatusOK handles this case with default header values.

This means that the status has been returned to the user in the response body.
*/
type CheckCompressToZipStatusOK struct {
	Payload *CheckCompressToZipStatusOKBody
}

func (o *CheckCompressToZipStatusOK) Error() string {
	return fmt.Sprintf("[GET /compress/zip/{request_id}][%d] checkCompressToZipStatusOK  %+v", 200, o.Payload)
}

func (o *CheckCompressToZipStatusOK) GetPayload() *CheckCompressToZipStatusOKBody {
	return o.Payload
}

func (o *CheckCompressToZipStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckCompressToZipStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCompressToZipStatusBadRequest creates a CheckCompressToZipStatusBadRequest with default headers values
func NewCheckCompressToZipStatusBadRequest() *CheckCompressToZipStatusBadRequest {
	return &CheckCompressToZipStatusBadRequest{}
}

/*CheckCompressToZipStatusBadRequest handles this case with default header values.

This means that the request ID given by the user is incorrect.
*/
type CheckCompressToZipStatusBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CheckCompressToZipStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /compress/zip/{request_id}][%d] checkCompressToZipStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckCompressToZipStatusBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCompressToZipStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCompressToZipStatusUnauthorized creates a CheckCompressToZipStatusUnauthorized with default headers values
func NewCheckCompressToZipStatusUnauthorized() *CheckCompressToZipStatusUnauthorized {
	return &CheckCompressToZipStatusUnauthorized{}
}

/*CheckCompressToZipStatusUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user
*/
type CheckCompressToZipStatusUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckCompressToZipStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compress/zip/{request_id}][%d] checkCompressToZipStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckCompressToZipStatusUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCompressToZipStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCompressToZipStatusNotFound creates a CheckCompressToZipStatusNotFound with default headers values
func NewCheckCompressToZipStatusNotFound() *CheckCompressToZipStatusNotFound {
	return &CheckCompressToZipStatusNotFound{}
}

/*CheckCompressToZipStatusNotFound handles this case with default header values.

This means that the ID doesn't exist and thus a status can't be returned.
*/
type CheckCompressToZipStatusNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CheckCompressToZipStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /compress/zip/{request_id}][%d] checkCompressToZipStatusNotFound  %+v", 404, o.Payload)
}

func (o *CheckCompressToZipStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCompressToZipStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCompressToZipStatusRequestURITooLong creates a CheckCompressToZipStatusRequestURITooLong with default headers values
func NewCheckCompressToZipStatusRequestURITooLong() *CheckCompressToZipStatusRequestURITooLong {
	return &CheckCompressToZipStatusRequestURITooLong{}
}

/*CheckCompressToZipStatusRequestURITooLong handles this case with default header values.

This means that the the request ID is longer than the server is willing to interpret.
*/
type CheckCompressToZipStatusRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *CheckCompressToZipStatusRequestURITooLong) Error() string {
	return fmt.Sprintf("[GET /compress/zip/{request_id}][%d] checkCompressToZipStatusRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *CheckCompressToZipStatusRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCompressToZipStatusRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCompressToZipStatusInternalServerError creates a CheckCompressToZipStatusInternalServerError with default headers values
func NewCheckCompressToZipStatusInternalServerError() *CheckCompressToZipStatusInternalServerError {
	return &CheckCompressToZipStatusInternalServerError{}
}

/*CheckCompressToZipStatusInternalServerError handles this case with default header values.

This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.
*/
type CheckCompressToZipStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CheckCompressToZipStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compress/zip/{request_id}][%d] checkCompressToZipStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckCompressToZipStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCompressToZipStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckCompressToZipStatusOKBody check compress to zip status o k body
swagger:model CheckCompressToZipStatusOKBody
*/
type CheckCompressToZipStatusOKBody struct {

	// status
	// Required: true
	Status *string `json:"status"`

	// Multiple paths for CompressAPI endpoints
	TargetPaths []string `json:"target_paths"`
}

// Validate validates this check compress to zip status o k body
func (o *CheckCompressToZipStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckCompressToZipStatusOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("checkCompressToZipStatusOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckCompressToZipStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckCompressToZipStatusOKBody) UnmarshalBinary(b []byte) error {
	var res CheckCompressToZipStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
