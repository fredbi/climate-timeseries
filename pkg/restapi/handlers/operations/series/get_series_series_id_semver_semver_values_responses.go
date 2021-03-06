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

// GetSeriesSeriesIDSemverSemverValuesOKCode is the HTTP code returned for type GetSeriesSeriesIDSemverSemverValuesOK
const GetSeriesSeriesIDSemverSemverValuesOKCode int = 200

/*GetSeriesSeriesIDSemverSemverValuesOK Values of a given version of a time series.


swagger:response getSeriesSeriesIdSemverSemverValuesOK
*/
type GetSeriesSeriesIDSemverSemverValuesOK struct {

	/*
	  In: Body
	*/
	Payload models.TsValues `json:"body,omitempty"`
}

// NewGetSeriesSeriesIDSemverSemverValuesOK creates GetSeriesSeriesIDSemverSemverValuesOK with default headers values
func NewGetSeriesSeriesIDSemverSemverValuesOK() *GetSeriesSeriesIDSemverSemverValuesOK {

	return &GetSeriesSeriesIDSemverSemverValuesOK{}
}

// WithPayload adds the payload to the get series series Id semver semver values o k response
func (o *GetSeriesSeriesIDSemverSemverValuesOK) WithPayload(payload models.TsValues) *GetSeriesSeriesIDSemverSemverValuesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series series Id semver semver values o k response
func (o *GetSeriesSeriesIDSemverSemverValuesOK) SetPayload(payload models.TsValues) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesSeriesIDSemverSemverValuesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

func (o *GetSeriesSeriesIDSemverSemverValuesOK) GetSeriesSeriesIDSemverSemverValuesResponder() {}

// GetSeriesSeriesIDSemverSemverValuesBadRequestCode is the HTTP code returned for type GetSeriesSeriesIDSemverSemverValuesBadRequest
const GetSeriesSeriesIDSemverSemverValuesBadRequestCode int = 400

/*GetSeriesSeriesIDSemverSemverValuesBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response getSeriesSeriesIdSemverSemverValuesBadRequest
*/
type GetSeriesSeriesIDSemverSemverValuesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesSeriesIDSemverSemverValuesBadRequest creates GetSeriesSeriesIDSemverSemverValuesBadRequest with default headers values
func NewGetSeriesSeriesIDSemverSemverValuesBadRequest() *GetSeriesSeriesIDSemverSemverValuesBadRequest {

	return &GetSeriesSeriesIDSemverSemverValuesBadRequest{}
}

// WithPayload adds the payload to the get series series Id semver semver values bad request response
func (o *GetSeriesSeriesIDSemverSemverValuesBadRequest) WithPayload(payload *models.APIError) *GetSeriesSeriesIDSemverSemverValuesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series series Id semver semver values bad request response
func (o *GetSeriesSeriesIDSemverSemverValuesBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesSeriesIDSemverSemverValuesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesSeriesIDSemverSemverValuesBadRequest) GetSeriesSeriesIDSemverSemverValuesResponder() {
}

// GetSeriesSeriesIDSemverSemverValuesUnauthorizedCode is the HTTP code returned for type GetSeriesSeriesIDSemverSemverValuesUnauthorized
const GetSeriesSeriesIDSemverSemverValuesUnauthorizedCode int = 401

/*GetSeriesSeriesIDSemverSemverValuesUnauthorized Unauthorized access for a lack of authentication


swagger:response getSeriesSeriesIdSemverSemverValuesUnauthorized
*/
type GetSeriesSeriesIDSemverSemverValuesUnauthorized struct {
}

// NewGetSeriesSeriesIDSemverSemverValuesUnauthorized creates GetSeriesSeriesIDSemverSemverValuesUnauthorized with default headers values
func NewGetSeriesSeriesIDSemverSemverValuesUnauthorized() *GetSeriesSeriesIDSemverSemverValuesUnauthorized {

	return &GetSeriesSeriesIDSemverSemverValuesUnauthorized{}
}

// WriteResponse to the client
func (o *GetSeriesSeriesIDSemverSemverValuesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *GetSeriesSeriesIDSemverSemverValuesUnauthorized) GetSeriesSeriesIDSemverSemverValuesResponder() {
}

// GetSeriesSeriesIDSemverSemverValuesForbiddenCode is the HTTP code returned for type GetSeriesSeriesIDSemverSemverValuesForbidden
const GetSeriesSeriesIDSemverSemverValuesForbiddenCode int = 403

/*GetSeriesSeriesIDSemverSemverValuesForbidden Forbidden access for a lack of sufficient privileges


swagger:response getSeriesSeriesIdSemverSemverValuesForbidden
*/
type GetSeriesSeriesIDSemverSemverValuesForbidden struct {
}

// NewGetSeriesSeriesIDSemverSemverValuesForbidden creates GetSeriesSeriesIDSemverSemverValuesForbidden with default headers values
func NewGetSeriesSeriesIDSemverSemverValuesForbidden() *GetSeriesSeriesIDSemverSemverValuesForbidden {

	return &GetSeriesSeriesIDSemverSemverValuesForbidden{}
}

// WriteResponse to the client
func (o *GetSeriesSeriesIDSemverSemverValuesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *GetSeriesSeriesIDSemverSemverValuesForbidden) GetSeriesSeriesIDSemverSemverValuesResponder() {
}

// GetSeriesSeriesIDSemverSemverValuesNotFoundCode is the HTTP code returned for type GetSeriesSeriesIDSemverSemverValuesNotFound
const GetSeriesSeriesIDSemverSemverValuesNotFoundCode int = 404

