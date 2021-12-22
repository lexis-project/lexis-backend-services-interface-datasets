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

// DecryptReader is a Reader for the Decrypt structure.
type DecryptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecryptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDecryptCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDecryptBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDecryptUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDecryptForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDecryptNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewDecryptRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDecryptTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDecryptInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDecryptCreated creates a DecryptCreated with default headers values
func NewDecryptCreated() *DecryptCreated {
	return &DecryptCreated{}
}

/*DecryptCreated handles this case with default header values.

The response code means that the decryption has been initiated. Status of the operation can be checked by querying the status.
*/
type DecryptCreated struct {
	Payload *models.RequestID
}

func (o *DecryptCreated) Error() string {
	return fmt.Sprintf("[POST /encryption/decrypt][%d] decryptCreated  %+v", 201, o.Payload)
}

func (o *DecryptCreated) GetPayload() *models.RequestID {
	return o.Payload
}

func (o *DecryptCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptBadRequest creates a DecryptBadRequest with default headers values
func NewDecryptBadRequest() *DecryptBadRequest {
	return &DecryptBadRequest{}
}

/*DecryptBadRequest handles this case with default header values.

This means that there's something wrong in the input parameters and the server couldn't understand the request.
*/
type DecryptBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DecryptBadRequest) Error() string {
	return fmt.Sprintf("[POST /encryption/decrypt][%d] decryptBadRequest  %+v", 400, o.Payload)
}

func (o *DecryptBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptUnauthorized creates a DecryptUnauthorized with default headers values
func NewDecryptUnauthorized() *DecryptUnauthorized {
	return &DecryptUnauthorized{}
}

/*DecryptUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and decryption can't be triggered unless the user first log in with a valid user
*/
type DecryptUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DecryptUnauthorized) Error() string {
	return fmt.Sprintf("[POST /encryption/decrypt][%d] decryptUnauthorized  %+v", 401, o.Payload)
}

func (o *DecryptUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptForbidden creates a DecryptForbidden with default headers values
func NewDecryptForbidden() *DecryptForbidden {
	return &DecryptForbidden{}
}

/*DecryptForbidden handles this case with default header values.

This means that the resource the user is trying to decrypt from or to is not readable. User doesn't have the correct rights to read the source file.
*/
type DecryptForbidden struct {
	Payload *models.ErrorResponse
}

func (o *DecryptForbidden) Error() string {
	return fmt.Sprintf("[POST /encryption/decrypt][%d] decryptForbidden  %+v", 403, o.Payload)
}

func (o *DecryptForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptNotFound creates a DecryptNotFound with default headers values
func NewDecryptNotFound() *DecryptNotFound {
	return &DecryptNotFound{}
}

/*DecryptNotFound handles this case with default header values.

This means that the source path on the system doesn't exist.
*/
type DecryptNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DecryptNotFound) Error() string {
	return fmt.Sprintf("[POST /encryption/decrypt][%d] decryptNotFound  %+v", 404, o.Payload)
}

func (o *DecryptNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptRequestURITooLong creates a DecryptRequestURITooLong with default headers values
func NewDecryptRequestURITooLong() *DecryptRequestURITooLong {
	return &DecryptRequestURITooLong{}
}

/*DecryptRequestURITooLong handles this case with default header values.

This means that the either the source path is longer than the server is willing to interpret.
*/
type DecryptRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *DecryptRequestURITooLong) Error() string {
	return fmt.Sprintf("[POST /encryption/decrypt][%d] decryptRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *DecryptRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptTooManyRequests creates a DecryptTooManyRequests with default headers values
func NewDecryptTooManyRequests() *DecryptTooManyRequests {
	return &DecryptTooManyRequests{}
}

/*DecryptTooManyRequests handles this case with default header values.

This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.
*/
type DecryptTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *DecryptTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /encryption/decrypt][%d] decryptTooManyRequests  %+v", 429, o.Payload)
}

func (o *DecryptTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptInternalServerError creates a DecryptInternalServerError with default headers values
func NewDecryptInternalServerError() *DecryptInternalServerError {
	return &DecryptInternalServerError{}
}

/*DecryptInternalServerError handles this case with default header values.

This means that something has gone wrong on the burst buffer. The user is advised to wait and send the request again.
*/
type DecryptInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DecryptInternalServerError) Error() string {
	return fmt.Sprintf("[POST /encryption/decrypt][%d] decryptInternalServerError  %+v", 500, o.Payload)
}

func (o *DecryptInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DecryptBody decrypt body
swagger:model DecryptBody
*/
type DecryptBody struct {

	// project
	// Required: true
	Project *string `json:"project"`

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// source system
	// Required: true
	SourceSystem *string `json:"source_system"`
}

// Validate validates this decrypt body
func (o *DecryptBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourceSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DecryptBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *DecryptBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *DecryptBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DecryptBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DecryptBody) UnmarshalBinary(b []byte) error {
	var res DecryptBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
