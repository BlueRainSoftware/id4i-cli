// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// RemoveAPIKeyPrivilegeForId4nsReader is a Reader for the RemoveAPIKeyPrivilegeForId4ns structure.
type RemoveAPIKeyPrivilegeForId4nsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAPIKeyPrivilegeForId4nsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveAPIKeyPrivilegeForId4nsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveAPIKeyPrivilegeForId4nsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveAPIKeyPrivilegeForId4nsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveAPIKeyPrivilegeForId4nsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveAPIKeyPrivilegeForId4nsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRemoveAPIKeyPrivilegeForId4nsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveAPIKeyPrivilegeForId4nsOK creates a RemoveAPIKeyPrivilegeForId4nsOK with default headers values
func NewRemoveAPIKeyPrivilegeForId4nsOK() *RemoveAPIKeyPrivilegeForId4nsOK {
	return &RemoveAPIKeyPrivilegeForId4nsOK{}
}

/*RemoveAPIKeyPrivilegeForId4nsOK handles this case with default header values.

OK
*/
type RemoveAPIKeyPrivilegeForId4nsOK struct {
}

func (o *RemoveAPIKeyPrivilegeForId4nsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges/{privilege}/id4ns][%d] removeApiKeyPrivilegeForId4nsOK ", 200)
}

func (o *RemoveAPIKeyPrivilegeForId4nsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAPIKeyPrivilegeForId4nsBadRequest creates a RemoveAPIKeyPrivilegeForId4nsBadRequest with default headers values
func NewRemoveAPIKeyPrivilegeForId4nsBadRequest() *RemoveAPIKeyPrivilegeForId4nsBadRequest {
	return &RemoveAPIKeyPrivilegeForId4nsBadRequest{}
}

/*RemoveAPIKeyPrivilegeForId4nsBadRequest handles this case with default header values.

Bad Request
*/
type RemoveAPIKeyPrivilegeForId4nsBadRequest struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeForId4nsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges/{privilege}/id4ns][%d] removeApiKeyPrivilegeForId4nsBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeForId4nsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAPIKeyPrivilegeForId4nsUnauthorized creates a RemoveAPIKeyPrivilegeForId4nsUnauthorized with default headers values
func NewRemoveAPIKeyPrivilegeForId4nsUnauthorized() *RemoveAPIKeyPrivilegeForId4nsUnauthorized {
	return &RemoveAPIKeyPrivilegeForId4nsUnauthorized{}
}

/*RemoveAPIKeyPrivilegeForId4nsUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveAPIKeyPrivilegeForId4nsUnauthorized struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeForId4nsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges/{privilege}/id4ns][%d] removeApiKeyPrivilegeForId4nsUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeForId4nsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAPIKeyPrivilegeForId4nsForbidden creates a RemoveAPIKeyPrivilegeForId4nsForbidden with default headers values
func NewRemoveAPIKeyPrivilegeForId4nsForbidden() *RemoveAPIKeyPrivilegeForId4nsForbidden {
	return &RemoveAPIKeyPrivilegeForId4nsForbidden{}
}

/*RemoveAPIKeyPrivilegeForId4nsForbidden handles this case with default header values.

Forbidden
*/
type RemoveAPIKeyPrivilegeForId4nsForbidden struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeForId4nsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges/{privilege}/id4ns][%d] removeApiKeyPrivilegeForId4nsForbidden  %+v", 403, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeForId4nsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAPIKeyPrivilegeForId4nsNotFound creates a RemoveAPIKeyPrivilegeForId4nsNotFound with default headers values
func NewRemoveAPIKeyPrivilegeForId4nsNotFound() *RemoveAPIKeyPrivilegeForId4nsNotFound {
	return &RemoveAPIKeyPrivilegeForId4nsNotFound{}
}

/*RemoveAPIKeyPrivilegeForId4nsNotFound handles this case with default header values.

Not Found
*/
type RemoveAPIKeyPrivilegeForId4nsNotFound struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeForId4nsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges/{privilege}/id4ns][%d] removeApiKeyPrivilegeForId4nsNotFound  %+v", 404, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeForId4nsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAPIKeyPrivilegeForId4nsInternalServerError creates a RemoveAPIKeyPrivilegeForId4nsInternalServerError with default headers values
func NewRemoveAPIKeyPrivilegeForId4nsInternalServerError() *RemoveAPIKeyPrivilegeForId4nsInternalServerError {
	return &RemoveAPIKeyPrivilegeForId4nsInternalServerError{}
}

/*RemoveAPIKeyPrivilegeForId4nsInternalServerError handles this case with default header values.

Internal Server Error
*/
type RemoveAPIKeyPrivilegeForId4nsInternalServerError struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeForId4nsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges/{privilege}/id4ns][%d] removeApiKeyPrivilegeForId4nsInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeForId4nsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
