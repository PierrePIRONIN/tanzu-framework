// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// ImportTKGConfigForDockerOKCode is the HTTP code returned for type ImportTKGConfigForDockerOK
const ImportTKGConfigForDockerOKCode int = 200

/*ImportTKGConfigForDockerOK Generated TKG configuration successfully

swagger:response importTKGConfigForDockerOK
*/
type ImportTKGConfigForDockerOK struct {

	/*
	  In: Body
	*/
	Payload *models.DockerRegionalClusterParams `json:"body,omitempty"`
}

// NewImportTKGConfigForDockerOK creates ImportTKGConfigForDockerOK with default headers values
func NewImportTKGConfigForDockerOK() *ImportTKGConfigForDockerOK {

	return &ImportTKGConfigForDockerOK{}
}

// WithPayload adds the payload to the import t k g config for docker o k response
func (o *ImportTKGConfigForDockerOK) WithPayload(payload *models.DockerRegionalClusterParams) *ImportTKGConfigForDockerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the import t k g config for docker o k response
func (o *ImportTKGConfigForDockerOK) SetPayload(payload *models.DockerRegionalClusterParams) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImportTKGConfigForDockerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImportTKGConfigForDockerBadRequestCode is the HTTP code returned for type ImportTKGConfigForDockerBadRequest
const ImportTKGConfigForDockerBadRequestCode int = 400

/*ImportTKGConfigForDockerBadRequest Bad request

swagger:response importTKGConfigForDockerBadRequest
*/
type ImportTKGConfigForDockerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewImportTKGConfigForDockerBadRequest creates ImportTKGConfigForDockerBadRequest with default headers values
func NewImportTKGConfigForDockerBadRequest() *ImportTKGConfigForDockerBadRequest {

	return &ImportTKGConfigForDockerBadRequest{}
}

// WithPayload adds the payload to the import t k g config for docker bad request response
func (o *ImportTKGConfigForDockerBadRequest) WithPayload(payload *models.Error) *ImportTKGConfigForDockerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the import t k g config for docker bad request response
func (o *ImportTKGConfigForDockerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImportTKGConfigForDockerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImportTKGConfigForDockerUnauthorizedCode is the HTTP code returned for type ImportTKGConfigForDockerUnauthorized
const ImportTKGConfigForDockerUnauthorizedCode int = 401

/*ImportTKGConfigForDockerUnauthorized Incorrect credentials

swagger:response importTKGConfigForDockerUnauthorized
*/
type ImportTKGConfigForDockerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewImportTKGConfigForDockerUnauthorized creates ImportTKGConfigForDockerUnauthorized with default headers values
func NewImportTKGConfigForDockerUnauthorized() *ImportTKGConfigForDockerUnauthorized {

	return &ImportTKGConfigForDockerUnauthorized{}
}

// WithPayload adds the payload to the import t k g config for docker unauthorized response
func (o *ImportTKGConfigForDockerUnauthorized) WithPayload(payload *models.Error) *ImportTKGConfigForDockerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the import t k g config for docker unauthorized response
func (o *ImportTKGConfigForDockerUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImportTKGConfigForDockerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImportTKGConfigForDockerInternalServerErrorCode is the HTTP code returned for type ImportTKGConfigForDockerInternalServerError
const ImportTKGConfigForDockerInternalServerErrorCode int = 500

/*ImportTKGConfigForDockerInternalServerError Internal server error

swagger:response importTKGConfigForDockerInternalServerError
*/
type ImportTKGConfigForDockerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewImportTKGConfigForDockerInternalServerError creates ImportTKGConfigForDockerInternalServerError with default headers values
func NewImportTKGConfigForDockerInternalServerError() *ImportTKGConfigForDockerInternalServerError {

	return &ImportTKGConfigForDockerInternalServerError{}
}

// WithPayload adds the payload to the import t k g config for docker internal server error response
func (o *ImportTKGConfigForDockerInternalServerError) WithPayload(payload *models.Error) *ImportTKGConfigForDockerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the import t k g config for docker internal server error response
func (o *ImportTKGConfigForDockerInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImportTKGConfigForDockerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
