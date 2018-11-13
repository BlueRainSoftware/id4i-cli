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

// NewInviteUsersParams creates a new InviteUsersParams object
// with the default values initialized.
func NewInviteUsersParams() *InviteUsersParams {
	var ()
	return &InviteUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInviteUsersParamsWithTimeout creates a new InviteUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInviteUsersParamsWithTimeout(timeout time.Duration) *InviteUsersParams {
	var ()
	return &InviteUsersParams{

		timeout: timeout,
	}
}

// NewInviteUsersParamsWithContext creates a new InviteUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewInviteUsersParamsWithContext(ctx context.Context) *InviteUsersParams {
	var ()
	return &InviteUsersParams{

		Context: ctx,
	}
}

// NewInviteUsersParamsWithHTTPClient creates a new InviteUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInviteUsersParamsWithHTTPClient(client *http.Client) *InviteUsersParams {
	var ()
	return &InviteUsersParams{
		HTTPClient: client,
	}
}

/*InviteUsersParams contains all the parameters to send to the API endpoint
for the invite users operation typically these are written to a http.Request
*/
type InviteUsersParams struct {

	/*InvitationList
	  invitationList

	*/
	InvitationList *api_models.OrganizationUserInvitationListRequest
	/*OrganizationID
	  The namespace of the organization where users should be invited

	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invite users params
func (o *InviteUsersParams) WithTimeout(timeout time.Duration) *InviteUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invite users params
func (o *InviteUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invite users params
func (o *InviteUsersParams) WithContext(ctx context.Context) *InviteUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invite users params
func (o *InviteUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invite users params
func (o *InviteUsersParams) WithHTTPClient(client *http.Client) *InviteUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invite users params
func (o *InviteUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationList adds the invitationList to the invite users params
func (o *InviteUsersParams) WithInvitationList(invitationList *api_models.OrganizationUserInvitationListRequest) *InviteUsersParams {
	o.SetInvitationList(invitationList)
	return o
}

// SetInvitationList adds the invitationList to the invite users params
func (o *InviteUsersParams) SetInvitationList(invitationList *api_models.OrganizationUserInvitationListRequest) {
	o.InvitationList = invitationList
}

// WithOrganizationID adds the organizationID to the invite users params
func (o *InviteUsersParams) WithOrganizationID(organizationID string) *InviteUsersParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the invite users params
func (o *InviteUsersParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *InviteUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvitationList != nil {
		if err := r.SetBodyParam(o.InvitationList); err != nil {
			return err
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}