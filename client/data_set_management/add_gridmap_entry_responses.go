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

// AddGridmapEntryReader is a Reader for the AddGridmapEntry structure.
type AddGridmapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddGridmapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddGridmapEntryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddGridmapEntryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddGridmapEntryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewAddGridmapEntryBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewAddGridmapEntryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddGridmapEntryCreated creates a AddGridmapEntryCreated with default headers values
func NewAddGridmapEntryCreated() *AddGridmapEntryCreated {
	return &AddGridmapEntryCreated{}
}

/*AddGridmapEntryCreated handles this case with default header values.

DN added
*/
type AddGridmapEntryCreated struct {
}

func (o *AddGridmapEntryCreated) Error() string {
	return fmt.Sprintf("[POST /gridftp/gridmap][%d] addGridmapEntryCreated ", 201)
}

func (o *AddGridmapEntryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddGridmapEntryUnauthorized creates a AddGridmapEntryUnauthorized with default headers values
func NewAddGridmapEntryUnauthorized() *AddGridmapEntryUnauthorized {
	return &AddGridmapEntryUnauthorized{}
}

/*AddGridmapEntryUnauthorized handles this case with default header values.

Unauthorized
*/
type AddGridmapEntryUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *AddGridmapEntryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /gridftp/gridmap][%d] addGridmapEntryUnauthorized  %+v", 401, o.Payload)
}

func (o *AddGridmapEntryUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddGridmapEntryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGridmapEntryForbidden creates a AddGridmapEntryForbidden with default headers values
func NewAddGridmapEntryForbidden() *AddGridmapEntryForbidden {
	return &AddGridmapEntryForbidden{}
}

/*AddGridmapEntryForbidden handles this case with default header values.

Forbidden
*/
type AddGridmapEntryForbidden struct {
	Payload *models.ErrorResponse
}

func (o *AddGridmapEntryForbidden) Error() string {
	return fmt.Sprintf("[POST /gridftp/gridmap][%d] addGridmapEntryForbidden  %+v", 403, o.Payload)
}

func (o *AddGridmapEntryForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddGridmapEntryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGridmapEntryBadGateway creates a AddGridmapEntryBadGateway with default headers values
func NewAddGridmapEntryBadGateway() *AddGridmapEntryBadGateway {
	return &AddGridmapEntryBadGateway{}
}

/*AddGridmapEntryBadGateway handles this case with default header values.

Bad Gateway
*/
type AddGridmapEntryBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *AddGridmapEntryBadGateway) Error() string {
	return fmt.Sprintf("[POST /gridftp/gridmap][%d] addGridmapEntryBadGateway  %+v", 502, o.Payload)
}

func (o *AddGridmapEntryBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddGridmapEntryBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGridmapEntryServiceUnavailable creates a AddGridmapEntryServiceUnavailable with default headers values
func NewAddGridmapEntryServiceUnavailable() *AddGridmapEntryServiceUnavailable {
	return &AddGridmapEntryServiceUnavailable{}
}

/*AddGridmapEntryServiceUnavailable handles this case with default header values.

Server error
*/
type AddGridmapEntryServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *AddGridmapEntryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /gridftp/gridmap][%d] addGridmapEntryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *AddGridmapEntryServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddGridmapEntryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddGridmapEntryBody add gridmap entry body
swagger:model AddGridmapEntryBody
*/
type AddGridmapEntryBody struct {

	// dn
	// Required: true
	Dn *string `json:"dn"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this add gridmap entry body
func (o *AddGridmapEntryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddGridmapEntryBody) validateDn(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"dn", "body", o.Dn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddGridmapEntryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddGridmapEntryBody) UnmarshalBinary(b []byte) error {
	var res AddGridmapEntryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
