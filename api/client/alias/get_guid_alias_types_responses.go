// Code generated by go-swagger; DO NOT EDIT.

package alias

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api/models"
)

// GETGUIDAliasTypesReader is a Reader for the GETGUIDAliasTypes structure.
type GETGUIDAliasTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETGUIDAliasTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETGUIDAliasTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETGUIDAliasTypesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETGUIDAliasTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETGUIDAliasTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETGUIDAliasTypesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETGUIDAliasTypesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETGUIDAliasTypesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETGUIDAliasTypesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETGUIDAliasTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETGUIDAliasTypesOK creates a GETGUIDAliasTypesOK with default headers values
func NewGETGUIDAliasTypesOK() *GETGUIDAliasTypesOK {
	return &GETGUIDAliasTypesOK{}
}

/*GETGUIDAliasTypesOK handles this case with default header values.

OK
*/
type GETGUIDAliasTypesOK struct {
	Payload []string
}

func (o *GETGUIDAliasTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesOK  %+v", 200, o.Payload)
}

func (o *GETGUIDAliasTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDAliasTypesAccepted creates a GETGUIDAliasTypesAccepted with default headers values
func NewGETGUIDAliasTypesAccepted() *GETGUIDAliasTypesAccepted {
	return &GETGUIDAliasTypesAccepted{}
}

/*GETGUIDAliasTypesAccepted handles this case with default header values.

Accepted
*/
type GETGUIDAliasTypesAccepted struct {
}

func (o *GETGUIDAliasTypesAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesAccepted ", 202)
}

func (o *GETGUIDAliasTypesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETGUIDAliasTypesUnauthorized creates a GETGUIDAliasTypesUnauthorized with default headers values
func NewGETGUIDAliasTypesUnauthorized() *GETGUIDAliasTypesUnauthorized {
	return &GETGUIDAliasTypesUnauthorized{}
}

/*GETGUIDAliasTypesUnauthorized handles this case with default header values.

Unauthorized
*/
type GETGUIDAliasTypesUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETGUIDAliasTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GETGUIDAliasTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDAliasTypesForbidden creates a GETGUIDAliasTypesForbidden with default headers values
func NewGETGUIDAliasTypesForbidden() *GETGUIDAliasTypesForbidden {
	return &GETGUIDAliasTypesForbidden{}
}

/*GETGUIDAliasTypesForbidden handles this case with default header values.

Forbidden
*/
type GETGUIDAliasTypesForbidden struct {
	Payload *api_models.APIError
}

func (o *GETGUIDAliasTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesForbidden  %+v", 403, o.Payload)
}

func (o *GETGUIDAliasTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDAliasTypesNotFound creates a GETGUIDAliasTypesNotFound with default headers values
func NewGETGUIDAliasTypesNotFound() *GETGUIDAliasTypesNotFound {
	return &GETGUIDAliasTypesNotFound{}
}

/*GETGUIDAliasTypesNotFound handles this case with default header values.

Not Found
*/
type GETGUIDAliasTypesNotFound struct {
	Payload *api_models.APIError
}

func (o *GETGUIDAliasTypesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesNotFound  %+v", 404, o.Payload)
}

func (o *GETGUIDAliasTypesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDAliasTypesMethodNotAllowed creates a GETGUIDAliasTypesMethodNotAllowed with default headers values
func NewGETGUIDAliasTypesMethodNotAllowed() *GETGUIDAliasTypesMethodNotAllowed {
	return &GETGUIDAliasTypesMethodNotAllowed{}
}

/*GETGUIDAliasTypesMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETGUIDAliasTypesMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETGUIDAliasTypesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETGUIDAliasTypesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDAliasTypesNotAcceptable creates a GETGUIDAliasTypesNotAcceptable with default headers values
func NewGETGUIDAliasTypesNotAcceptable() *GETGUIDAliasTypesNotAcceptable {
	return &GETGUIDAliasTypesNotAcceptable{}
}

/*GETGUIDAliasTypesNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETGUIDAliasTypesNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETGUIDAliasTypesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETGUIDAliasTypesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDAliasTypesUnsupportedMediaType creates a GETGUIDAliasTypesUnsupportedMediaType with default headers values
func NewGETGUIDAliasTypesUnsupportedMediaType() *GETGUIDAliasTypesUnsupportedMediaType {
	return &GETGUIDAliasTypesUnsupportedMediaType{}
}

/*GETGUIDAliasTypesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETGUIDAliasTypesUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETGUIDAliasTypesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETGUIDAliasTypesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGUIDAliasTypesInternalServerError creates a GETGUIDAliasTypesInternalServerError with default headers values
func NewGETGUIDAliasTypesInternalServerError() *GETGUIDAliasTypesInternalServerError {
	return &GETGUIDAliasTypesInternalServerError{}
}

/*GETGUIDAliasTypesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETGUIDAliasTypesInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETGUIDAliasTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/search/guids/aliases/types][%d] getGuidAliasTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETGUIDAliasTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
