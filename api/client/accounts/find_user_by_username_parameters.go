// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewFindUserByUsernameParams creates a new FindUserByUsernameParams object
// with the default values initialized.
func NewFindUserByUsernameParams() *FindUserByUsernameParams {
	var ()
	return &FindUserByUsernameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindUserByUsernameParamsWithTimeout creates a new FindUserByUsernameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindUserByUsernameParamsWithTimeout(timeout time.Duration) *FindUserByUsernameParams {
	var ()
	return &FindUserByUsernameParams{

		timeout: timeout,
	}
}

// NewFindUserByUsernameParamsWithContext creates a new FindUserByUsernameParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindUserByUsernameParamsWithContext(ctx context.Context) *FindUserByUsernameParams {
	var ()
	return &FindUserByUsernameParams{

		Context: ctx,
	}
}

// NewFindUserByUsernameParamsWithHTTPClient creates a new FindUserByUsernameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindUserByUsernameParamsWithHTTPClient(client *http.Client) *FindUserByUsernameParams {
	var ()
	return &FindUserByUsernameParams{
		HTTPClient: client,
	}
}

/*FindUserByUsernameParams contains all the parameters to send to the API endpoint
for the find user by username operation typically these are written to a http.Request
*/
type FindUserByUsernameParams struct {

	/*Username
	  username

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find user by username params
func (o *FindUserByUsernameParams) WithTimeout(timeout time.Duration) *FindUserByUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find user by username params
func (o *FindUserByUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find user by username params
func (o *FindUserByUsernameParams) WithContext(ctx context.Context) *FindUserByUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find user by username params
func (o *FindUserByUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find user by username params
func (o *FindUserByUsernameParams) WithHTTPClient(client *http.Client) *FindUserByUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find user by username params
func (o *FindUserByUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the find user by username params
func (o *FindUserByUsernameParams) WithUsername(username string) *FindUserByUsernameParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the find user by username params
func (o *FindUserByUsernameParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *FindUserByUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
