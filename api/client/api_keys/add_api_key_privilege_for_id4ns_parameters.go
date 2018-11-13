// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

// NewAddAPIKeyPrivilegeForId4nsParams creates a new AddAPIKeyPrivilegeForId4nsParams object
// with the default values initialized.
func NewAddAPIKeyPrivilegeForId4nsParams() *AddAPIKeyPrivilegeForId4nsParams {
	var ()
	return &AddAPIKeyPrivilegeForId4nsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAPIKeyPrivilegeForId4nsParamsWithTimeout creates a new AddAPIKeyPrivilegeForId4nsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAPIKeyPrivilegeForId4nsParamsWithTimeout(timeout time.Duration) *AddAPIKeyPrivilegeForId4nsParams {
	var ()
	return &AddAPIKeyPrivilegeForId4nsParams{

		timeout: timeout,
	}
}

// NewAddAPIKeyPrivilegeForId4nsParamsWithContext creates a new AddAPIKeyPrivilegeForId4nsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAPIKeyPrivilegeForId4nsParamsWithContext(ctx context.Context) *AddAPIKeyPrivilegeForId4nsParams {
	var ()
	return &AddAPIKeyPrivilegeForId4nsParams{

		Context: ctx,
	}
}

// NewAddAPIKeyPrivilegeForId4nsParamsWithHTTPClient creates a new AddAPIKeyPrivilegeForId4nsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAPIKeyPrivilegeForId4nsParamsWithHTTPClient(client *http.Client) *AddAPIKeyPrivilegeForId4nsParams {
	var ()
	return &AddAPIKeyPrivilegeForId4nsParams{
		HTTPClient: client,
	}
}

/*AddAPIKeyPrivilegeForId4nsParams contains all the parameters to send to the API endpoint
for the add Api key privilege for id4ns operation typically these are written to a http.Request
*/
type AddAPIKeyPrivilegeForId4nsParams struct {

	/*Id4ns
	  id4ns

	*/
	Id4ns *api_models.ListOfId4ns
	/*Key
	  key

	*/
	Key string
	/*Privilege
	  privilege

	*/
	Privilege string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) WithTimeout(timeout time.Duration) *AddAPIKeyPrivilegeForId4nsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) WithContext(ctx context.Context) *AddAPIKeyPrivilegeForId4nsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) WithHTTPClient(client *http.Client) *AddAPIKeyPrivilegeForId4nsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithId4ns adds the id4ns to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) WithId4ns(id4ns *api_models.ListOfId4ns) *AddAPIKeyPrivilegeForId4nsParams {
	o.SetId4ns(id4ns)
	return o
}

// SetId4ns adds the id4ns to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) SetId4ns(id4ns *api_models.ListOfId4ns) {
	o.Id4ns = id4ns
}

// WithKey adds the key to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) WithKey(key string) *AddAPIKeyPrivilegeForId4nsParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) SetKey(key string) {
	o.Key = key
}

// WithPrivilege adds the privilege to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) WithPrivilege(privilege string) *AddAPIKeyPrivilegeForId4nsParams {
	o.SetPrivilege(privilege)
	return o
}

// SetPrivilege adds the privilege to the add Api key privilege for id4ns params
func (o *AddAPIKeyPrivilegeForId4nsParams) SetPrivilege(privilege string) {
	o.Privilege = privilege
}

// WriteToRequest writes these params to a swagger request
func (o *AddAPIKeyPrivilegeForId4nsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Id4ns != nil {
		if err := r.SetBodyParam(o.Id4ns); err != nil {
			return err
		}
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	// path param privilege
	if err := r.SetPathParam("privilege", o.Privilege); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
