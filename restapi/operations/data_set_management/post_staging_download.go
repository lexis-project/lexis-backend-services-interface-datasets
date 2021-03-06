// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostStagingDownloadHandlerFunc turns a function with the right signature into a post staging download handler
type PostStagingDownloadHandlerFunc func(PostStagingDownloadParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostStagingDownloadHandlerFunc) Handle(params PostStagingDownloadParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostStagingDownloadHandler interface for that can handle valid post staging download params
type PostStagingDownloadHandler interface {
	Handle(PostStagingDownloadParams, interface{}) middleware.Responder
}

// NewPostStagingDownload creates a new http.Handler for the post staging download operation
func NewPostStagingDownload(ctx *middleware.Context, handler PostStagingDownloadHandler) *PostStagingDownload {
	return &PostStagingDownload{Context: ctx, Handler: handler}
}

/*PostStagingDownload swagger:route POST /staging/download dataSetManagement postStagingDownload

Download from staging zone

Download from staging zone

*/
type PostStagingDownload struct {
	Context *middleware.Context
	Handler PostStagingDownloadHandler
}

func (o *PostStagingDownload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostStagingDownloadParams()

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

// PostStagingDownloadBody post staging download body
//
// swagger:model PostStagingDownloadBody
type PostStagingDownloadBody struct {

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// source system
	// Required: true
	// Enum: [lrz_staging_area]
	SourceSystem *string `json:"source_system"`
}

// Validate validates this post staging download body
func (o *PostStagingDownloadBody) Validate(formats strfmt.Registry) error {
	var res []error

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

func (o *PostStagingDownloadBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

var postStagingDownloadBodyTypeSourceSystemPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["lrz_staging_area"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postStagingDownloadBodyTypeSourceSystemPropEnum = append(postStagingDownloadBodyTypeSourceSystemPropEnum, v)
	}
}

const (

	// PostStagingDownloadBodySourceSystemLrzStagingArea captures enum value "lrz_staging_area"
	PostStagingDownloadBodySourceSystemLrzStagingArea string = "lrz_staging_area"
)

// prop value enum
func (o *PostStagingDownloadBody) validateSourceSystemEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postStagingDownloadBodyTypeSourceSystemPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostStagingDownloadBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	// value enum
	if err := o.validateSourceSystemEnum("parameters"+"."+"source_system", "body", *o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostStagingDownloadBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostStagingDownloadBody) UnmarshalBinary(b []byte) error {
	var res PostStagingDownloadBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
