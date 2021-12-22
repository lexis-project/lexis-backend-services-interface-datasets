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

// NewDecryptDecompressParams creates a new DecryptDecompressParams object
// with the default values initialized.
func NewDecryptDecompressParams() *DecryptDecompressParams {
	var ()
	return &DecryptDecompressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDecryptDecompressParamsWithTimeout creates a new DecryptDecompressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDecryptDecompressParamsWithTimeout(timeout time.Duration) *DecryptDecompressParams {
	var ()
	return &DecryptDecompressParams{

		timeout: timeout,
	}
}

// NewDecryptDecompressParamsWithContext creates a new DecryptDecompressParams object
// with the default values initialized, and the ability to set a context for a request
func NewDecryptDecompressParamsWithContext(ctx context.Context) *DecryptDecompressParams {
	var ()
	return &DecryptDecompressParams{

		Context: ctx,
	}
}

// NewDecryptDecompressParamsWithHTTPClient creates a new DecryptDecompressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDecryptDecompressParamsWithHTTPClient(client *http.Client) *DecryptDecompressParams {
	var ()
	return &DecryptDecompressParams{
		HTTPClient: client,
	}
}

/*DecryptDecompressParams contains all the parameters to send to the API endpoint
for the decrypt decompress operation typically these are written to a http.Request
*/
type DecryptDecompressParams struct {

	/*Parameters
	  parameters

	*/
	Parameters DecryptDecompressBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the decrypt decompress params
func (o *DecryptDecompressParams) WithTimeout(timeout time.Duration) *DecryptDecompressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the decrypt decompress params
func (o *DecryptDecompressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the decrypt decompress params
func (o *DecryptDecompressParams) WithContext(ctx context.Context) *DecryptDecompressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the decrypt decompress params
func (o *DecryptDecompressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the decrypt decompress params
func (o *DecryptDecompressParams) WithHTTPClient(client *http.Client) *DecryptDecompressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the decrypt decompress params
func (o *DecryptDecompressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParameters adds the parameters to the decrypt decompress params
func (o *DecryptDecompressParams) WithParameters(parameters DecryptDecompressBody) *DecryptDecompressParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the decrypt decompress params
func (o *DecryptDecompressParams) SetParameters(parameters DecryptDecompressBody) {
	o.Parameters = parameters
}

// WriteToRequest writes these params to a swagger request
func (o *DecryptDecompressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Parameters); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}