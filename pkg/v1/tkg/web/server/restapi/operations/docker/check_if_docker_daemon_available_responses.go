// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// CheckIfDockerDaemonAvailableOKCode is the HTTP code returned for type CheckIfDockerDaemonAvailableOK
const CheckIfDockerDaemonAvailableOKCode int = 200

/*CheckIfDockerDaemonAvailableOK Checked the docker daemon status successfully.

swagger:response checkIfDockerDaemonAvailableOK
*/
type CheckIfDockerDaemonAvailableOK struct {

	/*
	  In: Body
	*/
	Payload *models.DockerDaemonStatus `json:"body,omitempty"`
}

// NewCheckIfDockerDaemonAvailableOK creates CheckIfDockerDaemonAvailableOK with default headers values
func NewCheckIfDockerDaemonAvailableOK() *CheckIfDockerDaemonAvailableOK {

	return &CheckIfDockerDaemonAvailableOK{}
}

// WithPayload adds the payload to the check if docker daemon available o k response
func (o *CheckIfDockerDaemonAvailableOK) WithPayload(payload *models.DockerDaemonStatus) *CheckIfDockerDaemonAvailableOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check if docker daemon available o k response
func (o *CheckIfDockerDaemonAvailableOK) SetPayload(payload *models.DockerDaemonStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckIfDockerDaemonAvailableOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckIfDockerDaemonAvailableBadRequestCode is the HTTP code returned for type CheckIfDockerDaemonAvailableBadRequest
const CheckIfDockerDaemonAvailableBadRequestCode int = 400

/*CheckIfDockerDaemonAvailableBadRequest Bad request

swagger:response checkIfDockerDaemonAvailableBadRequest
*/
type CheckIfDockerDaemonAvailableBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckIfDockerDaemonAvailableBadRequest creates CheckIfDockerDaemonAvailableBadRequest with default headers values
func NewCheckIfDockerDaemonAvailableBadRequest() *CheckIfDockerDaemonAvailableBadRequest {

	return &CheckIfDockerDaemonAvailableBadRequest{}
}

// WithPayload adds the payload to the check if docker daemon available bad request response
func (o *CheckIfDockerDaemonAvailableBadRequest) WithPayload(payload *models.Error) *CheckIfDockerDaemonAvailableBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check if docker daemon available bad request response
func (o *CheckIfDockerDaemonAvailableBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckIfDockerDaemonAvailableBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckIfDockerDaemonAvailableInternalServerErrorCode is the HTTP code returned for type CheckIfDockerDaemonAvailableInternalServerError
const CheckIfDockerDaemonAvailableInternalServerErrorCode int = 500

/*CheckIfDockerDaemonAvailableInternalServerError Internal server error

swagger:response checkIfDockerDaemonAvailableInternalServerError
*/
type CheckIfDockerDaemonAvailableInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckIfDockerDaemonAvailableInternalServerError creates CheckIfDockerDaemonAvailableInternalServerError with default headers values
func NewCheckIfDockerDaemonAvailableInternalServerError() *CheckIfDockerDaemonAvailableInternalServerError {

	return &CheckIfDockerDaemonAvailableInternalServerError{}
}

// WithPayload adds the payload to the check if docker daemon available internal server error response
func (o *CheckIfDockerDaemonAvailableInternalServerError) WithPayload(payload *models.Error) *CheckIfDockerDaemonAvailableInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check if docker daemon available internal server error response
func (o *CheckIfDockerDaemonAvailableInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckIfDockerDaemonAvailableInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
