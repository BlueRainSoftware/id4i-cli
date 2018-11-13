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

// NewRemoveGUIDAliasParams creates a new RemoveGUIDAliasParams object
// with the default values initialized.
func NewRemoveGUIDAliasParams() *RemoveGUIDAliasParams {
	var ()
	return &RemoveGUIDAliasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveGUIDAliasParamsWithTimeout creates a new RemoveGUIDAliasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveGUIDAliasParamsWithTimeout(timeout time.Duration) *RemoveGUIDAliasParams {
	var ()
	return &RemoveGUIDAliasParams{

		timeout: timeout,
	}
}

// NewRemoveGUIDAliasParamsWithContext creates a new RemoveGUIDAliasParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveGUIDAliasParamsWithContext(ctx context.Context) *RemoveGUIDAliasParams {
	var ()
	return &RemoveGUIDAliasParams{

		Context: ctx,
	}
}

// NewRemoveGUIDAliasParamsWithHTTPClient creates a new RemoveGUIDAliasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveGUIDAliasParamsWithHTTPClient(client *http.Client) *RemoveGUIDAliasParams {
	var ()
	return &RemoveGUIDAliasParams{
		HTTPClient: client,
	}
}

/*RemoveGUIDAliasParams contains all the parameters to send to the API endpoint
for the remove Guid alias operation typically these are written to a http.Request
*/
type RemoveGUIDAliasParams struct {

	/*AliasType
	  Alias type, see the corresponding API model

	*/
	AliasType string
	/*ID4N
	  The GUID or Collection to operate on

	*/
	ID4N string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove Guid alias params
func (o *RemoveGUIDAliasParams) WithTimeout(timeout time.Duration) *RemoveGUIDAliasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove Guid alias params
func (o *RemoveGUIDAliasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove Guid alias params
func (o *RemoveGUIDAliasParams) WithContext(ctx context.Context) *RemoveGUIDAliasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove Guid alias params
func (o *RemoveGUIDAliasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove Guid alias params
func (o *RemoveGUIDAliasParams) WithHTTPClient(client *http.Client) *RemoveGUIDAliasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove Guid alias params
func (o *RemoveGUIDAliasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAliasType adds the aliasType to the remove Guid alias params
func (o *RemoveGUIDAliasParams) WithAliasType(aliasType string) *RemoveGUIDAliasParams {
	o.SetAliasType(aliasType)
	return o
}

// SetAliasType adds the aliasType to the remove Guid alias params
func (o *RemoveGUIDAliasParams) SetAliasType(aliasType string) {
	o.AliasType = aliasType
}

// WithID4N adds the id4n to the remove Guid alias params
func (o *RemoveGUIDAliasParams) WithID4N(id4n string) *RemoveGUIDAliasParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the remove Guid alias params
func (o *RemoveGUIDAliasParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveGUIDAliasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aliasType
	if err := r.SetPathParam("aliasType", o.AliasType); err != nil {
		return err
	}

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}