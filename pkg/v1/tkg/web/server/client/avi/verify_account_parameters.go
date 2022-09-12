// Code generated by go-swagger; DO NOT EDIT.

package avi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// NewVerifyAccountParams creates a new VerifyAccountParams object
// with the default values initialized.
func NewVerifyAccountParams() *VerifyAccountParams {
	var ()
	return &VerifyAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyAccountParamsWithTimeout creates a new VerifyAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVerifyAccountParamsWithTimeout(timeout time.Duration) *VerifyAccountParams {
	var ()
	return &VerifyAccountParams{

		timeout: timeout,
	}
}

// NewVerifyAccountParamsWithContext creates a new VerifyAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewVerifyAccountParamsWithContext(ctx context.Context) *VerifyAccountParams {
	var ()
	return &VerifyAccountParams{

		Context: ctx,
	}
}

// NewVerifyAccountParamsWithHTTPClient creates a new VerifyAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVerifyAccountParamsWithHTTPClient(client *http.Client) *VerifyAccountParams {
	var ()
	return &VerifyAccountParams{
		HTTPClient: client,
	}
}

/*VerifyAccountParams contains all the parameters to send to the API endpoint
for the verify account operation typically these are written to a http.Request
*/
type VerifyAccountParams struct {

	/*Credentials
	  Avi controller credentials

	*/
	Credentials *models.AviControllerParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the verify account params
func (o *VerifyAccountParams) WithTimeout(timeout time.Duration) *VerifyAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify account params
func (o *VerifyAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify account params
func (o *VerifyAccountParams) WithContext(ctx context.Context) *VerifyAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify account params
func (o *VerifyAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify account params
func (o *VerifyAccountParams) WithHTTPClient(client *http.Client) *VerifyAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify account params
func (o *VerifyAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentials adds the credentials to the verify account params
func (o *VerifyAccountParams) WithCredentials(credentials *models.AviControllerParams) *VerifyAccountParams {
	o.SetCredentials(credentials)
	return o
}

// SetCredentials adds the credentials to the verify account params
func (o *VerifyAccountParams) SetCredentials(credentials *models.AviControllerParams) {
	o.Credentials = credentials
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credentials != nil {
		if err := r.SetBodyParam(o.Credentials); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
