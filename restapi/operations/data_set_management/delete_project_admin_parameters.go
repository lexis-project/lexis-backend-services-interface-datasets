// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// NewDeleteProjectAdminParams creates a new DeleteProjectAdminParams object
// no default values defined in spec.
func NewDeleteProjectAdminParams() DeleteProjectAdminParams {

	return DeleteProjectAdminParams{}
}

// DeleteProjectAdminParams contains all the bound params for the delete project admin operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteProjectAdmin
type DeleteProjectAdminParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	ProjectAndUserInfo *models.UserInProject
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteProjectAdminParams() beforehand.
func (o *DeleteProjectAdminParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UserInProject
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("projectAndUserInfo", "body", ""))
			} else {
				res = append(res, errors.NewParseError("projectAndUserInfo", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ProjectAndUserInfo = &body
			}
		}
	} else {
		res = append(res, errors.Required("projectAndUserInfo", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
