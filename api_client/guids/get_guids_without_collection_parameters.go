// Code generated by go-swagger; DO NOT EDIT.

package guids

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

// NewGETGuidsWithoutCollectionParams creates a new GETGuidsWithoutCollectionParams object
// with the default values initialized.
func NewGETGuidsWithoutCollectionParams() *GETGuidsWithoutCollectionParams {
	var ()
	return &GETGuidsWithoutCollectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETGuidsWithoutCollectionParamsWithTimeout creates a new GETGuidsWithoutCollectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETGuidsWithoutCollectionParamsWithTimeout(timeout time.Duration) *GETGuidsWithoutCollectionParams {
	var ()
	return &GETGuidsWithoutCollectionParams{

		timeout: timeout,
	}
}

// NewGETGuidsWithoutCollectionParamsWithContext creates a new GETGuidsWithoutCollectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETGuidsWithoutCollectionParamsWithContext(ctx context.Context) *GETGuidsWithoutCollectionParams {
	var ()
	return &GETGuidsWithoutCollectionParams{

		Context: ctx,
	}
}

// NewGETGuidsWithoutCollectionParamsWithHTTPClient creates a new GETGuidsWithoutCollectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETGuidsWithoutCollectionParamsWithHTTPClient(client *http.Client) *GETGuidsWithoutCollectionParams {
	var ()
	return &GETGuidsWithoutCollectionParams{
		HTTPClient: client,
	}
}

/*GETGuidsWithoutCollectionParams contains all the parameters to send to the API endpoint
for the get guids without collection operation typically these are written to a http.Request
*/
type GETGuidsWithoutCollectionParams struct {

	/*Limit
	  The maximum count of returned elements

	*/
	Limit *int32
	/*Offset
	  Start with the n-th element

	*/
	Offset *int32
	/*OrganizationID
	  The namespace of the organization to search GUIDs for

	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) WithTimeout(timeout time.Duration) *GETGuidsWithoutCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) WithContext(ctx context.Context) *GETGuidsWithoutCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) WithHTTPClient(client *http.Client) *GETGuidsWithoutCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) WithLimit(limit *int32) *GETGuidsWithoutCollectionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) WithOffset(offset *int32) *GETGuidsWithoutCollectionParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) WithOrganizationID(organizationID string) *GETGuidsWithoutCollectionParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get guids without collection params
func (o *GETGuidsWithoutCollectionParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GETGuidsWithoutCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param organizationId
	qrOrganizationID := o.OrganizationID
	qOrganizationID := qrOrganizationID
	if qOrganizationID != "" {
		if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
