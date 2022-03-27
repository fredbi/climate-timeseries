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

// GetSeriesVersionsVersionedSeriesIDValuesOKCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDValuesOK
const GetSeriesVersionsVersionedSeriesIDValuesOKCode int = 200

/*GetSeriesVersionsVersionedSeriesIDValuesOK Values of a given version of a time series.


swagger:response getSeriesVersionsVersionedSeriesIdValuesOK
*/
type GetSeriesVersionsVersionedSeriesIDValuesOK struct {

	/*
	  In: Body
	*/
	Payload models.TsValues `json:"body,omitempty"`
}

// NewGetSeriesVersionsVersionedSeriesIDValuesOK creates GetSeriesVersionsVersionedSeriesIDValuesOK with default headers values
func NewGetSeriesVersionsVersionedSeriesIDValuesOK() *GetSeriesVersionsVersionedSeriesIDValuesOK {

	return &GetSeriesVersionsVersionedSeriesIDValuesOK{}
}

// WithPayload adds the payload to the get series versions versioned series Id values o k response
func (o *GetSeriesVersionsVersionedSeriesIDValuesOK) WithPayload(payload models.TsValues) *GetSeriesVersionsVersionedSeriesIDValuesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series versions versioned series Id values o k response
func (o *GetSeriesVersionsVersionedSeriesIDValuesOK) SetPayload(payload models.TsValues) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDValuesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.TsValues{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *GetSeriesVersionsVersionedSeriesIDValuesOK) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// GetSeriesVersionsVersionedSeriesIDValuesBadRequestCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDValuesBadRequest
const GetSeriesVersionsVersionedSeriesIDValuesBadRequestCode int = 400

/*GetSeriesVersionsVersionedSeriesIDValuesBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response getSeriesVersionsVersionedSeriesIdValuesBadRequest
*/
type GetSeriesVersionsVersionedSeriesIDValuesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesVersionsVersionedSeriesIDValuesBadRequest creates GetSeriesVersionsVersionedSeriesIDValuesBadRequest with default headers values
func NewGetSeriesVersionsVersionedSeriesIDValuesBadRequest() *GetSeriesVersionsVersionedSeriesIDValuesBadRequest {

	return &GetSeriesVersionsVersionedSeriesIDValuesBadRequest{}
}

// WithPayload adds the payload to the get series versions versioned series Id values bad request response
func (o *GetSeriesVersionsVersionedSeriesIDValuesBadRequest) WithPayload(payload *models.APIError) *GetSeriesVersionsVersionedSeriesIDValuesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series versions versioned series Id values bad request response
func (o *GetSeriesVersionsVersionedSeriesIDValuesBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDValuesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesVersionsVersionedSeriesIDValuesBadRequest) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// GetSeriesVersionsVersionedSeriesIDValuesUnauthorizedCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDValuesUnauthorized
const GetSeriesVersionsVersionedSeriesIDValuesUnauthorizedCode int = 401

/*GetSeriesVersionsVersionedSeriesIDValuesUnauthorized Unauthorized access for a lack of authentication


swagger:response getSeriesVersionsVersionedSeriesIdValuesUnauthorized
*/
type GetSeriesVersionsVersionedSeriesIDValuesUnauthorized struct {
}

// NewGetSeriesVersionsVersionedSeriesIDValuesUnauthorized creates GetSeriesVersionsVersionedSeriesIDValuesUnauthorized with default headers values
func NewGetSeriesVersionsVersionedSeriesIDValuesUnauthorized() *GetSeriesVersionsVersionedSeriesIDValuesUnauthorized {

	return &GetSeriesVersionsVersionedSeriesIDValuesUnauthorized{}
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDValuesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *GetSeriesVersionsVersionedSeriesIDValuesUnauthorized) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// GetSeriesVersionsVersionedSeriesIDValuesForbiddenCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDValuesForbidden
const GetSeriesVersionsVersionedSeriesIDValuesForbiddenCode int = 403

/*GetSeriesVersionsVersionedSeriesIDValuesForbidden Forbidden access for a lack of sufficient privileges


swagger:response getSeriesVersionsVersionedSeriesIdValuesForbidden
*/
type GetSeriesVersionsVersionedSeriesIDValuesForbidden struct {
}

// NewGetSeriesVersionsVersionedSeriesIDValuesForbidden creates GetSeriesVersionsVersionedSeriesIDValuesForbidden with default headers values
func NewGetSeriesVersionsVersionedSeriesIDValuesForbidden() *GetSeriesVersionsVersionedSeriesIDValuesForbidden {

	return &GetSeriesVersionsVersionedSeriesIDValuesForbidden{}
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDValuesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *GetSeriesVersionsVersionedSeriesIDValuesForbidden) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// GetSeriesVersionsVersionedSeriesIDValuesNotFoundCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDValuesNotFound
const GetSeriesVersionsVersionedSeriesIDValuesNotFoundCode int = 404

