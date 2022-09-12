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

// SetAzureEndpointReader is a Reader for the SetAzureEndpoint structure.
type SetAzureEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAzureEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSetAzureEndpointCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetAzureEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetAzureEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetAzureEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAzureEndpointCreated creates a SetAzureEndpointCreated with default headers values
func NewSetAzureEndpointCreated() *SetAzureEndpointCreated {
	return &SetAzureEndpointCreated{}
}

/*SetAzureEndpointCreated handles this case with default header values.

Connection successful
*/
type SetAzureEndpointCreated struct {
}

func (o *SetAzureEndpointCreated) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure][%d] setAzureEndpointCreated ", 201)
}

func (o *SetAzureEndpointCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetAzureEndpointBadRequest creates a SetAzureEndpointBadRequest with default headers values
func NewSetAzureEndpointBadRequest() *SetAzureEndpointBadRequest {
	return &SetAzureEndpointBadRequest{}
}

/*SetAzureEndpointBadRequest handles this case with default header values.

Bad request
*/
type SetAzureEndpointBadRequest struct {
	Payload *models.Error
}

func (o *SetAzureEndpointBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure][%d] setAzureEndpointBadRequest  %+v", 400, o.Payload)
}

func (o *SetAzureEndpointBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAzureEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAzureEndpointUnauthorized creates a SetAzureEndpointUnauthorized with default headers values
func NewSetAzureEndpointUnauthorized() *SetAzureEndpointUnauthorized {
	return &SetAzureEndpointUnauthorized{}
}

/*SetAzureEndpointUnauthorized handles this case with default header values.

Incorrect credentials
*/
type SetAzureEndpointUnauthorized struct {
	Payload *models.Error
}

func (o *SetAzureEndpointUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure][%d] setAzureEndpointUnauthorized  %+v", 401, o.Payload)
}

func (o *SetAzureEndpointUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAzureEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAzureEndpointInternalServerError creates a SetAzureEndpointInternalServerError with default headers values
func NewSetAzureEndpointInternalServerError() *SetAzureEndpointInternalServerError {
	return &SetAzureEndpointInternalServerError{}
}

/*SetAzureEndpointInternalServerError handles this case with default header values.

Internal server error
*/
type SetAzureEndpointInternalServerError struct {
	Payload *models.Error
}

func (o *SetAzureEndpointInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure][%d] setAzureEndpointInternalServerError  %+v", 500, o.Payload)
}

func (o *SetAzureEndpointInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAzureEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
