// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// GETAllOrganizationRolesReader is a Reader for the GETAllOrganizationRoles structure.
type GETAllOrganizationRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETAllOrganizationRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETAllOrganizationRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETAllOrganizationRolesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETAllOrganizationRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETAllOrganizationRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETAllOrganizationRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETAllOrganizationRolesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETAllOrganizationRolesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETAllOrganizationRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETAllOrganizationRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETAllOrganizationRolesOK creates a GETAllOrganizationRolesOK with default headers values
func NewGETAllOrganizationRolesOK() *GETAllOrganizationRolesOK {
	return &GETAllOrganizationRolesOK{}
}

/*GETAllOrganizationRolesOK handles this case with default header values.

OK
*/
type GETAllOrganizationRolesOK struct {
	Payload *api_models.PaginatedUserRolesResponse
}

func (o *GETAllOrganizationRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesOK  %+v", 200, o.Payload)
}

func (o *GETAllOrganizationRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedUserRolesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllOrganizationRolesAccepted creates a GETAllOrganizationRolesAccepted with default headers values
func NewGETAllOrganizationRolesAccepted() *GETAllOrganizationRolesAccepted {
	return &GETAllOrganizationRolesAccepted{}
}

/*GETAllOrganizationRolesAccepted handles this case with default header values.

Accepted
*/
type GETAllOrganizationRolesAccepted struct {
}

func (o *GETAllOrganizationRolesAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesAccepted ", 202)
}

func (o *GETAllOrganizationRolesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETAllOrganizationRolesUnauthorized creates a GETAllOrganizationRolesUnauthorized with default headers values
func NewGETAllOrganizationRolesUnauthorized() *GETAllOrganizationRolesUnauthorized {
	return &GETAllOrganizationRolesUnauthorized{}
}

/*GETAllOrganizationRolesUnauthorized handles this case with default header values.

Unauthorized
*/
type GETAllOrganizationRolesUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETAllOrganizationRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GETAllOrganizationRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllOrganizationRolesForbidden creates a GETAllOrganizationRolesForbidden with default headers values
func NewGETAllOrganizationRolesForbidden() *GETAllOrganizationRolesForbidden {
	return &GETAllOrganizationRolesForbidden{}
}

/*GETAllOrganizationRolesForbidden handles this case with default header values.

Forbidden
*/
type GETAllOrganizationRolesForbidden struct {
	Payload *api_models.APIError
}

func (o *GETAllOrganizationRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesForbidden  %+v", 403, o.Payload)
}

func (o *GETAllOrganizationRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllOrganizationRolesNotFound creates a GETAllOrganizationRolesNotFound with default headers values
func NewGETAllOrganizationRolesNotFound() *GETAllOrganizationRolesNotFound {
	return &GETAllOrganizationRolesNotFound{}
}

/*GETAllOrganizationRolesNotFound handles this case with default header values.

Not Found
*/
type GETAllOrganizationRolesNotFound struct {
	Payload *api_models.APIError
}

func (o *GETAllOrganizationRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesNotFound  %+v", 404, o.Payload)
}

func (o *GETAllOrganizationRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllOrganizationRolesMethodNotAllowed creates a GETAllOrganizationRolesMethodNotAllowed with default headers values
func NewGETAllOrganizationRolesMethodNotAllowed() *GETAllOrganizationRolesMethodNotAllowed {
	return &GETAllOrganizationRolesMethodNotAllowed{}
}

/*GETAllOrganizationRolesMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETAllOrganizationRolesMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETAllOrganizationRolesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETAllOrganizationRolesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllOrganizationRolesNotAcceptable creates a GETAllOrganizationRolesNotAcceptable with default headers values
func NewGETAllOrganizationRolesNotAcceptable() *GETAllOrganizationRolesNotAcceptable {
	return &GETAllOrganizationRolesNotAcceptable{}
}

/*GETAllOrganizationRolesNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETAllOrganizationRolesNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETAllOrganizationRolesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETAllOrganizationRolesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllOrganizationRolesUnsupportedMediaType creates a GETAllOrganizationRolesUnsupportedMediaType with default headers values
func NewGETAllOrganizationRolesUnsupportedMediaType() *GETAllOrganizationRolesUnsupportedMediaType {
	return &GETAllOrganizationRolesUnsupportedMediaType{}
}

/*GETAllOrganizationRolesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETAllOrganizationRolesUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETAllOrganizationRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETAllOrganizationRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllOrganizationRolesInternalServerError creates a GETAllOrganizationRolesInternalServerError with default headers values
func NewGETAllOrganizationRolesInternalServerError() *GETAllOrganizationRolesInternalServerError {
	return &GETAllOrganizationRolesInternalServerError{}
}

/*GETAllOrganizationRolesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETAllOrganizationRolesInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETAllOrganizationRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/roles][%d] getAllOrganizationRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETAllOrganizationRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}