// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// FindUserByUsernameReader is a Reader for the FindUserByUsername structure.
type FindUserByUsernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindUserByUsernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindUserByUsernameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewFindUserByUsernameAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewFindUserByUsernameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewFindUserByUsernameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewFindUserByUsernameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewFindUserByUsernameMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewFindUserByUsernameNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewFindUserByUsernameUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewFindUserByUsernameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFindUserByUsernameOK creates a FindUserByUsernameOK with default headers values
func NewFindUserByUsernameOK() *FindUserByUsernameOK {
	return &FindUserByUsernameOK{}
}

/*FindUserByUsernameOK handles this case with default header values.

OK
*/
type FindUserByUsernameOK struct {
	Payload *api_models.UserPresentation
}

func (o *FindUserByUsernameOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameOK  %+v", 200, o.Payload)
}

func (o *FindUserByUsernameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.UserPresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByUsernameAccepted creates a FindUserByUsernameAccepted with default headers values
func NewFindUserByUsernameAccepted() *FindUserByUsernameAccepted {
	return &FindUserByUsernameAccepted{}
}

/*FindUserByUsernameAccepted handles this case with default header values.

Accepted
*/
type FindUserByUsernameAccepted struct {
}

func (o *FindUserByUsernameAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameAccepted ", 202)
}

func (o *FindUserByUsernameAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindUserByUsernameUnauthorized creates a FindUserByUsernameUnauthorized with default headers values
func NewFindUserByUsernameUnauthorized() *FindUserByUsernameUnauthorized {
	return &FindUserByUsernameUnauthorized{}
}

/*FindUserByUsernameUnauthorized handles this case with default header values.

Unauthorized
*/
type FindUserByUsernameUnauthorized struct {
	Payload *api_models.APIError
}

func (o *FindUserByUsernameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameUnauthorized  %+v", 401, o.Payload)
}

func (o *FindUserByUsernameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByUsernameForbidden creates a FindUserByUsernameForbidden with default headers values
func NewFindUserByUsernameForbidden() *FindUserByUsernameForbidden {
	return &FindUserByUsernameForbidden{}
}

/*FindUserByUsernameForbidden handles this case with default header values.

Forbidden
*/
type FindUserByUsernameForbidden struct {
	Payload *api_models.APIError
}

func (o *FindUserByUsernameForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameForbidden  %+v", 403, o.Payload)
}

func (o *FindUserByUsernameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByUsernameNotFound creates a FindUserByUsernameNotFound with default headers values
func NewFindUserByUsernameNotFound() *FindUserByUsernameNotFound {
	return &FindUserByUsernameNotFound{}
}

/*FindUserByUsernameNotFound handles this case with default header values.

Not Found
*/
type FindUserByUsernameNotFound struct {
	Payload *api_models.APIError
}

func (o *FindUserByUsernameNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameNotFound  %+v", 404, o.Payload)
}

func (o *FindUserByUsernameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByUsernameMethodNotAllowed creates a FindUserByUsernameMethodNotAllowed with default headers values
func NewFindUserByUsernameMethodNotAllowed() *FindUserByUsernameMethodNotAllowed {
	return &FindUserByUsernameMethodNotAllowed{}
}

/*FindUserByUsernameMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type FindUserByUsernameMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *FindUserByUsernameMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *FindUserByUsernameMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByUsernameNotAcceptable creates a FindUserByUsernameNotAcceptable with default headers values
func NewFindUserByUsernameNotAcceptable() *FindUserByUsernameNotAcceptable {
	return &FindUserByUsernameNotAcceptable{}
}

/*FindUserByUsernameNotAcceptable handles this case with default header values.

Not Acceptable
*/
type FindUserByUsernameNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *FindUserByUsernameNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameNotAcceptable  %+v", 406, o.Payload)
}

func (o *FindUserByUsernameNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByUsernameUnsupportedMediaType creates a FindUserByUsernameUnsupportedMediaType with default headers values
func NewFindUserByUsernameUnsupportedMediaType() *FindUserByUsernameUnsupportedMediaType {
	return &FindUserByUsernameUnsupportedMediaType{}
}

/*FindUserByUsernameUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type FindUserByUsernameUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *FindUserByUsernameUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *FindUserByUsernameUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByUsernameInternalServerError creates a FindUserByUsernameInternalServerError with default headers values
func NewFindUserByUsernameInternalServerError() *FindUserByUsernameInternalServerError {
	return &FindUserByUsernameInternalServerError{}
}

/*FindUserByUsernameInternalServerError handles this case with default header values.

Internal Server Error
*/
type FindUserByUsernameInternalServerError struct {
	Payload *api_models.APIError
}

func (o *FindUserByUsernameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/{username}][%d] findUserByUsernameInternalServerError  %+v", 500, o.Payload)
}

func (o *FindUserByUsernameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