/*GetSeriesVersionsVersionedSeriesIDValuesNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response getSeriesVersionsVersionedSeriesIdValuesNotFound
*/
type GetSeriesVersionsVersionedSeriesIDValuesNotFound struct {
}

// NewGetSeriesVersionsVersionedSeriesIDValuesNotFound creates GetSeriesVersionsVersionedSeriesIDValuesNotFound with default headers values
func NewGetSeriesVersionsVersionedSeriesIDValuesNotFound() *GetSeriesVersionsVersionedSeriesIDValuesNotFound {

	return &GetSeriesVersionsVersionedSeriesIDValuesNotFound{}
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDValuesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *GetSeriesVersionsVersionedSeriesIDValuesNotFound) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowedCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed
const GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowedCode int = 405

/*GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed MethodNotAllowed


swagger:response getSeriesVersionsVersionedSeriesIdValuesMethodNotAllowed
*/
type GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed struct {
}

// NewGetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed creates GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed with default headers values
func NewGetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed() *GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed {

	return &GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed{}
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *GetSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// GetSeriesVersionsVersionedSeriesIDValuesInternalServerErrorCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDValuesInternalServerError
const GetSeriesVersionsVersionedSeriesIDValuesInternalServerErrorCode int = 500

/*GetSeriesVersionsVersionedSeriesIDValuesInternalServerError An internal error has occured. See error details.


swagger:response getSeriesVersionsVersionedSeriesIdValuesInternalServerError
*/
type GetSeriesVersionsVersionedSeriesIDValuesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesVersionsVersionedSeriesIDValuesInternalServerError creates GetSeriesVersionsVersionedSeriesIDValuesInternalServerError with default headers values
func NewGetSeriesVersionsVersionedSeriesIDValuesInternalServerError() *GetSeriesVersionsVersionedSeriesIDValuesInternalServerError {

	return &GetSeriesVersionsVersionedSeriesIDValuesInternalServerError{}
}

// WithPayload adds the payload to the get series versions versioned series Id values internal server error response
func (o *GetSeriesVersionsVersionedSeriesIDValuesInternalServerError) WithPayload(payload *models.APIError) *GetSeriesVersionsVersionedSeriesIDValuesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series versions versioned series Id values internal server error response
func (o *GetSeriesVersionsVersionedSeriesIDValuesInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDValuesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesVersionsVersionedSeriesIDValuesInternalServerError) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

/*GetSeriesVersionsVersionedSeriesIDValuesDefault Other error. See error details.


swagger:response getSeriesVersionsVersionedSeriesIdValuesDefault
*/
type GetSeriesVersionsVersionedSeriesIDValuesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesVersionsVersionedSeriesIDValuesDefault creates GetSeriesVersionsVersionedSeriesIDValuesDefault with default headers values
func NewGetSeriesVersionsVersionedSeriesIDValuesDefault(code int) *GetSeriesVersionsVersionedSeriesIDValuesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSeriesVersionsVersionedSeriesIDValuesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get series versions versioned series ID values default response
func (o *GetSeriesVersionsVersionedSeriesIDValuesDefault) WithStatusCode(code int) *GetSeriesVersionsVersionedSeriesIDValuesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get series versions versioned series ID values default response
func (o *GetSeriesVersionsVersionedSeriesIDValuesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get series versions versioned series ID values default response
func (o *GetSeriesVersionsVersionedSeriesIDValuesDefault) WithPayload(payload *models.APIError) *GetSeriesVersionsVersionedSeriesIDValuesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series versions versioned series ID values default response
func (o *GetSeriesVersionsVersionedSeriesIDValuesDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDValuesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesVersionsVersionedSeriesIDValuesDefault) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

type GetSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder struct {
	middleware.Responder
}

func (*GetSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder) GetSeriesVersionsVersionedSeriesIDValuesResponder() {
}

func GetSeriesVersionsVersionedSeriesIDValuesNotImplemented() GetSeriesVersionsVersionedSeriesIDValuesResponder {
	return &GetSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetSeriesVersionsVersionedSeriesIDValues has not yet been implemented",
		),
	}
}

type GetSeriesVersionsVersionedSeriesIDValuesResponder interface {
	middleware.Responder
	GetSeriesVersionsVersionedSeriesIDValuesResponder()
}