// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new collections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for collections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddElementsToCollection adds elements to collection
*/
func (a *Client) AddElementsToCollection(params *AddElementsToCollectionParams, authInfo runtime.ClientAuthInfoWriter) (*AddElementsToCollectionOK, *AddElementsToCollectionCreated, *AddElementsToCollectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddElementsToCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addElementsToCollection",
		Method:             "POST",
		PathPattern:        "/api/v1/collections/{id4n}/elements",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddElementsToCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *AddElementsToCollectionOK:
		return value, nil, nil, nil
	case *AddElementsToCollectionCreated:
		return nil, value, nil, nil
	case *AddElementsToCollectionAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
CreateCollection creates collection
*/
func (a *Client) CreateCollection(params *CreateCollectionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCollectionOK, *CreateCollectionCreated, *CreateCollectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCollection",
		Method:             "POST",
		PathPattern:        "/api/v1/collections",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateCollectionOK:
		return value, nil, nil, nil
	case *CreateCollectionCreated:
		return nil, value, nil, nil
	case *CreateCollectionAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
DeleteCollection deletes collection
*/
func (a *Client) DeleteCollection(params *DeleteCollectionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCollection",
		Method:             "DELETE",
		PathPattern:        "/api/v1/collections/{id4n}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCollectionOK), nil

}

/*
FindCollection finds collection
*/
func (a *Client) FindCollection(params *FindCollectionParams, authInfo runtime.ClientAuthInfoWriter) (*FindCollectionOK, *FindCollectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findCollection",
		Method:             "GET",
		PathPattern:        "/api/v1/collections/{id4n}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *FindCollectionOK:
		return value, nil, nil
	case *FindCollectionAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ListElementsOfCollection lists contents of the collection
*/
func (a *Client) ListElementsOfCollection(params *ListElementsOfCollectionParams, authInfo runtime.ClientAuthInfoWriter) (*ListElementsOfCollectionOK, *ListElementsOfCollectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListElementsOfCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listElementsOfCollection",
		Method:             "GET",
		PathPattern:        "/api/v1/collections/{id4n}/elements",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListElementsOfCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListElementsOfCollectionOK:
		return value, nil, nil
	case *ListElementsOfCollectionAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
RemoveElementsFromCollection removes elements from collection
*/
func (a *Client) RemoveElementsFromCollection(params *RemoveElementsFromCollectionParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveElementsFromCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveElementsFromCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeElementsFromCollection",
		Method:             "DELETE",
		PathPattern:        "/api/v1/collections/{id4n}/elements",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveElementsFromCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveElementsFromCollectionOK), nil

}

/*
UpdateCollection updates collection

Update collection changing only the given values
*/
func (a *Client) UpdateCollection(params *UpdateCollectionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCollectionOK, *UpdateCollectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCollection",
		Method:             "PATCH",
		PathPattern:        "/api/v1/collections/{id4n}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateCollectionOK:
		return value, nil, nil
	case *UpdateCollectionNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}