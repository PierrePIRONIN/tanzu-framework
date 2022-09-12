// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// NewVerifyLdapConnectParams creates a new VerifyLdapConnectParams object
// no default values defined in spec.
func NewVerifyLdapConnectParams() VerifyLdapConnectParams {

	return VerifyLdapConnectParams{}
}

// VerifyLdapConnectParams contains all the bound params for the verify ldap connect operation
// typically these are obtained from a http.Request
//
// swagger:parameters verifyLdapConnect
type VerifyLdapConnectParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*LDAP configuration
	  In: body
	*/
	Credentials *models.LdapParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewVerifyLdapConnectParams() beforehand.
func (o *VerifyLdapConnectParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.LdapParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("credentials", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Credentials = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
