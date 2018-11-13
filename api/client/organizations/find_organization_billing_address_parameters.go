// Code generated by go-swagger; DO NOT EDIT.

package organizations

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
)

// NewFindOrganizationBillingAddressParams creates a new FindOrganizationBillingAddressParams object
// with the default values initialized.
func NewFindOrganizationBillingAddressParams() *FindOrganizationBillingAddressParams {
	var ()
	return &FindOrganizationBillingAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindOrganizationBillingAddressParamsWithTimeout creates a new FindOrganizationBillingAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindOrganizationBillingAddressParamsWithTimeout(timeout time.Duration) *FindOrganizationBillingAddressParams {
	var ()
	return &FindOrganizationBillingAddressParams{

		timeout: timeout,
	}
}

// NewFindOrganizationBillingAddressParamsWithContext creates a new FindOrganizationBillingAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindOrganizationBillingAddressParamsWithContext(ctx context.Context) *FindOrganizationBillingAddressParams {
	var ()
	return &FindOrganizationBillingAddressParams{

		Context: ctx,
	}
}

// NewFindOrganizationBillingAddressParamsWithHTTPClient creates a new FindOrganizationBillingAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindOrganizationBillingAddressParamsWithHTTPClient(client *http.Client) *FindOrganizationBillingAddressParams {
	var ()
	return &FindOrganizationBillingAddressParams{
		HTTPClient: client,
	}
}

/*FindOrganizationBillingAddressParams contains all the parameters to send to the API endpoint
for the find organization billing address operation typically these are written to a http.Request
*/
type FindOrganizationBillingAddressParams struct {

	/*OrganizationID
	  organizationId

	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find organization billing address params
func (o *FindOrganizationBillingAddressParams) WithTimeout(timeout time.Duration) *FindOrganizationBillingAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find organization billing address params
func (o *FindOrganizationBillingAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find organization billing address params
func (o *FindOrganizationBillingAddressParams) WithContext(ctx context.Context) *FindOrganizationBillingAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find organization billing address params
func (o *FindOrganizationBillingAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find organization billing address params
func (o *FindOrganizationBillingAddressParams) WithHTTPClient(client *http.Client) *FindOrganizationBillingAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find organization billing address params
func (o *FindOrganizationBillingAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the find organization billing address params
func (o *FindOrganizationBillingAddressParams) WithOrganizationID(organizationID string) *FindOrganizationBillingAddressParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the find organization billing address params
func (o *FindOrganizationBillingAddressParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *FindOrganizationBillingAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}