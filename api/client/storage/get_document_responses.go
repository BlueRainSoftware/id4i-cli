// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// GETDocumentReader is a Reader for the GETDocument structure.
type GETDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETDocumentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETDocumentMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETDocumentNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETDocumentOK creates a GETDocumentOK with default headers values
func NewGETDocumentOK() *GETDocumentOK {
	return &GETDocumentOK{}
}

/*GETDocumentOK handles this case with default header values.

OK
*/
type GETDocumentOK struct {
	Payload *api_models.Document
}

func (o *GETDocumentOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentOK  %+v", 200, o.Payload)
}

func (o *GETDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.Document)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDocumentAccepted creates a GETDocumentAccepted with default headers values
func NewGETDocumentAccepted() *GETDocumentAccepted {
	return &GETDocumentAccepted{}
}

/*GETDocumentAccepted handles this case with default header values.

Accepted
*/
type GETDocumentAccepted struct {
}

func (o *GETDocumentAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentAccepted ", 202)
}

func (o *GETDocumentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETDocumentUnauthorized creates a GETDocumentUnauthorized with default headers values
func NewGETDocumentUnauthorized() *GETDocumentUnauthorized {
	return &GETDocumentUnauthorized{}
}

/*GETDocumentUnauthorized handles this case with default header values.

Unauthorized
*/
type GETDocumentUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *GETDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDocumentForbidden creates a GETDocumentForbidden with default headers values
func NewGETDocumentForbidden() *GETDocumentForbidden {
	return &GETDocumentForbidden{}
}

/*GETDocumentForbidden handles this case with default header values.

Forbidden
*/
type GETDocumentForbidden struct {
	Payload *api_models.APIError
}

func (o *GETDocumentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GETDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDocumentNotFound creates a GETDocumentNotFound with default headers values
func NewGETDocumentNotFound() *GETDocumentNotFound {
	return &GETDocumentNotFound{}
}

/*GETDocumentNotFound handles this case with default header values.

Not Found
*/
type GETDocumentNotFound struct {
	Payload *api_models.APIError
}

func (o *GETDocumentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GETDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDocumentMethodNotAllowed creates a GETDocumentMethodNotAllowed with default headers values
func NewGETDocumentMethodNotAllowed() *GETDocumentMethodNotAllowed {
	return &GETDocumentMethodNotAllowed{}
}

/*GETDocumentMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETDocumentMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETDocumentMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETDocumentMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDocumentNotAcceptable creates a GETDocumentNotAcceptable with default headers values
func NewGETDocumentNotAcceptable() *GETDocumentNotAcceptable {
	return &GETDocumentNotAcceptable{}
}

/*GETDocumentNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETDocumentNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETDocumentNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETDocumentNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDocumentUnsupportedMediaType creates a GETDocumentUnsupportedMediaType with default headers values
func NewGETDocumentUnsupportedMediaType() *GETDocumentUnsupportedMediaType {
	return &GETDocumentUnsupportedMediaType{}
}

/*GETDocumentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETDocumentUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDocumentInternalServerError creates a GETDocumentInternalServerError with default headers values
func NewGETDocumentInternalServerError() *GETDocumentInternalServerError {
	return &GETDocumentInternalServerError{}
}

/*GETDocumentInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETDocumentInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}/{organizationId}/{fileName}/metadata][%d] getDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
