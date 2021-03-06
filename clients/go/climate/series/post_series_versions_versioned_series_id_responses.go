// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// PostSeriesVersionsVersionedSeriesIDReader is a Reader for the PostSeriesVersionsVersionedSeriesID structure.
type PostSeriesVersionsVersionedSeriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSeriesVersionsVersionedSeriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSeriesVersionsVersionedSeriesIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSeriesVersionsVersionedSeriesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSeriesVersionsVersionedSeriesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSeriesVersionsVersionedSeriesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSeriesVersionsVersionedSeriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostSeriesVersionsVersionedSeriesIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSeriesVersionsVersionedSeriesIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSeriesVersionsVersionedSeriesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostSeriesVersionsVersionedSeriesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSeriesVersionsVersionedSeriesIDCreated creates a PostSeriesVersionsVersionedSeriesIDCreated with default headers values
func NewPostSeriesVersionsVersionedSeriesIDCreated() *PostSeriesVersionsVersionedSeriesIDCreated {
	return &PostSeriesVersionsVersionedSeriesIDCreated{}
}

/* PostSeriesVersionsVersionedSeriesIDCreated describes a response with status code 201, with default header values.

 Time series version successfully created.

Check the response headers to retrieve this resource.

*/
type PostSeriesVersionsVersionedSeriesIDCreated struct {

	/* The URI of the newly created resource.


	   Format: uri
	*/
	Location strfmt.URI

	/* The ID of the updated resource.


	   Format: int64
	*/
	XID int64
}

// IsSuccess returns true when this post series versions versioned series Id created response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post series versions versioned series Id created response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id created response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post series versions versioned series Id created response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id created response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostSeriesVersionsVersionedSeriesIDCreated) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdCreated ", 201)
}

func (o *PostSeriesVersionsVersionedSeriesIDCreated) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdCreated ", 201)
}

func (o *PostSeriesVersionsVersionedSeriesIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		vallocation, err := formats.Parse("uri", hdrLocation)
		if err != nil {
			return errors.InvalidType("Location", "header", "strfmt.URI", hdrLocation)
		}
		o.Location = *(vallocation.(*strfmt.URI))
	}

	// hydrates response header X-ID
	hdrXID := response.GetHeader("X-ID")

	if hdrXID != "" {
		valxId, err := swag.ConvertInt64(hdrXID)
		if err != nil {
			return errors.InvalidType("X-ID", "header", "int64", hdrXID)
		}
		o.XID = valxId
	}

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDBadRequest creates a PostSeriesVersionsVersionedSeriesIDBadRequest with default headers values
func NewPostSeriesVersionsVersionedSeriesIDBadRequest() *PostSeriesVersionsVersionedSeriesIDBadRequest {
	return &PostSeriesVersionsVersionedSeriesIDBadRequest{}
}

/* PostSeriesVersionsVersionedSeriesIDBadRequest describes a response with status code 400, with default header values.

Client error in request. Input did not pass validations. See error details.

*/
type PostSeriesVersionsVersionedSeriesIDBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this post series versions versioned series Id bad request response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id bad request response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id bad request response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id bad request response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id bad request response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDUnauthorized creates a PostSeriesVersionsVersionedSeriesIDUnauthorized with default headers values
func NewPostSeriesVersionsVersionedSeriesIDUnauthorized() *PostSeriesVersionsVersionedSeriesIDUnauthorized {
	return &PostSeriesVersionsVersionedSeriesIDUnauthorized{}
}

/* PostSeriesVersionsVersionedSeriesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized access for a lack of authentication

*/
type PostSeriesVersionsVersionedSeriesIDUnauthorized struct {
}

// IsSuccess returns true when this post series versions versioned series Id unauthorized response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id unauthorized response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id unauthorized response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id unauthorized response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id unauthorized response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdUnauthorized ", 401)
}

func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdUnauthorized ", 401)
}

func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDForbidden creates a PostSeriesVersionsVersionedSeriesIDForbidden with default headers values
func NewPostSeriesVersionsVersionedSeriesIDForbidden() *PostSeriesVersionsVersionedSeriesIDForbidden {
	return &PostSeriesVersionsVersionedSeriesIDForbidden{}
}

/* PostSeriesVersionsVersionedSeriesIDForbidden describes a response with status code 403, with default header values.

Forbidden access for a lack of sufficient privileges

*/
type PostSeriesVersionsVersionedSeriesIDForbidden struct {
}

// IsSuccess returns true when this post series versions versioned series Id forbidden response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id forbidden response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id forbidden response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id forbidden response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id forbidden response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostSeriesVersionsVersionedSeriesIDForbidden) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdForbidden ", 403)
}

