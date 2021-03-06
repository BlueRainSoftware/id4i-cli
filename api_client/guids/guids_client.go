// Code generated by go-swagger; DO NOT EDIT.

package guids

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new guids API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for guids API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateGUID creates GUID s

Creating one or more GUIDs with a specified length.
*/
func (a *Client) CreateGUID(params *CreateGUIDParams, authInfo runtime.ClientAuthInfoWriter) (*CreateGUIDOK, *CreateGUIDCreated, *CreateGUIDAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createGuid",
		Method:             "POST",
		PathPattern:        "/api/v1/guids",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateGUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateGUIDOK:
		return value, nil, nil, nil
	case *CreateGUIDCreated:
		return nil, value, nil, nil
	case *CreateGUIDAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
GETCollections retrieves collections of an ID

Retrieving all owned or holding collections the specified id4n is assigned to.
*/
func (a *Client) GETCollections(params *GETCollectionsParams, authInfo runtime.ClientAuthInfoWriter) (*GETCollectionsOK, *GETCollectionsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCollectionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCollections",
		Method:             "GET",
		PathPattern:        "/api/v1/id4ns/{id4n}/collections",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETCollectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETCollectionsOK:
		return value, nil, nil
	case *GETCollectionsAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GETGUID retrieves GUID information
*/
func (a *Client) GETGUID(params *GETGUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GETGUIDOK, *GETGUIDAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGuid",
		Method:             "GET",
		PathPattern:        "/api/v1/guids/{id4n}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETGUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETGUIDOK:
		return value, nil, nil
	case *GETGUIDAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GETGuidsWithoutCollection retrieves guids not in any collection
*/
func (a *Client) GETGuidsWithoutCollection(params *GETGuidsWithoutCollectionParams, authInfo runtime.ClientAuthInfoWriter) (*GETGuidsWithoutCollectionOK, *GETGuidsWithoutCollectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETGuidsWithoutCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGuidsWithoutCollection",
		Method:             "GET",
		PathPattern:        "/api/v1/guids/withoutCollection",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETGuidsWithoutCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETGuidsWithoutCollectionOK:
		return value, nil, nil
	case *GETGuidsWithoutCollectionAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GETID4N retrieves ID4n information

Retrieving basic information about an ID like the type and the creation time.
*/
func (a *Client) GETID4N(params *GETID4NParams, authInfo runtime.ClientAuthInfoWriter) (*GETID4NOK, *GETID4NAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETID4NParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getId4n",
		Method:             "GET",
		PathPattern:        "/api/v1/id4ns/{id4n}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETID4NReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETID4NOK:
		return value, nil, nil
	case *GETID4NAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
UpdateGUID changes GUID information

Allows ownership transfer.
*/
func (a *Client) UpdateGUID(params *UpdateGUIDParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateGUIDOK, *UpdateGUIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateGuid",
		Method:             "PATCH",
		PathPattern:        "/api/v1/guids/{id4n}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateGUIDOK:
		return value, nil, nil
	case *UpdateGUIDNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
