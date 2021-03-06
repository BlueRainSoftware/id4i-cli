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

// NewGETAllCollectionsOfOrganizationParams creates a new GETAllCollectionsOfOrganizationParams object
// with the default values initialized.
func NewGETAllCollectionsOfOrganizationParams() *GETAllCollectionsOfOrganizationParams {
	var ()
	return &GETAllCollectionsOfOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETAllCollectionsOfOrganizationParamsWithTimeout creates a new GETAllCollectionsOfOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETAllCollectionsOfOrganizationParamsWithTimeout(timeout time.Duration) *GETAllCollectionsOfOrganizationParams {
	var ()
	return &GETAllCollectionsOfOrganizationParams{

		timeout: timeout,
	}
}

// NewGETAllCollectionsOfOrganizationParamsWithContext creates a new GETAllCollectionsOfOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETAllCollectionsOfOrganizationParamsWithContext(ctx context.Context) *GETAllCollectionsOfOrganizationParams {
	var ()
	return &GETAllCollectionsOfOrganizationParams{

		Context: ctx,
	}
}

// NewGETAllCollectionsOfOrganizationParamsWithHTTPClient creates a new GETAllCollectionsOfOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETAllCollectionsOfOrganizationParamsWithHTTPClient(client *http.Client) *GETAllCollectionsOfOrganizationParams {
	var ()
	return &GETAllCollectionsOfOrganizationParams{
		HTTPClient: client,
	}
}

/*GETAllCollectionsOfOrganizationParams contains all the parameters to send to the API endpoint
for the get all collections of organization operation typically these are written to a http.Request
*/
type GETAllCollectionsOfOrganizationParams struct {

	/*Label
	  Filter by this label

	*/
	Label *string
	/*LabelPrefix
	  Filter by this label prefix

	*/
	LabelPrefix *string
	/*Limit
	  The maximum count of returned elements

	*/
	Limit *int32
	/*Offset
	  Start with the n-th element

	*/
	Offset *int32
	/*OrganizationID
	  The namespace of the organization

	*/
	OrganizationID string
	/*Type
	  Filter by this type

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithTimeout(timeout time.Duration) *GETAllCollectionsOfOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithContext(ctx context.Context) *GETAllCollectionsOfOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithHTTPClient(client *http.Client) *GETAllCollectionsOfOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabel adds the label to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithLabel(label *string) *GETAllCollectionsOfOrganizationParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetLabel(label *string) {
	o.Label = label
}

// WithLabelPrefix adds the labelPrefix to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithLabelPrefix(labelPrefix *string) *GETAllCollectionsOfOrganizationParams {
	o.SetLabelPrefix(labelPrefix)
	return o
}

// SetLabelPrefix adds the labelPrefix to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetLabelPrefix(labelPrefix *string) {
	o.LabelPrefix = labelPrefix
}

// WithLimit adds the limit to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithLimit(limit *int32) *GETAllCollectionsOfOrganizationParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithOffset(offset *int32) *GETAllCollectionsOfOrganizationParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithOrganizationID(organizationID string) *GETAllCollectionsOfOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithType adds the typeVar to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) WithType(typeVar *string) *GETAllCollectionsOfOrganizationParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get all collections of organization params
func (o *GETAllCollectionsOfOrganizationParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GETAllCollectionsOfOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Label != nil {

		// query param label
		var qrLabel string
		if o.Label != nil {
			qrLabel = *o.Label
		}
		qLabel := qrLabel
		if qLabel != "" {
			if err := r.SetQueryParam("label", qLabel); err != nil {
				return err
			}
		}

	}

	if o.LabelPrefix != nil {

		// query param labelPrefix
		var qrLabelPrefix string
		if o.LabelPrefix != nil {
			qrLabelPrefix = *o.LabelPrefix
		}
		qLabelPrefix := qrLabelPrefix
		if qLabelPrefix != "" {
			if err := r.SetQueryParam("labelPrefix", qLabelPrefix); err != nil {
				return err
			}
		}

	}

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

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
