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

// WriteToMicrostorageReader is a Reader for the WriteToMicrostorage structure.
type WriteToMicrostorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WriteToMicrostorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWriteToMicrostorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewWriteToMicrostorageAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewWriteToMicrostorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewWriteToMicrostorageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWriteToMicrostorageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWriteToMicrostorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewWriteToMicrostorageMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewWriteToMicrostorageNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewWriteToMicrostorageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewWriteToMicrostorageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWriteToMicrostorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWriteToMicrostorageOK creates a WriteToMicrostorageOK with default headers values
func NewWriteToMicrostorageOK() *WriteToMicrostorageOK {
	return &WriteToMicrostorageOK{}
}

/*WriteToMicrostorageOK handles this case with default header values.

OK
*/
type WriteToMicrostorageOK struct {
	Payload interface{}
}

func (o *WriteToMicrostorageOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageOK  %+v", 200, o.Payload)
}

func (o *WriteToMicrostorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageAccepted creates a WriteToMicrostorageAccepted with default headers values
func NewWriteToMicrostorageAccepted() *WriteToMicrostorageAccepted {
	return &WriteToMicrostorageAccepted{}
}

/*WriteToMicrostorageAccepted handles this case with default header values.

Accepted
*/
type WriteToMicrostorageAccepted struct {
}

func (o *WriteToMicrostorageAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageAccepted ", 202)
}

func (o *WriteToMicrostorageAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWriteToMicrostorageBadRequest creates a WriteToMicrostorageBadRequest with default headers values
func NewWriteToMicrostorageBadRequest() *WriteToMicrostorageBadRequest {
	return &WriteToMicrostorageBadRequest{}
}

/*WriteToMicrostorageBadRequest handles this case with default header values.

Bad Request
*/
type WriteToMicrostorageBadRequest struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageBadRequest  %+v", 400, o.Payload)
}

func (o *WriteToMicrostorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageUnauthorized creates a WriteToMicrostorageUnauthorized with default headers values
func NewWriteToMicrostorageUnauthorized() *WriteToMicrostorageUnauthorized {
	return &WriteToMicrostorageUnauthorized{}
}

/*WriteToMicrostorageUnauthorized handles this case with default header values.

Unauthorized
*/
type WriteToMicrostorageUnauthorized struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageUnauthorized  %+v", 401, o.Payload)
}

func (o *WriteToMicrostorageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageForbidden creates a WriteToMicrostorageForbidden with default headers values
func NewWriteToMicrostorageForbidden() *WriteToMicrostorageForbidden {
	return &WriteToMicrostorageForbidden{}
}

/*WriteToMicrostorageForbidden handles this case with default header values.

Forbidden
*/
type WriteToMicrostorageForbidden struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageForbidden  %+v", 403, o.Payload)
}

func (o *WriteToMicrostorageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageNotFound creates a WriteToMicrostorageNotFound with default headers values
func NewWriteToMicrostorageNotFound() *WriteToMicrostorageNotFound {
	return &WriteToMicrostorageNotFound{}
}

/*WriteToMicrostorageNotFound handles this case with default header values.

Not Found
*/
type WriteToMicrostorageNotFound struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageNotFound  %+v", 404, o.Payload)
}

func (o *WriteToMicrostorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageMethodNotAllowed creates a WriteToMicrostorageMethodNotAllowed with default headers values
func NewWriteToMicrostorageMethodNotAllowed() *WriteToMicrostorageMethodNotAllowed {
	return &WriteToMicrostorageMethodNotAllowed{}
}

/*WriteToMicrostorageMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type WriteToMicrostorageMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *WriteToMicrostorageMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageNotAcceptable creates a WriteToMicrostorageNotAcceptable with default headers values
func NewWriteToMicrostorageNotAcceptable() *WriteToMicrostorageNotAcceptable {
	return &WriteToMicrostorageNotAcceptable{}
}

/*WriteToMicrostorageNotAcceptable handles this case with default header values.

Not Acceptable
*/
type WriteToMicrostorageNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageNotAcceptable  %+v", 406, o.Payload)
}

func (o *WriteToMicrostorageNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageConflict creates a WriteToMicrostorageConflict with default headers values
func NewWriteToMicrostorageConflict() *WriteToMicrostorageConflict {
	return &WriteToMicrostorageConflict{}
}

/*WriteToMicrostorageConflict handles this case with default header values.

Conflict
*/
type WriteToMicrostorageConflict struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageConflict  %+v", 409, o.Payload)
}

func (o *WriteToMicrostorageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageUnsupportedMediaType creates a WriteToMicrostorageUnsupportedMediaType with default headers values
func NewWriteToMicrostorageUnsupportedMediaType() *WriteToMicrostorageUnsupportedMediaType {
	return &WriteToMicrostorageUnsupportedMediaType{}
}

/*WriteToMicrostorageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type WriteToMicrostorageUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *WriteToMicrostorageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWriteToMicrostorageInternalServerError creates a WriteToMicrostorageInternalServerError with default headers values
func NewWriteToMicrostorageInternalServerError() *WriteToMicrostorageInternalServerError {
	return &WriteToMicrostorageInternalServerError{}
}

/*WriteToMicrostorageInternalServerError handles this case with default header values.

Internal Server Error
*/
type WriteToMicrostorageInternalServerError struct {
	Payload *api_models.APIError
}

func (o *WriteToMicrostorageInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/microstorage/{id4n}/{organization}][%d] writeToMicrostorageInternalServerError  %+v", 500, o.Payload)
}

func (o *WriteToMicrostorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
