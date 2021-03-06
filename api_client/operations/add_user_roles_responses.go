// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// AddUserRolesReader is a Reader for the AddUserRoles structure.
type AddUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewAddUserRolesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewAddUserRolesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddUserRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddUserRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddUserRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAddUserRolesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewAddUserRolesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddUserRolesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewAddUserRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddUserRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddUserRolesOK creates a AddUserRolesOK with default headers values
func NewAddUserRolesOK() *AddUserRolesOK {
	return &AddUserRolesOK{}
}

/*AddUserRolesOK handles this case with default header values.

OK
*/
type AddUserRolesOK struct {
}

func (o *AddUserRolesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesOK ", 200)
}

func (o *AddUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUserRolesCreated creates a AddUserRolesCreated with default headers values
func NewAddUserRolesCreated() *AddUserRolesCreated {
	return &AddUserRolesCreated{}
}

/*AddUserRolesCreated handles this case with default header values.

Created
*/
type AddUserRolesCreated struct {
}

func (o *AddUserRolesCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesCreated ", 201)
}

func (o *AddUserRolesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUserRolesAccepted creates a AddUserRolesAccepted with default headers values
func NewAddUserRolesAccepted() *AddUserRolesAccepted {
	return &AddUserRolesAccepted{}
}

/*AddUserRolesAccepted handles this case with default header values.

Accepted
*/
type AddUserRolesAccepted struct {
}

func (o *AddUserRolesAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesAccepted ", 202)
}

func (o *AddUserRolesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUserRolesBadRequest creates a AddUserRolesBadRequest with default headers values
func NewAddUserRolesBadRequest() *AddUserRolesBadRequest {
	return &AddUserRolesBadRequest{}
}

/*AddUserRolesBadRequest handles this case with default header values.

Bad Request
*/
type AddUserRolesBadRequest struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *AddUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRolesUnauthorized creates a AddUserRolesUnauthorized with default headers values
func NewAddUserRolesUnauthorized() *AddUserRolesUnauthorized {
	return &AddUserRolesUnauthorized{}
}

/*AddUserRolesUnauthorized handles this case with default header values.

Unauthorized
*/
type AddUserRolesUnauthorized struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *AddUserRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRolesForbidden creates a AddUserRolesForbidden with default headers values
func NewAddUserRolesForbidden() *AddUserRolesForbidden {
	return &AddUserRolesForbidden{}
}

/*AddUserRolesForbidden handles this case with default header values.

Forbidden
*/
type AddUserRolesForbidden struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *AddUserRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRolesNotFound creates a AddUserRolesNotFound with default headers values
func NewAddUserRolesNotFound() *AddUserRolesNotFound {
	return &AddUserRolesNotFound{}
}

/*AddUserRolesNotFound handles this case with default header values.

Not Found
*/
type AddUserRolesNotFound struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesNotFound  %+v", 404, o.Payload)
}

func (o *AddUserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRolesMethodNotAllowed creates a AddUserRolesMethodNotAllowed with default headers values
func NewAddUserRolesMethodNotAllowed() *AddUserRolesMethodNotAllowed {
	return &AddUserRolesMethodNotAllowed{}
}

/*AddUserRolesMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type AddUserRolesMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AddUserRolesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRolesNotAcceptable creates a AddUserRolesNotAcceptable with default headers values
func NewAddUserRolesNotAcceptable() *AddUserRolesNotAcceptable {
	return &AddUserRolesNotAcceptable{}
}

/*AddUserRolesNotAcceptable handles this case with default header values.

Not Acceptable
*/
type AddUserRolesNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesNotAcceptable  %+v", 406, o.Payload)
}

func (o *AddUserRolesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRolesConflict creates a AddUserRolesConflict with default headers values
func NewAddUserRolesConflict() *AddUserRolesConflict {
	return &AddUserRolesConflict{}
}

/*AddUserRolesConflict handles this case with default header values.

Conflict
*/
type AddUserRolesConflict struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesConflict  %+v", 409, o.Payload)
}

func (o *AddUserRolesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRolesUnsupportedMediaType creates a AddUserRolesUnsupportedMediaType with default headers values
func NewAddUserRolesUnsupportedMediaType() *AddUserRolesUnsupportedMediaType {
	return &AddUserRolesUnsupportedMediaType{}
}

/*AddUserRolesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type AddUserRolesUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *AddUserRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRolesInternalServerError creates a AddUserRolesInternalServerError with default headers values
func NewAddUserRolesInternalServerError() *AddUserRolesInternalServerError {
	return &AddUserRolesInternalServerError{}
}

/*AddUserRolesInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddUserRolesInternalServerError struct {
	Payload *api_models.APIError
}

func (o *AddUserRolesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations/{organizationId}/users/{username}/roles][%d] addUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
