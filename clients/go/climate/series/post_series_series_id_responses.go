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

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// PostSeriesSeriesIDReader is a Reader for the PostSeriesSeriesID structure.
type PostSeriesSeriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSeriesSeriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSeriesSeriesIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSeriesSeriesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSeriesSeriesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSeriesSeriesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSeriesSeriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostSeriesSeriesIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSeriesSeriesIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSeriesSeriesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostSeriesSeriesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSeriesSeriesIDCreated creates a PostSeriesSeriesIDCreated with default headers values
func NewPostSeriesSeriesIDCreated() *PostSeriesSeriesIDCreated {
	return &PostSeriesSeriesIDCreated{}
}

/* PostSeriesSeriesIDCreated describes a response with status code 201, with default header values.

 Series version successfully created.

Check the response headers to retrieve this resource.

*/
type PostSeriesSeriesIDCreated struct {

	/* The URI of the newly created versioned resource.


	   Format: uri
	*/
	Location strfmt.URI

	/* The URI of the newly created versioned resource.


	   Format: uri
	*/
	XLocationVersion strfmt.URI

	/* The UUID of the newly created versioned resource.


	   Format: uuid
	*/
	XVersionID strfmt.UUID
}

// IsSuccess returns true when this post series series Id created response has a 2xx status code
func (o *PostSeriesSeriesIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post series series Id created response has a 3xx status code
func (o *PostSeriesSeriesIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series series Id created response has a 4xx status code
func (o *PostSeriesSeriesIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post series series Id created response has a 5xx status code
func (o *PostSeriesSeriesIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post series series Id created response a status code equal to that given
func (o *PostSeriesSeriesIDCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostSeriesSeriesIDCreated) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdCreated ", 201)
}

func (o *PostSeriesSeriesIDCreated) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdCreated ", 201)
}

func (o *PostSeriesSeriesIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		vallocation, err := formats.Parse("uri", hdrLocation)
		if err != nil {
			return errors.InvalidType("Location", "header", "strfmt.URI", hdrLocation)
		}
		o.Location = *(vallocation.(*strfmt.URI))
	}

	// hydrates response header X-LocationVersion
	hdrXLocationVersion := response.GetHeader("X-LocationVersion")

	if hdrXLocationVersion != "" {
		valxLocationVersion, err := formats.Parse("uri", hdrXLocationVersion)
		if err != nil {
			return errors.InvalidType("X-LocationVersion", "header", "strfmt.URI", hdrXLocationVersion)
		}
		o.XLocationVersion = *(valxLocationVersion.(*strfmt.URI))
	}

	// hydrates response header X-VersionID
	hdrXVersionID := response.GetHeader("X-VersionID")

	if hdrXVersionID != "" {
		valxVersionId, err := formats.Parse("uuid", hdrXVersionID)
		if err != nil {
			return errors.InvalidType("X-VersionID", "header", "strfmt.UUID", hdrXVersionID)
		}
		o.XVersionID = *(valxVersionId.(*strfmt.UUID))
	}

	return nil
}

// NewPostSeriesSeriesIDBadRequest creates a PostSeriesSeriesIDBadRequest with default headers values
func NewPostSeriesSeriesIDBadRequest() *PostSeriesSeriesIDBadRequest {
	return &PostSeriesSeriesIDBadRequest{}
}

/* PostSeriesSeriesIDBadRequest describes a response with status code 400, with default header values.

Client error in request. Input did not pass validations. See error details.

*/
type PostSeriesSeriesIDBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this post series series Id bad request response has a 2xx status code
func (o *PostSeriesSeriesIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series series Id bad request response has a 3xx status code
func (o *PostSeriesSeriesIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series series Id bad request response has a 4xx status code
func (o *PostSeriesSeriesIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series series Id bad request response has a 5xx status code
func (o *PostSeriesSeriesIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post series series Id bad request response a status code equal to that given
func (o *PostSeriesSeriesIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostSeriesSeriesIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PostSeriesSeriesIDBadRequest) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PostSeriesSeriesIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesSeriesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSeriesSeriesIDUnauthorized creates a PostSeriesSeriesIDUnauthorized with default headers values
func NewPostSeriesSeriesIDUnauthorized() *PostSeriesSeriesIDUnauthorized {
	return &PostSeriesSeriesIDUnauthorized{}
}

/* PostSeriesSeriesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized access for a lack of authentication

*/
type PostSeriesSeriesIDUnauthorized struct {
}

// IsSuccess returns true when this post series series Id unauthorized response has a 2xx status code
func (o *PostSeriesSeriesIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series series Id unauthorized response has a 3xx status code
func (o *PostSeriesSeriesIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series series Id unauthorized response has a 4xx status code
func (o *PostSeriesSeriesIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series series Id unauthorized response has a 5xx status code
func (o *PostSeriesSeriesIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post series series Id unauthorized response a status code equal to that given
func (o *PostSeriesSeriesIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostSeriesSeriesIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdUnauthorized ", 401)
}

func (o *PostSeriesSeriesIDUnauthorized) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdUnauthorized ", 401)
}

func (o *PostSeriesSeriesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesSeriesIDForbidden creates a PostSeriesSeriesIDForbidden with default headers values
func NewPostSeriesSeriesIDForbidden() *PostSeriesSeriesIDForbidden {
	return &PostSeriesSeriesIDForbidden{}
}

/* PostSeriesSeriesIDForbidden describes a response with status code 403, with default header values.

Forbidden access for a lack of sufficient privileges

*/
type PostSeriesSeriesIDForbidden struct {
}

// IsSuccess returns true when this post series series Id forbidden response has a 2xx status code
func (o *PostSeriesSeriesIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series series Id forbidden response has a 3xx status code
func (o *PostSeriesSeriesIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series series Id forbidden response has a 4xx status code
func (o *PostSeriesSeriesIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series series Id forbidden response has a 5xx status code
func (o *PostSeriesSeriesIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post series series Id forbidden response a status code equal to that given
func (o *PostSeriesSeriesIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostSeriesSeriesIDForbidden) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdForbidden ", 403)
}

func (o *PostSeriesSeriesIDForbidden) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdForbidden ", 403)
}

func (o *PostSeriesSeriesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesSeriesIDNotFound creates a PostSeriesSeriesIDNotFound with default headers values
func NewPostSeriesSeriesIDNotFound() *PostSeriesSeriesIDNotFound {
	return &PostSeriesSeriesIDNotFound{}
}

/* PostSeriesSeriesIDNotFound describes a response with status code 404, with default header values.

Resource not found. The object requested does not exist or is not visible to the user.

*/
type PostSeriesSeriesIDNotFound struct {
}

// IsSuccess returns true when this post series series Id not found response has a 2xx status code
func (o *PostSeriesSeriesIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series series Id not found response has a 3xx status code
func (o *PostSeriesSeriesIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series series Id not found response has a 4xx status code
func (o *PostSeriesSeriesIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series series Id not found response has a 5xx status code
func (o *PostSeriesSeriesIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post series series Id not found response a status code equal to that given
func (o *PostSeriesSeriesIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostSeriesSeriesIDNotFound) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdNotFound ", 404)
}

func (o *PostSeriesSeriesIDNotFound) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdNotFound ", 404)
}

func (o *PostSeriesSeriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesSeriesIDMethodNotAllowed creates a PostSeriesSeriesIDMethodNotAllowed with default headers values
func NewPostSeriesSeriesIDMethodNotAllowed() *PostSeriesSeriesIDMethodNotAllowed {
	return &PostSeriesSeriesIDMethodNotAllowed{}
}

/* PostSeriesSeriesIDMethodNotAllowed describes a response with status code 405, with default header values.

MethodNotAllowed

*/
type PostSeriesSeriesIDMethodNotAllowed struct {
}

// IsSuccess returns true when this post series series Id method not allowed response has a 2xx status code
func (o *PostSeriesSeriesIDMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series series Id method not allowed response has a 3xx status code
func (o *PostSeriesSeriesIDMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series series Id method not allowed response has a 4xx status code
func (o *PostSeriesSeriesIDMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series series Id method not allowed response has a 5xx status code
func (o *PostSeriesSeriesIDMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this post series series Id method not allowed response a status code equal to that given
func (o *PostSeriesSeriesIDMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PostSeriesSeriesIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdMethodNotAllowed ", 405)
}

func (o *PostSeriesSeriesIDMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdMethodNotAllowed ", 405)
}

func (o *PostSeriesSeriesIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesSeriesIDConflict creates a PostSeriesSeriesIDConflict with default headers values
func NewPostSeriesSeriesIDConflict() *PostSeriesSeriesIDConflict {
	return &PostSeriesSeriesIDConflict{}
}

/* PostSeriesSeriesIDConflict describes a response with status code 409, with default header values.

Resource already exists. An object creation was requested, but this object was already existing.

*/
type PostSeriesSeriesIDConflict struct {
}

// IsSuccess returns true when this post series series Id conflict response has a 2xx status code
func (o *PostSeriesSeriesIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series series Id conflict response has a 3xx status code
func (o *PostSeriesSeriesIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series series Id conflict response has a 4xx status code
func (o *PostSeriesSeriesIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series series Id conflict response has a 5xx status code
func (o *PostSeriesSeriesIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post series series Id conflict response a status code equal to that given
func (o *PostSeriesSeriesIDConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostSeriesSeriesIDConflict) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdConflict ", 409)
}

func (o *PostSeriesSeriesIDConflict) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdConflict ", 409)
}

func (o *PostSeriesSeriesIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesSeriesIDInternalServerError creates a PostSeriesSeriesIDInternalServerError with default headers values
func NewPostSeriesSeriesIDInternalServerError() *PostSeriesSeriesIDInternalServerError {
	return &PostSeriesSeriesIDInternalServerError{}
}

/* PostSeriesSeriesIDInternalServerError describes a response with status code 500, with default header values.

An internal error has occured. See error details.

*/
type PostSeriesSeriesIDInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this post series series Id internal server error response has a 2xx status code
func (o *PostSeriesSeriesIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series series Id internal server error response has a 3xx status code
func (o *PostSeriesSeriesIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series series Id internal server error response has a 4xx status code
func (o *PostSeriesSeriesIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post series series Id internal server error response has a 5xx status code
func (o *PostSeriesSeriesIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post series series Id internal server error response a status code equal to that given
func (o *PostSeriesSeriesIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostSeriesSeriesIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSeriesSeriesIDInternalServerError) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] postSeriesSeriesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSeriesSeriesIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesSeriesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSeriesSeriesIDDefault creates a PostSeriesSeriesIDDefault with default headers values
func NewPostSeriesSeriesIDDefault(code int) *PostSeriesSeriesIDDefault {
	return &PostSeriesSeriesIDDefault{
		_statusCode: code,
	}
}

/* PostSeriesSeriesIDDefault describes a response with status code -1, with default header values.

Other error. See error details.

*/
type PostSeriesSeriesIDDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the post series series ID default response
func (o *PostSeriesSeriesIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post series series ID default response has a 2xx status code
func (o *PostSeriesSeriesIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post series series ID default response has a 3xx status code
func (o *PostSeriesSeriesIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post series series ID default response has a 4xx status code
func (o *PostSeriesSeriesIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post series series ID default response has a 5xx status code
func (o *PostSeriesSeriesIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post series series ID default response a status code equal to that given
func (o *PostSeriesSeriesIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostSeriesSeriesIDDefault) Error() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] PostSeriesSeriesID default  %+v", o._statusCode, o.Payload)
}

func (o *PostSeriesSeriesIDDefault) String() string {
	return fmt.Sprintf("[POST /series/{seriesId}][%d] PostSeriesSeriesID default  %+v", o._statusCode, o.Payload)
}

func (o *PostSeriesSeriesIDDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesSeriesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
