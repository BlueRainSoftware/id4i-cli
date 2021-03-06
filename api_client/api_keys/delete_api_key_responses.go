// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// DeleteAPIKeyReader is a Reader for the DeleteAPIKey structure.
type DeleteAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAPIKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteAPIKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteAPIKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteAPIKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPIKeyOK creates a DeleteAPIKeyOK with default headers values
func NewDeleteAPIKeyOK() *DeleteAPIKeyOK {
	return &DeleteAPIKeyOK{}
}

/*DeleteAPIKeyOK handles this case with default header values.

OK
*/
type DeleteAPIKeyOK struct {
}

func (o *DeleteAPIKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}][%d] deleteApiKeyOK ", 200)
}

func (o *DeleteAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIKeyBadRequest creates a DeleteAPIKeyBadRequest with default headers values
func NewDeleteAPIKeyBadRequest() *DeleteAPIKeyBadRequest {
	return &DeleteAPIKeyBadRequest{}
}

/*DeleteAPIKeyBadRequest handles this case with default header values.

Bad Request
*/
type DeleteAPIKeyBadRequest struct {
	Payload *api_models.APIError
}

func (o *DeleteAPIKeyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}][%d] deleteApiKeyBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIKeyUnauthorized creates a DeleteAPIKeyUnauthorized with default headers values
func NewDeleteAPIKeyUnauthorized() *DeleteAPIKeyUnauthorized {
	return &DeleteAPIKeyUnauthorized{}
}

/*DeleteAPIKeyUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAPIKeyUnauthorized struct {
	Payload *api_models.APIError
}

func (o *DeleteAPIKeyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}][%d] deleteApiKeyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAPIKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIKeyForbidden creates a DeleteAPIKeyForbidden with default headers values
func NewDeleteAPIKeyForbidden() *DeleteAPIKeyForbidden {
	return &DeleteAPIKeyForbidden{}
}

/*DeleteAPIKeyForbidden handles this case with default header values.

Forbidden
*/
type DeleteAPIKeyForbidden struct {
	Payload *api_models.APIError
}

func (o *DeleteAPIKeyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}][%d] deleteApiKeyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAPIKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIKeyNotFound creates a DeleteAPIKeyNotFound with default headers values
func NewDeleteAPIKeyNotFound() *DeleteAPIKeyNotFound {
	return &DeleteAPIKeyNotFound{}
}

/*DeleteAPIKeyNotFound handles this case with default header values.

Not Found
*/
type DeleteAPIKeyNotFound struct {
	Payload *api_models.APIError
}

func (o *DeleteAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}][%d] deleteApiKeyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIKeyInternalServerError creates a DeleteAPIKeyInternalServerError with default headers values
func NewDeleteAPIKeyInternalServerError() *DeleteAPIKeyInternalServerError {
	return &DeleteAPIKeyInternalServerError{}
}

/*DeleteAPIKeyInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteAPIKeyInternalServerError struct {
	Payload *api_models.APIError
}

func (o *DeleteAPIKeyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}][%d] deleteApiKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