func (o *PostSeriesVersionsVersionedSeriesIDForbidden) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdForbidden ", 403)
}

func (o *PostSeriesVersionsVersionedSeriesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDNotFound creates a PostSeriesVersionsVersionedSeriesIDNotFound with default headers values
func NewPostSeriesVersionsVersionedSeriesIDNotFound() *PostSeriesVersionsVersionedSeriesIDNotFound {
	return &PostSeriesVersionsVersionedSeriesIDNotFound{}
}

/* PostSeriesVersionsVersionedSeriesIDNotFound describes a response with status code 404, with default header values.

Resource not found. The object requested does not exist or is not visible to the user.

*/
type PostSeriesVersionsVersionedSeriesIDNotFound struct {
}

// IsSuccess returns true when this post series versions versioned series Id not found response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id not found response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id not found response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id not found response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id not found response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostSeriesVersionsVersionedSeriesIDNotFound) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdNotFound ", 404)
}

func (o *PostSeriesVersionsVersionedSeriesIDNotFound) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdNotFound ", 404)
}

func (o *PostSeriesVersionsVersionedSeriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDMethodNotAllowed creates a PostSeriesVersionsVersionedSeriesIDMethodNotAllowed with default headers values
func NewPostSeriesVersionsVersionedSeriesIDMethodNotAllowed() *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed {
	return &PostSeriesVersionsVersionedSeriesIDMethodNotAllowed{}
}

/* PostSeriesVersionsVersionedSeriesIDMethodNotAllowed describes a response with status code 405, with default header values.

MethodNotAllowed

*/
type PostSeriesVersionsVersionedSeriesIDMethodNotAllowed struct {
}

// IsSuccess returns true when this post series versions versioned series Id method not allowed response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id method not allowed response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id method not allowed response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id method not allowed response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id method not allowed response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdMethodNotAllowed ", 405)
}

func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdMethodNotAllowed ", 405)
}

func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDConflict creates a PostSeriesVersionsVersionedSeriesIDConflict with default headers values
func NewPostSeriesVersionsVersionedSeriesIDConflict() *PostSeriesVersionsVersionedSeriesIDConflict {
	return &PostSeriesVersionsVersionedSeriesIDConflict{}
}

/* PostSeriesVersionsVersionedSeriesIDConflict describes a response with status code 409, with default header values.

Resource already exists. An object creation was requested, but this object was already existing.

*/
type PostSeriesVersionsVersionedSeriesIDConflict struct {
}

// IsSuccess returns true when this post series versions versioned series Id conflict response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id conflict response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id conflict response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id conflict response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id conflict response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostSeriesVersionsVersionedSeriesIDConflict) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdConflict ", 409)
}

func (o *PostSeriesVersionsVersionedSeriesIDConflict) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdConflict ", 409)
}

func (o *PostSeriesVersionsVersionedSeriesIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDInternalServerError creates a PostSeriesVersionsVersionedSeriesIDInternalServerError with default headers values
func NewPostSeriesVersionsVersionedSeriesIDInternalServerError() *PostSeriesVersionsVersionedSeriesIDInternalServerError {
	return &PostSeriesVersionsVersionedSeriesIDInternalServerError{}
}

/* PostSeriesVersionsVersionedSeriesIDInternalServerError describes a response with status code 500, with default header values.

An internal error has occured. See error details.

*/
type PostSeriesVersionsVersionedSeriesIDInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this post series versions versioned series Id internal server error response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id internal server error response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id internal server error response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post series versions versioned series Id internal server error response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post series versions versioned series Id internal server error response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] postSeriesVersionsVersionedSeriesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDDefault creates a PostSeriesVersionsVersionedSeriesIDDefault with default headers values
func NewPostSeriesVersionsVersionedSeriesIDDefault(code int) *PostSeriesVersionsVersionedSeriesIDDefault {
	return &PostSeriesVersionsVersionedSeriesIDDefault{
		_statusCode: code,
	}
}

/* PostSeriesVersionsVersionedSeriesIDDefault describes a response with status code -1, with default header values.

Other error. See error details.

*/
type PostSeriesVersionsVersionedSeriesIDDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the post series versions versioned series ID default response
func (o *PostSeriesVersionsVersionedSeriesIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post series versions versioned series ID default response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post series versions versioned series ID default response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post series versions versioned series ID default response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post series versions versioned series ID default response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post series versions versioned series ID default response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostSeriesVersionsVersionedSeriesIDDefault) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] PostSeriesVersionsVersionedSeriesID default  %+v", o._statusCode, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDDefault) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}][%d] PostSeriesVersionsVersionedSeriesID default  %+v", o._statusCode, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesVersionsVersionedSeriesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
