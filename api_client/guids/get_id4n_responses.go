// Code generated by go-swagger; DO NOT EDIT.

package guids

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// GETID4NReader is a Reader for the GETID4N structure.
type GETID4NReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETID4NReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETID4NOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETID4NAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETID4NUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETID4NForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETID4NNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETID4NMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETID4NNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETID4NUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETID4NInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETID4NOK creates a GETID4NOK with default headers values
func NewGETID4NOK() *GETID4NOK {
	return &GETID4NOK{}
}

/*GETID4NOK handles this case with default header values.

OK
*/
type GETID4NOK struct {
	Payload *api_models.ID4NPresentation
}

func (o *GETID4NOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nOK  %+v", 200, o.Payload)
}

func (o *GETID4NOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.ID4NPresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETID4NAccepted creates a GETID4NAccepted with default headers values
func NewGETID4NAccepted() *GETID4NAccepted {
	return &GETID4NAccepted{}
}

/*GETID4NAccepted handles this case with default header values.

Accepted
*/
type GETID4NAccepted struct {
}

func (o *GETID4NAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nAccepted ", 202)
}

func (o *GETID4NAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETID4NUnauthorized creates a GETID4NUnauthorized with default headers values
func NewGETID4NUnauthorized() *GETID4NUnauthorized {
	return &GETID4NUnauthorized{}
}

/*GETID4NUnauthorized handles this case with default header values.

Unauthorized
*/
type GETID4NUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETID4NUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nUnauthorized  %+v", 401, o.Payload)
}

func (o *GETID4NUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETID4NForbidden creates a GETID4NForbidden with default headers values
func NewGETID4NForbidden() *GETID4NForbidden {
	return &GETID4NForbidden{}
}

/*GETID4NForbidden handles this case with default header values.

Forbidden
*/
type GETID4NForbidden struct {
	Payload *api_models.APIError
}

func (o *GETID4NForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nForbidden  %+v", 403, o.Payload)
}

func (o *GETID4NForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETID4NNotFound creates a GETID4NNotFound with default headers values
func NewGETID4NNotFound() *GETID4NNotFound {
	return &GETID4NNotFound{}
}

/*GETID4NNotFound handles this case with default header values.

Not Found
*/
type GETID4NNotFound struct {
	Payload *api_models.APIError
}

func (o *GETID4NNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nNotFound  %+v", 404, o.Payload)
}

func (o *GETID4NNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETID4NMethodNotAllowed creates a GETID4NMethodNotAllowed with default headers values
func NewGETID4NMethodNotAllowed() *GETID4NMethodNotAllowed {
	return &GETID4NMethodNotAllowed{}
}

/*GETID4NMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETID4NMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETID4NMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETID4NMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETID4NNotAcceptable creates a GETID4NNotAcceptable with default headers values
func NewGETID4NNotAcceptable() *GETID4NNotAcceptable {
	return &GETID4NNotAcceptable{}
}

/*GETID4NNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETID4NNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETID4NNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETID4NNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETID4NUnsupportedMediaType creates a GETID4NUnsupportedMediaType with default headers values
func NewGETID4NUnsupportedMediaType() *GETID4NUnsupportedMediaType {
	return &GETID4NUnsupportedMediaType{}
}

/*GETID4NUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETID4NUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETID4NUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETID4NUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETID4NInternalServerError creates a GETID4NInternalServerError with default headers values
func NewGETID4NInternalServerError() *GETID4NInternalServerError {
	return &GETID4NInternalServerError{}
}

/*GETID4NInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETID4NInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETID4NInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}][%d] getId4nInternalServerError  %+v", 500, o.Payload)
}

func (o *GETID4NInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}