// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// ListAllRolesReader is a Reader for the ListAllRoles structure.
type ListAllRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAllRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewListAllRolesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListAllRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListAllRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListAllRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListAllRolesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewListAllRolesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewListAllRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAllRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAllRolesOK creates a ListAllRolesOK with default headers values
func NewListAllRolesOK() *ListAllRolesOK {
	return &ListAllRolesOK{}
}

/*ListAllRolesOK handles this case with default header values.

OK
*/
type ListAllRolesOK struct {
	Payload *api_models.RoleResponse
}

func (o *ListAllRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesOK  %+v", 200, o.Payload)
}

func (o *ListAllRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.RoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllRolesAccepted creates a ListAllRolesAccepted with default headers values
func NewListAllRolesAccepted() *ListAllRolesAccepted {
	return &ListAllRolesAccepted{}
}

/*ListAllRolesAccepted handles this case with default header values.

Accepted
*/
type ListAllRolesAccepted struct {
}

func (o *ListAllRolesAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesAccepted ", 202)
}

func (o *ListAllRolesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllRolesUnauthorized creates a ListAllRolesUnauthorized with default headers values
func NewListAllRolesUnauthorized() *ListAllRolesUnauthorized {
	return &ListAllRolesUnauthorized{}
}

/*ListAllRolesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAllRolesUnauthorized struct {
	Payload *api_models.APIError
}

func (o *ListAllRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAllRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllRolesForbidden creates a ListAllRolesForbidden with default headers values
func NewListAllRolesForbidden() *ListAllRolesForbidden {
	return &ListAllRolesForbidden{}
}

/*ListAllRolesForbidden handles this case with default header values.

Forbidden
*/
type ListAllRolesForbidden struct {
	Payload *api_models.APIError
}

func (o *ListAllRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesForbidden  %+v", 403, o.Payload)
}

func (o *ListAllRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllRolesNotFound creates a ListAllRolesNotFound with default headers values
func NewListAllRolesNotFound() *ListAllRolesNotFound {
	return &ListAllRolesNotFound{}
}

/*ListAllRolesNotFound handles this case with default header values.

Not Found
*/
type ListAllRolesNotFound struct {
	Payload *api_models.APIError
}

func (o *ListAllRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesNotFound  %+v", 404, o.Payload)
}

func (o *ListAllRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllRolesMethodNotAllowed creates a ListAllRolesMethodNotAllowed with default headers values
func NewListAllRolesMethodNotAllowed() *ListAllRolesMethodNotAllowed {
	return &ListAllRolesMethodNotAllowed{}
}

/*ListAllRolesMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type ListAllRolesMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *ListAllRolesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListAllRolesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllRolesNotAcceptable creates a ListAllRolesNotAcceptable with default headers values
func NewListAllRolesNotAcceptable() *ListAllRolesNotAcceptable {
	return &ListAllRolesNotAcceptable{}
}

/*ListAllRolesNotAcceptable handles this case with default header values.

Not Acceptable
*/
type ListAllRolesNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *ListAllRolesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesNotAcceptable  %+v", 406, o.Payload)
}

func (o *ListAllRolesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllRolesUnsupportedMediaType creates a ListAllRolesUnsupportedMediaType with default headers values
func NewListAllRolesUnsupportedMediaType() *ListAllRolesUnsupportedMediaType {
	return &ListAllRolesUnsupportedMediaType{}
}

/*ListAllRolesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type ListAllRolesUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *ListAllRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListAllRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllRolesInternalServerError creates a ListAllRolesInternalServerError with default headers values
func NewListAllRolesInternalServerError() *ListAllRolesInternalServerError {
	return &ListAllRolesInternalServerError{}
}

/*ListAllRolesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAllRolesInternalServerError struct {
	Payload *api_models.APIError
}

func (o *ListAllRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] listAllRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}