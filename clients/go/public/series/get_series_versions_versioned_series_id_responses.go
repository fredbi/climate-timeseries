// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// GetSeriesVersionsVersionedSeriesIDReader is a Reader for the GetSeriesVersionsVersionedSeriesID structure.
type GetSeriesVersionsVersionedSeriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSeriesVersionsVersionedSeriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSeriesVersionsVersionedSeriesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSeriesVersionsVersionedSeriesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSeriesVersionsVersionedSeriesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSeriesVersionsVersionedSeriesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSeriesVersionsVersionedSeriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetSeriesVersionsVersionedSeriesIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSeriesVersionsVersionedSeriesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSeriesVersionsVersionedSeriesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSeriesVersionsVersionedSeriesIDOK creates a GetSeriesVersionsVersionedSeriesIDOK with default headers values
func NewGetSeriesVersionsVersionedSeriesIDOK() *GetSeriesVersionsVersionedSeriesIDOK {
	return &GetSeriesVersionsVersionedSeriesIDOK{}
}

/* GetSeriesVersionsVersionedSeriesIDOK describes a response with status code 200, with default header values.

Description of a given version of a time series.

*/
type GetSeriesVersionsVersionedSeriesIDOK struct {
	Payload *models.VersionedSeries
}

// IsSuccess returns true when this get series versions versioned series Id o k response has a 2xx status code
func (o *GetSeriesVersionsVersionedSeriesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get series versions versioned series Id o k response has a 3xx status code
func (o *GetSeriesVersionsVersionedSeriesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get series versions versioned series Id o k response has a 4xx status code
func (o *GetSeriesVersionsVersionedSeriesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get series versions versioned series Id o k response has a 5xx status code
func (o *GetSeriesVersionsVersionedSeriesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get series versions versioned series Id o k response a status code equal to that given
func (o *GetSeriesVersionsVersionedSeriesIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSeriesVersionsVersionedSeriesIDOK) Error() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdOK  %+v", 200, o.Payload)
}

func (o *GetSeriesVersionsVersionedSeriesIDOK) String() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdOK  %+v", 200, o.Payload)
}

func (o *GetSeriesVersionsVersionedSeriesIDOK) GetPayload() *models.VersionedSeries {
	return o.Payload
}

func (o *GetSeriesVersionsVersionedSeriesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedSeries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesVersionsVersionedSeriesIDBadRequest creates a GetSeriesVersionsVersionedSeriesIDBadRequest with default headers values
func NewGetSeriesVersionsVersionedSeriesIDBadRequest() *GetSeriesVersionsVersionedSeriesIDBadRequest {
	return &GetSeriesVersionsVersionedSeriesIDBadRequest{}
}

/* GetSeriesVersionsVersionedSeriesIDBadRequest describes a response with status code 400, with default header values.

Client error in request. Input did not pass validations. See error details.

*/
type GetSeriesVersionsVersionedSeriesIDBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this get series versions versioned series Id bad request response has a 2xx status code
func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get series versions versioned series Id bad request response has a 3xx status code
func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get series versions versioned series Id bad request response has a 4xx status code
func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get series versions versioned series Id bad request response has a 5xx status code
func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get series versions versioned series Id bad request response a status code equal to that given
func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) String() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesVersionsVersionedSeriesIDUnauthorized creates a GetSeriesVersionsVersionedSeriesIDUnauthorized with default headers values
func NewGetSeriesVersionsVersionedSeriesIDUnauthorized() *GetSeriesVersionsVersionedSeriesIDUnauthorized {
	return &GetSeriesVersionsVersionedSeriesIDUnauthorized{}
}

/* GetSeriesVersionsVersionedSeriesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized access for a lack of authentication

*/
type GetSeriesVersionsVersionedSeriesIDUnauthorized struct {
}

// IsSuccess returns true when this get series versions versioned series Id unauthorized response has a 2xx status code
func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get series versions versioned series Id unauthorized response has a 3xx status code
func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get series versions versioned series Id unauthorized response has a 4xx status code
func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get series versions versioned series Id unauthorized response has a 5xx status code
func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get series versions versioned series Id unauthorized response a status code equal to that given
func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdUnauthorized ", 401)
}

func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdUnauthorized ", 401)
}

func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeriesVersionsVersionedSeriesIDForbidden creates a GetSeriesVersionsVersionedSeriesIDForbidden with default headers values
func NewGetSeriesVersionsVersionedSeriesIDForbidden() *GetSeriesVersionsVersionedSeriesIDForbidden {
	return &GetSeriesVersionsVersionedSeriesIDForbidden{}
}

/* GetSeriesVersionsVersionedSeriesIDForbidden describes a response with status code 403, with default header values.

Forbidden access for a lack of sufficient privileges

*/
type GetSeriesVersionsVersionedSeriesIDForbidden struct {
}

