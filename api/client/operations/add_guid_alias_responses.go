// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// AddGUIDAliasReader is a Reader for the AddGUIDAlias structure.
type AddGUIDAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddGUIDAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddGUIDAliasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewAddGUIDAliasCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewAddGUIDAliasAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddGUIDAliasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddGUIDAliasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddGUIDAliasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddGUIDAliasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAddGUIDAliasMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewAddGUIDAliasNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddGUIDAliasConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewAddGUIDAliasUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddGUIDAliasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddGUIDAliasOK creates a AddGUIDAliasOK with default headers values
func NewAddGUIDAliasOK() *AddGUIDAliasOK {
	return &AddGUIDAliasOK{}
}

/*AddGUIDAliasOK handles this case with default header values.

OK
*/
type AddGUIDAliasOK struct {
}

func (o *AddGUIDAliasOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasOK ", 200)
}

func (o *AddGUIDAliasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddGUIDAliasCreated creates a AddGUIDAliasCreated with default headers values
func NewAddGUIDAliasCreated() *AddGUIDAliasCreated {
	return &AddGUIDAliasCreated{}
}

/*AddGUIDAliasCreated handles this case with default header values.

Created
*/
type AddGUIDAliasCreated struct {
}

func (o *AddGUIDAliasCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasCreated ", 201)
}

func (o *AddGUIDAliasCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddGUIDAliasAccepted creates a AddGUIDAliasAccepted with default headers values
func NewAddGUIDAliasAccepted() *AddGUIDAliasAccepted {
	return &AddGUIDAliasAccepted{}
}

/*AddGUIDAliasAccepted handles this case with default header values.

Accepted
*/
type AddGUIDAliasAccepted struct {
}

func (o *AddGUIDAliasAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasAccepted ", 202)
}

func (o *AddGUIDAliasAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddGUIDAliasBadRequest creates a AddGUIDAliasBadRequest with default headers values
func NewAddGUIDAliasBadRequest() *AddGUIDAliasBadRequest {
	return &AddGUIDAliasBadRequest{}
}

/*AddGUIDAliasBadRequest handles this case with default header values.

Bad Request
*/
type AddGUIDAliasBadRequest struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasBadRequest  %+v", 400, o.Payload)
}

func (o *AddGUIDAliasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGUIDAliasUnauthorized creates a AddGUIDAliasUnauthorized with default headers values
func NewAddGUIDAliasUnauthorized() *AddGUIDAliasUnauthorized {
	return &AddGUIDAliasUnauthorized{}
}

/*AddGUIDAliasUnauthorized handles this case with default header values.

Unauthorized
*/
type AddGUIDAliasUnauthorized struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasUnauthorized  %+v", 401, o.Payload)
}

func (o *AddGUIDAliasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGUIDAliasForbidden creates a AddGUIDAliasForbidden with default headers values
func NewAddGUIDAliasForbidden() *AddGUIDAliasForbidden {
	return &AddGUIDAliasForbidden{}
}

/*AddGUIDAliasForbidden handles this case with default header values.

Forbidden
*/
type AddGUIDAliasForbidden struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasForbidden  %+v", 403, o.Payload)
}

func (o *AddGUIDAliasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGUIDAliasNotFound creates a AddGUIDAliasNotFound with default headers values
func NewAddGUIDAliasNotFound() *AddGUIDAliasNotFound {
	return &AddGUIDAliasNotFound{}
}

/*AddGUIDAliasNotFound handles this case with default header values.

Not Found
*/
type AddGUIDAliasNotFound struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasNotFound  %+v", 404, o.Payload)
}

func (o *AddGUIDAliasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGUIDAliasMethodNotAllowed creates a AddGUIDAliasMethodNotAllowed with default headers values
func NewAddGUIDAliasMethodNotAllowed() *AddGUIDAliasMethodNotAllowed {
	return &AddGUIDAliasMethodNotAllowed{}
}

/*AddGUIDAliasMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type AddGUIDAliasMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AddGUIDAliasMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGUIDAliasNotAcceptable creates a AddGUIDAliasNotAcceptable with default headers values
func NewAddGUIDAliasNotAcceptable() *AddGUIDAliasNotAcceptable {
	return &AddGUIDAliasNotAcceptable{}
}

/*AddGUIDAliasNotAcceptable handles this case with default header values.

Not Acceptable
*/
type AddGUIDAliasNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasNotAcceptable  %+v", 406, o.Payload)
}

func (o *AddGUIDAliasNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGUIDAliasConflict creates a AddGUIDAliasConflict with default headers values
func NewAddGUIDAliasConflict() *AddGUIDAliasConflict {
	return &AddGUIDAliasConflict{}
}

/*AddGUIDAliasConflict handles this case with default header values.

Conflict
*/
type AddGUIDAliasConflict struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasConflict  %+v", 409, o.Payload)
}

func (o *AddGUIDAliasConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGUIDAliasUnsupportedMediaType creates a AddGUIDAliasUnsupportedMediaType with default headers values
func NewAddGUIDAliasUnsupportedMediaType() *AddGUIDAliasUnsupportedMediaType {
	return &AddGUIDAliasUnsupportedMediaType{}
}

/*AddGUIDAliasUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type AddGUIDAliasUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *AddGUIDAliasUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGUIDAliasInternalServerError creates a AddGUIDAliasInternalServerError with default headers values
func NewAddGUIDAliasInternalServerError() *AddGUIDAliasInternalServerError {
	return &AddGUIDAliasInternalServerError{}
}

/*AddGUIDAliasInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddGUIDAliasInternalServerError struct {
	Payload *api_models.APIError
}

func (o *AddGUIDAliasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] addGuidAliasInternalServerError  %+v", 500, o.Payload)
}

func (o *AddGUIDAliasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
