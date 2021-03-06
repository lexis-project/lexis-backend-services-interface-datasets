// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewCheckDecompressionStatusParams creates a new CheckDecompressionStatusParams object
// no default values defined in spec.
func NewCheckDecompressionStatusParams() CheckDecompressionStatusParams {

	return CheckDecompressionStatusParams{}
}

// CheckDecompressionStatusParams contains all the bound params for the check decompression status operation
// typically these are obtained from a http.Request
//
// swagger:parameters CheckDecompressionStatus
type CheckDecompressionStatusParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	RequestID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCheckDecompressionStatusParams() beforehand.
func (o *CheckDecompressionStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rRequestID, rhkRequestID, _ := route.Params.GetOK("request_id")
	if err := o.bindRequestID(rRequestID, rhkRequestID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindRequestID binds and validates parameter RequestID from path.
func (o *CheckDecompressionStatusParams) bindRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("request_id", "path", "strfmt.UUID", raw)
	}
	o.RequestID = *(value.(*strfmt.UUID))

	if err := o.validateRequestID(formats); err != nil {
		return err
	}

	return nil
}

// validateRequestID carries on validations for parameter RequestID
func (o *CheckDecompressionStatusParams) validateRequestID(formats strfmt.Registry) error {

	if err := validate.FormatOf("request_id", "path", "uuid", o.RequestID.String(), formats); err != nil {
		return err
	}
	return nil
}