// IsSuccess returns true when this get series versions versioned series Id forbidden response has a 2xx status code
func (o *GetSeriesVersionsVersionedSeriesIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get series versions versioned series Id forbidden response has a 3xx status code
func (o *GetSeriesVersionsVersionedSeriesIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get series versions versioned series Id forbidden response has a 4xx status code
func (o *GetSeriesVersionsVersionedSeriesIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get series versions versioned series Id forbidden response has a 5xx status code
func (o *GetSeriesVersionsVersionedSeriesIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get series versions versioned series Id forbidden response a status code equal to that given
func (o *GetSeriesVersionsVersionedSeriesIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSeriesVersionsVersionedSeriesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdForbidden ", 403)
}

func (o *GetSeriesVersionsVersionedSeriesIDForbidden) String() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdForbidden ", 403)
}

func (o *GetSeriesVersionsVersionedSeriesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeriesVersionsVersionedSeriesIDNotFound creates a GetSeriesVersionsVersionedSeriesIDNotFound with default headers values
func NewGetSeriesVersionsVersionedSeriesIDNotFound() *GetSeriesVersionsVersionedSeriesIDNotFound {
	return &GetSeriesVersionsVersionedSeriesIDNotFound{}
}

/* GetSeriesVersionsVersionedSeriesIDNotFound describes a response with status code 404, with default header values.

Resource not found. The object requested does not exist or is not visible to the user.

*/
type GetSeriesVersionsVersionedSeriesIDNotFound struct {
}

// IsSuccess returns true when this get series versions versioned series Id not found response has a 2xx status code
func (o *GetSeriesVersionsVersionedSeriesIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get series versions versioned series Id not found response has a 3xx status code
func (o *GetSeriesVersionsVersionedSeriesIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get series versions versioned series Id not found response has a 4xx status code
func (o *GetSeriesVersionsVersionedSeriesIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get series versions versioned series Id not found response has a 5xx status code
func (o *GetSeriesVersionsVersionedSeriesIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get series versions versioned series Id not found response a status code equal to that given
func (o *GetSeriesVersionsVersionedSeriesIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSeriesVersionsVersionedSeriesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdNotFound ", 404)
}

func (o *GetSeriesVersionsVersionedSeriesIDNotFound) String() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdNotFound ", 404)
}

func (o *GetSeriesVersionsVersionedSeriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeriesVersionsVersionedSeriesIDMethodNotAllowed creates a GetSeriesVersionsVersionedSeriesIDMethodNotAllowed with default headers values
func NewGetSeriesVersionsVersionedSeriesIDMethodNotAllowed() *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed {
	return &GetSeriesVersionsVersionedSeriesIDMethodNotAllowed{}
}

/* GetSeriesVersionsVersionedSeriesIDMethodNotAllowed describes a response with status code 405, with default header values.

MethodNotAllowed

*/
type GetSeriesVersionsVersionedSeriesIDMethodNotAllowed struct {
}

// IsSuccess returns true when this get series versions versioned series Id method not allowed response has a 2xx status code
func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get series versions versioned series Id method not allowed response has a 3xx status code
func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get series versions versioned series Id method not allowed response has a 4xx status code
func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this get series versions versioned series Id method not allowed response has a 5xx status code
func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this get series versions versioned series Id method not allowed response a status code equal to that given
func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdMethodNotAllowed ", 405)
}

func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdMethodNotAllowed ", 405)
}

func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeriesVersionsVersionedSeriesIDInternalServerError creates a GetSeriesVersionsVersionedSeriesIDInternalServerError with default headers values
func NewGetSeriesVersionsVersionedSeriesIDInternalServerError() *GetSeriesVersionsVersionedSeriesIDInternalServerError {
	return &GetSeriesVersionsVersionedSeriesIDInternalServerError{}
}

/* GetSeriesVersionsVersionedSeriesIDInternalServerError describes a response with status code 500, with default header values.

An internal error has occured. See error details.

*/
type GetSeriesVersionsVersionedSeriesIDInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this get series versions versioned series Id internal server error response has a 2xx status code
func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get series versions versioned series Id internal server error response has a 3xx status code
func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get series versions versioned series Id internal server error response has a 4xx status code
func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get series versions versioned series Id internal server error response has a 5xx status code
func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get series versions versioned series Id internal server error response a status code equal to that given
func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] getSeriesVersionsVersionedSeriesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesVersionsVersionedSeriesIDDefault creates a GetSeriesVersionsVersionedSeriesIDDefault with default headers values
func NewGetSeriesVersionsVersionedSeriesIDDefault(code int) *GetSeriesVersionsVersionedSeriesIDDefault {
	return &GetSeriesVersionsVersionedSeriesIDDefault{
		_statusCode: code,
	}
}

/* GetSeriesVersionsVersionedSeriesIDDefault describes a response with status code -1, with default header values.

Other error. See error details.

*/
type GetSeriesVersionsVersionedSeriesIDDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the get series versions versioned series ID default response
func (o *GetSeriesVersionsVersionedSeriesIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get series versions versioned series ID default response has a 2xx status code
func (o *GetSeriesVersionsVersionedSeriesIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get series versions versioned series ID default response has a 3xx status code
func (o *GetSeriesVersionsVersionedSeriesIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get series versions versioned series ID default response has a 4xx status code
func (o *GetSeriesVersionsVersionedSeriesIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get series versions versioned series ID default response has a 5xx status code
func (o *GetSeriesVersionsVersionedSeriesIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get series versions versioned series ID default response a status code equal to that given
func (o *GetSeriesVersionsVersionedSeriesIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetSeriesVersionsVersionedSeriesIDDefault) Error() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] GetSeriesVersionsVersionedSeriesID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSeriesVersionsVersionedSeriesIDDefault) String() string {
	return fmt.Sprintf("[GET /series/versions/{versionedSeriesId}][%d] GetSeriesVersionsVersionedSeriesID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSeriesVersionsVersionedSeriesIDDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSeriesVersionsVersionedSeriesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
