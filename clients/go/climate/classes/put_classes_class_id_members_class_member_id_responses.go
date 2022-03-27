// Code generated by go-swagger; DO NOT EDIT.

package classes

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

// PutClassesClassIDMembersClassMemberIDReader is a Reader for the PutClassesClassIDMembersClassMemberID structure.
type PutClassesClassIDMembersClassMemberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClassesClassIDMembersClassMemberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutClassesClassIDMembersClassMemberIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutClassesClassIDMembersClassMemberIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutClassesClassIDMembersClassMemberIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutClassesClassIDMembersClassMemberIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutClassesClassIDMembersClassMemberIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutClassesClassIDMembersClassMemberIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutClassesClassIDMembersClassMemberIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutClassesClassIDMembersClassMemberIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutClassesClassIDMembersClassMemberIDNoContent creates a PutClassesClassIDMembersClassMemberIDNoContent with default headers values
func NewPutClassesClassIDMembersClassMemberIDNoContent() *PutClassesClassIDMembersClassMemberIDNoContent {
	return &PutClassesClassIDMembersClassMemberIDNoContent{}
}

/* PutClassesClassIDMembersClassMemberIDNoContent describes a response with status code 204, with default header values.

Class members updated.

*/
type PutClassesClassIDMembersClassMemberIDNoContent struct {

	/* The URI of the updated resource.


	   Format: uri
	*/
	Location strfmt.URI

	/* The ID of the updated resource.


	   Format: int64
	*/
	XID int64
}

