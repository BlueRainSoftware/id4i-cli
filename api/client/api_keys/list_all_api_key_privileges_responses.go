// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// ListAllAPIKeyPrivilegesReader is a Reader for the ListAllAPIKeyPrivileges structure.
type ListAllAPIKeyPrivilegesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllAPIKeyPrivilegesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAllAPIKeyPrivilegesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewListAllAPIKeyPrivilegesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListAllAPIKeyPrivilegesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListAllAPIKeyPrivilegesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListAllAPIKeyPrivilegesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListAllAPIKeyPrivilegesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewListAllAPIKeyPrivilegesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewListAllAPIKeyPrivilegesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAllAPIKeyPrivilegesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAllAPIKeyPrivilegesOK creates a ListAllAPIKeyPrivilegesOK with default headers values
func NewListAllAPIKeyPrivilegesOK() *ListAllAPIKeyPrivilegesOK {
	return &ListAllAPIKeyPrivilegesOK{}
}

/*ListAllAPIKeyPrivilegesOK handles this case with default header values.

OK
*/
type ListAllAPIKeyPrivilegesOK struct {
	Payload *api_models.APIKeyPrivilegeInfoResponse
}

func (o *ListAllAPIKeyPrivilegesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesOK  %+v", 200, o.Payload)
}

func (o *ListAllAPIKeyPrivilegesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIKeyPrivilegeInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllAPIKeyPrivilegesAccepted creates a ListAllAPIKeyPrivilegesAccepted with default headers values
func NewListAllAPIKeyPrivilegesAccepted() *ListAllAPIKeyPrivilegesAccepted {
	return &ListAllAPIKeyPrivilegesAccepted{}
}

/*ListAllAPIKeyPrivilegesAccepted handles this case with default header values.

Accepted
*/
type ListAllAPIKeyPrivilegesAccepted struct {
}

func (o *ListAllAPIKeyPrivilegesAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesAccepted ", 202)
}

func (o *ListAllAPIKeyPrivilegesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllAPIKeyPrivilegesUnauthorized creates a ListAllAPIKeyPrivilegesUnauthorized with default headers values
func NewListAllAPIKeyPrivilegesUnauthorized() *ListAllAPIKeyPrivilegesUnauthorized {
	return &ListAllAPIKeyPrivilegesUnauthorized{}
}

/*ListAllAPIKeyPrivilegesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAllAPIKeyPrivilegesUnauthorized struct {
	Payload *api_models.APIError
}

func (o *ListAllAPIKeyPrivilegesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAllAPIKeyPrivilegesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllAPIKeyPrivilegesForbidden creates a ListAllAPIKeyPrivilegesForbidden with default headers values
func NewListAllAPIKeyPrivilegesForbidden() *ListAllAPIKeyPrivilegesForbidden {
	return &ListAllAPIKeyPrivilegesForbidden{}
}

/*ListAllAPIKeyPrivilegesForbidden handles this case with default header values.

Forbidden
*/
type ListAllAPIKeyPrivilegesForbidden struct {
	Payload *api_models.APIError
}

func (o *ListAllAPIKeyPrivilegesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesForbidden  %+v", 403, o.Payload)
}

func (o *ListAllAPIKeyPrivilegesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllAPIKeyPrivilegesNotFound creates a ListAllAPIKeyPrivilegesNotFound with default headers values
func NewListAllAPIKeyPrivilegesNotFound() *ListAllAPIKeyPrivilegesNotFound {
	return &ListAllAPIKeyPrivilegesNotFound{}
}

/*ListAllAPIKeyPrivilegesNotFound handles this case with default header values.

Not Found
*/
type ListAllAPIKeyPrivilegesNotFound struct {
	Payload *api_models.APIError
}

func (o *ListAllAPIKeyPrivilegesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesNotFound  %+v", 404, o.Payload)
}

func (o *ListAllAPIKeyPrivilegesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllAPIKeyPrivilegesMethodNotAllowed creates a ListAllAPIKeyPrivilegesMethodNotAllowed with default headers values
func NewListAllAPIKeyPrivilegesMethodNotAllowed() *ListAllAPIKeyPrivilegesMethodNotAllowed {
	return &ListAllAPIKeyPrivilegesMethodNotAllowed{}
}

/*ListAllAPIKeyPrivilegesMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type ListAllAPIKeyPrivilegesMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *ListAllAPIKeyPrivilegesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListAllAPIKeyPrivilegesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllAPIKeyPrivilegesNotAcceptable creates a ListAllAPIKeyPrivilegesNotAcceptable with default headers values
func NewListAllAPIKeyPrivilegesNotAcceptable() *ListAllAPIKeyPrivilegesNotAcceptable {
	return &ListAllAPIKeyPrivilegesNotAcceptable{}
}

/*ListAllAPIKeyPrivilegesNotAcceptable handles this case with default header values.

Not Acceptable
*/
type ListAllAPIKeyPrivilegesNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *ListAllAPIKeyPrivilegesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesNotAcceptable  %+v", 406, o.Payload)
}

func (o *ListAllAPIKeyPrivilegesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllAPIKeyPrivilegesUnsupportedMediaType creates a ListAllAPIKeyPrivilegesUnsupportedMediaType with default headers values
func NewListAllAPIKeyPrivilegesUnsupportedMediaType() *ListAllAPIKeyPrivilegesUnsupportedMediaType {
	return &ListAllAPIKeyPrivilegesUnsupportedMediaType{}
}

/*ListAllAPIKeyPrivilegesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type ListAllAPIKeyPrivilegesUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *ListAllAPIKeyPrivilegesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListAllAPIKeyPrivilegesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllAPIKeyPrivilegesInternalServerError creates a ListAllAPIKeyPrivilegesInternalServerError with default headers values
func NewListAllAPIKeyPrivilegesInternalServerError() *ListAllAPIKeyPrivilegesInternalServerError {
	return &ListAllAPIKeyPrivilegesInternalServerError{}
}

/*ListAllAPIKeyPrivilegesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAllAPIKeyPrivilegesInternalServerError struct {
	Payload *api_models.APIError
}

func (o *ListAllAPIKeyPrivilegesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/privileges][%d] listAllApiKeyPrivilegesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllAPIKeyPrivilegesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
