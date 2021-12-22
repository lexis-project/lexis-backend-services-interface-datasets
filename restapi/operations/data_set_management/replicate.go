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

// ReplicateHandlerFunc turns a function with the right signature into a replicate handler
type ReplicateHandlerFunc func(ReplicateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplicateHandlerFunc) Handle(params ReplicateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplicateHandler interface for that can handle valid replicate params
type ReplicateHandler interface {
	Handle(ReplicateParams, interface{}) middleware.Responder
}

// NewReplicate creates a new http.Handler for the replicate operation
func NewReplicate(ctx *middleware.Context, handler ReplicateHandler) *Replicate {
	return &Replicate{Context: ctx, Handler: handler}
}

/*Replicate swagger:route POST /replicate dataSetManagement EUDATReplication replicate

This is called when a user requests data to be replicated between different systems. The request will be added to the jobs queue.

This is called when a user requests data to be replicated between different systems. The request will be added to the jobs queue.

*/
type Replicate struct {
	Context *middleware.Context
	Handler ReplicateHandler
}

func (o *Replicate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplicateParams()

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

// ReplicateBody replicate body
//
// swagger:model ReplicateBody
type ReplicateBody struct {

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// source system
	// Required: true
	SourceSystem *string `json:"source_system"`

	// target path
	TargetPath string `json:"target_path,omitempty"`

	// target system
	// Required: true
	TargetSystem *string `json:"target_system"`
}

// Validate validates this replicate body
func (o *ReplicateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourceSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTargetSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplicateBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *ReplicateBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

func (o *ReplicateBody) validateTargetSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"target_system", "body", o.TargetSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReplicateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplicateBody) UnmarshalBinary(b []byte) error {
	var res ReplicateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
