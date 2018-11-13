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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETUsersOfOrganizationParams creates a new GETUsersOfOrganizationParams object
// with the default values initialized.
func NewGETUsersOfOrganizationParams() *GETUsersOfOrganizationParams {
	var ()
	return &GETUsersOfOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETUsersOfOrganizationParamsWithTimeout creates a new GETUsersOfOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETUsersOfOrganizationParamsWithTimeout(timeout time.Duration) *GETUsersOfOrganizationParams {
	var ()
	return &GETUsersOfOrganizationParams{

		timeout: timeout,
	}
}

// NewGETUsersOfOrganizationParamsWithContext creates a new GETUsersOfOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETUsersOfOrganizationParamsWithContext(ctx context.Context) *GETUsersOfOrganizationParams {
	var ()
	return &GETUsersOfOrganizationParams{

		Context: ctx,
	}
}

// NewGETUsersOfOrganizationParamsWithHTTPClient creates a new GETUsersOfOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETUsersOfOrganizationParamsWithHTTPClient(client *http.Client) *GETUsersOfOrganizationParams {
	var ()
	return &GETUsersOfOrganizationParams{
		HTTPClient: client,
	}
}

/*GETUsersOfOrganizationParams contains all the parameters to send to the API endpoint
for the get users of organization operation typically these are written to a http.Request
*/
type GETUsersOfOrganizationParams struct {

	/*Limit
	  The maximum count of returned elements

	*/
	Limit *int32
	/*Offset
	  Start with the n-th element

	*/
	Offset *int32
	/*OrganizationID
	  organizationId

	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users of organization params
func (o *GETUsersOfOrganizationParams) WithTimeout(timeout time.Duration) *GETUsersOfOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users of organization params
func (o *GETUsersOfOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users of organization params
func (o *GETUsersOfOrganizationParams) WithContext(ctx context.Context) *GETUsersOfOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users of organization params
func (o *GETUsersOfOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users of organization params
func (o *GETUsersOfOrganizationParams) WithHTTPClient(client *http.Client) *GETUsersOfOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users of organization params
func (o *GETUsersOfOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get users of organization params
func (o *GETUsersOfOrganizationParams) WithLimit(limit *int32) *GETUsersOfOrganizationParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get users of organization params
func (o *GETUsersOfOrganizationParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get users of organization params
func (o *GETUsersOfOrganizationParams) WithOffset(offset *int32) *GETUsersOfOrganizationParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get users of organization params
func (o *GETUsersOfOrganizationParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get users of organization params
func (o *GETUsersOfOrganizationParams) WithOrganizationID(organizationID string) *GETUsersOfOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get users of organization params
func (o *GETUsersOfOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GETUsersOfOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
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