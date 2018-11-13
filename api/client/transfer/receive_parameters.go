// Code generated by go-swagger; DO NOT EDIT.

package transfer

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

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// NewReceiveParams creates a new ReceiveParams object
// with the default values initialized.
func NewReceiveParams() *ReceiveParams {
	var ()
	return &ReceiveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReceiveParamsWithTimeout creates a new ReceiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReceiveParamsWithTimeout(timeout time.Duration) *ReceiveParams {
	var ()
	return &ReceiveParams{

		timeout: timeout,
	}
}

// NewReceiveParamsWithContext creates a new ReceiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewReceiveParamsWithContext(ctx context.Context) *ReceiveParams {
	var ()
	return &ReceiveParams{

		Context: ctx,
	}
}

// NewReceiveParamsWithHTTPClient creates a new ReceiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReceiveParamsWithHTTPClient(client *http.Client) *ReceiveParams {
	var ()
	return &ReceiveParams{
		HTTPClient: client,
	}
}

/*ReceiveParams contains all the parameters to send to the API endpoint
for the receive operation typically these are written to a http.Request
*/
type ReceiveParams struct {

	/*ID4N
	  This ID4N identifies the object to take hold of

	*/
	ID4N string
	/*Request
	  Required information to receive an id4n object

	*/
	Request *api_models.TransferReceiveInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the receive params
func (o *ReceiveParams) WithTimeout(timeout time.Duration) *ReceiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the receive params
func (o *ReceiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the receive params
func (o *ReceiveParams) WithContext(ctx context.Context) *ReceiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the receive params
func (o *ReceiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the receive params
func (o *ReceiveParams) WithHTTPClient(client *http.Client) *ReceiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the receive params
func (o *ReceiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the receive params
func (o *ReceiveParams) WithID4N(id4n string) *ReceiveParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the receive params
func (o *ReceiveParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WithRequest adds the request to the receive params
func (o *ReceiveParams) WithRequest(request *api_models.TransferReceiveInfo) *ReceiveParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the receive params
func (o *ReceiveParams) SetRequest(request *api_models.TransferReceiveInfo) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ReceiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