// IsSuccess returns true when this put classes class Id members class member Id no content response has a 2xx status code
func (o *PutClassesClassIDMembersClassMemberIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put classes class Id members class member Id no content response has a 3xx status code
func (o *PutClassesClassIDMembersClassMemberIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put classes class Id members class member Id no content response has a 4xx status code
func (o *PutClassesClassIDMembersClassMemberIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put classes class Id members class member Id no content response has a 5xx status code
func (o *PutClassesClassIDMembersClassMemberIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put classes class Id members class member Id no content response a status code equal to that given
func (o *PutClassesClassIDMembersClassMemberIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PutClassesClassIDMembersClassMemberIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdNoContent ", 204)
}

func (o *PutClassesClassIDMembersClassMemberIDNoContent) String() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdNoContent ", 204)
}

func (o *PutClassesClassIDMembersClassMemberIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPutClassesClassIDMembersClassMemberIDBadRequest creates a PutClassesClassIDMembersClassMemberIDBadRequest with default headers values
func NewPutClassesClassIDMembersClassMemberIDBadRequest() *PutClassesClassIDMembersClassMemberIDBadRequest {
	return &PutClassesClassIDMembersClassMemberIDBadRequest{}
}

/* PutClassesClassIDMembersClassMemberIDBadRequest describes a response with status code 400, with default header values.

Client error in request. Input did not pass validations. See error details.

*/
type PutClassesClassIDMembersClassMemberIDBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this put classes class Id members class member Id bad request response has a 2xx status code
func (o *PutClassesClassIDMembersClassMemberIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put classes class Id members class member Id bad request response has a 3xx status code
func (o *PutClassesClassIDMembersClassMemberIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put classes class Id members class member Id bad request response has a 4xx status code
func (o *PutClassesClassIDMembersClassMemberIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put classes class Id members class member Id bad request response has a 5xx status code
func (o *PutClassesClassIDMembersClassMemberIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put classes class Id members class member Id bad request response a status code equal to that given
func (o *PutClassesClassIDMembersClassMemberIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutClassesClassIDMembersClassMemberIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutClassesClassIDMembersClassMemberIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutClassesClassIDMembersClassMemberIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PutClassesClassIDMembersClassMemberIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClassesClassIDMembersClassMemberIDUnauthorized creates a PutClassesClassIDMembersClassMemberIDUnauthorized with default headers values
func NewPutClassesClassIDMembersClassMemberIDUnauthorized() *PutClassesClassIDMembersClassMemberIDUnauthorized {
	return &PutClassesClassIDMembersClassMemberIDUnauthorized{}
}

/* PutClassesClassIDMembersClassMemberIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized access for a lack of authentication

*/
type PutClassesClassIDMembersClassMemberIDUnauthorized struct {
}

// IsSuccess returns true when this put classes class Id members class member Id unauthorized response has a 2xx status code
func (o *PutClassesClassIDMembersClassMemberIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put classes class Id members class member Id unauthorized response has a 3xx status code
func (o *PutClassesClassIDMembersClassMemberIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put classes class Id members class member Id unauthorized response has a 4xx status code
func (o *PutClassesClassIDMembersClassMemberIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put classes class Id members class member Id unauthorized response has a 5xx status code
func (o *PutClassesClassIDMembersClassMemberIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put classes class Id members class member Id unauthorized response a status code equal to that given
func (o *PutClassesClassIDMembersClassMemberIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutClassesClassIDMembersClassMemberIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdUnauthorized ", 401)
}

func (o *PutClassesClassIDMembersClassMemberIDUnauthorized) String() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdUnauthorized ", 401)
}

func (o *PutClassesClassIDMembersClassMemberIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClassesClassIDMembersClassMemberIDForbidden creates a PutClassesClassIDMembersClassMemberIDForbidden with default headers values
func NewPutClassesClassIDMembersClassMemberIDForbidden() *PutClassesClassIDMembersClassMemberIDForbidden {
	return &PutClassesClassIDMembersClassMemberIDForbidden{}
}

/* PutClassesClassIDMembersClassMemberIDForbidden describes a response with status code 403, with default header values.

Forbidden access for a lack of sufficient privileges

*/
type PutClassesClassIDMembersClassMemberIDForbidden struct {
}

// IsSuccess returns true when this put classes class Id members class member Id forbidden response has a 2xx status code
func (o *PutClassesClassIDMembersClassMemberIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put classes class Id members class member Id forbidden response has a 3xx status code
func (o *PutClassesClassIDMembersClassMemberIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put classes class Id members class member Id forbidden response has a 4xx status code
func (o *PutClassesClassIDMembersClassMemberIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put classes class Id members class member Id forbidden response has a 5xx status code
func (o *PutClassesClassIDMembersClassMemberIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put classes class Id members class member Id forbidden response a status code equal to that given
func (o *PutClassesClassIDMembersClassMemberIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutClassesClassIDMembersClassMemberIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdForbidden ", 403)
}

func (o *PutClassesClassIDMembersClassMemberIDForbidden) String() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdForbidden ", 403)
}

func (o *PutClassesClassIDMembersClassMemberIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClassesClassIDMembersClassMemberIDNotFound creates a PutClassesClassIDMembersClassMemberIDNotFound with default headers values
func NewPutClassesClassIDMembersClassMemberIDNotFound() *PutClassesClassIDMembersClassMemberIDNotFound {
	return &PutClassesClassIDMembersClassMemberIDNotFound{}
}

/* PutClassesClassIDMembersClassMemberIDNotFound describes a response with status code 404, with default header values.

Resource not found. The object requested does not exist or is not visible to the user.

*/
type PutClassesClassIDMembersClassMemberIDNotFound struct {
}

// IsSuccess returns true when this put classes class Id members class member Id not found response has a 2xx status code
func (o *PutClassesClassIDMembersClassMemberIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put classes class Id members class member Id not found response has a 3xx status code
func (o *PutClassesClassIDMembersClassMemberIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put classes class Id members class member Id not found response has a 4xx status code
func (o *PutClassesClassIDMembersClassMemberIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put classes class Id members class member Id not found response has a 5xx status code
func (o *PutClassesClassIDMembersClassMemberIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put classes class Id members class member Id not found response a status code equal to that given
func (o *PutClassesClassIDMembersClassMemberIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutClassesClassIDMembersClassMemberIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdNotFound ", 404)
}

func (o *PutClassesClassIDMembersClassMemberIDNotFound) String() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdNotFound ", 404)
}

func (o *PutClassesClassIDMembersClassMemberIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClassesClassIDMembersClassMemberIDMethodNotAllowed creates a PutClassesClassIDMembersClassMemberIDMethodNotAllowed with default headers values
func NewPutClassesClassIDMembersClassMemberIDMethodNotAllowed() *PutClassesClassIDMembersClassMemberIDMethodNotAllowed {
	return &PutClassesClassIDMembersClassMemberIDMethodNotAllowed{}
}

/* PutClassesClassIDMembersClassMemberIDMethodNotAllowed describes a response with status code 405, with default header values.

MethodNotAllowed

*/
type PutClassesClassIDMembersClassMemberIDMethodNotAllowed struct {
}

// IsSuccess returns true when this put classes class Id members class member Id method not allowed response has a 2xx status code
func (o *PutClassesClassIDMembersClassMemberIDMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put classes class Id members class member Id method not allowed response has a 3xx status code
func (o *PutClassesClassIDMembersClassMemberIDMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put classes class Id members class member Id method not allowed response has a 4xx status code
func (o *PutClassesClassIDMembersClassMemberIDMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this put classes class Id members class member Id method not allowed response has a 5xx status code
func (o *PutClassesClassIDMembersClassMemberIDMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this put classes class Id members class member Id method not allowed response a status code equal to that given
func (o *PutClassesClassIDMembersClassMemberIDMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PutClassesClassIDMembersClassMemberIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdMethodNotAllowed ", 405)
}

func (o *PutClassesClassIDMembersClassMemberIDMethodNotAllowed) String() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdMethodNotAllowed ", 405)
}

func (o *PutClassesClassIDMembersClassMemberIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClassesClassIDMembersClassMemberIDInternalServerError creates a PutClassesClassIDMembersClassMemberIDInternalServerError with default headers values
func NewPutClassesClassIDMembersClassMemberIDInternalServerError() *PutClassesClassIDMembersClassMemberIDInternalServerError {
	return &PutClassesClassIDMembersClassMemberIDInternalServerError{}
}

/* PutClassesClassIDMembersClassMemberIDInternalServerError describes a response with status code 500, with default header values.

An internal error has occured. See error details.

*/
type PutClassesClassIDMembersClassMemberIDInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this put classes class Id members class member Id internal server error response has a 2xx status code
func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put classes class Id members class member Id internal server error response has a 3xx status code
func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put classes class Id members class member Id internal server error response has a 4xx status code
func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put classes class Id members class member Id internal server error response has a 5xx status code
func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put classes class Id members class member Id internal server error response a status code equal to that given
func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] putClassesClassIdMembersClassMemberIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PutClassesClassIDMembersClassMemberIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClassesClassIDMembersClassMemberIDDefault creates a PutClassesClassIDMembersClassMemberIDDefault with default headers values
func NewPutClassesClassIDMembersClassMemberIDDefault(code int) *PutClassesClassIDMembersClassMemberIDDefault {
	return &PutClassesClassIDMembersClassMemberIDDefault{
		_statusCode: code,
	}
}

/* PutClassesClassIDMembersClassMemberIDDefault describes a response with status code -1, with default header values.

Other error. See error details.

*/
type PutClassesClassIDMembersClassMemberIDDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the put classes class ID members class member ID default response
func (o *PutClassesClassIDMembersClassMemberIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this put classes class ID members class member ID default response has a 2xx status code
func (o *PutClassesClassIDMembersClassMemberIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put classes class ID members class member ID default response has a 3xx status code
func (o *PutClassesClassIDMembersClassMemberIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put classes class ID members class member ID default response has a 4xx status code
func (o *PutClassesClassIDMembersClassMemberIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put classes class ID members class member ID default response has a 5xx status code
func (o *PutClassesClassIDMembersClassMemberIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put classes class ID members class member ID default response a status code equal to that given
func (o *PutClassesClassIDMembersClassMemberIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PutClassesClassIDMembersClassMemberIDDefault) Error() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] PutClassesClassIDMembersClassMemberID default  %+v", o._statusCode, o.Payload)
}

func (o *PutClassesClassIDMembersClassMemberIDDefault) String() string {
	return fmt.Sprintf("[PUT /classes/{classId}/members/{classMemberId}][%d] PutClassesClassIDMembersClassMemberID default  %+v", o._statusCode, o.Payload)
}

func (o *PutClassesClassIDMembersClassMemberIDDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PutClassesClassIDMembersClassMemberIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}