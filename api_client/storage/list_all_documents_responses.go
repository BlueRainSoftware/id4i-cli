// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// ListAllDocumentsReader is a Reader for the ListAllDocuments structure.
type ListAllDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAllDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewListAllDocumentsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListAllDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListAllDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListAllDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListAllDocumentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewListAllDocumentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewListAllDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAllDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAllDocumentsOK creates a ListAllDocumentsOK with default headers values
func NewListAllDocumentsOK() *ListAllDocumentsOK {
	return &ListAllDocumentsOK{}
}

/*ListAllDocumentsOK handles this case with default header values.

OK
*/
type ListAllDocumentsOK struct {
	Payload *api_models.PaginatedResponseOfDocument
}

func (o *ListAllDocumentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsOK  %+v", 200, o.Payload)
}

func (o *ListAllDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedResponseOfDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllDocumentsAccepted creates a ListAllDocumentsAccepted with default headers values
func NewListAllDocumentsAccepted() *ListAllDocumentsAccepted {
	return &ListAllDocumentsAccepted{}
}

/*ListAllDocumentsAccepted handles this case with default header values.

Accepted
*/
type ListAllDocumentsAccepted struct {
}

func (o *ListAllDocumentsAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsAccepted ", 202)
}

func (o *ListAllDocumentsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllDocumentsUnauthorized creates a ListAllDocumentsUnauthorized with default headers values
func NewListAllDocumentsUnauthorized() *ListAllDocumentsUnauthorized {
	return &ListAllDocumentsUnauthorized{}
}

/*ListAllDocumentsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAllDocumentsUnauthorized struct {
	Payload *api_models.APIError
}

func (o *ListAllDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAllDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllDocumentsForbidden creates a ListAllDocumentsForbidden with default headers values
func NewListAllDocumentsForbidden() *ListAllDocumentsForbidden {
	return &ListAllDocumentsForbidden{}
}

/*ListAllDocumentsForbidden handles this case with default header values.

Forbidden
*/
type ListAllDocumentsForbidden struct {
	Payload *api_models.APIError
}

func (o *ListAllDocumentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *ListAllDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllDocumentsNotFound creates a ListAllDocumentsNotFound with default headers values
func NewListAllDocumentsNotFound() *ListAllDocumentsNotFound {
	return &ListAllDocumentsNotFound{}
}

/*ListAllDocumentsNotFound handles this case with default header values.

Not Found
*/
type ListAllDocumentsNotFound struct {
	Payload *api_models.APIError
}

func (o *ListAllDocumentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *ListAllDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllDocumentsMethodNotAllowed creates a ListAllDocumentsMethodNotAllowed with default headers values
func NewListAllDocumentsMethodNotAllowed() *ListAllDocumentsMethodNotAllowed {
	return &ListAllDocumentsMethodNotAllowed{}
}

/*ListAllDocumentsMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type ListAllDocumentsMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *ListAllDocumentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListAllDocumentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllDocumentsNotAcceptable creates a ListAllDocumentsNotAcceptable with default headers values
func NewListAllDocumentsNotAcceptable() *ListAllDocumentsNotAcceptable {
	return &ListAllDocumentsNotAcceptable{}
}

/*ListAllDocumentsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type ListAllDocumentsNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *ListAllDocumentsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsNotAcceptable  %+v", 406, o.Payload)
}

func (o *ListAllDocumentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllDocumentsUnsupportedMediaType creates a ListAllDocumentsUnsupportedMediaType with default headers values
func NewListAllDocumentsUnsupportedMediaType() *ListAllDocumentsUnsupportedMediaType {
	return &ListAllDocumentsUnsupportedMediaType{}
}

/*ListAllDocumentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type ListAllDocumentsUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *ListAllDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListAllDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllDocumentsInternalServerError creates a ListAllDocumentsInternalServerError with default headers values
func NewListAllDocumentsInternalServerError() *ListAllDocumentsInternalServerError {
	return &ListAllDocumentsInternalServerError{}
}

/*ListAllDocumentsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAllDocumentsInternalServerError struct {
	Payload *api_models.APIError
}

func (o *ListAllDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/documents/{id4n}][%d] listAllDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
