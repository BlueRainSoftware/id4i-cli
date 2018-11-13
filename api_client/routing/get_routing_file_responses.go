// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// GETRoutingFileReader is a Reader for the GETRoutingFile structure.
type GETRoutingFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETRoutingFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETRoutingFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETRoutingFileAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETRoutingFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETRoutingFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETRoutingFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETRoutingFileMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETRoutingFileNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETRoutingFileUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETRoutingFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETRoutingFileOK creates a GETRoutingFileOK with default headers values
func NewGETRoutingFileOK() *GETRoutingFileOK {
	return &GETRoutingFileOK{}
}

/*GETRoutingFileOK handles this case with default header values.

OK
*/
type GETRoutingFileOK struct {
	Payload *api_models.RoutingFile
}

func (o *GETRoutingFileOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileOK  %+v", 200, o.Payload)
}

func (o *GETRoutingFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.RoutingFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRoutingFileAccepted creates a GETRoutingFileAccepted with default headers values
func NewGETRoutingFileAccepted() *GETRoutingFileAccepted {
	return &GETRoutingFileAccepted{}
}

/*GETRoutingFileAccepted handles this case with default header values.

Accepted
*/
type GETRoutingFileAccepted struct {
}

func (o *GETRoutingFileAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileAccepted ", 202)
}

func (o *GETRoutingFileAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETRoutingFileUnauthorized creates a GETRoutingFileUnauthorized with default headers values
func NewGETRoutingFileUnauthorized() *GETRoutingFileUnauthorized {
	return &GETRoutingFileUnauthorized{}
}

/*GETRoutingFileUnauthorized handles this case with default header values.

Unauthorized
*/
type GETRoutingFileUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETRoutingFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileUnauthorized  %+v", 401, o.Payload)
}

func (o *GETRoutingFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRoutingFileForbidden creates a GETRoutingFileForbidden with default headers values
func NewGETRoutingFileForbidden() *GETRoutingFileForbidden {
	return &GETRoutingFileForbidden{}
}

/*GETRoutingFileForbidden handles this case with default header values.

Forbidden
*/
type GETRoutingFileForbidden struct {
	Payload *api_models.APIError
}

func (o *GETRoutingFileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileForbidden  %+v", 403, o.Payload)
}

func (o *GETRoutingFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRoutingFileNotFound creates a GETRoutingFileNotFound with default headers values
func NewGETRoutingFileNotFound() *GETRoutingFileNotFound {
	return &GETRoutingFileNotFound{}
}

/*GETRoutingFileNotFound handles this case with default header values.

Not Found
*/
type GETRoutingFileNotFound struct {
	Payload *api_models.APIError
}

func (o *GETRoutingFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileNotFound  %+v", 404, o.Payload)
}

func (o *GETRoutingFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRoutingFileMethodNotAllowed creates a GETRoutingFileMethodNotAllowed with default headers values
func NewGETRoutingFileMethodNotAllowed() *GETRoutingFileMethodNotAllowed {
	return &GETRoutingFileMethodNotAllowed{}
}

/*GETRoutingFileMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETRoutingFileMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETRoutingFileMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETRoutingFileMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRoutingFileNotAcceptable creates a GETRoutingFileNotAcceptable with default headers values
func NewGETRoutingFileNotAcceptable() *GETRoutingFileNotAcceptable {
	return &GETRoutingFileNotAcceptable{}
}

/*GETRoutingFileNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETRoutingFileNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETRoutingFileNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETRoutingFileNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRoutingFileUnsupportedMediaType creates a GETRoutingFileUnsupportedMediaType with default headers values
func NewGETRoutingFileUnsupportedMediaType() *GETRoutingFileUnsupportedMediaType {
	return &GETRoutingFileUnsupportedMediaType{}
}

/*GETRoutingFileUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETRoutingFileUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETRoutingFileUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETRoutingFileUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRoutingFileInternalServerError creates a GETRoutingFileInternalServerError with default headers values
func NewGETRoutingFileInternalServerError() *GETRoutingFileInternalServerError {
	return &GETRoutingFileInternalServerError{}
}

/*GETRoutingFileInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETRoutingFileInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETRoutingFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/routingfiles/{id4n}][%d] getRoutingFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GETRoutingFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}