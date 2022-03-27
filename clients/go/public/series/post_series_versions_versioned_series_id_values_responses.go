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

// PostSeriesVersionsVersionedSeriesIDValuesReader is a Reader for the PostSeriesVersionsVersionedSeriesIDValues structure.
type PostSeriesVersionsVersionedSeriesIDValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSeriesVersionsVersionedSeriesIDValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostSeriesVersionsVersionedSeriesIDValuesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSeriesVersionsVersionedSeriesIDValuesCreated creates a PostSeriesVersionsVersionedSeriesIDValuesCreated with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesCreated() *PostSeriesVersionsVersionedSeriesIDValuesCreated {
	return &PostSeriesVersionsVersionedSeriesIDValuesCreated{}
}

/* PostSeriesVersionsVersionedSeriesIDValuesCreated describes a response with status code 201, with default header values.

 Time series values successfully created.

Check the response headers to retrieve this resource.

*/
type PostSeriesVersionsVersionedSeriesIDValuesCreated struct {

	/* The URI of the newly created resource.


	   Format: uri
	*/
	Location strfmt.URI

	/* The ID of the updated resource.


	   Format: int64
	*/
	XID int64
}

// IsSuccess returns true when this post series versions versioned series Id values created response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post series versions versioned series Id values created response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id values created response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post series versions versioned series Id values created response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id values created response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesCreated) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesCreated ", 201)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesCreated) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesCreated ", 201)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostSeriesVersionsVersionedSeriesIDValuesBadRequest creates a PostSeriesVersionsVersionedSeriesIDValuesBadRequest with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesBadRequest() *PostSeriesVersionsVersionedSeriesIDValuesBadRequest {
	return &PostSeriesVersionsVersionedSeriesIDValuesBadRequest{}
}

/* PostSeriesVersionsVersionedSeriesIDValuesBadRequest describes a response with status code 400, with default header values.

Client error in request. Input did not pass validations. See error details.

*/
type PostSeriesVersionsVersionedSeriesIDValuesBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this post series versions versioned series Id values bad request response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id values bad request response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id values bad request response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id values bad request response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id values bad request response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDValuesUnauthorized creates a PostSeriesVersionsVersionedSeriesIDValuesUnauthorized with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesUnauthorized() *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized {
	return &PostSeriesVersionsVersionedSeriesIDValuesUnauthorized{}
}

/* PostSeriesVersionsVersionedSeriesIDValuesUnauthorized describes a response with status code 401, with default header values.

Unauthorized access for a lack of authentication

*/
type PostSeriesVersionsVersionedSeriesIDValuesUnauthorized struct {
}

// IsSuccess returns true when this post series versions versioned series Id values unauthorized response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id values unauthorized response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id values unauthorized response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id values unauthorized response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id values unauthorized response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesUnauthorized ", 401)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesUnauthorized ", 401)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDValuesForbidden creates a PostSeriesVersionsVersionedSeriesIDValuesForbidden with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesForbidden() *PostSeriesVersionsVersionedSeriesIDValuesForbidden {
	return &PostSeriesVersionsVersionedSeriesIDValuesForbidden{}
}

/* PostSeriesVersionsVersionedSeriesIDValuesForbidden describes a response with status code 403, with default header values.

Forbidden access for a lack of sufficient privileges

*/
type PostSeriesVersionsVersionedSeriesIDValuesForbidden struct {
}

// IsSuccess returns true when this post series versions versioned series Id values forbidden response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id values forbidden response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id values forbidden response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id values forbidden response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id values forbidden response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesForbidden) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesForbidden ", 403)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesForbidden) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesForbidden ", 403)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDValuesNotFound creates a PostSeriesVersionsVersionedSeriesIDValuesNotFound with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesNotFound() *PostSeriesVersionsVersionedSeriesIDValuesNotFound {
	return &PostSeriesVersionsVersionedSeriesIDValuesNotFound{}
}

/* PostSeriesVersionsVersionedSeriesIDValuesNotFound describes a response with status code 404, with default header values.

Resource not found. The object requested does not exist or is not visible to the user.

*/
type PostSeriesVersionsVersionedSeriesIDValuesNotFound struct {
}

// IsSuccess returns true when this post series versions versioned series Id values not found response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id values not found response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id values not found response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id values not found response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id values not found response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesNotFound) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesNotFound ", 404)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesNotFound) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesNotFound ", 404)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed creates a PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed() *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed {
	return &PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed{}
}

/* PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed describes a response with status code 405, with default header values.

MethodNotAllowed

*/
type PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed struct {
}

// IsSuccess returns true when this post series versions versioned series Id values method not allowed response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id values method not allowed response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id values method not allowed response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id values method not allowed response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id values method not allowed response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesMethodNotAllowed ", 405)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesMethodNotAllowed ", 405)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDValuesConflict creates a PostSeriesVersionsVersionedSeriesIDValuesConflict with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesConflict() *PostSeriesVersionsVersionedSeriesIDValuesConflict {
	return &PostSeriesVersionsVersionedSeriesIDValuesConflict{}
}

/* PostSeriesVersionsVersionedSeriesIDValuesConflict describes a response with status code 409, with default header values.

Resource already exists. An object creation was requested, but this object was already existing.

*/
type PostSeriesVersionsVersionedSeriesIDValuesConflict struct {
}

// IsSuccess returns true when this post series versions versioned series Id values conflict response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id values conflict response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id values conflict response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series versions versioned series Id values conflict response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post series versions versioned series Id values conflict response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesConflict) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesConflict ", 409)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesConflict) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesConflict ", 409)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDValuesInternalServerError creates a PostSeriesVersionsVersionedSeriesIDValuesInternalServerError with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesInternalServerError() *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError {
	return &PostSeriesVersionsVersionedSeriesIDValuesInternalServerError{}
}

/* PostSeriesVersionsVersionedSeriesIDValuesInternalServerError describes a response with status code 500, with default header values.

An internal error has occured. See error details.

*/
type PostSeriesVersionsVersionedSeriesIDValuesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this post series versions versioned series Id values internal server error response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series versions versioned series Id values internal server error response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series versions versioned series Id values internal server error response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post series versions versioned series Id values internal server error response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post series versions versioned series Id values internal server error response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] postSeriesVersionsVersionedSeriesIdValuesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSeriesVersionsVersionedSeriesIDValuesDefault creates a PostSeriesVersionsVersionedSeriesIDValuesDefault with default headers values
func NewPostSeriesVersionsVersionedSeriesIDValuesDefault(code int) *PostSeriesVersionsVersionedSeriesIDValuesDefault {
	return &PostSeriesVersionsVersionedSeriesIDValuesDefault{
		_statusCode: code,
	}
}

/* PostSeriesVersionsVersionedSeriesIDValuesDefault describes a response with status code -1, with default header values.

Other error. See error details.

*/
type PostSeriesVersionsVersionedSeriesIDValuesDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the post series versions versioned series ID values default response
func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post series versions versioned series ID values default response has a 2xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post series versions versioned series ID values default response has a 3xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post series versions versioned series ID values default response has a 4xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post series versions versioned series ID values default response has a 5xx status code
func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post series versions versioned series ID values default response a status code equal to that given
func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) Error() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] PostSeriesVersionsVersionedSeriesIDValues default  %+v", o._statusCode, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) String() string {
	return fmt.Sprintf("[POST /series/versions/{versionedSeriesId}/values][%d] PostSeriesVersionsVersionedSeriesIDValues default  %+v", o._statusCode, o.Payload)
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesVersionsVersionedSeriesIDValuesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
