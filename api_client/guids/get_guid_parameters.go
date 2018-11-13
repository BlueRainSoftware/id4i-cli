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

// NewGETGUIDParams creates a new GETGUIDParams object
// with the default values initialized.
func NewGETGUIDParams() *GETGUIDParams {
	var ()
	return &GETGUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETGUIDParamsWithTimeout creates a new GETGUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETGUIDParamsWithTimeout(timeout time.Duration) *GETGUIDParams {
	var ()
	return &GETGUIDParams{

		timeout: timeout,
	}
}

// NewGETGUIDParamsWithContext creates a new GETGUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETGUIDParamsWithContext(ctx context.Context) *GETGUIDParams {
	var ()
	return &GETGUIDParams{

		Context: ctx,
	}
}

// NewGETGUIDParamsWithHTTPClient creates a new GETGUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETGUIDParamsWithHTTPClient(client *http.Client) *GETGUIDParams {
	var ()
	return &GETGUIDParams{
		HTTPClient: client,
	}
}

/*GETGUIDParams contains all the parameters to send to the API endpoint
for the get Guid operation typically these are written to a http.Request
*/
type GETGUIDParams struct {

	/*ID4N
	  The GUID number

	*/
	ID4N string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Guid params
func (o *GETGUIDParams) WithTimeout(timeout time.Duration) *GETGUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Guid params
func (o *GETGUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Guid params
func (o *GETGUIDParams) WithContext(ctx context.Context) *GETGUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Guid params
func (o *GETGUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Guid params
func (o *GETGUIDParams) WithHTTPClient(client *http.Client) *GETGUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Guid params
func (o *GETGUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the get Guid params
func (o *GETGUIDParams) WithID4N(id4n string) *GETGUIDParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the get Guid params
func (o *GETGUIDParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WriteToRequest writes these params to a swagger request
func (o *GETGUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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