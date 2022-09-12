// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// GetVSphereNodeTypesOKCode is the HTTP code returned for type GetVSphereNodeTypesOK
const GetVSphereNodeTypesOKCode int = 200

/*GetVSphereNodeTypesOK Successful retrieval of node types supported by vSphere

swagger:response getVSphereNodeTypesOK
*/
type GetVSphereNodeTypesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.NodeType `json:"body,omitempty"`
}

// NewGetVSphereNodeTypesOK creates GetVSphereNodeTypesOK with default headers values
func NewGetVSphereNodeTypesOK() *GetVSphereNodeTypesOK {

	return &GetVSphereNodeTypesOK{}
}

// WithPayload adds the payload to the get v sphere node types o k response
func (o *GetVSphereNodeTypesOK) WithPayload(payload []*models.NodeType) *GetVSphereNodeTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere node types o k response
func (o *GetVSphereNodeTypesOK) SetPayload(payload []*models.NodeType) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereNodeTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.NodeType, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetVSphereNodeTypesBadRequestCode is the HTTP code returned for type GetVSphereNodeTypesBadRequest
const GetVSphereNodeTypesBadRequestCode int = 400

/*GetVSphereNodeTypesBadRequest Bad request

swagger:response getVSphereNodeTypesBadRequest
*/
type GetVSphereNodeTypesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereNodeTypesBadRequest creates GetVSphereNodeTypesBadRequest with default headers values
func NewGetVSphereNodeTypesBadRequest() *GetVSphereNodeTypesBadRequest {

	return &GetVSphereNodeTypesBadRequest{}
}

// WithPayload adds the payload to the get v sphere node types bad request response
func (o *GetVSphereNodeTypesBadRequest) WithPayload(payload *models.Error) *GetVSphereNodeTypesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere node types bad request response
func (o *GetVSphereNodeTypesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereNodeTypesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVSphereNodeTypesUnauthorizedCode is the HTTP code returned for type GetVSphereNodeTypesUnauthorized
const GetVSphereNodeTypesUnauthorizedCode int = 401

/*GetVSphereNodeTypesUnauthorized Incorrect credentials

swagger:response getVSphereNodeTypesUnauthorized
*/
type GetVSphereNodeTypesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereNodeTypesUnauthorized creates GetVSphereNodeTypesUnauthorized with default headers values
func NewGetVSphereNodeTypesUnauthorized() *GetVSphereNodeTypesUnauthorized {

	return &GetVSphereNodeTypesUnauthorized{}
}

// WithPayload adds the payload to the get v sphere node types unauthorized response
func (o *GetVSphereNodeTypesUnauthorized) WithPayload(payload *models.Error) *GetVSphereNodeTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere node types unauthorized response
func (o *GetVSphereNodeTypesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereNodeTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVSphereNodeTypesInternalServerErrorCode is the HTTP code returned for type GetVSphereNodeTypesInternalServerError
const GetVSphereNodeTypesInternalServerErrorCode int = 500

/*GetVSphereNodeTypesInternalServerError Internal server error

swagger:response getVSphereNodeTypesInternalServerError
*/
type GetVSphereNodeTypesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereNodeTypesInternalServerError creates GetVSphereNodeTypesInternalServerError with default headers values
func NewGetVSphereNodeTypesInternalServerError() *GetVSphereNodeTypesInternalServerError {

	return &GetVSphereNodeTypesInternalServerError{}
}

// WithPayload adds the payload to the get v sphere node types internal server error response
func (o *GetVSphereNodeTypesInternalServerError) WithPayload(payload *models.Error) *GetVSphereNodeTypesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere node types internal server error response
func (o *GetVSphereNodeTypesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereNodeTypesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
