// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// GetVSphereOSImagesOKCode is the HTTP code returned for type GetVSphereOSImagesOK
const GetVSphereOSImagesOKCode int = 200

/*GetVSphereOSImagesOK Successful retrieval of node images supported by vSphere

swagger:response getVSphereOSImagesOK
*/
type GetVSphereOSImagesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.VSphereVirtualMachine `json:"body,omitempty"`
}

// NewGetVSphereOSImagesOK creates GetVSphereOSImagesOK with default headers values
func NewGetVSphereOSImagesOK() *GetVSphereOSImagesOK {

	return &GetVSphereOSImagesOK{}
}

// WithPayload adds the payload to the get v sphere o s images o k response
func (o *GetVSphereOSImagesOK) WithPayload(payload []*models.VSphereVirtualMachine) *GetVSphereOSImagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere o s images o k response
func (o *GetVSphereOSImagesOK) SetPayload(payload []*models.VSphereVirtualMachine) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereOSImagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.VSphereVirtualMachine, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetVSphereOSImagesBadRequestCode is the HTTP code returned for type GetVSphereOSImagesBadRequest
const GetVSphereOSImagesBadRequestCode int = 400

/*GetVSphereOSImagesBadRequest Bad request

swagger:response getVSphereOSImagesBadRequest
*/
type GetVSphereOSImagesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereOSImagesBadRequest creates GetVSphereOSImagesBadRequest with default headers values
func NewGetVSphereOSImagesBadRequest() *GetVSphereOSImagesBadRequest {

	return &GetVSphereOSImagesBadRequest{}
}

// WithPayload adds the payload to the get v sphere o s images bad request response
func (o *GetVSphereOSImagesBadRequest) WithPayload(payload *models.Error) *GetVSphereOSImagesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere o s images bad request response
func (o *GetVSphereOSImagesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereOSImagesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVSphereOSImagesUnauthorizedCode is the HTTP code returned for type GetVSphereOSImagesUnauthorized
const GetVSphereOSImagesUnauthorizedCode int = 401

/*GetVSphereOSImagesUnauthorized Incorrect credentials

swagger:response getVSphereOSImagesUnauthorized
*/
type GetVSphereOSImagesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereOSImagesUnauthorized creates GetVSphereOSImagesUnauthorized with default headers values
func NewGetVSphereOSImagesUnauthorized() *GetVSphereOSImagesUnauthorized {

	return &GetVSphereOSImagesUnauthorized{}
}

// WithPayload adds the payload to the get v sphere o s images unauthorized response
func (o *GetVSphereOSImagesUnauthorized) WithPayload(payload *models.Error) *GetVSphereOSImagesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere o s images unauthorized response
func (o *GetVSphereOSImagesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereOSImagesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVSphereOSImagesInternalServerErrorCode is the HTTP code returned for type GetVSphereOSImagesInternalServerError
const GetVSphereOSImagesInternalServerErrorCode int = 500

/*GetVSphereOSImagesInternalServerError Internal server error

swagger:response getVSphereOSImagesInternalServerError
*/
type GetVSphereOSImagesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereOSImagesInternalServerError creates GetVSphereOSImagesInternalServerError with default headers values
func NewGetVSphereOSImagesInternalServerError() *GetVSphereOSImagesInternalServerError {

	return &GetVSphereOSImagesInternalServerError{}
}

// WithPayload adds the payload to the get v sphere o s images internal server error response
func (o *GetVSphereOSImagesInternalServerError) WithPayload(payload *models.Error) *GetVSphereOSImagesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere o s images internal server error response
func (o *GetVSphereOSImagesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereOSImagesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
