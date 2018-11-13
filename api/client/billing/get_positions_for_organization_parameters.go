// Code generated by go-swagger; DO NOT EDIT.

package billing

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

// NewGETPositionsForOrganizationParams creates a new GETPositionsForOrganizationParams object
// with the default values initialized.
func NewGETPositionsForOrganizationParams() *GETPositionsForOrganizationParams {
	var ()
	return &GETPositionsForOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETPositionsForOrganizationParamsWithTimeout creates a new GETPositionsForOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETPositionsForOrganizationParamsWithTimeout(timeout time.Duration) *GETPositionsForOrganizationParams {
	var ()
	return &GETPositionsForOrganizationParams{

		timeout: timeout,
	}
}

// NewGETPositionsForOrganizationParamsWithContext creates a new GETPositionsForOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETPositionsForOrganizationParamsWithContext(ctx context.Context) *GETPositionsForOrganizationParams {
	var ()
	return &GETPositionsForOrganizationParams{

		Context: ctx,
	}
}

// NewGETPositionsForOrganizationParamsWithHTTPClient creates a new GETPositionsForOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETPositionsForOrganizationParamsWithHTTPClient(client *http.Client) *GETPositionsForOrganizationParams {
	var ()
	return &GETPositionsForOrganizationParams{
		HTTPClient: client,
	}
}

/*GETPositionsForOrganizationParams contains all the parameters to send to the API endpoint
for the get positions for organization operation typically these are written to a http.Request
*/
type GETPositionsForOrganizationParams struct {

	/*FromDate
	  Billing start date

	*/
	FromDate *strfmt.DateTime
	/*OrganizationID
	  The organization to compute the billing information for

	*/
	OrganizationID string
	/*ToDate
	  Billing end date

	*/
	ToDate *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get positions for organization params
func (o *GETPositionsForOrganizationParams) WithTimeout(timeout time.Duration) *GETPositionsForOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get positions for organization params
func (o *GETPositionsForOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get positions for organization params
func (o *GETPositionsForOrganizationParams) WithContext(ctx context.Context) *GETPositionsForOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get positions for organization params
func (o *GETPositionsForOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get positions for organization params
func (o *GETPositionsForOrganizationParams) WithHTTPClient(client *http.Client) *GETPositionsForOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get positions for organization params
func (o *GETPositionsForOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFromDate adds the fromDate to the get positions for organization params
func (o *GETPositionsForOrganizationParams) WithFromDate(fromDate *strfmt.DateTime) *GETPositionsForOrganizationParams {
	o.SetFromDate(fromDate)
	return o
}

// SetFromDate adds the fromDate to the get positions for organization params
func (o *GETPositionsForOrganizationParams) SetFromDate(fromDate *strfmt.DateTime) {
	o.FromDate = fromDate
}

// WithOrganizationID adds the organizationID to the get positions for organization params
func (o *GETPositionsForOrganizationParams) WithOrganizationID(organizationID string) *GETPositionsForOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get positions for organization params
func (o *GETPositionsForOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithToDate adds the toDate to the get positions for organization params
func (o *GETPositionsForOrganizationParams) WithToDate(toDate *strfmt.DateTime) *GETPositionsForOrganizationParams {
	o.SetToDate(toDate)
	return o
}

// SetToDate adds the toDate to the get positions for organization params
func (o *GETPositionsForOrganizationParams) SetToDate(toDate *strfmt.DateTime) {
	o.ToDate = toDate
}

// WriteToRequest writes these params to a swagger request
func (o *GETPositionsForOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FromDate != nil {

		// query param fromDate
		var qrFromDate strfmt.DateTime
		if o.FromDate != nil {
			qrFromDate = *o.FromDate
		}
		qFromDate := qrFromDate.String()
		if qFromDate != "" {
			if err := r.SetQueryParam("fromDate", qFromDate); err != nil {
				return err
			}
		}

	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.ToDate != nil {

		// query param toDate
		var qrToDate strfmt.DateTime
		if o.ToDate != nil {
			qrToDate = *o.ToDate
		}
		qToDate := qrToDate.String()
		if qToDate != "" {
			if err := r.SetQueryParam("toDate", qToDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
