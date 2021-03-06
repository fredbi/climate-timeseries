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

// GetSeriesVersionsVersionedSeriesIDOKCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDOK
const GetSeriesVersionsVersionedSeriesIDOKCode int = 200

/*GetSeriesVersionsVersionedSeriesIDOK Description of a given version of a time series.


swagger:response getSeriesVersionsVersionedSeriesIdOK
*/
type GetSeriesVersionsVersionedSeriesIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.VersionedSeries `json:"body,omitempty"`
}

// NewGetSeriesVersionsVersionedSeriesIDOK creates GetSeriesVersionsVersionedSeriesIDOK with default headers values
func NewGetSeriesVersionsVersionedSeriesIDOK() *GetSeriesVersionsVersionedSeriesIDOK {

	return &GetSeriesVersionsVersionedSeriesIDOK{}
}

// WithPayload adds the payload to the get series versions versioned series Id o k response
func (o *GetSeriesVersionsVersionedSeriesIDOK) WithPayload(payload *models.VersionedSeries) *GetSeriesVersionsVersionedSeriesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series versions versioned series Id o k response
func (o *GetSeriesVersionsVersionedSeriesIDOK) SetPayload(payload *models.VersionedSeries) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesVersionsVersionedSeriesIDOK) GetSeriesVersionsVersionedSeriesIDResponder() {}

// GetSeriesVersionsVersionedSeriesIDBadRequestCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDBadRequest
const GetSeriesVersionsVersionedSeriesIDBadRequestCode int = 400

/*GetSeriesVersionsVersionedSeriesIDBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response getSeriesVersionsVersionedSeriesIdBadRequest
*/
type GetSeriesVersionsVersionedSeriesIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesVersionsVersionedSeriesIDBadRequest creates GetSeriesVersionsVersionedSeriesIDBadRequest with default headers values
func NewGetSeriesVersionsVersionedSeriesIDBadRequest() *GetSeriesVersionsVersionedSeriesIDBadRequest {

	return &GetSeriesVersionsVersionedSeriesIDBadRequest{}
}

// WithPayload adds the payload to the get series versions versioned series Id bad request response
func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) WithPayload(payload *models.APIError) *GetSeriesVersionsVersionedSeriesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series versions versioned series Id bad request response
func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesVersionsVersionedSeriesIDBadRequest) GetSeriesVersionsVersionedSeriesIDResponder() {
}

// GetSeriesVersionsVersionedSeriesIDUnauthorizedCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDUnauthorized
const GetSeriesVersionsVersionedSeriesIDUnauthorizedCode int = 401

/*GetSeriesVersionsVersionedSeriesIDUnauthorized Unauthorized access for a lack of authentication


swagger:response getSeriesVersionsVersionedSeriesIdUnauthorized
*/
type GetSeriesVersionsVersionedSeriesIDUnauthorized struct {
}

// NewGetSeriesVersionsVersionedSeriesIDUnauthorized creates GetSeriesVersionsVersionedSeriesIDUnauthorized with default headers values
func NewGetSeriesVersionsVersionedSeriesIDUnauthorized() *GetSeriesVersionsVersionedSeriesIDUnauthorized {

	return &GetSeriesVersionsVersionedSeriesIDUnauthorized{}
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *GetSeriesVersionsVersionedSeriesIDUnauthorized) GetSeriesVersionsVersionedSeriesIDResponder() {
}

// GetSeriesVersionsVersionedSeriesIDForbiddenCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDForbidden
const GetSeriesVersionsVersionedSeriesIDForbiddenCode int = 403

/*GetSeriesVersionsVersionedSeriesIDForbidden Forbidden access for a lack of sufficient privileges


swagger:response getSeriesVersionsVersionedSeriesIdForbidden
*/
type GetSeriesVersionsVersionedSeriesIDForbidden struct {
}

// NewGetSeriesVersionsVersionedSeriesIDForbidden creates GetSeriesVersionsVersionedSeriesIDForbidden with default headers values
func NewGetSeriesVersionsVersionedSeriesIDForbidden() *GetSeriesVersionsVersionedSeriesIDForbidden {

	return &GetSeriesVersionsVersionedSeriesIDForbidden{}
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *GetSeriesVersionsVersionedSeriesIDForbidden) GetSeriesVersionsVersionedSeriesIDResponder() {}

// GetSeriesVersionsVersionedSeriesIDNotFoundCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDNotFound
const GetSeriesVersionsVersionedSeriesIDNotFoundCode int = 404

/*GetSeriesVersionsVersionedSeriesIDNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response getSeriesVersionsVersionedSeriesIdNotFound
*/
type GetSeriesVersionsVersionedSeriesIDNotFound struct {
}

