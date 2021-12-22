// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddProjectUserHandlerFunc turns a function with the right signature into a add project user handler
type AddProjectUserHandlerFunc func(AddProjectUserParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddProjectUserHandlerFunc) Handle(params AddProjectUserParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddProjectUserHandler interface for that can handle valid add project user params
type AddProjectUserHandler interface {
	Handle(AddProjectUserParams, interface{}) middleware.Responder
}

// NewAddProjectUser creates a new http.Handler for the add project user operation
func NewAddProjectUser(ctx *middleware.Context, handler AddProjectUserHandler) *AddProjectUser {
	return &AddProjectUser{Context: ctx, Handler: handler}
}

/*AddProjectUser swagger:route POST /project/user dataSetManagement iRODSPermissionManagement addProjectUser

Add a user to a project in iRODS WP3 DDI

Add a user to a project in iRODS WP3 DDI

*/
type AddProjectUser struct {
	Context *middleware.Context
	Handler AddProjectUserHandler
}

func (o *AddProjectUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddProjectUserParams()

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