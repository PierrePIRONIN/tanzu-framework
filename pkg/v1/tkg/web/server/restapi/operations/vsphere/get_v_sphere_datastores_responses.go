// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// GetVSphereDatastoresOKCode is the HTTP code returned for type GetVSphereDatastoresOK
const GetVSphereDatastoresOKCode int = 200

/*GetVSphereDatastoresOK Successful retrieval of vSphere datastores

swagger:response getVSphereDatastoresOK
*/
type GetVSphereDatastoresOK struct {

	/*
	  In: Body
	*/
	Payload []*models.VSphereDatastore `json:"body,omitempty"`
}

// NewGetVSphereDatastoresOK creates GetVSphereDatastoresOK with default headers values
func NewGetVSphereDatastoresOK() *GetVSphereDatastoresOK {

	return &GetVSphereDatastoresOK{}
}

// WithPayload adds the payload to the get v sphere datastores o k response
func (o *GetVSphereDatastoresOK) WithPayload(payload []*models.VSphereDatastore) *GetVSphereDatastoresOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere datastores o k response
func (o *GetVSphereDatastoresOK) SetPayload(payload []*models.VSphereDatastore) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereDatastoresOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.VSphereDatastore, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetVSphereDatastoresBadRequestCode is the HTTP code returned for type GetVSphereDatastoresBadRequest
const GetVSphereDatastoresBadRequestCode int = 400

/*GetVSphereDatastoresBadRequest Bad request

swagger:response getVSphereDatastoresBadRequest
*/
type GetVSphereDatastoresBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereDatastoresBadRequest creates GetVSphereDatastoresBadRequest with default headers values
func NewGetVSphereDatastoresBadRequest() *GetVSphereDatastoresBadRequest {

	return &GetVSphereDatastoresBadRequest{}
}

// WithPayload adds the payload to the get v sphere datastores bad request response
func (o *GetVSphereDatastoresBadRequest) WithPayload(payload *models.Error) *GetVSphereDatastoresBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere datastores bad request response
func (o *GetVSphereDatastoresBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereDatastoresBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVSphereDatastoresUnauthorizedCode is the HTTP code returned for type GetVSphereDatastoresUnauthorized
const GetVSphereDatastoresUnauthorizedCode int = 401

/*GetVSphereDatastoresUnauthorized Incorrect credentials

swagger:response getVSphereDatastoresUnauthorized
*/
type GetVSphereDatastoresUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereDatastoresUnauthorized creates GetVSphereDatastoresUnauthorized with default headers values
func NewGetVSphereDatastoresUnauthorized() *GetVSphereDatastoresUnauthorized {

	return &GetVSphereDatastoresUnauthorized{}
}

// WithPayload adds the payload to the get v sphere datastores unauthorized response
func (o *GetVSphereDatastoresUnauthorized) WithPayload(payload *models.Error) *GetVSphereDatastoresUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere datastores unauthorized response
func (o *GetVSphereDatastoresUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereDatastoresUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVSphereDatastoresInternalServerErrorCode is the HTTP code returned for type GetVSphereDatastoresInternalServerError
const GetVSphereDatastoresInternalServerErrorCode int = 500

/*GetVSphereDatastoresInternalServerError Internal server error

swagger:response getVSphereDatastoresInternalServerError
*/
type GetVSphereDatastoresInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereDatastoresInternalServerError creates GetVSphereDatastoresInternalServerError with default headers values
func NewGetVSphereDatastoresInternalServerError() *GetVSphereDatastoresInternalServerError {

	return &GetVSphereDatastoresInternalServerError{}
}

// WithPayload adds the payload to the get v sphere datastores internal server error response
func (o *GetVSphereDatastoresInternalServerError) WithPayload(payload *models.Error) *GetVSphereDatastoresInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere datastores internal server error response
func (o *GetVSphereDatastoresInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereDatastoresInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