// NewGetSeriesVersionsVersionedSeriesIDNotFound creates GetSeriesVersionsVersionedSeriesIDNotFound with default headers values
func NewGetSeriesVersionsVersionedSeriesIDNotFound() *GetSeriesVersionsVersionedSeriesIDNotFound {

	return &GetSeriesVersionsVersionedSeriesIDNotFound{}
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *GetSeriesVersionsVersionedSeriesIDNotFound) GetSeriesVersionsVersionedSeriesIDResponder() {}

// GetSeriesVersionsVersionedSeriesIDMethodNotAllowedCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDMethodNotAllowed
const GetSeriesVersionsVersionedSeriesIDMethodNotAllowedCode int = 405

/*GetSeriesVersionsVersionedSeriesIDMethodNotAllowed MethodNotAllowed


swagger:response getSeriesVersionsVersionedSeriesIdMethodNotAllowed
*/
type GetSeriesVersionsVersionedSeriesIDMethodNotAllowed struct {
}

// NewGetSeriesVersionsVersionedSeriesIDMethodNotAllowed creates GetSeriesVersionsVersionedSeriesIDMethodNotAllowed with default headers values
func NewGetSeriesVersionsVersionedSeriesIDMethodNotAllowed() *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed {

	return &GetSeriesVersionsVersionedSeriesIDMethodNotAllowed{}
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *GetSeriesVersionsVersionedSeriesIDMethodNotAllowed) GetSeriesVersionsVersionedSeriesIDResponder() {
}

// GetSeriesVersionsVersionedSeriesIDInternalServerErrorCode is the HTTP code returned for type GetSeriesVersionsVersionedSeriesIDInternalServerError
const GetSeriesVersionsVersionedSeriesIDInternalServerErrorCode int = 500

/*GetSeriesVersionsVersionedSeriesIDInternalServerError An internal error has occured. See error details.


swagger:response getSeriesVersionsVersionedSeriesIdInternalServerError
*/
type GetSeriesVersionsVersionedSeriesIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesVersionsVersionedSeriesIDInternalServerError creates GetSeriesVersionsVersionedSeriesIDInternalServerError with default headers values
func NewGetSeriesVersionsVersionedSeriesIDInternalServerError() *GetSeriesVersionsVersionedSeriesIDInternalServerError {

	return &GetSeriesVersionsVersionedSeriesIDInternalServerError{}
}

// WithPayload adds the payload to the get series versions versioned series Id internal server error response
func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) WithPayload(payload *models.APIError) *GetSeriesVersionsVersionedSeriesIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series versions versioned series Id internal server error response
func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesVersionsVersionedSeriesIDInternalServerError) GetSeriesVersionsVersionedSeriesIDResponder() {
}

/*GetSeriesVersionsVersionedSeriesIDDefault Other error. See error details.


swagger:response getSeriesVersionsVersionedSeriesIdDefault
*/
type GetSeriesVersionsVersionedSeriesIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesVersionsVersionedSeriesIDDefault creates GetSeriesVersionsVersionedSeriesIDDefault with default headers values
func NewGetSeriesVersionsVersionedSeriesIDDefault(code int) *GetSeriesVersionsVersionedSeriesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSeriesVersionsVersionedSeriesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get series versions versioned series ID default response
func (o *GetSeriesVersionsVersionedSeriesIDDefault) WithStatusCode(code int) *GetSeriesVersionsVersionedSeriesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get series versions versioned series ID default response
func (o *GetSeriesVersionsVersionedSeriesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get series versions versioned series ID default response
func (o *GetSeriesVersionsVersionedSeriesIDDefault) WithPayload(payload *models.APIError) *GetSeriesVersionsVersionedSeriesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series versions versioned series ID default response
func (o *GetSeriesVersionsVersionedSeriesIDDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesVersionsVersionedSeriesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesVersionsVersionedSeriesIDDefault) GetSeriesVersionsVersionedSeriesIDResponder() {}

type GetSeriesVersionsVersionedSeriesIDNotImplementedResponder struct {
	middleware.Responder
}

func (*GetSeriesVersionsVersionedSeriesIDNotImplementedResponder) GetSeriesVersionsVersionedSeriesIDResponder() {
}

func GetSeriesVersionsVersionedSeriesIDNotImplemented() GetSeriesVersionsVersionedSeriesIDResponder {
	return &GetSeriesVersionsVersionedSeriesIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetSeriesVersionsVersionedSeriesID has not yet been implemented",
		),
	}
}

type GetSeriesVersionsVersionedSeriesIDResponder interface {
	middleware.Responder
	GetSeriesVersionsVersionedSeriesIDResponder()
}
