// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// SetOrganizationLogoReader is a Reader for the SetOrganizationLogo structure.
type SetOrganizationLogoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetOrganizationLogoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetOrganizationLogoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewSetOrganizationLogoCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewSetOrganizationLogoAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetOrganizationLogoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSetOrganizationLogoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSetOrganizationLogoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSetOrganizationLogoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSetOrganizationLogoMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewSetOrganizationLogoNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewSetOrganizationLogoConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewSetOrganizationLogoUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSetOrganizationLogoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetOrganizationLogoOK creates a SetOrganizationLogoOK with default headers values
func NewSetOrganizationLogoOK() *SetOrganizationLogoOK {
	return &SetOrganizationLogoOK{}
}

/*SetOrganizationLogoOK handles this case with default header values.

OK
*/
type SetOrganizationLogoOK struct {
	Payload *api_models.PublicImagePresentation
}

func (o *SetOrganizationLogoOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoOK  %+v", 200, o.Payload)
}

func (o *SetOrganizationLogoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PublicImagePresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoCreated creates a SetOrganizationLogoCreated with default headers values
func NewSetOrganizationLogoCreated() *SetOrganizationLogoCreated {
	return &SetOrganizationLogoCreated{}
}

/*SetOrganizationLogoCreated handles this case with default header values.

Created
*/
type SetOrganizationLogoCreated struct {
}

func (o *SetOrganizationLogoCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoCreated ", 201)
}

func (o *SetOrganizationLogoCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetOrganizationLogoAccepted creates a SetOrganizationLogoAccepted with default headers values
func NewSetOrganizationLogoAccepted() *SetOrganizationLogoAccepted {
	return &SetOrganizationLogoAccepted{}
}

/*SetOrganizationLogoAccepted handles this case with default header values.

Accepted
*/
type SetOrganizationLogoAccepted struct {
}

func (o *SetOrganizationLogoAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoAccepted ", 202)
}

func (o *SetOrganizationLogoAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetOrganizationLogoBadRequest creates a SetOrganizationLogoBadRequest with default headers values
func NewSetOrganizationLogoBadRequest() *SetOrganizationLogoBadRequest {
	return &SetOrganizationLogoBadRequest{}
}

/*SetOrganizationLogoBadRequest handles this case with default header values.

Bad Request
*/
type SetOrganizationLogoBadRequest struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoBadRequest  %+v", 400, o.Payload)
}

func (o *SetOrganizationLogoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoUnauthorized creates a SetOrganizationLogoUnauthorized with default headers values
func NewSetOrganizationLogoUnauthorized() *SetOrganizationLogoUnauthorized {
	return &SetOrganizationLogoUnauthorized{}
}

/*SetOrganizationLogoUnauthorized handles this case with default header values.

Unauthorized
*/
type SetOrganizationLogoUnauthorized struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoUnauthorized  %+v", 401, o.Payload)
}

func (o *SetOrganizationLogoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoForbidden creates a SetOrganizationLogoForbidden with default headers values
func NewSetOrganizationLogoForbidden() *SetOrganizationLogoForbidden {
	return &SetOrganizationLogoForbidden{}
}

/*SetOrganizationLogoForbidden handles this case with default header values.

Forbidden
*/
type SetOrganizationLogoForbidden struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoForbidden  %+v", 403, o.Payload)
}

func (o *SetOrganizationLogoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoNotFound creates a SetOrganizationLogoNotFound with default headers values
func NewSetOrganizationLogoNotFound() *SetOrganizationLogoNotFound {
	return &SetOrganizationLogoNotFound{}
}

/*SetOrganizationLogoNotFound handles this case with default header values.

Not Found
*/
type SetOrganizationLogoNotFound struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoNotFound  %+v", 404, o.Payload)
}

func (o *SetOrganizationLogoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoMethodNotAllowed creates a SetOrganizationLogoMethodNotAllowed with default headers values
func NewSetOrganizationLogoMethodNotAllowed() *SetOrganizationLogoMethodNotAllowed {
	return &SetOrganizationLogoMethodNotAllowed{}
}

/*SetOrganizationLogoMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type SetOrganizationLogoMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SetOrganizationLogoMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoNotAcceptable creates a SetOrganizationLogoNotAcceptable with default headers values
func NewSetOrganizationLogoNotAcceptable() *SetOrganizationLogoNotAcceptable {
	return &SetOrganizationLogoNotAcceptable{}
}

/*SetOrganizationLogoNotAcceptable handles this case with default header values.

Not Acceptable
*/
type SetOrganizationLogoNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoNotAcceptable  %+v", 406, o.Payload)
}

func (o *SetOrganizationLogoNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoConflict creates a SetOrganizationLogoConflict with default headers values
func NewSetOrganizationLogoConflict() *SetOrganizationLogoConflict {
	return &SetOrganizationLogoConflict{}
}

/*SetOrganizationLogoConflict handles this case with default header values.

Conflict
*/
type SetOrganizationLogoConflict struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoConflict  %+v", 409, o.Payload)
}

func (o *SetOrganizationLogoConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoUnsupportedMediaType creates a SetOrganizationLogoUnsupportedMediaType with default headers values
func NewSetOrganizationLogoUnsupportedMediaType() *SetOrganizationLogoUnsupportedMediaType {
	return &SetOrganizationLogoUnsupportedMediaType{}
}

/*SetOrganizationLogoUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type SetOrganizationLogoUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SetOrganizationLogoUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrganizationLogoInternalServerError creates a SetOrganizationLogoInternalServerError with default headers values
func NewSetOrganizationLogoInternalServerError() *SetOrganizationLogoInternalServerError {
	return &SetOrganizationLogoInternalServerError{}
}

/*SetOrganizationLogoInternalServerError handles this case with default header values.

Internal Server Error
*/
type SetOrganizationLogoInternalServerError struct {
	Payload *api_models.APIError
}

func (o *SetOrganizationLogoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/logo][%d] setOrganizationLogoInternalServerError  %+v", 500, o.Payload)
}

func (o *SetOrganizationLogoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