/*GetSeriesSeriesIDSemverSemverValuesNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response getSeriesSeriesIdSemverSemverValuesNotFound
*/
type GetSeriesSeriesIDSemverSemverValuesNotFound struct {
}

// NewGetSeriesSeriesIDSemverSemverValuesNotFound creates GetSeriesSeriesIDSemverSemverValuesNotFound with default headers values
func NewGetSeriesSeriesIDSemverSemverValuesNotFound() *GetSeriesSeriesIDSemverSemverValuesNotFound {

	return &GetSeriesSeriesIDSemverSemverValuesNotFound{}
}

// WriteResponse to the client
func (o *GetSeriesSeriesIDSemverSemverValuesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *GetSeriesSeriesIDSemverSemverValuesNotFound) GetSeriesSeriesIDSemverSemverValuesResponder() {
}

// GetSeriesSeriesIDSemverSemverValuesMethodNotAllowedCode is the HTTP code returned for type GetSeriesSeriesIDSemverSemverValuesMethodNotAllowed
const GetSeriesSeriesIDSemverSemverValuesMethodNotAllowedCode int = 405

/*GetSeriesSeriesIDSemverSemverValuesMethodNotAllowed MethodNotAllowed


swagger:response getSeriesSeriesIdSemverSemverValuesMethodNotAllowed
*/
type GetSeriesSeriesIDSemverSemverValuesMethodNotAllowed struct {
}

// NewGetSeriesSeriesIDSemverSemverValuesMethodNotAllowed creates GetSeriesSeriesIDSemverSemverValuesMethodNotAllowed with default headers values
func NewGetSeriesSeriesIDSemverSemverValuesMethodNotAllowed() *GetSeriesSeriesIDSemverSemverValuesMethodNotAllowed {

	return &GetSeriesSeriesIDSemverSemverValuesMethodNotAllowed{}
}

// WriteResponse to the client
func (o *GetSeriesSeriesIDSemverSemverValuesMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *GetSeriesSeriesIDSemverSemverValuesMethodNotAllowed) GetSeriesSeriesIDSemverSemverValuesResponder() {
}

// GetSeriesSeriesIDSemverSemverValuesInternalServerErrorCode is the HTTP code returned for type GetSeriesSeriesIDSemverSemverValuesInternalServerError
const GetSeriesSeriesIDSemverSemverValuesInternalServerErrorCode int = 500

/*GetSeriesSeriesIDSemverSemverValuesInternalServerError An internal error has occured. See error details.


swagger:response getSeriesSeriesIdSemverSemverValuesInternalServerError
*/
type GetSeriesSeriesIDSemverSemverValuesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesSeriesIDSemverSemverValuesInternalServerError creates GetSeriesSeriesIDSemverSemverValuesInternalServerError with default headers values
func NewGetSeriesSeriesIDSemverSemverValuesInternalServerError() *GetSeriesSeriesIDSemverSemverValuesInternalServerError {

	return &GetSeriesSeriesIDSemverSemverValuesInternalServerError{}
}

// WithPayload adds the payload to the get series series Id semver semver values internal server error response
func (o *GetSeriesSeriesIDSemverSemverValuesInternalServerError) WithPayload(payload *models.APIError) *GetSeriesSeriesIDSemverSemverValuesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series series Id semver semver values internal server error response
func (o *GetSeriesSeriesIDSemverSemverValuesInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesSeriesIDSemverSemverValuesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesSeriesIDSemverSemverValuesInternalServerError) GetSeriesSeriesIDSemverSemverValuesResponder() {
}

/*GetSeriesSeriesIDSemverSemverValuesDefault Other error. See error details.


swagger:response getSeriesSeriesIdSemverSemverValuesDefault
*/
type GetSeriesSeriesIDSemverSemverValuesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSeriesSeriesIDSemverSemverValuesDefault creates GetSeriesSeriesIDSemverSemverValuesDefault with default headers values
func NewGetSeriesSeriesIDSemverSemverValuesDefault(code int) *GetSeriesSeriesIDSemverSemverValuesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSeriesSeriesIDSemverSemverValuesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get series series ID semver semver values default response
func (o *GetSeriesSeriesIDSemverSemverValuesDefault) WithStatusCode(code int) *GetSeriesSeriesIDSemverSemverValuesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get series series ID semver semver values default response
func (o *GetSeriesSeriesIDSemverSemverValuesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get series series ID semver semver values default response
func (o *GetSeriesSeriesIDSemverSemverValuesDefault) WithPayload(payload *models.APIError) *GetSeriesSeriesIDSemverSemverValuesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get series series ID semver semver values default response
func (o *GetSeriesSeriesIDSemverSemverValuesDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSeriesSeriesIDSemverSemverValuesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSeriesSeriesIDSemverSemverValuesDefault) GetSeriesSeriesIDSemverSemverValuesResponder() {}

type GetSeriesSeriesIDSemverSemverValuesNotImplementedResponder struct {
	middleware.Responder
}

func (*GetSeriesSeriesIDSemverSemverValuesNotImplementedResponder) GetSeriesSeriesIDSemverSemverValuesResponder() {
}

func GetSeriesSeriesIDSemverSemverValuesNotImplemented() GetSeriesSeriesIDSemverSemverValuesResponder {
	return &GetSeriesSeriesIDSemverSemverValuesNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetSeriesSeriesIDSemverSemverValues has not yet been implemented",
		),
	}
}

type GetSeriesSeriesIDSemverSemverValuesResponder interface {
	middleware.Responder
	GetSeriesSeriesIDSemverSemverValuesResponder()
}
