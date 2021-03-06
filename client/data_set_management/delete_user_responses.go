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

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewDeleteUserBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserNoContent creates a DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

/*DeleteUserNoContent handles this case with default header values.

User was deleted
*/
type DeleteUserNoContent struct {
}

func (o *DeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user][%d] deleteUserNoContent ", 204)
}

func (o *DeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserBadRequest creates a DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {
	return &DeleteUserBadRequest{}
}

/*DeleteUserBadRequest handles this case with default header values.

Malformed request
*/
type DeleteUserBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DeleteUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserUnauthorized creates a DeleteUserUnauthorized with default headers values
func NewDeleteUserUnauthorized() *DeleteUserUnauthorized {
	return &DeleteUserUnauthorized{}
}

/*DeleteUserUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteUserUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DeleteUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /user][%d] deleteUserUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserInternalServerError creates a DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {
	return &DeleteUserInternalServerError{}
}

/*DeleteUserInternalServerError handles this case with default header values.

Internal error processing request
*/
type DeleteUserInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /user][%d] deleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserBadGateway creates a DeleteUserBadGateway with default headers values
func NewDeleteUserBadGateway() *DeleteUserBadGateway {
	return &DeleteUserBadGateway{}
}

/*DeleteUserBadGateway handles this case with default header values.

Bad Gateway
*/
type DeleteUserBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *DeleteUserBadGateway) Error() string {
	return fmt.Sprintf("[DELETE /user][%d] deleteUserBadGateway  %+v", 502, o.Payload)
}

func (o *DeleteUserBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteUserBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteUserBody delete user body
swagger:model DeleteUserBody
*/
type DeleteUserBody struct {

	// Keycloak Username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this delete user body
func (o *DeleteUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteUserBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("User"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteUserBody) UnmarshalBinary(b []byte) error {
	var res DeleteUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
