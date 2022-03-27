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

// DeleteSeriesVersionsVersionedSeriesIDNoContentCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDNoContent
const DeleteSeriesVersionsVersionedSeriesIDNoContentCode int = 204

/*DeleteSeriesVersionsVersionedSeriesIDNoContent Version of the time series successfully deleted.


swagger:response deleteSeriesVersionsVersionedSeriesIdNoContent
*/
type DeleteSeriesVersionsVersionedSeriesIDNoContent struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDNoContent creates DeleteSeriesVersionsVersionedSeriesIDNoContent with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDNoContent() *DeleteSeriesVersionsVersionedSeriesIDNoContent {

	return &DeleteSeriesVersionsVersionedSeriesIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDNoContent) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDBadRequestCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDBadRequest
const DeleteSeriesVersionsVersionedSeriesIDBadRequestCode int = 400

/*DeleteSeriesVersionsVersionedSeriesIDBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response deleteSeriesVersionsVersionedSeriesIdBadRequest
*/
type DeleteSeriesVersionsVersionedSeriesIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesVersionsVersionedSeriesIDBadRequest creates DeleteSeriesVersionsVersionedSeriesIDBadRequest with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDBadRequest() *DeleteSeriesVersionsVersionedSeriesIDBadRequest {

	return &DeleteSeriesVersionsVersionedSeriesIDBadRequest{}
}

// WithPayload adds the payload to the delete series versions versioned series Id bad request response
func (o *DeleteSeriesVersionsVersionedSeriesIDBadRequest) WithPayload(payload *models.APIError) *DeleteSeriesVersionsVersionedSeriesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series versions versioned series Id bad request response
func (o *DeleteSeriesVersionsVersionedSeriesIDBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesVersionsVersionedSeriesIDBadRequest) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDUnauthorizedCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDUnauthorized
const DeleteSeriesVersionsVersionedSeriesIDUnauthorizedCode int = 401

/*DeleteSeriesVersionsVersionedSeriesIDUnauthorized Unauthorized access for a lack of authentication


swagger:response deleteSeriesVersionsVersionedSeriesIdUnauthorized
*/
type DeleteSeriesVersionsVersionedSeriesIDUnauthorized struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDUnauthorized creates DeleteSeriesVersionsVersionedSeriesIDUnauthorized with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDUnauthorized() *DeleteSeriesVersionsVersionedSeriesIDUnauthorized {

	return &DeleteSeriesVersionsVersionedSeriesIDUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDUnauthorized) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDForbiddenCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDForbidden
const DeleteSeriesVersionsVersionedSeriesIDForbiddenCode int = 403

/*DeleteSeriesVersionsVersionedSeriesIDForbidden Forbidden access for a lack of sufficient privileges


swagger:response deleteSeriesVersionsVersionedSeriesIdForbidden
*/
type DeleteSeriesVersionsVersionedSeriesIDForbidden struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDForbidden creates DeleteSeriesVersionsVersionedSeriesIDForbidden with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDForbidden() *DeleteSeriesVersionsVersionedSeriesIDForbidden {

	return &DeleteSeriesVersionsVersionedSeriesIDForbidden{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDForbidden) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDNotFoundCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDNotFound
const DeleteSeriesVersionsVersionedSeriesIDNotFoundCode int = 404

/*DeleteSeriesVersionsVersionedSeriesIDNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response deleteSeriesVersionsVersionedSeriesIdNotFound
*/
type DeleteSeriesVersionsVersionedSeriesIDNotFound struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDNotFound creates DeleteSeriesVersionsVersionedSeriesIDNotFound with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDNotFound() *DeleteSeriesVersionsVersionedSeriesIDNotFound {

	return &DeleteSeriesVersionsVersionedSeriesIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDNotFound) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowedCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed
const DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowedCode int = 405

/*DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed MethodNotAllowed


swagger:response deleteSeriesVersionsVersionedSeriesIdMethodNotAllowed
*/
type DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed creates DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed() *DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed {

	return &DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDMethodNotAllowed) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDInternalServerErrorCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDInternalServerError
const DeleteSeriesVersionsVersionedSeriesIDInternalServerErrorCode int = 500

/*DeleteSeriesVersionsVersionedSeriesIDInternalServerError An internal error has occured. See error details.


swagger:response deleteSeriesVersionsVersionedSeriesIdInternalServerError
*/
type DeleteSeriesVersionsVersionedSeriesIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesVersionsVersionedSeriesIDInternalServerError creates DeleteSeriesVersionsVersionedSeriesIDInternalServerError with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDInternalServerError() *DeleteSeriesVersionsVersionedSeriesIDInternalServerError {

	return &DeleteSeriesVersionsVersionedSeriesIDInternalServerError{}
}

// WithPayload adds the payload to the delete series versions versioned series Id internal server error response
func (o *DeleteSeriesVersionsVersionedSeriesIDInternalServerError) WithPayload(payload *models.APIError) *DeleteSeriesVersionsVersionedSeriesIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series versions versioned series Id internal server error response
func (o *DeleteSeriesVersionsVersionedSeriesIDInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesVersionsVersionedSeriesIDInternalServerError) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

/*DeleteSeriesVersionsVersionedSeriesIDDefault Other error. See error details.


swagger:response deleteSeriesVersionsVersionedSeriesIdDefault
*/
type DeleteSeriesVersionsVersionedSeriesIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesVersionsVersionedSeriesIDDefault creates DeleteSeriesVersionsVersionedSeriesIDDefault with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDDefault(code int) *DeleteSeriesVersionsVersionedSeriesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSeriesVersionsVersionedSeriesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete series versions versioned series ID default response
func (o *DeleteSeriesVersionsVersionedSeriesIDDefault) WithStatusCode(code int) *DeleteSeriesVersionsVersionedSeriesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete series versions versioned series ID default response
func (o *DeleteSeriesVersionsVersionedSeriesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete series versions versioned series ID default response
func (o *DeleteSeriesVersionsVersionedSeriesIDDefault) WithPayload(payload *models.APIError) *DeleteSeriesVersionsVersionedSeriesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series versions versioned series ID default response
func (o *DeleteSeriesVersionsVersionedSeriesIDDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesVersionsVersionedSeriesIDDefault) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

type DeleteSeriesVersionsVersionedSeriesIDNotImplementedResponder struct {
	middleware.Responder
}

func (*DeleteSeriesVersionsVersionedSeriesIDNotImplementedResponder) DeleteSeriesVersionsVersionedSeriesIDResponder() {
}

func DeleteSeriesVersionsVersionedSeriesIDNotImplemented() DeleteSeriesVersionsVersionedSeriesIDResponder {
	return &DeleteSeriesVersionsVersionedSeriesIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.DeleteSeriesVersionsVersionedSeriesID has not yet been implemented",
		),
	}
}

type DeleteSeriesVersionsVersionedSeriesIDResponder interface {
	middleware.Responder
	DeleteSeriesVersionsVersionedSeriesIDResponder()
}
