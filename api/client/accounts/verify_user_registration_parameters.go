// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// NewVerifyUserRegistrationParams creates a new VerifyUserRegistrationParams object
// with the default values initialized.
func NewVerifyUserRegistrationParams() *VerifyUserRegistrationParams {
	var ()
	return &VerifyUserRegistrationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyUserRegistrationParamsWithTimeout creates a new VerifyUserRegistrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVerifyUserRegistrationParamsWithTimeout(timeout time.Duration) *VerifyUserRegistrationParams {
	var ()
	return &VerifyUserRegistrationParams{

		timeout: timeout,
	}
}

// NewVerifyUserRegistrationParamsWithContext creates a new VerifyUserRegistrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewVerifyUserRegistrationParamsWithContext(ctx context.Context) *VerifyUserRegistrationParams {
	var ()
	return &VerifyUserRegistrationParams{

		Context: ctx,
	}
}

// NewVerifyUserRegistrationParamsWithHTTPClient creates a new VerifyUserRegistrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVerifyUserRegistrationParamsWithHTTPClient(client *http.Client) *VerifyUserRegistrationParams {
	var ()
	return &VerifyUserRegistrationParams{
		HTTPClient: client,
	}
}

/*VerifyUserRegistrationParams contains all the parameters to send to the API endpoint
for the verify user registration operation typically these are written to a http.Request
*/
type VerifyUserRegistrationParams struct {

	/*Token
	  The token for user verification.

	*/
	Token *api_models.RegistrationVerificationTokenPresentation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the verify user registration params
func (o *VerifyUserRegistrationParams) WithTimeout(timeout time.Duration) *VerifyUserRegistrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify user registration params
func (o *VerifyUserRegistrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify user registration params
func (o *VerifyUserRegistrationParams) WithContext(ctx context.Context) *VerifyUserRegistrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify user registration params
func (o *VerifyUserRegistrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify user registration params
func (o *VerifyUserRegistrationParams) WithHTTPClient(client *http.Client) *VerifyUserRegistrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify user registration params
func (o *VerifyUserRegistrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the verify user registration params
func (o *VerifyUserRegistrationParams) WithToken(token *api_models.RegistrationVerificationTokenPresentation) *VerifyUserRegistrationParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the verify user registration params
func (o *VerifyUserRegistrationParams) SetToken(token *api_models.RegistrationVerificationTokenPresentation) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyUserRegistrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Token != nil {
		if err := r.SetBodyParam(o.Token); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
