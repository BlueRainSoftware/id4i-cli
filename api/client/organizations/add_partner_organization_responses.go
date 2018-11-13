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

// AddPartnerOrganizationReader is a Reader for the AddPartnerOrganization structure.
type AddPartnerOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPartnerOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddPartnerOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewAddPartnerOrganizationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddPartnerOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddPartnerOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddPartnerOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddPartnerOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAddPartnerOrganizationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewAddPartnerOrganizationNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddPartnerOrganizationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewAddPartnerOrganizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddPartnerOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddPartnerOrganizationOK creates a AddPartnerOrganizationOK with default headers values
func NewAddPartnerOrganizationOK() *AddPartnerOrganizationOK {
	return &AddPartnerOrganizationOK{}
}

/*AddPartnerOrganizationOK handles this case with default header values.

OK
*/
type AddPartnerOrganizationOK struct {
}

func (o *AddPartnerOrganizationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationOK ", 200)
}

func (o *AddPartnerOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPartnerOrganizationAccepted creates a AddPartnerOrganizationAccepted with default headers values
func NewAddPartnerOrganizationAccepted() *AddPartnerOrganizationAccepted {
	return &AddPartnerOrganizationAccepted{}
}

/*AddPartnerOrganizationAccepted handles this case with default header values.

Accepted
*/
type AddPartnerOrganizationAccepted struct {
}

func (o *AddPartnerOrganizationAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationAccepted ", 202)
}

func (o *AddPartnerOrganizationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPartnerOrganizationBadRequest creates a AddPartnerOrganizationBadRequest with default headers values
func NewAddPartnerOrganizationBadRequest() *AddPartnerOrganizationBadRequest {
	return &AddPartnerOrganizationBadRequest{}
}

/*AddPartnerOrganizationBadRequest handles this case with default header values.

Bad Request
*/
type AddPartnerOrganizationBadRequest struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *AddPartnerOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartnerOrganizationUnauthorized creates a AddPartnerOrganizationUnauthorized with default headers values
func NewAddPartnerOrganizationUnauthorized() *AddPartnerOrganizationUnauthorized {
	return &AddPartnerOrganizationUnauthorized{}
}

/*AddPartnerOrganizationUnauthorized handles this case with default header values.

Unauthorized
*/
type AddPartnerOrganizationUnauthorized struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *AddPartnerOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartnerOrganizationForbidden creates a AddPartnerOrganizationForbidden with default headers values
func NewAddPartnerOrganizationForbidden() *AddPartnerOrganizationForbidden {
	return &AddPartnerOrganizationForbidden{}
}

/*AddPartnerOrganizationForbidden handles this case with default header values.

Forbidden
*/
type AddPartnerOrganizationForbidden struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *AddPartnerOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartnerOrganizationNotFound creates a AddPartnerOrganizationNotFound with default headers values
func NewAddPartnerOrganizationNotFound() *AddPartnerOrganizationNotFound {
	return &AddPartnerOrganizationNotFound{}
}

/*AddPartnerOrganizationNotFound handles this case with default header values.

Not Found
*/
type AddPartnerOrganizationNotFound struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *AddPartnerOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartnerOrganizationMethodNotAllowed creates a AddPartnerOrganizationMethodNotAllowed with default headers values
func NewAddPartnerOrganizationMethodNotAllowed() *AddPartnerOrganizationMethodNotAllowed {
	return &AddPartnerOrganizationMethodNotAllowed{}
}

/*AddPartnerOrganizationMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type AddPartnerOrganizationMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AddPartnerOrganizationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartnerOrganizationNotAcceptable creates a AddPartnerOrganizationNotAcceptable with default headers values
func NewAddPartnerOrganizationNotAcceptable() *AddPartnerOrganizationNotAcceptable {
	return &AddPartnerOrganizationNotAcceptable{}
}

/*AddPartnerOrganizationNotAcceptable handles this case with default header values.

Not Acceptable
*/
type AddPartnerOrganizationNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationNotAcceptable  %+v", 406, o.Payload)
}

func (o *AddPartnerOrganizationNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartnerOrganizationConflict creates a AddPartnerOrganizationConflict with default headers values
func NewAddPartnerOrganizationConflict() *AddPartnerOrganizationConflict {
	return &AddPartnerOrganizationConflict{}
}

/*AddPartnerOrganizationConflict handles this case with default header values.

Conflict
*/
type AddPartnerOrganizationConflict struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationConflict  %+v", 409, o.Payload)
}

func (o *AddPartnerOrganizationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartnerOrganizationUnsupportedMediaType creates a AddPartnerOrganizationUnsupportedMediaType with default headers values
func NewAddPartnerOrganizationUnsupportedMediaType() *AddPartnerOrganizationUnsupportedMediaType {
	return &AddPartnerOrganizationUnsupportedMediaType{}
}

/*AddPartnerOrganizationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type AddPartnerOrganizationUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *AddPartnerOrganizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartnerOrganizationInternalServerError creates a AddPartnerOrganizationInternalServerError with default headers values
func NewAddPartnerOrganizationInternalServerError() *AddPartnerOrganizationInternalServerError {
	return &AddPartnerOrganizationInternalServerError{}
}

/*AddPartnerOrganizationInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddPartnerOrganizationInternalServerError struct {
	Payload *api_models.APIError
}

func (o *AddPartnerOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organizations/{organizationId}/partner][%d] addPartnerOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *AddPartnerOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
