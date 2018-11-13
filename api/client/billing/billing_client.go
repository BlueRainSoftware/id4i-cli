// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new billing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GETPositionsForOrganization Gets billing positions for a given organization
*/
func (a *Client) GETPositionsForOrganization(params *GETPositionsForOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*GETPositionsForOrganizationOK, *GETPositionsForOrganizationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETPositionsForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPositionsForOrganization",
		Method:             "GET",
		PathPattern:        "/api/v1/billing/{organizationId}/positions",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETPositionsForOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETPositionsForOrganizationOK:
		return value, nil, nil
	case *GETPositionsForOrganizationAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GETSumForOrganization Gets billing amount of services for a given organization
*/
func (a *Client) GETSumForOrganization(params *GETSumForOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*GETSumForOrganizationOK, *GETSumForOrganizationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETSumForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSumForOrganization",
		Method:             "GET",
		PathPattern:        "/api/v1/billing/{organizationId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETSumForOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETSumForOrganizationOK:
		return value, nil, nil
	case *GETSumForOrganizationAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
