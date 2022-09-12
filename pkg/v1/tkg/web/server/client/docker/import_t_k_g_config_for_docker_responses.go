// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// ImportTKGConfigForDockerReader is a Reader for the ImportTKGConfigForDocker structure.
type ImportTKGConfigForDockerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportTKGConfigForDockerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportTKGConfigForDockerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportTKGConfigForDockerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportTKGConfigForDockerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImportTKGConfigForDockerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportTKGConfigForDockerOK creates a ImportTKGConfigForDockerOK with default headers values
func NewImportTKGConfigForDockerOK() *ImportTKGConfigForDockerOK {
	return &ImportTKGConfigForDockerOK{}
}

/*ImportTKGConfigForDockerOK handles this case with default header values.

Generated TKG configuration successfully
*/
type ImportTKGConfigForDockerOK struct {
	Payload *models.DockerRegionalClusterParams
}

func (o *ImportTKGConfigForDockerOK) Error() string {
	return fmt.Sprintf("[POST /api/providers/docker/config/import][%d] importTKGConfigForDockerOK  %+v", 200, o.Payload)
}

func (o *ImportTKGConfigForDockerOK) GetPayload() *models.DockerRegionalClusterParams {
	return o.Payload
}

func (o *ImportTKGConfigForDockerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DockerRegionalClusterParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForDockerBadRequest creates a ImportTKGConfigForDockerBadRequest with default headers values
func NewImportTKGConfigForDockerBadRequest() *ImportTKGConfigForDockerBadRequest {
	return &ImportTKGConfigForDockerBadRequest{}
}

/*ImportTKGConfigForDockerBadRequest handles this case with default header values.

Bad request
*/
type ImportTKGConfigForDockerBadRequest struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForDockerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/docker/config/import][%d] importTKGConfigForDockerBadRequest  %+v", 400, o.Payload)
}

func (o *ImportTKGConfigForDockerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForDockerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForDockerUnauthorized creates a ImportTKGConfigForDockerUnauthorized with default headers values
func NewImportTKGConfigForDockerUnauthorized() *ImportTKGConfigForDockerUnauthorized {
	return &ImportTKGConfigForDockerUnauthorized{}
}

/*ImportTKGConfigForDockerUnauthorized handles this case with default header values.

Incorrect credentials
*/
type ImportTKGConfigForDockerUnauthorized struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForDockerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/docker/config/import][%d] importTKGConfigForDockerUnauthorized  %+v", 401, o.Payload)
}

func (o *ImportTKGConfigForDockerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForDockerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForDockerInternalServerError creates a ImportTKGConfigForDockerInternalServerError with default headers values
func NewImportTKGConfigForDockerInternalServerError() *ImportTKGConfigForDockerInternalServerError {
	return &ImportTKGConfigForDockerInternalServerError{}
}

/*ImportTKGConfigForDockerInternalServerError handles this case with default header values.

Internal server error
*/
type ImportTKGConfigForDockerInternalServerError struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForDockerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/docker/config/import][%d] importTKGConfigForDockerInternalServerError  %+v", 500, o.Payload)
}

func (o *ImportTKGConfigForDockerInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForDockerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
