// Code generated by go-swagger; DO NOT EDIT.

package transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// PrepareReader is a Reader for the Prepare structure.
type PrepareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrepareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPrepareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewPrepareAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPrepareBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPrepareUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPrepareForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPrepareNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPrepareMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewPrepareNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPrepareConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewPrepareUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPrepareInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPrepareOK creates a PrepareOK with default headers values
func NewPrepareOK() *PrepareOK {
	return &PrepareOK{}
}

/*PrepareOK handles this case with default header values.

OK
*/
type PrepareOK struct {
	Payload interface{}
}

func (o *PrepareOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareOK  %+v", 200, o.Payload)
}

func (o *PrepareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareAccepted creates a PrepareAccepted with default headers values
func NewPrepareAccepted() *PrepareAccepted {
	return &PrepareAccepted{}
}

/*PrepareAccepted handles this case with default header values.

Accepted
*/
type PrepareAccepted struct {
}

func (o *PrepareAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareAccepted ", 202)
}

func (o *PrepareAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPrepareBadRequest creates a PrepareBadRequest with default headers values
func NewPrepareBadRequest() *PrepareBadRequest {
	return &PrepareBadRequest{}
}

/*PrepareBadRequest handles this case with default header values.

Bad Request
*/
type PrepareBadRequest struct {
	Payload *api_models.APIError
}

func (o *PrepareBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareBadRequest  %+v", 400, o.Payload)
}

func (o *PrepareBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareUnauthorized creates a PrepareUnauthorized with default headers values
func NewPrepareUnauthorized() *PrepareUnauthorized {
	return &PrepareUnauthorized{}
}

/*PrepareUnauthorized handles this case with default header values.

Unauthorized
*/
type PrepareUnauthorized struct {
	Payload *api_models.APIError
}

func (o *PrepareUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareUnauthorized  %+v", 401, o.Payload)
}

func (o *PrepareUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareForbidden creates a PrepareForbidden with default headers values
func NewPrepareForbidden() *PrepareForbidden {
	return &PrepareForbidden{}
}

/*PrepareForbidden handles this case with default header values.

Forbidden
*/
type PrepareForbidden struct {
	Payload *api_models.APIError
}

func (o *PrepareForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareForbidden  %+v", 403, o.Payload)
}

func (o *PrepareForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareNotFound creates a PrepareNotFound with default headers values
func NewPrepareNotFound() *PrepareNotFound {
	return &PrepareNotFound{}
}

/*PrepareNotFound handles this case with default header values.

Not Found
*/
type PrepareNotFound struct {
	Payload *api_models.APIError
}

func (o *PrepareNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareNotFound  %+v", 404, o.Payload)
}

func (o *PrepareNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareMethodNotAllowed creates a PrepareMethodNotAllowed with default headers values
func NewPrepareMethodNotAllowed() *PrepareMethodNotAllowed {
	return &PrepareMethodNotAllowed{}
}

/*PrepareMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type PrepareMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *PrepareMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PrepareMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareNotAcceptable creates a PrepareNotAcceptable with default headers values
func NewPrepareNotAcceptable() *PrepareNotAcceptable {
	return &PrepareNotAcceptable{}
}

/*PrepareNotAcceptable handles this case with default header values.

Not Acceptable
*/
type PrepareNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *PrepareNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareNotAcceptable  %+v", 406, o.Payload)
}

func (o *PrepareNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareConflict creates a PrepareConflict with default headers values
func NewPrepareConflict() *PrepareConflict {
	return &PrepareConflict{}
}

/*PrepareConflict handles this case with default header values.

Conflict
*/
type PrepareConflict struct {
	Payload *api_models.APIError
}

func (o *PrepareConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareConflict  %+v", 409, o.Payload)
}

func (o *PrepareConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareUnsupportedMediaType creates a PrepareUnsupportedMediaType with default headers values
func NewPrepareUnsupportedMediaType() *PrepareUnsupportedMediaType {
	return &PrepareUnsupportedMediaType{}
}

/*PrepareUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type PrepareUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *PrepareUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PrepareUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareInternalServerError creates a PrepareInternalServerError with default headers values
func NewPrepareInternalServerError() *PrepareInternalServerError {
	return &PrepareInternalServerError{}
}

/*PrepareInternalServerError handles this case with default header values.

Internal Server Error
*/
type PrepareInternalServerError struct {
	Payload *api_models.APIError
}

func (o *PrepareInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/transfers/{id4n}/sendInfo][%d] prepareInternalServerError  %+v", 500, o.Payload)
}

func (o *PrepareInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
