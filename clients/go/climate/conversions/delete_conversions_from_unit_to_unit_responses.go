// Code generated by go-swagger; DO NOT EDIT.

package conversions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// DeleteConversionsFromUnitToUnitReader is a Reader for the DeleteConversionsFromUnitToUnit structure.
type DeleteConversionsFromUnitToUnitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversionsFromUnitToUnitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConversionsFromUnitToUnitNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConversionsFromUnitToUnitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversionsFromUnitToUnitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversionsFromUnitToUnitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversionsFromUnitToUnitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteConversionsFromUnitToUnitMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversionsFromUnitToUnitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteConversionsFromUnitToUnitDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteConversionsFromUnitToUnitNoContent creates a DeleteConversionsFromUnitToUnitNoContent with default headers values
func NewDeleteConversionsFromUnitToUnitNoContent() *DeleteConversionsFromUnitToUnitNoContent {
	return &DeleteConversionsFromUnitToUnitNoContent{}
}

/* DeleteConversionsFromUnitToUnitNoContent describes a response with status code 204, with default header values.

Unit conversion deleted.

*/
type DeleteConversionsFromUnitToUnitNoContent struct {
}

// IsSuccess returns true when this delete conversions from unit to unit no content response has a 2xx status code
func (o *DeleteConversionsFromUnitToUnitNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete conversions from unit to unit no content response has a 3xx status code
func (o *DeleteConversionsFromUnitToUnitNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversions from unit to unit no content response has a 4xx status code
func (o *DeleteConversionsFromUnitToUnitNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversions from unit to unit no content response has a 5xx status code
func (o *DeleteConversionsFromUnitToUnitNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversions from unit to unit no content response a status code equal to that given
func (o *DeleteConversionsFromUnitToUnitNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteConversionsFromUnitToUnitNoContent) Error() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitNoContent ", 204)
}

func (o *DeleteConversionsFromUnitToUnitNoContent) String() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitNoContent ", 204)
}

func (o *DeleteConversionsFromUnitToUnitNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversionsFromUnitToUnitBadRequest creates a DeleteConversionsFromUnitToUnitBadRequest with default headers values
func NewDeleteConversionsFromUnitToUnitBadRequest() *DeleteConversionsFromUnitToUnitBadRequest {
	return &DeleteConversionsFromUnitToUnitBadRequest{}
}

/* DeleteConversionsFromUnitToUnitBadRequest describes a response with status code 400, with default header values.

Client error in request. Input did not pass validations. See error details.

*/
type DeleteConversionsFromUnitToUnitBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete conversions from unit to unit bad request response has a 2xx status code
func (o *DeleteConversionsFromUnitToUnitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversions from unit to unit bad request response has a 3xx status code
func (o *DeleteConversionsFromUnitToUnitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversions from unit to unit bad request response has a 4xx status code
func (o *DeleteConversionsFromUnitToUnitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversions from unit to unit bad request response has a 5xx status code
func (o *DeleteConversionsFromUnitToUnitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversions from unit to unit bad request response a status code equal to that given
func (o *DeleteConversionsFromUnitToUnitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteConversionsFromUnitToUnitBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversionsFromUnitToUnitBadRequest) String() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversionsFromUnitToUnitBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConversionsFromUnitToUnitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversionsFromUnitToUnitUnauthorized creates a DeleteConversionsFromUnitToUnitUnauthorized with default headers values
func NewDeleteConversionsFromUnitToUnitUnauthorized() *DeleteConversionsFromUnitToUnitUnauthorized {
	return &DeleteConversionsFromUnitToUnitUnauthorized{}
}

/* DeleteConversionsFromUnitToUnitUnauthorized describes a response with status code 401, with default header values.

Unauthorized access for a lack of authentication

*/
type DeleteConversionsFromUnitToUnitUnauthorized struct {
}

// IsSuccess returns true when this delete conversions from unit to unit unauthorized response has a 2xx status code
func (o *DeleteConversionsFromUnitToUnitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversions from unit to unit unauthorized response has a 3xx status code
func (o *DeleteConversionsFromUnitToUnitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversions from unit to unit unauthorized response has a 4xx status code
func (o *DeleteConversionsFromUnitToUnitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversions from unit to unit unauthorized response has a 5xx status code
func (o *DeleteConversionsFromUnitToUnitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversions from unit to unit unauthorized response a status code equal to that given
func (o *DeleteConversionsFromUnitToUnitUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteConversionsFromUnitToUnitUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitUnauthorized ", 401)
}

func (o *DeleteConversionsFromUnitToUnitUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitUnauthorized ", 401)
}

func (o *DeleteConversionsFromUnitToUnitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversionsFromUnitToUnitForbidden creates a DeleteConversionsFromUnitToUnitForbidden with default headers values
func NewDeleteConversionsFromUnitToUnitForbidden() *DeleteConversionsFromUnitToUnitForbidden {
	return &DeleteConversionsFromUnitToUnitForbidden{}
}

/* DeleteConversionsFromUnitToUnitForbidden describes a response with status code 403, with default header values.

Forbidden access for a lack of sufficient privileges

*/
type DeleteConversionsFromUnitToUnitForbidden struct {
}

// IsSuccess returns true when this delete conversions from unit to unit forbidden response has a 2xx status code
func (o *DeleteConversionsFromUnitToUnitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversions from unit to unit forbidden response has a 3xx status code
func (o *DeleteConversionsFromUnitToUnitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversions from unit to unit forbidden response has a 4xx status code
func (o *DeleteConversionsFromUnitToUnitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversions from unit to unit forbidden response has a 5xx status code
func (o *DeleteConversionsFromUnitToUnitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversions from unit to unit forbidden response a status code equal to that given
func (o *DeleteConversionsFromUnitToUnitForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteConversionsFromUnitToUnitForbidden) Error() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitForbidden ", 403)
}

func (o *DeleteConversionsFromUnitToUnitForbidden) String() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitForbidden ", 403)
}

func (o *DeleteConversionsFromUnitToUnitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversionsFromUnitToUnitNotFound creates a DeleteConversionsFromUnitToUnitNotFound with default headers values
func NewDeleteConversionsFromUnitToUnitNotFound() *DeleteConversionsFromUnitToUnitNotFound {
	return &DeleteConversionsFromUnitToUnitNotFound{}
}

/* DeleteConversionsFromUnitToUnitNotFound describes a response with status code 404, with default header values.

Resource not found. The object requested does not exist or is not visible to the user.

*/
type DeleteConversionsFromUnitToUnitNotFound struct {
}

// IsSuccess returns true when this delete conversions from unit to unit not found response has a 2xx status code
func (o *DeleteConversionsFromUnitToUnitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversions from unit to unit not found response has a 3xx status code
func (o *DeleteConversionsFromUnitToUnitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversions from unit to unit not found response has a 4xx status code
func (o *DeleteConversionsFromUnitToUnitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversions from unit to unit not found response has a 5xx status code
func (o *DeleteConversionsFromUnitToUnitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversions from unit to unit not found response a status code equal to that given
func (o *DeleteConversionsFromUnitToUnitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteConversionsFromUnitToUnitNotFound) Error() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitNotFound ", 404)
}

func (o *DeleteConversionsFromUnitToUnitNotFound) String() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitNotFound ", 404)
}

func (o *DeleteConversionsFromUnitToUnitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversionsFromUnitToUnitMethodNotAllowed creates a DeleteConversionsFromUnitToUnitMethodNotAllowed with default headers values
func NewDeleteConversionsFromUnitToUnitMethodNotAllowed() *DeleteConversionsFromUnitToUnitMethodNotAllowed {
	return &DeleteConversionsFromUnitToUnitMethodNotAllowed{}
}

/* DeleteConversionsFromUnitToUnitMethodNotAllowed describes a response with status code 405, with default header values.

MethodNotAllowed

*/
type DeleteConversionsFromUnitToUnitMethodNotAllowed struct {
}

// IsSuccess returns true when this delete conversions from unit to unit method not allowed response has a 2xx status code
func (o *DeleteConversionsFromUnitToUnitMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversions from unit to unit method not allowed response has a 3xx status code
func (o *DeleteConversionsFromUnitToUnitMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversions from unit to unit method not allowed response has a 4xx status code
func (o *DeleteConversionsFromUnitToUnitMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversions from unit to unit method not allowed response has a 5xx status code
func (o *DeleteConversionsFromUnitToUnitMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversions from unit to unit method not allowed response a status code equal to that given
func (o *DeleteConversionsFromUnitToUnitMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *DeleteConversionsFromUnitToUnitMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitMethodNotAllowed ", 405)
}

func (o *DeleteConversionsFromUnitToUnitMethodNotAllowed) String() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitMethodNotAllowed ", 405)
}

func (o *DeleteConversionsFromUnitToUnitMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversionsFromUnitToUnitInternalServerError creates a DeleteConversionsFromUnitToUnitInternalServerError with default headers values
func NewDeleteConversionsFromUnitToUnitInternalServerError() *DeleteConversionsFromUnitToUnitInternalServerError {
	return &DeleteConversionsFromUnitToUnitInternalServerError{}
}

/* DeleteConversionsFromUnitToUnitInternalServerError describes a response with status code 500, with default header values.

An internal error has occured. See error details.

*/
type DeleteConversionsFromUnitToUnitInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete conversions from unit to unit internal server error response has a 2xx status code
func (o *DeleteConversionsFromUnitToUnitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversions from unit to unit internal server error response has a 3xx status code
func (o *DeleteConversionsFromUnitToUnitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversions from unit to unit internal server error response has a 4xx status code
func (o *DeleteConversionsFromUnitToUnitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversions from unit to unit internal server error response has a 5xx status code
func (o *DeleteConversionsFromUnitToUnitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete conversions from unit to unit internal server error response a status code equal to that given
func (o *DeleteConversionsFromUnitToUnitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteConversionsFromUnitToUnitInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversionsFromUnitToUnitInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] deleteConversionsFromUnitToUnitInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversionsFromUnitToUnitInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConversionsFromUnitToUnitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversionsFromUnitToUnitDefault creates a DeleteConversionsFromUnitToUnitDefault with default headers values
func NewDeleteConversionsFromUnitToUnitDefault(code int) *DeleteConversionsFromUnitToUnitDefault {
	return &DeleteConversionsFromUnitToUnitDefault{
		_statusCode: code,
	}
}

/* DeleteConversionsFromUnitToUnitDefault describes a response with status code -1, with default header values.

Other error. See error details.

*/
type DeleteConversionsFromUnitToUnitDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the delete conversions from unit to unit default response
func (o *DeleteConversionsFromUnitToUnitDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete conversions from unit to unit default response has a 2xx status code
func (o *DeleteConversionsFromUnitToUnitDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete conversions from unit to unit default response has a 3xx status code
func (o *DeleteConversionsFromUnitToUnitDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete conversions from unit to unit default response has a 4xx status code
func (o *DeleteConversionsFromUnitToUnitDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete conversions from unit to unit default response has a 5xx status code
func (o *DeleteConversionsFromUnitToUnitDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete conversions from unit to unit default response a status code equal to that given
func (o *DeleteConversionsFromUnitToUnitDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteConversionsFromUnitToUnitDefault) Error() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] DeleteConversionsFromUnitToUnit default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteConversionsFromUnitToUnitDefault) String() string {
	return fmt.Sprintf("[DELETE /conversions/{fromUnit}/{toUnit}][%d] DeleteConversionsFromUnitToUnit default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteConversionsFromUnitToUnitDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConversionsFromUnitToUnitDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}