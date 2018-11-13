// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

// NewListId4nsParams creates a new ListId4nsParams object
// with the default values initialized.
func NewListId4nsParams() *ListId4nsParams {
	var ()
	return &ListId4nsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListId4nsParamsWithTimeout creates a new ListId4nsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListId4nsParamsWithTimeout(timeout time.Duration) *ListId4nsParams {
	var ()
	return &ListId4nsParams{

		timeout: timeout,
	}
}

// NewListId4nsParamsWithContext creates a new ListId4nsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListId4nsParamsWithContext(ctx context.Context) *ListId4nsParams {
	var ()
	return &ListId4nsParams{

		Context: ctx,
	}
}

// NewListId4nsParamsWithHTTPClient creates a new ListId4nsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListId4nsParamsWithHTTPClient(client *http.Client) *ListId4nsParams {
	var ()
	return &ListId4nsParams{
		HTTPClient: client,
	}
}

/*ListId4nsParams contains all the parameters to send to the API endpoint
for the list id4ns operation typically these are written to a http.Request
*/
type ListId4nsParams struct {

	/*Key
	  key

	*/
	Key string
	/*Limit
	  The maximum count of returned elements

	*/
	Limit *int32
	/*Offset
	  Start with the n-th element

	*/
	Offset *int32
	/*Privilege
	  privilege

	*/
	Privilege string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list id4ns params
func (o *ListId4nsParams) WithTimeout(timeout time.Duration) *ListId4nsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list id4ns params
func (o *ListId4nsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list id4ns params
func (o *ListId4nsParams) WithContext(ctx context.Context) *ListId4nsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list id4ns params
func (o *ListId4nsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list id4ns params
func (o *ListId4nsParams) WithHTTPClient(client *http.Client) *ListId4nsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list id4ns params
func (o *ListId4nsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the list id4ns params
func (o *ListId4nsParams) WithKey(key string) *ListId4nsParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the list id4ns params
func (o *ListId4nsParams) SetKey(key string) {
	o.Key = key
}

// WithLimit adds the limit to the list id4ns params
func (o *ListId4nsParams) WithLimit(limit *int32) *ListId4nsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list id4ns params
func (o *ListId4nsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list id4ns params
func (o *ListId4nsParams) WithOffset(offset *int32) *ListId4nsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list id4ns params
func (o *ListId4nsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithPrivilege adds the privilege to the list id4ns params
func (o *ListId4nsParams) WithPrivilege(privilege string) *ListId4nsParams {
	o.SetPrivilege(privilege)
	return o
}

// SetPrivilege adds the privilege to the list id4ns params
func (o *ListId4nsParams) SetPrivilege(privilege string) {
	o.Privilege = privilege
}

// WriteToRequest writes these params to a swagger request
func (o *ListId4nsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
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

	// path param privilege
	if err := r.SetPathParam("privilege", o.Privilege); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
