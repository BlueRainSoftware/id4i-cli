// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// ListCountriesReader is a Reader for the ListCountries structure.
type ListCountriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCountriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCountriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewListCountriesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListCountriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListCountriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListCountriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListCountriesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewListCountriesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewListCountriesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListCountriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCountriesOK creates a ListCountriesOK with default headers values
func NewListCountriesOK() *ListCountriesOK {
	return &ListCountriesOK{}
}

/*ListCountriesOK handles this case with default header values.

OK
*/
type ListCountriesOK struct {
	Payload *api_models.PaginatedResponseOfCountry
}

func (o *ListCountriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesOK  %+v", 200, o.Payload)
}

func (o *ListCountriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedResponseOfCountry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCountriesAccepted creates a ListCountriesAccepted with default headers values
func NewListCountriesAccepted() *ListCountriesAccepted {
	return &ListCountriesAccepted{}
}

/*ListCountriesAccepted handles this case with default header values.

Accepted
*/
type ListCountriesAccepted struct {
}

func (o *ListCountriesAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesAccepted ", 202)
}

func (o *ListCountriesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCountriesUnauthorized creates a ListCountriesUnauthorized with default headers values
func NewListCountriesUnauthorized() *ListCountriesUnauthorized {
	return &ListCountriesUnauthorized{}
}

/*ListCountriesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCountriesUnauthorized struct {
	Payload *api_models.APIError
}

func (o *ListCountriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCountriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCountriesForbidden creates a ListCountriesForbidden with default headers values
func NewListCountriesForbidden() *ListCountriesForbidden {
	return &ListCountriesForbidden{}
}

/*ListCountriesForbidden handles this case with default header values.

Forbidden
*/
type ListCountriesForbidden struct {
	Payload *api_models.APIError
}

func (o *ListCountriesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesForbidden  %+v", 403, o.Payload)
}

func (o *ListCountriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCountriesNotFound creates a ListCountriesNotFound with default headers values
func NewListCountriesNotFound() *ListCountriesNotFound {
	return &ListCountriesNotFound{}
}

/*ListCountriesNotFound handles this case with default header values.

Not Found
*/
type ListCountriesNotFound struct {
	Payload *api_models.APIError
}

func (o *ListCountriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesNotFound  %+v", 404, o.Payload)
}

func (o *ListCountriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCountriesMethodNotAllowed creates a ListCountriesMethodNotAllowed with default headers values
func NewListCountriesMethodNotAllowed() *ListCountriesMethodNotAllowed {
	return &ListCountriesMethodNotAllowed{}
}

/*ListCountriesMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type ListCountriesMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *ListCountriesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListCountriesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCountriesNotAcceptable creates a ListCountriesNotAcceptable with default headers values
func NewListCountriesNotAcceptable() *ListCountriesNotAcceptable {
	return &ListCountriesNotAcceptable{}
}

/*ListCountriesNotAcceptable handles this case with default header values.

Not Acceptable
*/
type ListCountriesNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *ListCountriesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesNotAcceptable  %+v", 406, o.Payload)
}

func (o *ListCountriesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCountriesUnsupportedMediaType creates a ListCountriesUnsupportedMediaType with default headers values
func NewListCountriesUnsupportedMediaType() *ListCountriesUnsupportedMediaType {
	return &ListCountriesUnsupportedMediaType{}
}

/*ListCountriesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type ListCountriesUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *ListCountriesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListCountriesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCountriesInternalServerError creates a ListCountriesInternalServerError with default headers values
func NewListCountriesInternalServerError() *ListCountriesInternalServerError {
	return &ListCountriesInternalServerError{}
}

/*ListCountriesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListCountriesInternalServerError struct {
	Payload *api_models.APIError
}

func (o *ListCountriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/countries][%d] listCountriesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCountriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
