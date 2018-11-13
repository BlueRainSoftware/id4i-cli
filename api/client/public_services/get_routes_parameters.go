// Code generated by go-swagger; DO NOT EDIT.

package public_services

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

// NewGETRoutesParams creates a new GETRoutesParams object
// with the default values initialized.
func NewGETRoutesParams() *GETRoutesParams {
	var (
		interpolateDefault = bool(true)
		typeVarDefault     = string("web")
	)
	return &GETRoutesParams{
		Interpolate: &interpolateDefault,
		Type:        typeVarDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGETRoutesParamsWithTimeout creates a new GETRoutesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETRoutesParamsWithTimeout(timeout time.Duration) *GETRoutesParams {
	var (
		interpolateDefault = bool(true)
		typeVarDefault     = string("web")
	)
	return &GETRoutesParams{
		Interpolate: &interpolateDefault,
		Type:        typeVarDefault,

		timeout: timeout,
	}
}

// NewGETRoutesParamsWithContext creates a new GETRoutesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETRoutesParamsWithContext(ctx context.Context) *GETRoutesParams {
	var (
		interpolateDefault = bool(true)
		typeDefault        = string("web")
	)
	return &GETRoutesParams{
		Interpolate: &interpolateDefault,
		Type:        typeDefault,

		Context: ctx,
	}
}

// NewGETRoutesParamsWithHTTPClient creates a new GETRoutesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETRoutesParamsWithHTTPClient(client *http.Client) *GETRoutesParams {
	var (
		interpolateDefault = bool(true)
		typeDefault        = string("web")
	)
	return &GETRoutesParams{
		Interpolate: &interpolateDefault,
		Type:        typeDefault,
		HTTPClient:  client,
	}
}

/*GETRoutesParams contains all the parameters to send to the API endpoint
for the get routes operation typically these are written to a http.Request
*/
type GETRoutesParams struct {

	/*ID4N
	  id4n

	*/
	ID4N string
	/*Interpolate
	  interpolate

	*/
	Interpolate *bool
	/*Type
	  type

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routes params
func (o *GETRoutesParams) WithTimeout(timeout time.Duration) *GETRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routes params
func (o *GETRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routes params
func (o *GETRoutesParams) WithContext(ctx context.Context) *GETRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routes params
func (o *GETRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routes params
func (o *GETRoutesParams) WithHTTPClient(client *http.Client) *GETRoutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routes params
func (o *GETRoutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the get routes params
func (o *GETRoutesParams) WithID4N(id4n string) *GETRoutesParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the get routes params
func (o *GETRoutesParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WithInterpolate adds the interpolate to the get routes params
func (o *GETRoutesParams) WithInterpolate(interpolate *bool) *GETRoutesParams {
	o.SetInterpolate(interpolate)
	return o
}

// SetInterpolate adds the interpolate to the get routes params
func (o *GETRoutesParams) SetInterpolate(interpolate *bool) {
	o.Interpolate = interpolate
}

// WithType adds the typeVar to the get routes params
func (o *GETRoutesParams) WithType(typeVar string) *GETRoutesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get routes params
func (o *GETRoutesParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GETRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	if o.Interpolate != nil {

		// query param interpolate
		var qrInterpolate bool
		if o.Interpolate != nil {
			qrInterpolate = *o.Interpolate
		}
		qInterpolate := swag.FormatBool(qrInterpolate)
		if qInterpolate != "" {
			if err := r.SetQueryParam("interpolate", qInterpolate); err != nil {
				return err
			}
		}

	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
