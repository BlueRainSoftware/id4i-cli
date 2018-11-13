// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// GETAllCollectionsOfOrganizationReader is a Reader for the GETAllCollectionsOfOrganization structure.
type GETAllCollectionsOfOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETAllCollectionsOfOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETAllCollectionsOfOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETAllCollectionsOfOrganizationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETAllCollectionsOfOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETAllCollectionsOfOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETAllCollectionsOfOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETAllCollectionsOfOrganizationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETAllCollectionsOfOrganizationNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETAllCollectionsOfOrganizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETAllCollectionsOfOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETAllCollectionsOfOrganizationOK creates a GETAllCollectionsOfOrganizationOK with default headers values
func NewGETAllCollectionsOfOrganizationOK() *GETAllCollectionsOfOrganizationOK {
	return &GETAllCollectionsOfOrganizationOK{}
}

/*GETAllCollectionsOfOrganizationOK handles this case with default header values.

OK
*/
type GETAllCollectionsOfOrganizationOK struct {
	Payload *api_models.PaginatedGUIDCollection
}

func (o *GETAllCollectionsOfOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationOK  %+v", 200, o.Payload)
}

func (o *GETAllCollectionsOfOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedGUIDCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllCollectionsOfOrganizationAccepted creates a GETAllCollectionsOfOrganizationAccepted with default headers values
func NewGETAllCollectionsOfOrganizationAccepted() *GETAllCollectionsOfOrganizationAccepted {
	return &GETAllCollectionsOfOrganizationAccepted{}
}

/*GETAllCollectionsOfOrganizationAccepted handles this case with default header values.

Accepted
*/
type GETAllCollectionsOfOrganizationAccepted struct {
}

func (o *GETAllCollectionsOfOrganizationAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationAccepted ", 202)
}

func (o *GETAllCollectionsOfOrganizationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETAllCollectionsOfOrganizationUnauthorized creates a GETAllCollectionsOfOrganizationUnauthorized with default headers values
func NewGETAllCollectionsOfOrganizationUnauthorized() *GETAllCollectionsOfOrganizationUnauthorized {
	return &GETAllCollectionsOfOrganizationUnauthorized{}
}

/*GETAllCollectionsOfOrganizationUnauthorized handles this case with default header values.

Unauthorized
*/
type GETAllCollectionsOfOrganizationUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETAllCollectionsOfOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GETAllCollectionsOfOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllCollectionsOfOrganizationForbidden creates a GETAllCollectionsOfOrganizationForbidden with default headers values
func NewGETAllCollectionsOfOrganizationForbidden() *GETAllCollectionsOfOrganizationForbidden {
	return &GETAllCollectionsOfOrganizationForbidden{}
}

/*GETAllCollectionsOfOrganizationForbidden handles this case with default header values.

Forbidden
*/
type GETAllCollectionsOfOrganizationForbidden struct {
	Payload *api_models.APIError
}

func (o *GETAllCollectionsOfOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *GETAllCollectionsOfOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllCollectionsOfOrganizationNotFound creates a GETAllCollectionsOfOrganizationNotFound with default headers values
func NewGETAllCollectionsOfOrganizationNotFound() *GETAllCollectionsOfOrganizationNotFound {
	return &GETAllCollectionsOfOrganizationNotFound{}
}

/*GETAllCollectionsOfOrganizationNotFound handles this case with default header values.

Not Found
*/
type GETAllCollectionsOfOrganizationNotFound struct {
	Payload *api_models.APIError
}

func (o *GETAllCollectionsOfOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *GETAllCollectionsOfOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllCollectionsOfOrganizationMethodNotAllowed creates a GETAllCollectionsOfOrganizationMethodNotAllowed with default headers values
func NewGETAllCollectionsOfOrganizationMethodNotAllowed() *GETAllCollectionsOfOrganizationMethodNotAllowed {
	return &GETAllCollectionsOfOrganizationMethodNotAllowed{}
}

/*GETAllCollectionsOfOrganizationMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETAllCollectionsOfOrganizationMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETAllCollectionsOfOrganizationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETAllCollectionsOfOrganizationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllCollectionsOfOrganizationNotAcceptable creates a GETAllCollectionsOfOrganizationNotAcceptable with default headers values
func NewGETAllCollectionsOfOrganizationNotAcceptable() *GETAllCollectionsOfOrganizationNotAcceptable {
	return &GETAllCollectionsOfOrganizationNotAcceptable{}
}

/*GETAllCollectionsOfOrganizationNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETAllCollectionsOfOrganizationNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETAllCollectionsOfOrganizationNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETAllCollectionsOfOrganizationNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllCollectionsOfOrganizationUnsupportedMediaType creates a GETAllCollectionsOfOrganizationUnsupportedMediaType with default headers values
func NewGETAllCollectionsOfOrganizationUnsupportedMediaType() *GETAllCollectionsOfOrganizationUnsupportedMediaType {
	return &GETAllCollectionsOfOrganizationUnsupportedMediaType{}
}

/*GETAllCollectionsOfOrganizationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETAllCollectionsOfOrganizationUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETAllCollectionsOfOrganizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETAllCollectionsOfOrganizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllCollectionsOfOrganizationInternalServerError creates a GETAllCollectionsOfOrganizationInternalServerError with default headers values
func NewGETAllCollectionsOfOrganizationInternalServerError() *GETAllCollectionsOfOrganizationInternalServerError {
	return &GETAllCollectionsOfOrganizationInternalServerError{}
}

/*GETAllCollectionsOfOrganizationInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETAllCollectionsOfOrganizationInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETAllCollectionsOfOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/collections][%d] getAllCollectionsOfOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GETAllCollectionsOfOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}