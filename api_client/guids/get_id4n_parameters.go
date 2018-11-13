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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETID4NParams creates a new GETID4NParams object
// with the default values initialized.
func NewGETID4NParams() *GETID4NParams {
	var ()
	return &GETID4NParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETID4NParamsWithTimeout creates a new GETID4NParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETID4NParamsWithTimeout(timeout time.Duration) *GETID4NParams {
	var ()
	return &GETID4NParams{

		timeout: timeout,
	}
}

// NewGETID4NParamsWithContext creates a new GETID4NParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETID4NParamsWithContext(ctx context.Context) *GETID4NParams {
	var ()
	return &GETID4NParams{

		Context: ctx,
	}
}

// NewGETID4NParamsWithHTTPClient creates a new GETID4NParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETID4NParamsWithHTTPClient(client *http.Client) *GETID4NParams {
	var ()
	return &GETID4NParams{
		HTTPClient: client,
	}
}

/*GETID4NParams contains all the parameters to send to the API endpoint
for the get Id4n operation typically these are written to a http.Request
*/
type GETID4NParams struct {

	/*ID4N
	  The ID to resolve to

	*/
	ID4N string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Id4n params
func (o *GETID4NParams) WithTimeout(timeout time.Duration) *GETID4NParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Id4n params
func (o *GETID4NParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Id4n params
func (o *GETID4NParams) WithContext(ctx context.Context) *GETID4NParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Id4n params
func (o *GETID4NParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Id4n params
func (o *GETID4NParams) WithHTTPClient(client *http.Client) *GETID4NParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Id4n params
func (o *GETID4NParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the get Id4n params
func (o *GETID4NParams) WithID4N(id4n string) *GETID4NParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the get Id4n params
func (o *GETID4NParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WriteToRequest writes these params to a swagger request
func (o *GETID4NParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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