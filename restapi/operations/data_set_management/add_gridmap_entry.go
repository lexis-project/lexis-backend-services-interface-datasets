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

// AddGridmapEntryHandlerFunc turns a function with the right signature into a add gridmap entry handler
type AddGridmapEntryHandlerFunc func(AddGridmapEntryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddGridmapEntryHandlerFunc) Handle(params AddGridmapEntryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddGridmapEntryHandler interface for that can handle valid add gridmap entry params
type AddGridmapEntryHandler interface {
	Handle(AddGridmapEntryParams, interface{}) middleware.Responder
}

// NewAddGridmapEntry creates a new http.Handler for the add gridmap entry operation
func NewAddGridmapEntry(ctx *middleware.Context, handler AddGridmapEntryHandler) *AddGridmapEntry {
	return &AddGridmapEntry{Context: ctx, Handler: handler}
}

/*AddGridmapEntry swagger:route POST /gridftp/gridmap dataSetManagement gridftp addGridmapEntry

Add a DN entry to the DDI B2STAGE GridFTP service

Add a DN entry to the DDI B2STAGE GridFTP service

*/
type AddGridmapEntry struct {
	Context *middleware.Context
	Handler AddGridmapEntryHandler
}

func (o *AddGridmapEntry) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddGridmapEntryParams()

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

// AddGridmapEntryBody add gridmap entry body
//
// swagger:model AddGridmapEntryBody
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
