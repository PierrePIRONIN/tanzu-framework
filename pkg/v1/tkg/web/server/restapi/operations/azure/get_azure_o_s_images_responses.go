// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// GetAzureOSImagesOKCode is the HTTP code returned for type GetAzureOSImagesOK
const GetAzureOSImagesOKCode int = 200

/*GetAzureOSImagesOK Successful retrieval of Azure supported os images

swagger:response getAzureOSImagesOK
*/
type GetAzureOSImagesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AzureVirtualMachine `json:"body,omitempty"`
}

// NewGetAzureOSImagesOK creates GetAzureOSImagesOK with default headers values
func NewGetAzureOSImagesOK() *GetAzureOSImagesOK {

	return &GetAzureOSImagesOK{}
}

// WithPayload adds the payload to the get azure o s images o k response
func (o *GetAzureOSImagesOK) WithPayload(payload []*models.AzureVirtualMachine) *GetAzureOSImagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure o s images o k response
func (o *GetAzureOSImagesOK) SetPayload(payload []*models.AzureVirtualMachine) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureOSImagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AzureVirtualMachine, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAzureOSImagesBadRequestCode is the HTTP code returned for type GetAzureOSImagesBadRequest
const GetAzureOSImagesBadRequestCode int = 400

/*GetAzureOSImagesBadRequest Bad request

swagger:response getAzureOSImagesBadRequest
*/
type GetAzureOSImagesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureOSImagesBadRequest creates GetAzureOSImagesBadRequest with default headers values
func NewGetAzureOSImagesBadRequest() *GetAzureOSImagesBadRequest {

	return &GetAzureOSImagesBadRequest{}
}

// WithPayload adds the payload to the get azure o s images bad request response
func (o *GetAzureOSImagesBadRequest) WithPayload(payload *models.Error) *GetAzureOSImagesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure o s images bad request response
func (o *GetAzureOSImagesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureOSImagesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureOSImagesUnauthorizedCode is the HTTP code returned for type GetAzureOSImagesUnauthorized
const GetAzureOSImagesUnauthorizedCode int = 401

/*GetAzureOSImagesUnauthorized Incorrect credentials

swagger:response getAzureOSImagesUnauthorized
*/
type GetAzureOSImagesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureOSImagesUnauthorized creates GetAzureOSImagesUnauthorized with default headers values
func NewGetAzureOSImagesUnauthorized() *GetAzureOSImagesUnauthorized {

	return &GetAzureOSImagesUnauthorized{}
}

// WithPayload adds the payload to the get azure o s images unauthorized response
func (o *GetAzureOSImagesUnauthorized) WithPayload(payload *models.Error) *GetAzureOSImagesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure o s images unauthorized response
func (o *GetAzureOSImagesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureOSImagesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureOSImagesInternalServerErrorCode is the HTTP code returned for type GetAzureOSImagesInternalServerError
const GetAzureOSImagesInternalServerErrorCode int = 500

/*GetAzureOSImagesInternalServerError Internal server error

swagger:response getAzureOSImagesInternalServerError
*/
type GetAzureOSImagesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureOSImagesInternalServerError creates GetAzureOSImagesInternalServerError with default headers values
func NewGetAzureOSImagesInternalServerError() *GetAzureOSImagesInternalServerError {

	return &GetAzureOSImagesInternalServerError{}
}

// WithPayload adds the payload to the get azure o s images internal server error response
func (o *GetAzureOSImagesInternalServerError) WithPayload(payload *models.Error) *GetAzureOSImagesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure o s images internal server error response
func (o *GetAzureOSImagesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureOSImagesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
