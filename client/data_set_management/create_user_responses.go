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

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewCreateUserBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/*CreateUserCreated handles this case with default header values.

User was created
*/
type CreateUserCreated struct {
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserCreated ", 201)
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/*CreateUserBadRequest handles this case with default header values.

Malformed request
*/
type CreateUserBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserUnauthorized creates a CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {
	return &CreateUserUnauthorized{}
}

/*CreateUserUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateUserUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUserUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserForbidden creates a CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

/*CreateUserForbidden handles this case with default header values.

Permission denied
*/
type CreateUserForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserConflict creates a CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {
	return &CreateUserConflict{}
}

/*CreateUserConflict handles this case with default header values.

User already exists
*/
type CreateUserConflict struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserConflict  %+v", 409, o.Payload)
}

func (o *CreateUserConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserInternalServerError creates a CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {
	return &CreateUserInternalServerError{}
}

/*CreateUserInternalServerError handles this case with default header values.

Internal error processing request
*/
type CreateUserInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserBadGateway creates a CreateUserBadGateway with default headers values
func NewCreateUserBadGateway() *CreateUserBadGateway {
	return &CreateUserBadGateway{}
}

/*CreateUserBadGateway handles this case with default header values.

Bad Gateway
*/
type CreateUserBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserBadGateway) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserBadGateway  %+v", 502, o.Payload)
}

func (o *CreateUserBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserServiceUnavailable creates a CreateUserServiceUnavailable with default headers values
func NewCreateUserServiceUnavailable() *CreateUserServiceUnavailable {
	return &CreateUserServiceUnavailable{}
}

/*CreateUserServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type CreateUserServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateUserServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateUserBody create user body
swagger:model CreateUserBody
*/
type CreateUserBody struct {

	// Keycloak ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Keycloak Username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this create user body
func (o *CreateUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateUserBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("UserInfo"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("UserInfo"+"."+"id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateUserBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("UserInfo"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateUserBody) UnmarshalBinary(b []byte) error {
	var res CreateUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}