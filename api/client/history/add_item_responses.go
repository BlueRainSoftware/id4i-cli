// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// AddItemReader is a Reader for the AddItem structure.
type AddItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewAddItemCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewAddItemAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddItemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAddItemMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewAddItemNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddItemConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewAddItemUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddItemOK creates a AddItemOK with default headers values
func NewAddItemOK() *AddItemOK {
	return &AddItemOK{}
}

/*AddItemOK handles this case with default header values.

OK
*/
type AddItemOK struct {
}

func (o *AddItemOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemOK ", 200)
}

func (o *AddItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddItemCreated creates a AddItemCreated with default headers values
func NewAddItemCreated() *AddItemCreated {
	return &AddItemCreated{}
}

/*AddItemCreated handles this case with default header values.

Created
*/
type AddItemCreated struct {
}

func (o *AddItemCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemCreated ", 201)
}

func (o *AddItemCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddItemAccepted creates a AddItemAccepted with default headers values
func NewAddItemAccepted() *AddItemAccepted {
	return &AddItemAccepted{}
}

/*AddItemAccepted handles this case with default header values.

Accepted
*/
type AddItemAccepted struct {
}

func (o *AddItemAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemAccepted ", 202)
}

func (o *AddItemAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddItemBadRequest creates a AddItemBadRequest with default headers values
func NewAddItemBadRequest() *AddItemBadRequest {
	return &AddItemBadRequest{}
}

/*AddItemBadRequest handles this case with default header values.

Bad Request
*/
type AddItemBadRequest struct {
	Payload *api_models.APIError
}

func (o *AddItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemBadRequest  %+v", 400, o.Payload)
}

func (o *AddItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemUnauthorized creates a AddItemUnauthorized with default headers values
func NewAddItemUnauthorized() *AddItemUnauthorized {
	return &AddItemUnauthorized{}
}

/*AddItemUnauthorized handles this case with default header values.

Unauthorized
*/
type AddItemUnauthorized struct {
	Payload *api_models.APIError
}

func (o *AddItemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemUnauthorized  %+v", 401, o.Payload)
}

func (o *AddItemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemForbidden creates a AddItemForbidden with default headers values
func NewAddItemForbidden() *AddItemForbidden {
	return &AddItemForbidden{}
}

/*AddItemForbidden handles this case with default header values.

Forbidden
*/
type AddItemForbidden struct {
	Payload *api_models.APIError
}

func (o *AddItemForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemForbidden  %+v", 403, o.Payload)
}

func (o *AddItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemNotFound creates a AddItemNotFound with default headers values
func NewAddItemNotFound() *AddItemNotFound {
	return &AddItemNotFound{}
}

/*AddItemNotFound handles this case with default header values.

Not Found
*/
type AddItemNotFound struct {
	Payload *api_models.APIError
}

func (o *AddItemNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemNotFound  %+v", 404, o.Payload)
}

func (o *AddItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemMethodNotAllowed creates a AddItemMethodNotAllowed with default headers values
func NewAddItemMethodNotAllowed() *AddItemMethodNotAllowed {
	return &AddItemMethodNotAllowed{}
}

/*AddItemMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type AddItemMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *AddItemMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AddItemMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemNotAcceptable creates a AddItemNotAcceptable with default headers values
func NewAddItemNotAcceptable() *AddItemNotAcceptable {
	return &AddItemNotAcceptable{}
}

/*AddItemNotAcceptable handles this case with default header values.

Not Acceptable
*/
type AddItemNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *AddItemNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemNotAcceptable  %+v", 406, o.Payload)
}

func (o *AddItemNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemConflict creates a AddItemConflict with default headers values
func NewAddItemConflict() *AddItemConflict {
	return &AddItemConflict{}
}

/*AddItemConflict handles this case with default header values.

Conflict
*/
type AddItemConflict struct {
	Payload *api_models.APIError
}

func (o *AddItemConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemConflict  %+v", 409, o.Payload)
}

func (o *AddItemConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemUnsupportedMediaType creates a AddItemUnsupportedMediaType with default headers values
func NewAddItemUnsupportedMediaType() *AddItemUnsupportedMediaType {
	return &AddItemUnsupportedMediaType{}
}

/*AddItemUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type AddItemUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *AddItemUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *AddItemUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemInternalServerError creates a AddItemInternalServerError with default headers values
func NewAddItemInternalServerError() *AddItemInternalServerError {
	return &AddItemInternalServerError{}
}

/*AddItemInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddItemInternalServerError struct {
	Payload *api_models.APIError
}

func (o *AddItemInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/history/{id4n}][%d] addItemInternalServerError  %+v", 500, o.Payload)
}

func (o *AddItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
