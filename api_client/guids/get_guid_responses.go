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

// GETGUIDReader is a Reader for the GETGUID structure.
type GETGUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETGUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETGUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETGUIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETGUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETGUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETGUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETGUIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETGUIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETGUIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETGUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETGUIDOK creates a GETGUIDOK with default headers values
func NewGETGUIDOK() *GETGUIDOK {
	return &GETGUIDOK{}
}

/*GETGUIDOK handles this case with default header values.

OK
*/
type GETGUIDOK struct {
	Payload *api_models.GUID
}

func (o *GETGUIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidOK  %+v", 200, o.Payload)
}

func (o *GETGUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.GUID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDAccepted creates a GETGUIDAccepted with default headers values
func NewGETGUIDAccepted() *GETGUIDAccepted {
	return &GETGUIDAccepted{}
}

/*GETGUIDAccepted handles this case with default header values.

Accepted
*/
type GETGUIDAccepted struct {
}

func (o *GETGUIDAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidAccepted ", 202)
}

func (o *GETGUIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETGUIDUnauthorized creates a GETGUIDUnauthorized with default headers values
func NewGETGUIDUnauthorized() *GETGUIDUnauthorized {
	return &GETGUIDUnauthorized{}
}

/*GETGUIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GETGUIDUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETGUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidUnauthorized  %+v", 401, o.Payload)
}

func (o *GETGUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDForbidden creates a GETGUIDForbidden with default headers values
func NewGETGUIDForbidden() *GETGUIDForbidden {
	return &GETGUIDForbidden{}
}

/*GETGUIDForbidden handles this case with default header values.

Forbidden
*/
type GETGUIDForbidden struct {
	Payload *api_models.APIError
}

func (o *GETGUIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidForbidden  %+v", 403, o.Payload)
}

func (o *GETGUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDNotFound creates a GETGUIDNotFound with default headers values
func NewGETGUIDNotFound() *GETGUIDNotFound {
	return &GETGUIDNotFound{}
}

/*GETGUIDNotFound handles this case with default header values.

Not Found
*/
type GETGUIDNotFound struct {
	Payload *api_models.APIError
}

func (o *GETGUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidNotFound  %+v", 404, o.Payload)
}

func (o *GETGUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDMethodNotAllowed creates a GETGUIDMethodNotAllowed with default headers values
func NewGETGUIDMethodNotAllowed() *GETGUIDMethodNotAllowed {
	return &GETGUIDMethodNotAllowed{}
}

/*GETGUIDMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETGUIDMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETGUIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETGUIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDNotAcceptable creates a GETGUIDNotAcceptable with default headers values
func NewGETGUIDNotAcceptable() *GETGUIDNotAcceptable {
	return &GETGUIDNotAcceptable{}
}

/*GETGUIDNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETGUIDNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETGUIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETGUIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDUnsupportedMediaType creates a GETGUIDUnsupportedMediaType with default headers values
func NewGETGUIDUnsupportedMediaType() *GETGUIDUnsupportedMediaType {
	return &GETGUIDUnsupportedMediaType{}
}

/*GETGUIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETGUIDUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETGUIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETGUIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDInternalServerError creates a GETGUIDInternalServerError with default headers values
func NewGETGUIDInternalServerError() *GETGUIDInternalServerError {
	return &GETGUIDInternalServerError{}
}

/*GETGUIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETGUIDInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETGUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/{id4n}][%d] getGuidInternalServerError  %+v", 500, o.Payload)
}

func (o *GETGUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
