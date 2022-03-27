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

// DeleteSeriesVersionsVersionedSeriesIDValuesNoContentCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDValuesNoContent
const DeleteSeriesVersionsVersionedSeriesIDValuesNoContentCode int = 204

/*DeleteSeriesVersionsVersionedSeriesIDValuesNoContent Values of the versioned time series successfully deleted.


swagger:response deleteSeriesVersionsVersionedSeriesIdValuesNoContent
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesNoContent struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesNoContent creates DeleteSeriesVersionsVersionedSeriesIDValuesNoContent with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDValuesNoContent() *DeleteSeriesVersionsVersionedSeriesIDValuesNoContent {

	return &DeleteSeriesVersionsVersionedSeriesIDValuesNoContent{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValuesNoContent) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDValuesBadRequestCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest
const DeleteSeriesVersionsVersionedSeriesIDValuesBadRequestCode int = 400

/*DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response deleteSeriesVersionsVersionedSeriesIdValuesBadRequest
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesBadRequest creates DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDValuesBadRequest() *DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest {

	return &DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest{}
}

// WithPayload adds the payload to the delete series versions versioned series Id values bad request response
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest) WithPayload(payload *models.APIError) *DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series versions versioned series Id values bad request response
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValuesBadRequest) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorizedCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized
const DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorizedCode int = 401

/*DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized Unauthorized access for a lack of authentication


swagger:response deleteSeriesVersionsVersionedSeriesIdValuesUnauthorized
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized creates DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized() *DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized {

	return &DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValuesUnauthorized) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDValuesForbiddenCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDValuesForbidden
const DeleteSeriesVersionsVersionedSeriesIDValuesForbiddenCode int = 403

/*DeleteSeriesVersionsVersionedSeriesIDValuesForbidden Forbidden access for a lack of sufficient privileges


swagger:response deleteSeriesVersionsVersionedSeriesIdValuesForbidden
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesForbidden struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesForbidden creates DeleteSeriesVersionsVersionedSeriesIDValuesForbidden with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDValuesForbidden() *DeleteSeriesVersionsVersionedSeriesIDValuesForbidden {

	return &DeleteSeriesVersionsVersionedSeriesIDValuesForbidden{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValuesForbidden) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDValuesNotFoundCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDValuesNotFound
const DeleteSeriesVersionsVersionedSeriesIDValuesNotFoundCode int = 404

/*DeleteSeriesVersionsVersionedSeriesIDValuesNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response deleteSeriesVersionsVersionedSeriesIdValuesNotFound
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesNotFound struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesNotFound creates DeleteSeriesVersionsVersionedSeriesIDValuesNotFound with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDValuesNotFound() *DeleteSeriesVersionsVersionedSeriesIDValuesNotFound {

	return &DeleteSeriesVersionsVersionedSeriesIDValuesNotFound{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValuesNotFound) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowedCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed
const DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowedCode int = 405

/*DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed MethodNotAllowed


swagger:response deleteSeriesVersionsVersionedSeriesIdValuesMethodNotAllowed
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed struct {
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed creates DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed() *DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed {

	return &DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerErrorCode is the HTTP code returned for type DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError
const DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerErrorCode int = 500

/*DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError An internal error has occured. See error details.


swagger:response deleteSeriesVersionsVersionedSeriesIdValuesInternalServerError
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError creates DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError() *DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError {

	return &DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError{}
}

// WithPayload adds the payload to the delete series versions versioned series Id values internal server error response
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError) WithPayload(payload *models.APIError) *DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series versions versioned series Id values internal server error response
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValuesInternalServerError) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

/*DeleteSeriesVersionsVersionedSeriesIDValuesDefault Other error. See error details.


swagger:response deleteSeriesVersionsVersionedSeriesIdValuesDefault
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesDefault creates DeleteSeriesVersionsVersionedSeriesIDValuesDefault with default headers values
func NewDeleteSeriesVersionsVersionedSeriesIDValuesDefault(code int) *DeleteSeriesVersionsVersionedSeriesIDValuesDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSeriesVersionsVersionedSeriesIDValuesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete series versions versioned series ID values default response
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesDefault) WithStatusCode(code int) *DeleteSeriesVersionsVersionedSeriesIDValuesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete series versions versioned series ID values default response
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete series versions versioned series ID values default response
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesDefault) WithPayload(payload *models.APIError) *DeleteSeriesVersionsVersionedSeriesIDValuesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete series versions versioned series ID values default response
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValuesDefault) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

type DeleteSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder struct {
	middleware.Responder
}

func (*DeleteSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder) DeleteSeriesVersionsVersionedSeriesIDValuesResponder() {
}

func DeleteSeriesVersionsVersionedSeriesIDValuesNotImplemented() DeleteSeriesVersionsVersionedSeriesIDValuesResponder {
	return &DeleteSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.DeleteSeriesVersionsVersionedSeriesIDValues has not yet been implemented",
		),
	}
}

type DeleteSeriesVersionsVersionedSeriesIDValuesResponder interface {
	middleware.Responder
	DeleteSeriesVersionsVersionedSeriesIDValuesResponder()
}
