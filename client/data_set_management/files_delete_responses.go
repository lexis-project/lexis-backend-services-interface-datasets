// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/models"
)

// FilesDeleteReader is a Reader for the FilesDelete structure.
type FilesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewFilesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFilesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewFilesDeletePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewFilesDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFilesDeleteNoContent creates a FilesDeleteNoContent with default headers values
func NewFilesDeleteNoContent() *FilesDeleteNoContent {
	return &FilesDeleteNoContent{}
}

/*FilesDeleteNoContent handles this case with default header values.

Upload was terminated
*/
type FilesDeleteNoContent struct {
	/*Protocol version
	 */
	TusResumable string
}

func (o *FilesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dataset/upload/{id}][%d] filesDeleteNoContent ", 204)
}

func (o *FilesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Tus-Resumable
	o.TusResumable = response.GetHeader("Tus-Resumable")

	return nil
}

// NewFilesDeleteUnauthorized creates a FilesDeleteUnauthorized with default headers values
func NewFilesDeleteUnauthorized() *FilesDeleteUnauthorized {
	return &FilesDeleteUnauthorized{}
}

/*FilesDeleteUnauthorized handles this case with default header values.

Authorization failed
*/
type FilesDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *FilesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /dataset/upload/{id}][%d] filesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *FilesDeleteUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FilesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesDeletePreconditionFailed creates a FilesDeletePreconditionFailed with default headers values
func NewFilesDeletePreconditionFailed() *FilesDeletePreconditionFailed {
	return &FilesDeletePreconditionFailed{}
}

/*FilesDeletePreconditionFailed handles this case with default header values.

Precondition Failed
*/
type FilesDeletePreconditionFailed struct {
	/*Protocol version
	 */
	TusResumable string
	/*The Tus-Version response header MUST be a comma-separated list of protocol versions supported by the Server. The list MUST be sorted by Server's preference where the first one is the most preferred one.
	 */
	TusVersion string
}

func (o *FilesDeletePreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /dataset/upload/{id}][%d] filesDeletePreconditionFailed ", 412)
}

func (o *FilesDeletePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Tus-Resumable
	o.TusResumable = response.GetHeader("Tus-Resumable")

	// response header Tus-Version
	o.TusVersion = response.GetHeader("Tus-Version")

	return nil
}

// NewFilesDeleteServiceUnavailable creates a FilesDeleteServiceUnavailable with default headers values
func NewFilesDeleteServiceUnavailable() *FilesDeleteServiceUnavailable {
	return &FilesDeleteServiceUnavailable{}
}

/*FilesDeleteServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type FilesDeleteServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *FilesDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /dataset/upload/{id}][%d] filesDeleteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *FilesDeleteServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FilesDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
