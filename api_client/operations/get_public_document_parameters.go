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
)

// NewGETPublicDocumentParams creates a new GETPublicDocumentParams object
// with the default values initialized.
func NewGETPublicDocumentParams() *GETPublicDocumentParams {
	var ()
	return &GETPublicDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETPublicDocumentParamsWithTimeout creates a new GETPublicDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETPublicDocumentParamsWithTimeout(timeout time.Duration) *GETPublicDocumentParams {
	var ()
	return &GETPublicDocumentParams{

		timeout: timeout,
	}
}

// NewGETPublicDocumentParamsWithContext creates a new GETPublicDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETPublicDocumentParamsWithContext(ctx context.Context) *GETPublicDocumentParams {
	var ()
	return &GETPublicDocumentParams{

		Context: ctx,
	}
}

// NewGETPublicDocumentParamsWithHTTPClient creates a new GETPublicDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETPublicDocumentParamsWithHTTPClient(client *http.Client) *GETPublicDocumentParams {
	var ()
	return &GETPublicDocumentParams{
		HTTPClient: client,
	}
}

/*GETPublicDocumentParams contains all the parameters to send to the API endpoint
for the get public document operation typically these are written to a http.Request
*/
type GETPublicDocumentParams struct {

	/*FileName
	  fileName

	*/
	FileName string
	/*ID4N
	  id4n

	*/
	ID4N string
	/*OrganizationID
	  organizationId

	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public document params
func (o *GETPublicDocumentParams) WithTimeout(timeout time.Duration) *GETPublicDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public document params
func (o *GETPublicDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public document params
func (o *GETPublicDocumentParams) WithContext(ctx context.Context) *GETPublicDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public document params
func (o *GETPublicDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public document params
func (o *GETPublicDocumentParams) WithHTTPClient(client *http.Client) *GETPublicDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public document params
func (o *GETPublicDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileName adds the fileName to the get public document params
func (o *GETPublicDocumentParams) WithFileName(fileName string) *GETPublicDocumentParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the get public document params
func (o *GETPublicDocumentParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithID4N adds the id4n to the get public document params
func (o *GETPublicDocumentParams) WithID4N(id4n string) *GETPublicDocumentParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the get public document params
func (o *GETPublicDocumentParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WithOrganizationID adds the organizationID to the get public document params
func (o *GETPublicDocumentParams) WithOrganizationID(organizationID string) *GETPublicDocumentParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get public document params
func (o *GETPublicDocumentParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GETPublicDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fileName
	if err := r.SetPathParam("fileName", o.FileName); err != nil {
		return err
	}

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
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
