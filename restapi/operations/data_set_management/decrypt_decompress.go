// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DecryptDecompressHandlerFunc turns a function with the right signature into a decrypt decompress handler
type DecryptDecompressHandlerFunc func(DecryptDecompressParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DecryptDecompressHandlerFunc) Handle(params DecryptDecompressParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DecryptDecompressHandler interface for that can handle valid decrypt decompress params
type DecryptDecompressHandler interface {
	Handle(DecryptDecompressParams, interface{}) middleware.Responder
}

// NewDecryptDecompress creates a new http.Handler for the decrypt decompress operation
func NewDecryptDecompress(ctx *middleware.Context, handler DecryptDecompressHandler) *DecryptDecompress {
	return &DecryptDecompress{Context: ctx, Handler: handler}
}

/*DecryptDecompress swagger:route POST /encryption/decrypt_decompress dataSetManagement decryptDecompress

Decrypt and decompress a dataset or subdataset (by enqueuing the request for latter processing)

Decrypt and decompress a dataset or subdataset (by enqueuing the request for latter processing)

*/
type DecryptDecompress struct {
	Context *middleware.Context
	Handler DecryptDecompressHandler
}

func (o *DecryptDecompress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDecryptDecompressParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DecryptDecompressBody decrypt decompress body
//
// swagger:model DecryptDecompressBody
type DecryptDecompressBody struct {

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

// Validate validates this decrypt decompress body
func (o *DecryptDecompressBody) Validate(formats strfmt.Registry) error {
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

func (o *DecryptDecompressBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *DecryptDecompressBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *DecryptDecompressBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DecryptDecompressBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DecryptDecompressBody) UnmarshalBinary(b []byte) error {
	var res DecryptDecompressBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
