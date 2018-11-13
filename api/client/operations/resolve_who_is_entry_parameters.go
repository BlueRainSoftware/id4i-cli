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

// NewResolveWhoIsEntryParams creates a new ResolveWhoIsEntryParams object
// with the default values initialized.
func NewResolveWhoIsEntryParams() *ResolveWhoIsEntryParams {
	var ()
	return &ResolveWhoIsEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResolveWhoIsEntryParamsWithTimeout creates a new ResolveWhoIsEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResolveWhoIsEntryParamsWithTimeout(timeout time.Duration) *ResolveWhoIsEntryParams {
	var ()
	return &ResolveWhoIsEntryParams{

		timeout: timeout,
	}
}

// NewResolveWhoIsEntryParamsWithContext creates a new ResolveWhoIsEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewResolveWhoIsEntryParamsWithContext(ctx context.Context) *ResolveWhoIsEntryParams {
	var ()
	return &ResolveWhoIsEntryParams{

		Context: ctx,
	}
}

// NewResolveWhoIsEntryParamsWithHTTPClient creates a new ResolveWhoIsEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResolveWhoIsEntryParamsWithHTTPClient(client *http.Client) *ResolveWhoIsEntryParams {
	var ()
	return &ResolveWhoIsEntryParams{
		HTTPClient: client,
	}
}

/*ResolveWhoIsEntryParams contains all the parameters to send to the API endpoint
for the resolve who is entry operation typically these are written to a http.Request
*/
type ResolveWhoIsEntryParams struct {

	/*ID4N
	  id4n

	*/
	ID4N string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resolve who is entry params
func (o *ResolveWhoIsEntryParams) WithTimeout(timeout time.Duration) *ResolveWhoIsEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resolve who is entry params
func (o *ResolveWhoIsEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resolve who is entry params
func (o *ResolveWhoIsEntryParams) WithContext(ctx context.Context) *ResolveWhoIsEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resolve who is entry params
func (o *ResolveWhoIsEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resolve who is entry params
func (o *ResolveWhoIsEntryParams) WithHTTPClient(client *http.Client) *ResolveWhoIsEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resolve who is entry params
func (o *ResolveWhoIsEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the resolve who is entry params
func (o *ResolveWhoIsEntryParams) WithID4N(id4n string) *ResolveWhoIsEntryParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the resolve who is entry params
func (o *ResolveWhoIsEntryParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WriteToRequest writes these params to a swagger request
func (o *ResolveWhoIsEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}