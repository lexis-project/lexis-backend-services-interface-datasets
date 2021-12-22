// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CloudNFSExportAddHandlerFunc turns a function with the right signature into a cloud n f s export add handler
type CloudNFSExportAddHandlerFunc func(CloudNFSExportAddParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CloudNFSExportAddHandlerFunc) Handle(params CloudNFSExportAddParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CloudNFSExportAddHandler interface for that can handle valid cloud n f s export add params
type CloudNFSExportAddHandler interface {
	Handle(CloudNFSExportAddParams, interface{}) middleware.Responder
}

// NewCloudNFSExportAdd creates a new http.Handler for the cloud n f s export add operation
func NewCloudNFSExportAdd(ctx *middleware.Context, handler CloudNFSExportAddHandler) *CloudNFSExportAdd {
	return &CloudNFSExportAdd{Context: ctx, Handler: handler}
}

/*CloudNFSExportAdd swagger:route POST /cloud/add/{param} staging cloudNFSExportAdd

Request that an nfs export be created for an LRZ cloud instance

Request that an nfs export be created for an LRZ cloud instance

*/
type CloudNFSExportAdd struct {
	Context *middleware.Context
	Handler CloudNFSExportAddHandler
}

func (o *CloudNFSExportAdd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCloudNFSExportAddParams()

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
