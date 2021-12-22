// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewOptionsDatasetUploadParams creates a new OptionsDatasetUploadParams object
// with the default values initialized.
func NewOptionsDatasetUploadParams() *OptionsDatasetUploadParams {
	var ()
	return &OptionsDatasetUploadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOptionsDatasetUploadParamsWithTimeout creates a new OptionsDatasetUploadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOptionsDatasetUploadParamsWithTimeout(timeout time.Duration) *OptionsDatasetUploadParams {
	var ()
	return &OptionsDatasetUploadParams{

		timeout: timeout,
	}
}

// NewOptionsDatasetUploadParamsWithContext creates a new OptionsDatasetUploadParams object
// with the default values initialized, and the ability to set a context for a request
func NewOptionsDatasetUploadParamsWithContext(ctx context.Context) *OptionsDatasetUploadParams {
	var ()
	return &OptionsDatasetUploadParams{

		Context: ctx,
	}
}

// NewOptionsDatasetUploadParamsWithHTTPClient creates a new OptionsDatasetUploadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOptionsDatasetUploadParamsWithHTTPClient(client *http.Client) *OptionsDatasetUploadParams {
	var ()
	return &OptionsDatasetUploadParams{
		HTTPClient: client,
	}
}

/*OptionsDatasetUploadParams contains all the parameters to send to the API endpoint
for the options dataset upload operation typically these are written to a http.Request
*/
type OptionsDatasetUploadParams struct {

	/*TusResumable
	  Protocol version

	*/
	TusResumable string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the options dataset upload params
func (o *OptionsDatasetUploadParams) WithTimeout(timeout time.Duration) *OptionsDatasetUploadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the options dataset upload params
func (o *OptionsDatasetUploadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the options dataset upload params
func (o *OptionsDatasetUploadParams) WithContext(ctx context.Context) *OptionsDatasetUploadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the options dataset upload params
func (o *OptionsDatasetUploadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the options dataset upload params
func (o *OptionsDatasetUploadParams) WithHTTPClient(client *http.Client) *OptionsDatasetUploadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the options dataset upload params
func (o *OptionsDatasetUploadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTusResumable adds the tusResumable to the options dataset upload params
func (o *OptionsDatasetUploadParams) WithTusResumable(tusResumable string) *OptionsDatasetUploadParams {
	o.SetTusResumable(tusResumable)
	return o
}

// SetTusResumable adds the tusResumable to the options dataset upload params
func (o *OptionsDatasetUploadParams) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WriteToRequest writes these params to a swagger request
func (o *OptionsDatasetUploadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Tus-Resumable
	if err := r.SetHeaderParam("Tus-Resumable", o.TusResumable); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
