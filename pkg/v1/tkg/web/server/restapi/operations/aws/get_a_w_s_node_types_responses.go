// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// GetAWSNodeTypesOKCode is the HTTP code returned for type GetAWSNodeTypesOK
const GetAWSNodeTypesOKCode int = 200

/*GetAWSNodeTypesOK Successful retrieval of AWS node types

swagger:response getAWSNodeTypesOK
*/
type GetAWSNodeTypesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetAWSNodeTypesOK creates GetAWSNodeTypesOK with default headers values
func NewGetAWSNodeTypesOK() *GetAWSNodeTypesOK {

	return &GetAWSNodeTypesOK{}
}

// WithPayload adds the payload to the get a w s node types o k response
func (o *GetAWSNodeTypesOK) WithPayload(payload []string) *GetAWSNodeTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s node types o k response
func (o *GetAWSNodeTypesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSNodeTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAWSNodeTypesBadRequestCode is the HTTP code returned for type GetAWSNodeTypesBadRequest
const GetAWSNodeTypesBadRequestCode int = 400

/*GetAWSNodeTypesBadRequest Bad request

swagger:response getAWSNodeTypesBadRequest
*/
type GetAWSNodeTypesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSNodeTypesBadRequest creates GetAWSNodeTypesBadRequest with default headers values
func NewGetAWSNodeTypesBadRequest() *GetAWSNodeTypesBadRequest {

	return &GetAWSNodeTypesBadRequest{}
}

// WithPayload adds the payload to the get a w s node types bad request response
func (o *GetAWSNodeTypesBadRequest) WithPayload(payload *models.Error) *GetAWSNodeTypesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s node types bad request response
func (o *GetAWSNodeTypesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSNodeTypesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAWSNodeTypesUnauthorizedCode is the HTTP code returned for type GetAWSNodeTypesUnauthorized
const GetAWSNodeTypesUnauthorizedCode int = 401

/*GetAWSNodeTypesUnauthorized Incorrect credentials

swagger:response getAWSNodeTypesUnauthorized
*/
type GetAWSNodeTypesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSNodeTypesUnauthorized creates GetAWSNodeTypesUnauthorized with default headers values
func NewGetAWSNodeTypesUnauthorized() *GetAWSNodeTypesUnauthorized {

	return &GetAWSNodeTypesUnauthorized{}
}

// WithPayload adds the payload to the get a w s node types unauthorized response
func (o *GetAWSNodeTypesUnauthorized) WithPayload(payload *models.Error) *GetAWSNodeTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s node types unauthorized response
func (o *GetAWSNodeTypesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSNodeTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAWSNodeTypesInternalServerErrorCode is the HTTP code returned for type GetAWSNodeTypesInternalServerError
const GetAWSNodeTypesInternalServerErrorCode int = 500

/*GetAWSNodeTypesInternalServerError Internal server error

swagger:response getAWSNodeTypesInternalServerError
*/
type GetAWSNodeTypesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSNodeTypesInternalServerError creates GetAWSNodeTypesInternalServerError with default headers values
func NewGetAWSNodeTypesInternalServerError() *GetAWSNodeTypesInternalServerError {

	return &GetAWSNodeTypesInternalServerError{}
}

// WithPayload adds the payload to the get a w s node types internal server error response
func (o *GetAWSNodeTypesInternalServerError) WithPayload(payload *models.Error) *GetAWSNodeTypesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s node types internal server error response
func (o *GetAWSNodeTypesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSNodeTypesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
