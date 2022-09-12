// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// GetAzureResourceGroupsOKCode is the HTTP code returned for type GetAzureResourceGroupsOK
const GetAzureResourceGroupsOKCode int = 200

/*GetAzureResourceGroupsOK Successful retrieval of Azure resource groups

swagger:response getAzureResourceGroupsOK
*/
type GetAzureResourceGroupsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AzureResourceGroup `json:"body,omitempty"`
}

// NewGetAzureResourceGroupsOK creates GetAzureResourceGroupsOK with default headers values
func NewGetAzureResourceGroupsOK() *GetAzureResourceGroupsOK {

	return &GetAzureResourceGroupsOK{}
}

// WithPayload adds the payload to the get azure resource groups o k response
func (o *GetAzureResourceGroupsOK) WithPayload(payload []*models.AzureResourceGroup) *GetAzureResourceGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure resource groups o k response
func (o *GetAzureResourceGroupsOK) SetPayload(payload []*models.AzureResourceGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureResourceGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AzureResourceGroup, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAzureResourceGroupsBadRequestCode is the HTTP code returned for type GetAzureResourceGroupsBadRequest
const GetAzureResourceGroupsBadRequestCode int = 400

/*GetAzureResourceGroupsBadRequest Bad Request

swagger:response getAzureResourceGroupsBadRequest
*/
type GetAzureResourceGroupsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureResourceGroupsBadRequest creates GetAzureResourceGroupsBadRequest with default headers values
func NewGetAzureResourceGroupsBadRequest() *GetAzureResourceGroupsBadRequest {

	return &GetAzureResourceGroupsBadRequest{}
}

// WithPayload adds the payload to the get azure resource groups bad request response
func (o *GetAzureResourceGroupsBadRequest) WithPayload(payload *models.Error) *GetAzureResourceGroupsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure resource groups bad request response
func (o *GetAzureResourceGroupsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureResourceGroupsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureResourceGroupsUnauthorizedCode is the HTTP code returned for type GetAzureResourceGroupsUnauthorized
const GetAzureResourceGroupsUnauthorizedCode int = 401

/*GetAzureResourceGroupsUnauthorized Incorrect credentials

swagger:response getAzureResourceGroupsUnauthorized
*/
type GetAzureResourceGroupsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureResourceGroupsUnauthorized creates GetAzureResourceGroupsUnauthorized with default headers values
func NewGetAzureResourceGroupsUnauthorized() *GetAzureResourceGroupsUnauthorized {

	return &GetAzureResourceGroupsUnauthorized{}
}

// WithPayload adds the payload to the get azure resource groups unauthorized response
func (o *GetAzureResourceGroupsUnauthorized) WithPayload(payload *models.Error) *GetAzureResourceGroupsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure resource groups unauthorized response
func (o *GetAzureResourceGroupsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureResourceGroupsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureResourceGroupsInternalServerErrorCode is the HTTP code returned for type GetAzureResourceGroupsInternalServerError
const GetAzureResourceGroupsInternalServerErrorCode int = 500

/*GetAzureResourceGroupsInternalServerError Internal server error

swagger:response getAzureResourceGroupsInternalServerError
*/
type GetAzureResourceGroupsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureResourceGroupsInternalServerError creates GetAzureResourceGroupsInternalServerError with default headers values
func NewGetAzureResourceGroupsInternalServerError() *GetAzureResourceGroupsInternalServerError {

	return &GetAzureResourceGroupsInternalServerError{}
}

// WithPayload adds the payload to the get azure resource groups internal server error response
func (o *GetAzureResourceGroupsInternalServerError) WithPayload(payload *models.Error) *GetAzureResourceGroupsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure resource groups internal server error response
func (o *GetAzureResourceGroupsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureResourceGroupsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
