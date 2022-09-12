// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// SetVSphereEndpointReader is a Reader for the SetVSphereEndpoint structure.
type SetVSphereEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetVSphereEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSetVSphereEndpointCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetVSphereEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetVSphereEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetVSphereEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetVSphereEndpointCreated creates a SetVSphereEndpointCreated with default headers values
func NewSetVSphereEndpointCreated() *SetVSphereEndpointCreated {
	return &SetVSphereEndpointCreated{}
}

/*SetVSphereEndpointCreated handles this case with default header values.

Connection successful
*/
type SetVSphereEndpointCreated struct {
	Payload *models.VsphereInfo
}

func (o *SetVSphereEndpointCreated) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere][%d] setVSphereEndpointCreated  %+v", 201, o.Payload)
}

func (o *SetVSphereEndpointCreated) GetPayload() *models.VsphereInfo {
	return o.Payload
}

func (o *SetVSphereEndpointCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VsphereInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetVSphereEndpointBadRequest creates a SetVSphereEndpointBadRequest with default headers values
func NewSetVSphereEndpointBadRequest() *SetVSphereEndpointBadRequest {
	return &SetVSphereEndpointBadRequest{}
}

/*SetVSphereEndpointBadRequest handles this case with default header values.

Bad request
*/
type SetVSphereEndpointBadRequest struct {
	Payload *models.Error
}

func (o *SetVSphereEndpointBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere][%d] setVSphereEndpointBadRequest  %+v", 400, o.Payload)
}

func (o *SetVSphereEndpointBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetVSphereEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetVSphereEndpointUnauthorized creates a SetVSphereEndpointUnauthorized with default headers values
func NewSetVSphereEndpointUnauthorized() *SetVSphereEndpointUnauthorized {
	return &SetVSphereEndpointUnauthorized{}
}

/*SetVSphereEndpointUnauthorized handles this case with default header values.

Incorrect credentials
*/
type SetVSphereEndpointUnauthorized struct {
	Payload *models.Error
}

func (o *SetVSphereEndpointUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere][%d] setVSphereEndpointUnauthorized  %+v", 401, o.Payload)
}

func (o *SetVSphereEndpointUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetVSphereEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetVSphereEndpointInternalServerError creates a SetVSphereEndpointInternalServerError with default headers values
func NewSetVSphereEndpointInternalServerError() *SetVSphereEndpointInternalServerError {
	return &SetVSphereEndpointInternalServerError{}
}

/*SetVSphereEndpointInternalServerError handles this case with default header values.

Internal server error
*/
type SetVSphereEndpointInternalServerError struct {
	Payload *models.Error
}

func (o *SetVSphereEndpointInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere][%d] setVSphereEndpointInternalServerError  %+v", 500, o.Payload)
}

func (o *SetVSphereEndpointInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetVSphereEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
