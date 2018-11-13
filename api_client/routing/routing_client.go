// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new routing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for routing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GETAllRoutes retrieves all routes of a GUID or ID4N
*/
func (a *Client) GETAllRoutes(params *GETAllRoutesParams, authInfo runtime.ClientAuthInfoWriter) (*GETAllRoutesOK, *GETAllRoutesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETAllRoutesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllRoutes",
		Method:             "GET",
		PathPattern:        "/api/v1/routingfiles/{id4n}/routes/{type}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETAllRoutesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETAllRoutesOK:
		return value, nil, nil
	case *GETAllRoutesAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GETRoute retrieves current route of a GUID or ID4N
*/
func (a *Client) GETRoute(params *GETRouteParams, authInfo runtime.ClientAuthInfoWriter) (*GETRouteOK, *GETRouteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoute",
		Method:             "GET",
		PathPattern:        "/api/v1/routingfiles/{id4n}/route/{type}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETRouteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETRouteOK:
		return value, nil, nil
	case *GETRouteAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GETRoutingFile retrieves routing file
*/
func (a *Client) GETRoutingFile(params *GETRoutingFileParams, authInfo runtime.ClientAuthInfoWriter) (*GETRoutingFileOK, *GETRoutingFileAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETRoutingFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoutingFile",
		Method:             "GET",
		PathPattern:        "/api/v1/routingfiles/{id4n}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETRoutingFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETRoutingFileOK:
		return value, nil, nil
	case *GETRoutingFileAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
UpdateRoutingFile stores routing file
*/
func (a *Client) UpdateRoutingFile(params *UpdateRoutingFileParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRoutingFileOK, *UpdateRoutingFileAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoutingFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRoutingFile",
		Method:             "PUT",
		PathPattern:        "/api/v1/routingfiles/{id4n}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRoutingFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateRoutingFileOK:
		return value, nil, nil
	case *UpdateRoutingFileAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}