// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewAddUserRolesParams creates a new AddUserRolesParams object
// with the default values initialized.
func NewAddUserRolesParams() *AddUserRolesParams {
	var ()
	return &AddUserRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddUserRolesParamsWithTimeout creates a new AddUserRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddUserRolesParamsWithTimeout(timeout time.Duration) *AddUserRolesParams {
	var ()
	return &AddUserRolesParams{

		timeout: timeout,
	}
}

// NewAddUserRolesParamsWithContext creates a new AddUserRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddUserRolesParamsWithContext(ctx context.Context) *AddUserRolesParams {
	var ()
	return &AddUserRolesParams{

		Context: ctx,
	}
}

// NewAddUserRolesParamsWithHTTPClient creates a new AddUserRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddUserRolesParamsWithHTTPClient(client *http.Client) *AddUserRolesParams {
	var ()
	return &AddUserRolesParams{
		HTTPClient: client,
	}
}

/*AddUserRolesParams contains all the parameters to send to the API endpoint
for the add user roles operation typically these are written to a http.Request
*/
type AddUserRolesParams struct {

	/*ChangeRoleRequest
	  changeRoleRequest

	*/
	ChangeRoleRequest *api_models.ChangeRoleRequest
	/*OrganizationID
	  The namespace of the organization

	*/
	OrganizationID string
	/*Username
	  username

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add user roles params
func (o *AddUserRolesParams) WithTimeout(timeout time.Duration) *AddUserRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add user roles params
func (o *AddUserRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add user roles params
func (o *AddUserRolesParams) WithContext(ctx context.Context) *AddUserRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add user roles params
func (o *AddUserRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add user roles params
func (o *AddUserRolesParams) WithHTTPClient(client *http.Client) *AddUserRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add user roles params
func (o *AddUserRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeRoleRequest adds the changeRoleRequest to the add user roles params
func (o *AddUserRolesParams) WithChangeRoleRequest(changeRoleRequest *api_models.ChangeRoleRequest) *AddUserRolesParams {
	o.SetChangeRoleRequest(changeRoleRequest)
	return o
}

// SetChangeRoleRequest adds the changeRoleRequest to the add user roles params
func (o *AddUserRolesParams) SetChangeRoleRequest(changeRoleRequest *api_models.ChangeRoleRequest) {
	o.ChangeRoleRequest = changeRoleRequest
}

// WithOrganizationID adds the organizationID to the add user roles params
func (o *AddUserRolesParams) WithOrganizationID(organizationID string) *AddUserRolesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the add user roles params
func (o *AddUserRolesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUsername adds the username to the add user roles params
func (o *AddUserRolesParams) WithUsername(username string) *AddUserRolesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the add user roles params
func (o *AddUserRolesParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *AddUserRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangeRoleRequest != nil {
		if err := r.SetBodyParam(o.ChangeRoleRequest); err != nil {
			return err
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
