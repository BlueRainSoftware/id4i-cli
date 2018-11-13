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

// NewVerifyPasswordResetParams creates a new VerifyPasswordResetParams object
// with the default values initialized.
func NewVerifyPasswordResetParams() *VerifyPasswordResetParams {
	var ()
	return &VerifyPasswordResetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyPasswordResetParamsWithTimeout creates a new VerifyPasswordResetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVerifyPasswordResetParamsWithTimeout(timeout time.Duration) *VerifyPasswordResetParams {
	var ()
	return &VerifyPasswordResetParams{

		timeout: timeout,
	}
}

// NewVerifyPasswordResetParamsWithContext creates a new VerifyPasswordResetParams object
// with the default values initialized, and the ability to set a context for a request
func NewVerifyPasswordResetParamsWithContext(ctx context.Context) *VerifyPasswordResetParams {
	var ()
	return &VerifyPasswordResetParams{

		Context: ctx,
	}
}

// NewVerifyPasswordResetParamsWithHTTPClient creates a new VerifyPasswordResetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVerifyPasswordResetParamsWithHTTPClient(client *http.Client) *VerifyPasswordResetParams {
	var ()
	return &VerifyPasswordResetParams{
		HTTPClient: client,
	}
}

/*VerifyPasswordResetParams contains all the parameters to send to the API endpoint
for the verify password reset operation typically these are written to a http.Request
*/
type VerifyPasswordResetParams struct {

	/*VerificationRequest
	  Contains the new password and the verification token to set the new password.

	*/
	VerificationRequest *api_models.PasswordResetVerificationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the verify password reset params
func (o *VerifyPasswordResetParams) WithTimeout(timeout time.Duration) *VerifyPasswordResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify password reset params
func (o *VerifyPasswordResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify password reset params
func (o *VerifyPasswordResetParams) WithContext(ctx context.Context) *VerifyPasswordResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify password reset params
func (o *VerifyPasswordResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify password reset params
func (o *VerifyPasswordResetParams) WithHTTPClient(client *http.Client) *VerifyPasswordResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify password reset params
func (o *VerifyPasswordResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVerificationRequest adds the verificationRequest to the verify password reset params
func (o *VerifyPasswordResetParams) WithVerificationRequest(verificationRequest *api_models.PasswordResetVerificationRequest) *VerifyPasswordResetParams {
	o.SetVerificationRequest(verificationRequest)
	return o
}

// SetVerificationRequest adds the verificationRequest to the verify password reset params
func (o *VerifyPasswordResetParams) SetVerificationRequest(verificationRequest *api_models.PasswordResetVerificationRequest) {
	o.VerificationRequest = verificationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyPasswordResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VerificationRequest != nil {
		if err := r.SetBodyParam(o.VerificationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
