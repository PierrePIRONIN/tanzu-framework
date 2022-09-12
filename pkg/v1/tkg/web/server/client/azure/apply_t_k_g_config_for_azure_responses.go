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

// ApplyTKGConfigForAzureReader is a Reader for the ApplyTKGConfigForAzure structure.
type ApplyTKGConfigForAzureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplyTKGConfigForAzureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplyTKGConfigForAzureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApplyTKGConfigForAzureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewApplyTKGConfigForAzureUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApplyTKGConfigForAzureInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewApplyTKGConfigForAzureOK creates a ApplyTKGConfigForAzureOK with default headers values
func NewApplyTKGConfigForAzureOK() *ApplyTKGConfigForAzureOK {
	return &ApplyTKGConfigForAzureOK{}
}

/*ApplyTKGConfigForAzureOK handles this case with default header values.

Apply change to TKG configuration successfully
*/
type ApplyTKGConfigForAzureOK struct {
	Payload *models.ConfigFileInfo
}

func (o *ApplyTKGConfigForAzureOK) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure/tkgconfig][%d] applyTKGConfigForAzureOK  %+v", 200, o.Payload)
}

func (o *ApplyTKGConfigForAzureOK) GetPayload() *models.ConfigFileInfo {
	return o.Payload
}

func (o *ApplyTKGConfigForAzureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigFileInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForAzureBadRequest creates a ApplyTKGConfigForAzureBadRequest with default headers values
func NewApplyTKGConfigForAzureBadRequest() *ApplyTKGConfigForAzureBadRequest {
	return &ApplyTKGConfigForAzureBadRequest{}
}

/*ApplyTKGConfigForAzureBadRequest handles this case with default header values.

Bad request
*/
type ApplyTKGConfigForAzureBadRequest struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForAzureBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure/tkgconfig][%d] applyTKGConfigForAzureBadRequest  %+v", 400, o.Payload)
}

func (o *ApplyTKGConfigForAzureBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForAzureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForAzureUnauthorized creates a ApplyTKGConfigForAzureUnauthorized with default headers values
func NewApplyTKGConfigForAzureUnauthorized() *ApplyTKGConfigForAzureUnauthorized {
	return &ApplyTKGConfigForAzureUnauthorized{}
}

/*ApplyTKGConfigForAzureUnauthorized handles this case with default header values.

Incorrect credentials
*/
type ApplyTKGConfigForAzureUnauthorized struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForAzureUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure/tkgconfig][%d] applyTKGConfigForAzureUnauthorized  %+v", 401, o.Payload)
}

func (o *ApplyTKGConfigForAzureUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForAzureUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForAzureInternalServerError creates a ApplyTKGConfigForAzureInternalServerError with default headers values
func NewApplyTKGConfigForAzureInternalServerError() *ApplyTKGConfigForAzureInternalServerError {
	return &ApplyTKGConfigForAzureInternalServerError{}
}

/*ApplyTKGConfigForAzureInternalServerError handles this case with default header values.

Internal server error
*/
type ApplyTKGConfigForAzureInternalServerError struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForAzureInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/azure/tkgconfig][%d] applyTKGConfigForAzureInternalServerError  %+v", 500, o.Payload)
}

func (o *ApplyTKGConfigForAzureInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForAzureInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
