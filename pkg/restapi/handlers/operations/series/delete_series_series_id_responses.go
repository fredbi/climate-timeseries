// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// DeleteSeriesSeriesIDNoContentCode is the HTTP code returned for type DeleteSeriesSeriesIDNoContent
const DeleteSeriesSeriesIDNoContentCode int = 204

/*DeleteSeriesSeriesIDNoContent Series successfully deleted.


swagger:response deleteSeriesSeriesIdNoContent
*/
type DeleteSeriesSeriesIDNoContent struct {
}

// NewDeleteSeriesSeriesIDNoContent creates DeleteSeriesSeriesIDNoContent with default headers values
func NewDeleteSeriesSeriesIDNoContent() *DeleteSeriesSeriesIDNoContent {

	return &DeleteSeriesSeriesIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteSeriesSeriesIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *DeleteSeriesSeriesIDNoContent) DeleteSeriesSeriesIDResponder() {}

// DeleteSeriesSeriesIDBadRequestCode is the HTTP code returned for type DeleteSeriesSeriesIDBadRequest
const DeleteSeriesSeriesIDBadRequestCode int = 400

/*DeleteSeriesSeriesIDBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response deleteSeriesSeriesIdBadRequest
*/
type DeleteSeriesSeriesIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesSeriesIDBadRequest creates DeleteSeriesSeriesIDBadRequest with default headers values
func NewDeleteSeriesSeriesIDBadRequest() *DeleteSeriesSeriesIDBadRequest {

	return &DeleteSeriesSeriesIDBadRequest{}
}

// WithPayload adds the payload to the delete series series Id bad request response
func (o *DeleteSeriesSeriesIDBadRequest) WithPayload(payload *models.APIError) *DeleteSeriesSeriesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series series Id bad request response
func (o *DeleteSeriesSeriesIDBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesSeriesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesSeriesIDBadRequest) DeleteSeriesSeriesIDResponder() {}

// DeleteSeriesSeriesIDUnauthorizedCode is the HTTP code returned for type DeleteSeriesSeriesIDUnauthorized
const DeleteSeriesSeriesIDUnauthorizedCode int = 401

/*DeleteSeriesSeriesIDUnauthorized Unauthorized access for a lack of authentication


swagger:response deleteSeriesSeriesIdUnauthorized
*/
type DeleteSeriesSeriesIDUnauthorized struct {
}

// NewDeleteSeriesSeriesIDUnauthorized creates DeleteSeriesSeriesIDUnauthorized with default headers values
func NewDeleteSeriesSeriesIDUnauthorized() *DeleteSeriesSeriesIDUnauthorized {

	return &DeleteSeriesSeriesIDUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteSeriesSeriesIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *DeleteSeriesSeriesIDUnauthorized) DeleteSeriesSeriesIDResponder() {}

// DeleteSeriesSeriesIDForbiddenCode is the HTTP code returned for type DeleteSeriesSeriesIDForbidden
const DeleteSeriesSeriesIDForbiddenCode int = 403

/*DeleteSeriesSeriesIDForbidden Forbidden access for a lack of sufficient privileges


swagger:response deleteSeriesSeriesIdForbidden
*/
type DeleteSeriesSeriesIDForbidden struct {
}

// NewDeleteSeriesSeriesIDForbidden creates DeleteSeriesSeriesIDForbidden with default headers values
func NewDeleteSeriesSeriesIDForbidden() *DeleteSeriesSeriesIDForbidden {

	return &DeleteSeriesSeriesIDForbidden{}
}

// WriteResponse to the client
func (o *DeleteSeriesSeriesIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *DeleteSeriesSeriesIDForbidden) DeleteSeriesSeriesIDResponder() {}

// DeleteSeriesSeriesIDNotFoundCode is the HTTP code returned for type DeleteSeriesSeriesIDNotFound
const DeleteSeriesSeriesIDNotFoundCode int = 404

/*DeleteSeriesSeriesIDNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response deleteSeriesSeriesIdNotFound
*/
type DeleteSeriesSeriesIDNotFound struct {
}

