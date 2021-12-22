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
	"github.com/go-openapi/swag"
)

// NewFilePatchParams creates a new FilePatchParams object
// with the default values initialized.
func NewFilePatchParams() *FilePatchParams {
	var ()
	return &FilePatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFilePatchParamsWithTimeout creates a new FilePatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFilePatchParamsWithTimeout(timeout time.Duration) *FilePatchParams {
	var ()
	return &FilePatchParams{

		timeout: timeout,
	}
}

// NewFilePatchParamsWithContext creates a new FilePatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewFilePatchParamsWithContext(ctx context.Context) *FilePatchParams {
	var ()
	return &FilePatchParams{

		Context: ctx,
	}
}

// NewFilePatchParamsWithHTTPClient creates a new FilePatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFilePatchParamsWithHTTPClient(client *http.Client) *FilePatchParams {
	var ()
	return &FilePatchParams{
		HTTPClient: client,
	}
}

/*FilePatchParams contains all the parameters to send to the API endpoint
for the file patch operation typically these are written to a http.Request
*/
type FilePatchParams struct {

	/*Body*/
	Body string
	/*ContentLength*/
	ContentLength int64
	/*TusResumable
	  Protocol version

	*/
	TusResumable string
	/*UploadChecksum
	  Added by the checksum extension. The Upload-Checksum request header contains information about the checksum of the current body payload. The header MUST consist of the name of the used checksum algorithm and the Base64 encoded checksum separated by a space.

	*/
	UploadChecksum *string
	/*UploadOffset
	  The Upload-Offset request and response header indicates a byte offset within a resource. The value MUST be a non-negative integer.

	*/
	UploadOffset int64
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the file patch params
func (o *FilePatchParams) WithTimeout(timeout time.Duration) *FilePatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file patch params
func (o *FilePatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file patch params
func (o *FilePatchParams) WithContext(ctx context.Context) *FilePatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file patch params
func (o *FilePatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file patch params
func (o *FilePatchParams) WithHTTPClient(client *http.Client) *FilePatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file patch params
func (o *FilePatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the file patch params
func (o *FilePatchParams) WithBody(body string) *FilePatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the file patch params
func (o *FilePatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentLength adds the contentLength to the file patch params
func (o *FilePatchParams) WithContentLength(contentLength int64) *FilePatchParams {
	o.SetContentLength(contentLength)
	return o
}

// SetContentLength adds the contentLength to the file patch params
func (o *FilePatchParams) SetContentLength(contentLength int64) {
	o.ContentLength = contentLength
}

// WithTusResumable adds the tusResumable to the file patch params
func (o *FilePatchParams) WithTusResumable(tusResumable string) *FilePatchParams {
	o.SetTusResumable(tusResumable)
	return o
}

// SetTusResumable adds the tusResumable to the file patch params
func (o *FilePatchParams) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WithUploadChecksum adds the uploadChecksum to the file patch params
func (o *FilePatchParams) WithUploadChecksum(uploadChecksum *string) *FilePatchParams {
	o.SetUploadChecksum(uploadChecksum)
	return o
}

// SetUploadChecksum adds the uploadChecksum to the file patch params
func (o *FilePatchParams) SetUploadChecksum(uploadChecksum *string) {
	o.UploadChecksum = uploadChecksum
}

// WithUploadOffset adds the uploadOffset to the file patch params
func (o *FilePatchParams) WithUploadOffset(uploadOffset int64) *FilePatchParams {
	o.SetUploadOffset(uploadOffset)
	return o
}

// SetUploadOffset adds the uploadOffset to the file patch params
func (o *FilePatchParams) SetUploadOffset(uploadOffset int64) {
	o.UploadOffset = uploadOffset
}

// WithID adds the id to the file patch params
func (o *FilePatchParams) WithID(id string) *FilePatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the file patch params
func (o *FilePatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FilePatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// header param Content-Length
	if err := r.SetHeaderParam("Content-Length", swag.FormatInt64(o.ContentLength)); err != nil {
		return err
	}

	// header param Tus-Resumable
	if err := r.SetHeaderParam("Tus-Resumable", o.TusResumable); err != nil {
		return err
	}

	if o.UploadChecksum != nil {

		// header param Upload-Checksum
		if err := r.SetHeaderParam("Upload-Checksum", *o.UploadChecksum); err != nil {
			return err
		}

	}

	// header param Upload-offset
	if err := r.SetHeaderParam("Upload-offset", swag.FormatInt64(o.UploadOffset)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}