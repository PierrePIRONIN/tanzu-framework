// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// CreateAzureVirtualNetworkCreatedCode is the HTTP code returned for type CreateAzureVirtualNetworkCreated
const CreateAzureVirtualNetworkCreatedCode int = 201

/*CreateAzureVirtualNetworkCreated Successfully created Azure Virtual network

swagger:response createAzureVirtualNetworkCreated
*/
type CreateAzureVirtualNetworkCreated struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateAzureVirtualNetworkCreated creates CreateAzureVirtualNetworkCreated with default headers values
func NewCreateAzureVirtualNetworkCreated() *CreateAzureVirtualNetworkCreated {

	return &CreateAzureVirtualNetworkCreated{}
}

// WithPayload adds the payload to the create azure virtual network created response
func (o *CreateAzureVirtualNetworkCreated) WithPayload(payload string) *CreateAzureVirtualNetworkCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create azure virtual network created response
func (o *CreateAzureVirtualNetworkCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAzureVirtualNetworkCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateAzureVirtualNetworkBadRequestCode is the HTTP code returned for type CreateAzureVirtualNetworkBadRequest
const CreateAzureVirtualNetworkBadRequestCode int = 400

/*CreateAzureVirtualNetworkBadRequest Bad request

swagger:response createAzureVirtualNetworkBadRequest
*/
type CreateAzureVirtualNetworkBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateAzureVirtualNetworkBadRequest creates CreateAzureVirtualNetworkBadRequest with default headers values
func NewCreateAzureVirtualNetworkBadRequest() *CreateAzureVirtualNetworkBadRequest {

	return &CreateAzureVirtualNetworkBadRequest{}
}

// WithPayload adds the payload to the create azure virtual network bad request response
func (o *CreateAzureVirtualNetworkBadRequest) WithPayload(payload *models.Error) *CreateAzureVirtualNetworkBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create azure virtual network bad request response
func (o *CreateAzureVirtualNetworkBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAzureVirtualNetworkBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAzureVirtualNetworkUnauthorizedCode is the HTTP code returned for type CreateAzureVirtualNetworkUnauthorized
const CreateAzureVirtualNetworkUnauthorizedCode int = 401

/*CreateAzureVirtualNetworkUnauthorized Incorrect credentials

swagger:response createAzureVirtualNetworkUnauthorized
*/
type CreateAzureVirtualNetworkUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateAzureVirtualNetworkUnauthorized creates CreateAzureVirtualNetworkUnauthorized with default headers values
func NewCreateAzureVirtualNetworkUnauthorized() *CreateAzureVirtualNetworkUnauthorized {

	return &CreateAzureVirtualNetworkUnauthorized{}
}

// WithPayload adds the payload to the create azure virtual network unauthorized response
func (o *CreateAzureVirtualNetworkUnauthorized) WithPayload(payload *models.Error) *CreateAzureVirtualNetworkUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create azure virtual network unauthorized response
func (o *CreateAzureVirtualNetworkUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAzureVirtualNetworkUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAzureVirtualNetworkInternalServerErrorCode is the HTTP code returned for type CreateAzureVirtualNetworkInternalServerError
const CreateAzureVirtualNetworkInternalServerErrorCode int = 500

/*CreateAzureVirtualNetworkInternalServerError Internal server error

swagger:response createAzureVirtualNetworkInternalServerError
*/
type CreateAzureVirtualNetworkInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateAzureVirtualNetworkInternalServerError creates CreateAzureVirtualNetworkInternalServerError with default headers values
func NewCreateAzureVirtualNetworkInternalServerError() *CreateAzureVirtualNetworkInternalServerError {

	return &CreateAzureVirtualNetworkInternalServerError{}
}

// WithPayload adds the payload to the create azure virtual network internal server error response
func (o *CreateAzureVirtualNetworkInternalServerError) WithPayload(payload *models.Error) *CreateAzureVirtualNetworkInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create azure virtual network internal server error response
func (o *CreateAzureVirtualNetworkInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAzureVirtualNetworkInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
