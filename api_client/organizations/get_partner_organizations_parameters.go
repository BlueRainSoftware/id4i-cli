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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETPartnerOrganizationsParams creates a new GETPartnerOrganizationsParams object
// with the default values initialized.
func NewGETPartnerOrganizationsParams() *GETPartnerOrganizationsParams {
	var ()
	return &GETPartnerOrganizationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETPartnerOrganizationsParamsWithTimeout creates a new GETPartnerOrganizationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETPartnerOrganizationsParamsWithTimeout(timeout time.Duration) *GETPartnerOrganizationsParams {
	var ()
	return &GETPartnerOrganizationsParams{

		timeout: timeout,
	}
}

// NewGETPartnerOrganizationsParamsWithContext creates a new GETPartnerOrganizationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETPartnerOrganizationsParamsWithContext(ctx context.Context) *GETPartnerOrganizationsParams {
	var ()
	return &GETPartnerOrganizationsParams{

		Context: ctx,
	}
}

// NewGETPartnerOrganizationsParamsWithHTTPClient creates a new GETPartnerOrganizationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETPartnerOrganizationsParamsWithHTTPClient(client *http.Client) *GETPartnerOrganizationsParams {
	var ()
	return &GETPartnerOrganizationsParams{
		HTTPClient: client,
	}
}

/*GETPartnerOrganizationsParams contains all the parameters to send to the API endpoint
for the get partner organizations operation typically these are written to a http.Request
*/
type GETPartnerOrganizationsParams struct {

	/*Limit
	  The maximum count of returned elements

	*/
	Limit *int32
	/*Offset
	  Start with the n-th element

	*/
	Offset *int32
	/*OrganizationID
	  The namespace of the organization to query partner organizations

	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get partner organizations params
func (o *GETPartnerOrganizationsParams) WithTimeout(timeout time.Duration) *GETPartnerOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get partner organizations params
func (o *GETPartnerOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get partner organizations params
func (o *GETPartnerOrganizationsParams) WithContext(ctx context.Context) *GETPartnerOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get partner organizations params
func (o *GETPartnerOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get partner organizations params
func (o *GETPartnerOrganizationsParams) WithHTTPClient(client *http.Client) *GETPartnerOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get partner organizations params
func (o *GETPartnerOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get partner organizations params
func (o *GETPartnerOrganizationsParams) WithLimit(limit *int32) *GETPartnerOrganizationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get partner organizations params
func (o *GETPartnerOrganizationsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get partner organizations params
func (o *GETPartnerOrganizationsParams) WithOffset(offset *int32) *GETPartnerOrganizationsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get partner organizations params
func (o *GETPartnerOrganizationsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get partner organizations params
func (o *GETPartnerOrganizationsParams) WithOrganizationID(organizationID string) *GETPartnerOrganizationsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get partner organizations params
func (o *GETPartnerOrganizationsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GETPartnerOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
