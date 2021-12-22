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

// DeleteProjectHandlerFunc turns a function with the right signature into a delete project handler
type DeleteProjectHandlerFunc func(DeleteProjectParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteProjectHandlerFunc) Handle(params DeleteProjectParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteProjectHandler interface for that can handle valid delete project params
type DeleteProjectHandler interface {
	Handle(DeleteProjectParams, interface{}) middleware.Responder
}

// NewDeleteProject creates a new http.Handler for the delete project operation
func NewDeleteProject(ctx *middleware.Context, handler DeleteProjectHandler) *DeleteProject {
	return &DeleteProject{Context: ctx, Handler: handler}
}

/*DeleteProject swagger:route DELETE /project dataSetManagement deleteProject

Delete a project in iRODS WP3 DDI

Delete a project in iRODS WP3 DDI.

*/
type DeleteProject struct {
	Context *middleware.Context
	Handler DeleteProjectHandler
}

func (o *DeleteProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteProjectParams()

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

// DeleteProjectBody delete project body
//
// swagger:model DeleteProjectBody
type DeleteProjectBody struct {

	// Project name (Keycloak group)
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this delete project body
func (o *DeleteProjectBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteProjectBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("ProjectInfo"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteProjectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteProjectBody) UnmarshalBinary(b []byte) error {
	var res DeleteProjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}