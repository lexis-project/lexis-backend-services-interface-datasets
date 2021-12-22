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

// NewDeleteDatasetByMetadataParams creates a new DeleteDatasetByMetadataParams object
// no default values defined in spec.
func NewDeleteDatasetByMetadataParams() DeleteDatasetByMetadataParams {

	return DeleteDatasetByMetadataParams{}
}

// DeleteDatasetByMetadataParams contains all the bound params for the delete dataset by metadata operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteDatasetByMetadata
type DeleteDatasetByMetadataParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*metadata relating to dataset which is being searched for
	  Required: true
	  In: body
	*/
	MetadataQuery *models.MetadataQuery
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteDatasetByMetadataParams() beforehand.
func (o *DeleteDatasetByMetadataParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.MetadataQuery
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("metadataQuery", "body", ""))
			} else {
				res = append(res, errors.NewParseError("metadataQuery", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.MetadataQuery = &body
			}
		}
	} else {
		res = append(res, errors.Required("metadataQuery", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}