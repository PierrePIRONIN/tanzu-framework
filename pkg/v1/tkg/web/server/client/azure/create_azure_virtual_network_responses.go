// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// CreateAzureVirtualNetworkReader is a Reader for the CreateAzureVirtualNetwork structure.
type CreateAzureVirtualNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAzureVirtualNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAzureVirtualNetworkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAzureVirtualNetworkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAzureVirtualNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAzureVirtualNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAzureVirtualNetworkCreated creates a CreateAzureVirtualNetworkCreated with default headers values
func NewCreateAzureVirtualNetworkCreated() *CreateAzureVirtualNetworkCreated {
	return &CreateAzureVirtualNetworkCreated{}
}

/*CreateAzureVirtualNetworkCreated handles this case with default header values.

Successfully created Azure Virtual network
*/
type CreateAzureVirtualNetworkCreated struct {
	Payload string
}

func (o *CreateAzureVirtualNetworkCreated) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure/resourcegroups/{resourceGroupName}/vnets][%d] createAzureVirtualNetworkCreated  %+v", 201, o.Payload)
}

func (o *CreateAzureVirtualNetworkCreated) GetPayload() string {
	return o.Payload
}

func (o *CreateAzureVirtualNetworkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureVirtualNetworkBadRequest creates a CreateAzureVirtualNetworkBadRequest with default headers values
func NewCreateAzureVirtualNetworkBadRequest() *CreateAzureVirtualNetworkBadRequest {
	return &CreateAzureVirtualNetworkBadRequest{}
}

/*CreateAzureVirtualNetworkBadRequest handles this case with default header values.

Bad request
*/
type CreateAzureVirtualNetworkBadRequest struct {
	Payload *models.Error
}

func (o *CreateAzureVirtualNetworkBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure/resourcegroups/{resourceGroupName}/vnets][%d] createAzureVirtualNetworkBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAzureVirtualNetworkBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAzureVirtualNetworkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureVirtualNetworkUnauthorized creates a CreateAzureVirtualNetworkUnauthorized with default headers values
func NewCreateAzureVirtualNetworkUnauthorized() *CreateAzureVirtualNetworkUnauthorized {
	return &CreateAzureVirtualNetworkUnauthorized{}
}

/*CreateAzureVirtualNetworkUnauthorized handles this case with default header values.

Incorrect credentials
*/
type CreateAzureVirtualNetworkUnauthorized struct {
	Payload *models.Error
}

func (o *CreateAzureVirtualNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure/resourcegroups/{resourceGroupName}/vnets][%d] createAzureVirtualNetworkUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAzureVirtualNetworkUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAzureVirtualNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureVirtualNetworkInternalServerError creates a CreateAzureVirtualNetworkInternalServerError with default headers values
func NewCreateAzureVirtualNetworkInternalServerError() *CreateAzureVirtualNetworkInternalServerError {
	return &CreateAzureVirtualNetworkInternalServerError{}
}

/*CreateAzureVirtualNetworkInternalServerError handles this case with default header values.

Internal server error
*/
type CreateAzureVirtualNetworkInternalServerError struct {
	Payload *models.Error
}

func (o *CreateAzureVirtualNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure/resourcegroups/{resourceGroupName}/vnets][%d] createAzureVirtualNetworkInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAzureVirtualNetworkInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAzureVirtualNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
