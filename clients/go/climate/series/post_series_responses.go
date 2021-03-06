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

// PostSeriesReader is a Reader for the PostSeries structure.
type PostSeriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSeriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSeriesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSeriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSeriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSeriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSeriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostSeriesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSeriesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSeriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostSeriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSeriesCreated creates a PostSeriesCreated with default headers values
func NewPostSeriesCreated() *PostSeriesCreated {
	return &PostSeriesCreated{}
}

/* PostSeriesCreated describes a response with status code 201, with default header values.

 Series successfully created.

Check the response headers to retrieve this resource.

*/
type PostSeriesCreated struct {

	/* The URI of the newly created resource.


	   Format: uri
	*/
	Location strfmt.URI

	/* The ID of the newly created resource.


	   Format: int64
	*/
	XID int64

	/* The URI of the newly created versioned resource, whenever applicable.


	   Format: uri
	*/
	XLocationVersion strfmt.URI

	/* The UUID of the newly created versioned resource, whenever applicable.


	   Format: uuid
	*/
	XVersionID strfmt.UUID
}

// IsSuccess returns true when this post series created response has a 2xx status code
func (o *PostSeriesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post series created response has a 3xx status code
func (o *PostSeriesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series created response has a 4xx status code
func (o *PostSeriesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post series created response has a 5xx status code
func (o *PostSeriesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post series created response a status code equal to that given
func (o *PostSeriesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostSeriesCreated) Error() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesCreated ", 201)
}

func (o *PostSeriesCreated) String() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesCreated ", 201)
}

func (o *PostSeriesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostSeriesBadRequest creates a PostSeriesBadRequest with default headers values
func NewPostSeriesBadRequest() *PostSeriesBadRequest {
	return &PostSeriesBadRequest{}
}

/* PostSeriesBadRequest describes a response with status code 400, with default header values.

Client error in request. Input did not pass validations. See error details.

*/
type PostSeriesBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this post series bad request response has a 2xx status code
func (o *PostSeriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series bad request response has a 3xx status code
func (o *PostSeriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series bad request response has a 4xx status code
func (o *PostSeriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series bad request response has a 5xx status code
func (o *PostSeriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post series bad request response a status code equal to that given
func (o *PostSeriesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostSeriesBadRequest) Error() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSeriesBadRequest) String() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSeriesBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSeriesUnauthorized creates a PostSeriesUnauthorized with default headers values
func NewPostSeriesUnauthorized() *PostSeriesUnauthorized {
	return &PostSeriesUnauthorized{}
}

/* PostSeriesUnauthorized describes a response with status code 401, with default header values.

Unauthorized access for a lack of authentication

*/
type PostSeriesUnauthorized struct {
}

// IsSuccess returns true when this post series unauthorized response has a 2xx status code
func (o *PostSeriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series unauthorized response has a 3xx status code
func (o *PostSeriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series unauthorized response has a 4xx status code
func (o *PostSeriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series unauthorized response has a 5xx status code
func (o *PostSeriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post series unauthorized response a status code equal to that given
func (o *PostSeriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostSeriesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesUnauthorized ", 401)
}

func (o *PostSeriesUnauthorized) String() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesUnauthorized ", 401)
}

func (o *PostSeriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesForbidden creates a PostSeriesForbidden with default headers values
func NewPostSeriesForbidden() *PostSeriesForbidden {
	return &PostSeriesForbidden{}
}

/* PostSeriesForbidden describes a response with status code 403, with default header values.

Forbidden access for a lack of sufficient privileges

*/
type PostSeriesForbidden struct {
}

// IsSuccess returns true when this post series forbidden response has a 2xx status code
func (o *PostSeriesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series forbidden response has a 3xx status code
func (o *PostSeriesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series forbidden response has a 4xx status code
func (o *PostSeriesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series forbidden response has a 5xx status code
func (o *PostSeriesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post series forbidden response a status code equal to that given
func (o *PostSeriesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostSeriesForbidden) Error() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesForbidden ", 403)
}

func (o *PostSeriesForbidden) String() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesForbidden ", 403)
}

func (o *PostSeriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesNotFound creates a PostSeriesNotFound with default headers values
func NewPostSeriesNotFound() *PostSeriesNotFound {
	return &PostSeriesNotFound{}
}

/* PostSeriesNotFound describes a response with status code 404, with default header values.

Resource not found. The object requested does not exist or is not visible to the user.

*/
type PostSeriesNotFound struct {
}

// IsSuccess returns true when this post series not found response has a 2xx status code
func (o *PostSeriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series not found response has a 3xx status code
func (o *PostSeriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series not found response has a 4xx status code
func (o *PostSeriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series not found response has a 5xx status code
func (o *PostSeriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post series not found response a status code equal to that given
func (o *PostSeriesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostSeriesNotFound) Error() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesNotFound ", 404)
}

func (o *PostSeriesNotFound) String() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesNotFound ", 404)
}

func (o *PostSeriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesMethodNotAllowed creates a PostSeriesMethodNotAllowed with default headers values
func NewPostSeriesMethodNotAllowed() *PostSeriesMethodNotAllowed {
	return &PostSeriesMethodNotAllowed{}
}

/* PostSeriesMethodNotAllowed describes a response with status code 405, with default header values.

MethodNotAllowed

*/
type PostSeriesMethodNotAllowed struct {
}

// IsSuccess returns true when this post series method not allowed response has a 2xx status code
func (o *PostSeriesMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series method not allowed response has a 3xx status code
func (o *PostSeriesMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series method not allowed response has a 4xx status code
func (o *PostSeriesMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series method not allowed response has a 5xx status code
func (o *PostSeriesMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this post series method not allowed response a status code equal to that given
func (o *PostSeriesMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PostSeriesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesMethodNotAllowed ", 405)
}

func (o *PostSeriesMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesMethodNotAllowed ", 405)
}

func (o *PostSeriesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesConflict creates a PostSeriesConflict with default headers values
func NewPostSeriesConflict() *PostSeriesConflict {
	return &PostSeriesConflict{}
}

/* PostSeriesConflict describes a response with status code 409, with default header values.

Resource already exists. An object creation was requested, but this object was already existing.

*/
type PostSeriesConflict struct {
}

// IsSuccess returns true when this post series conflict response has a 2xx status code
func (o *PostSeriesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series conflict response has a 3xx status code
func (o *PostSeriesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series conflict response has a 4xx status code
func (o *PostSeriesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post series conflict response has a 5xx status code
func (o *PostSeriesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post series conflict response a status code equal to that given
func (o *PostSeriesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostSeriesConflict) Error() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesConflict ", 409)
}

func (o *PostSeriesConflict) String() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesConflict ", 409)
}

func (o *PostSeriesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSeriesInternalServerError creates a PostSeriesInternalServerError with default headers values
func NewPostSeriesInternalServerError() *PostSeriesInternalServerError {
	return &PostSeriesInternalServerError{}
}

/* PostSeriesInternalServerError describes a response with status code 500, with default header values.

An internal error has occured. See error details.

*/
type PostSeriesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this post series internal server error response has a 2xx status code
func (o *PostSeriesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post series internal server error response has a 3xx status code
func (o *PostSeriesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post series internal server error response has a 4xx status code
func (o *PostSeriesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post series internal server error response has a 5xx status code
func (o *PostSeriesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post series internal server error response a status code equal to that given
func (o *PostSeriesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostSeriesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSeriesInternalServerError) String() string {
	return fmt.Sprintf("[POST /series][%d] postSeriesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSeriesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSeriesDefault creates a PostSeriesDefault with default headers values
func NewPostSeriesDefault(code int) *PostSeriesDefault {
	return &PostSeriesDefault{
		_statusCode: code,
	}
}

/* PostSeriesDefault describes a response with status code -1, with default header values.

Other error. See error details.

*/
type PostSeriesDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the post series default response
func (o *PostSeriesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post series default response has a 2xx status code
func (o *PostSeriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post series default response has a 3xx status code
func (o *PostSeriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post series default response has a 4xx status code
func (o *PostSeriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post series default response has a 5xx status code
func (o *PostSeriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post series default response a status code equal to that given
func (o *PostSeriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostSeriesDefault) Error() string {
	return fmt.Sprintf("[POST /series][%d] PostSeries default  %+v", o._statusCode, o.Payload)
}

func (o *PostSeriesDefault) String() string {
	return fmt.Sprintf("[POST /series][%d] PostSeries default  %+v", o._statusCode, o.Payload)
}

func (o *PostSeriesDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSeriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
