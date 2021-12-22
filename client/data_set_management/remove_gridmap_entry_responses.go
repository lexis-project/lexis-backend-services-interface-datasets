// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// RemoveGridmapEntryReader is a Reader for the RemoveGridmapEntry structure.
type RemoveGridmapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveGridmapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveGridmapEntryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveGridmapEntryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveGridmapEntryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewRemoveGridmapEntryBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewRemoveGridmapEntryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveGridmapEntryNoContent creates a RemoveGridmapEntryNoContent with default headers values
func NewRemoveGridmapEntryNoContent() *RemoveGridmapEntryNoContent {
	return &RemoveGridmapEntryNoContent{}
}

/*RemoveGridmapEntryNoContent handles this case with default header values.

DN removed
*/
type RemoveGridmapEntryNoContent struct {
}

func (o *RemoveGridmapEntryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /gridftp/gridmap][%d] removeGridmapEntryNoContent ", 204)
}

func (o *RemoveGridmapEntryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveGridmapEntryUnauthorized creates a RemoveGridmapEntryUnauthorized with default headers values
func NewRemoveGridmapEntryUnauthorized() *RemoveGridmapEntryUnauthorized {
	return &RemoveGridmapEntryUnauthorized{}
}

/*RemoveGridmapEntryUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveGridmapEntryUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *RemoveGridmapEntryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /gridftp/gridmap][%d] removeGridmapEntryUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveGridmapEntryUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RemoveGridmapEntryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveGridmapEntryForbidden creates a RemoveGridmapEntryForbidden with default headers values
func NewRemoveGridmapEntryForbidden() *RemoveGridmapEntryForbidden {
	return &RemoveGridmapEntryForbidden{}
}

/*RemoveGridmapEntryForbidden handles this case with default header values.

Forbidden
*/
type RemoveGridmapEntryForbidden struct {
	Payload *models.ErrorResponse
}

func (o *RemoveGridmapEntryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /gridftp/gridmap][%d] removeGridmapEntryForbidden  %+v", 403, o.Payload)
}

func (o *RemoveGridmapEntryForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RemoveGridmapEntryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveGridmapEntryBadGateway creates a RemoveGridmapEntryBadGateway with default headers values
func NewRemoveGridmapEntryBadGateway() *RemoveGridmapEntryBadGateway {
	return &RemoveGridmapEntryBadGateway{}
}

/*RemoveGridmapEntryBadGateway handles this case with default header values.

Bad Gateway
*/
type RemoveGridmapEntryBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *RemoveGridmapEntryBadGateway) Error() string {
	return fmt.Sprintf("[DELETE /gridftp/gridmap][%d] removeGridmapEntryBadGateway  %+v", 502, o.Payload)
}

func (o *RemoveGridmapEntryBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RemoveGridmapEntryBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveGridmapEntryServiceUnavailable creates a RemoveGridmapEntryServiceUnavailable with default headers values
func NewRemoveGridmapEntryServiceUnavailable() *RemoveGridmapEntryServiceUnavailable {
	return &RemoveGridmapEntryServiceUnavailable{}
}

/*RemoveGridmapEntryServiceUnavailable handles this case with default header values.

Server error
*/
type RemoveGridmapEntryServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *RemoveGridmapEntryServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /gridftp/gridmap][%d] removeGridmapEntryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *RemoveGridmapEntryServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RemoveGridmapEntryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveGridmapEntryBody remove gridmap entry body
swagger:model RemoveGridmapEntryBody
*/
type RemoveGridmapEntryBody struct {

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this remove gridmap entry body
func (o *RemoveGridmapEntryBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveGridmapEntryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveGridmapEntryBody) UnmarshalBinary(b []byte) error {
	var res RemoveGridmapEntryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
