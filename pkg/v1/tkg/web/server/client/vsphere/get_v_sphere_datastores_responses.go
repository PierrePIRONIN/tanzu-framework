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

// GetVSphereDatastoresReader is a Reader for the GetVSphereDatastores structure.
type GetVSphereDatastoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVSphereDatastoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVSphereDatastoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVSphereDatastoresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVSphereDatastoresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVSphereDatastoresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVSphereDatastoresOK creates a GetVSphereDatastoresOK with default headers values
func NewGetVSphereDatastoresOK() *GetVSphereDatastoresOK {
	return &GetVSphereDatastoresOK{}
}

/*GetVSphereDatastoresOK handles this case with default header values.

Successful retrieval of vSphere datastores
*/
type GetVSphereDatastoresOK struct {
	Payload []*models.VSphereDatastore
}

func (o *GetVSphereDatastoresOK) Error() string {
	return fmt.Sprintf("[GET /api/providers/vsphere/datastores][%d] getVSphereDatastoresOK  %+v", 200, o.Payload)
}

func (o *GetVSphereDatastoresOK) GetPayload() []*models.VSphereDatastore {
	return o.Payload
}

func (o *GetVSphereDatastoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVSphereDatastoresBadRequest creates a GetVSphereDatastoresBadRequest with default headers values
func NewGetVSphereDatastoresBadRequest() *GetVSphereDatastoresBadRequest {
	return &GetVSphereDatastoresBadRequest{}
}

/*GetVSphereDatastoresBadRequest handles this case with default header values.

Bad request
*/
type GetVSphereDatastoresBadRequest struct {
	Payload *models.Error
}

func (o *GetVSphereDatastoresBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/providers/vsphere/datastores][%d] getVSphereDatastoresBadRequest  %+v", 400, o.Payload)
}

func (o *GetVSphereDatastoresBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVSphereDatastoresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVSphereDatastoresUnauthorized creates a GetVSphereDatastoresUnauthorized with default headers values
func NewGetVSphereDatastoresUnauthorized() *GetVSphereDatastoresUnauthorized {
	return &GetVSphereDatastoresUnauthorized{}
}

/*GetVSphereDatastoresUnauthorized handles this case with default header values.

Incorrect credentials
*/
type GetVSphereDatastoresUnauthorized struct {
	Payload *models.Error
}

func (o *GetVSphereDatastoresUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/providers/vsphere/datastores][%d] getVSphereDatastoresUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVSphereDatastoresUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVSphereDatastoresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVSphereDatastoresInternalServerError creates a GetVSphereDatastoresInternalServerError with default headers values
func NewGetVSphereDatastoresInternalServerError() *GetVSphereDatastoresInternalServerError {
	return &GetVSphereDatastoresInternalServerError{}
}

/*GetVSphereDatastoresInternalServerError handles this case with default header values.

Internal server error
*/
type GetVSphereDatastoresInternalServerError struct {
	Payload *models.Error
}

func (o *GetVSphereDatastoresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/providers/vsphere/datastores][%d] getVSphereDatastoresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVSphereDatastoresInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVSphereDatastoresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