// NewDeleteSeriesSeriesIDNotFound creates DeleteSeriesSeriesIDNotFound with default headers values
func NewDeleteSeriesSeriesIDNotFound() *DeleteSeriesSeriesIDNotFound {

	return &DeleteSeriesSeriesIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteSeriesSeriesIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *DeleteSeriesSeriesIDNotFound) DeleteSeriesSeriesIDResponder() {}

// DeleteSeriesSeriesIDMethodNotAllowedCode is the HTTP code returned for type DeleteSeriesSeriesIDMethodNotAllowed
const DeleteSeriesSeriesIDMethodNotAllowedCode int = 405

/*DeleteSeriesSeriesIDMethodNotAllowed MethodNotAllowed


swagger:response deleteSeriesSeriesIdMethodNotAllowed
*/
type DeleteSeriesSeriesIDMethodNotAllowed struct {
}

// NewDeleteSeriesSeriesIDMethodNotAllowed creates DeleteSeriesSeriesIDMethodNotAllowed with default headers values
func NewDeleteSeriesSeriesIDMethodNotAllowed() *DeleteSeriesSeriesIDMethodNotAllowed {

	return &DeleteSeriesSeriesIDMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DeleteSeriesSeriesIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *DeleteSeriesSeriesIDMethodNotAllowed) DeleteSeriesSeriesIDResponder() {}

// DeleteSeriesSeriesIDInternalServerErrorCode is the HTTP code returned for type DeleteSeriesSeriesIDInternalServerError
const DeleteSeriesSeriesIDInternalServerErrorCode int = 500

/*DeleteSeriesSeriesIDInternalServerError An internal error has occured. See error details.


swagger:response deleteSeriesSeriesIdInternalServerError
*/
type DeleteSeriesSeriesIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesSeriesIDInternalServerError creates DeleteSeriesSeriesIDInternalServerError with default headers values
func NewDeleteSeriesSeriesIDInternalServerError() *DeleteSeriesSeriesIDInternalServerError {

	return &DeleteSeriesSeriesIDInternalServerError{}
}

// WithPayload adds the payload to the delete series series Id internal server error response
func (o *DeleteSeriesSeriesIDInternalServerError) WithPayload(payload *models.APIError) *DeleteSeriesSeriesIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series series Id internal server error response
func (o *DeleteSeriesSeriesIDInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesSeriesIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesSeriesIDInternalServerError) DeleteSeriesSeriesIDResponder() {}

/*DeleteSeriesSeriesIDDefault Other error. See error details.


swagger:response deleteSeriesSeriesIdDefault
*/
type DeleteSeriesSeriesIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesSeriesIDDefault creates DeleteSeriesSeriesIDDefault with default headers values
func NewDeleteSeriesSeriesIDDefault(code int) *DeleteSeriesSeriesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSeriesSeriesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete series series ID default response
func (o *DeleteSeriesSeriesIDDefault) WithStatusCode(code int) *DeleteSeriesSeriesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete series series ID default response
func (o *DeleteSeriesSeriesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete series series ID default response
func (o *DeleteSeriesSeriesIDDefault) WithPayload(payload *models.APIError) *DeleteSeriesSeriesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series series ID default response
func (o *DeleteSeriesSeriesIDDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesSeriesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesSeriesIDDefault) DeleteSeriesSeriesIDResponder() {}

type DeleteSeriesSeriesIDNotImplementedResponder struct {
	middleware.Responder
}

func (*DeleteSeriesSeriesIDNotImplementedResponder) DeleteSeriesSeriesIDResponder() {}

func DeleteSeriesSeriesIDNotImplemented() DeleteSeriesSeriesIDResponder {
	return &DeleteSeriesSeriesIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.DeleteSeriesSeriesID has not yet been implemented",
		),
	}
}

type DeleteSeriesSeriesIDResponder interface {
	middleware.Responder
	DeleteSeriesSeriesIDResponder()
}