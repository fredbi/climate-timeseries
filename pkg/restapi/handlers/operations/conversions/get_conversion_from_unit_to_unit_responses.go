// Code generated by go-swagger; DO NOT EDIT.

package conversions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// GetConversionFromUnitToUnitOKCode is the HTTP code returned for type GetConversionFromUnitToUnitOK
const GetConversionFromUnitToUnitOKCode int = 200

/*GetConversionFromUnitToUnitOK A unit conversion specification


swagger:response getConversionFromUnitToUnitOK
*/
type GetConversionFromUnitToUnitOK struct {

	/*
	  In: Body
	*/
	Payload *models.ConversionSpec `json:"body,omitempty"`
}

// NewGetConversionFromUnitToUnitOK creates GetConversionFromUnitToUnitOK with default headers values
func NewGetConversionFromUnitToUnitOK() *GetConversionFromUnitToUnitOK {

	return &GetConversionFromUnitToUnitOK{}
}

// WithPayload adds the payload to the get conversion from unit to unit o k response
func (o *GetConversionFromUnitToUnitOK) WithPayload(payload *models.ConversionSpec) *GetConversionFromUnitToUnitOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get conversion from unit to unit o k response
func (o *GetConversionFromUnitToUnitOK) SetPayload(payload *models.ConversionSpec) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConversionFromUnitToUnitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetConversionFromUnitToUnitOK) GetConversionFromUnitToUnitResponder() {}

// GetConversionFromUnitToUnitBadRequestCode is the HTTP code returned for type GetConversionFromUnitToUnitBadRequest
const GetConversionFromUnitToUnitBadRequestCode int = 400

/*GetConversionFromUnitToUnitBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response getConversionFromUnitToUnitBadRequest
*/
type GetConversionFromUnitToUnitBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetConversionFromUnitToUnitBadRequest creates GetConversionFromUnitToUnitBadRequest with default headers values
func NewGetConversionFromUnitToUnitBadRequest() *GetConversionFromUnitToUnitBadRequest {

	return &GetConversionFromUnitToUnitBadRequest{}
}

// WithPayload adds the payload to the get conversion from unit to unit bad request response
func (o *GetConversionFromUnitToUnitBadRequest) WithPayload(payload *models.APIError) *GetConversionFromUnitToUnitBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get conversion from unit to unit bad request response
func (o *GetConversionFromUnitToUnitBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConversionFromUnitToUnitBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetConversionFromUnitToUnitBadRequest) GetConversionFromUnitToUnitResponder() {}

// GetConversionFromUnitToUnitUnauthorizedCode is the HTTP code returned for type GetConversionFromUnitToUnitUnauthorized
const GetConversionFromUnitToUnitUnauthorizedCode int = 401

/*GetConversionFromUnitToUnitUnauthorized Unauthorized access for a lack of authentication


swagger:response getConversionFromUnitToUnitUnauthorized
*/
type GetConversionFromUnitToUnitUnauthorized struct {
}

// NewGetConversionFromUnitToUnitUnauthorized creates GetConversionFromUnitToUnitUnauthorized with default headers values
func NewGetConversionFromUnitToUnitUnauthorized() *GetConversionFromUnitToUnitUnauthorized {

	return &GetConversionFromUnitToUnitUnauthorized{}
}

// WriteResponse to the client
func (o *GetConversionFromUnitToUnitUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *GetConversionFromUnitToUnitUnauthorized) GetConversionFromUnitToUnitResponder() {}

// GetConversionFromUnitToUnitForbiddenCode is the HTTP code returned for type GetConversionFromUnitToUnitForbidden
const GetConversionFromUnitToUnitForbiddenCode int = 403

/*GetConversionFromUnitToUnitForbidden Forbidden access for a lack of sufficient privileges


swagger:response getConversionFromUnitToUnitForbidden
*/
type GetConversionFromUnitToUnitForbidden struct {
}

// NewGetConversionFromUnitToUnitForbidden creates GetConversionFromUnitToUnitForbidden with default headers values
func NewGetConversionFromUnitToUnitForbidden() *GetConversionFromUnitToUnitForbidden {

	return &GetConversionFromUnitToUnitForbidden{}
}

// WriteResponse to the client
func (o *GetConversionFromUnitToUnitForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *GetConversionFromUnitToUnitForbidden) GetConversionFromUnitToUnitResponder() {}

// GetConversionFromUnitToUnitNotFoundCode is the HTTP code returned for type GetConversionFromUnitToUnitNotFound
const GetConversionFromUnitToUnitNotFoundCode int = 404

/*GetConversionFromUnitToUnitNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response getConversionFromUnitToUnitNotFound
*/
type GetConversionFromUnitToUnitNotFound struct {
}

// NewGetConversionFromUnitToUnitNotFound creates GetConversionFromUnitToUnitNotFound with default headers values
func NewGetConversionFromUnitToUnitNotFound() *GetConversionFromUnitToUnitNotFound {

	return &GetConversionFromUnitToUnitNotFound{}
}

// WriteResponse to the client
func (o *GetConversionFromUnitToUnitNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *GetConversionFromUnitToUnitNotFound) GetConversionFromUnitToUnitResponder() {}

// GetConversionFromUnitToUnitMethodNotAllowedCode is the HTTP code returned for type GetConversionFromUnitToUnitMethodNotAllowed
const GetConversionFromUnitToUnitMethodNotAllowedCode int = 405

/*GetConversionFromUnitToUnitMethodNotAllowed MethodNotAllowed


swagger:response getConversionFromUnitToUnitMethodNotAllowed
*/
type GetConversionFromUnitToUnitMethodNotAllowed struct {
}

// NewGetConversionFromUnitToUnitMethodNotAllowed creates GetConversionFromUnitToUnitMethodNotAllowed with default headers values
func NewGetConversionFromUnitToUnitMethodNotAllowed() *GetConversionFromUnitToUnitMethodNotAllowed {

	return &GetConversionFromUnitToUnitMethodNotAllowed{}
}

// WriteResponse to the client
func (o *GetConversionFromUnitToUnitMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *GetConversionFromUnitToUnitMethodNotAllowed) GetConversionFromUnitToUnitResponder() {}

// GetConversionFromUnitToUnitInternalServerErrorCode is the HTTP code returned for type GetConversionFromUnitToUnitInternalServerError
const GetConversionFromUnitToUnitInternalServerErrorCode int = 500

/*GetConversionFromUnitToUnitInternalServerError An internal error has occured. See error details.


swagger:response getConversionFromUnitToUnitInternalServerError
*/
type GetConversionFromUnitToUnitInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetConversionFromUnitToUnitInternalServerError creates GetConversionFromUnitToUnitInternalServerError with default headers values
func NewGetConversionFromUnitToUnitInternalServerError() *GetConversionFromUnitToUnitInternalServerError {

	return &GetConversionFromUnitToUnitInternalServerError{}
}

// WithPayload adds the payload to the get conversion from unit to unit internal server error response
func (o *GetConversionFromUnitToUnitInternalServerError) WithPayload(payload *models.APIError) *GetConversionFromUnitToUnitInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get conversion from unit to unit internal server error response
func (o *GetConversionFromUnitToUnitInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConversionFromUnitToUnitInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetConversionFromUnitToUnitInternalServerError) GetConversionFromUnitToUnitResponder() {}

/*GetConversionFromUnitToUnitDefault Other error. See error details.


swagger:response getConversionFromUnitToUnitDefault
*/
type GetConversionFromUnitToUnitDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetConversionFromUnitToUnitDefault creates GetConversionFromUnitToUnitDefault with default headers values
func NewGetConversionFromUnitToUnitDefault(code int) *GetConversionFromUnitToUnitDefault {
	if code <= 0 {
		code = 500
	}

	return &GetConversionFromUnitToUnitDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get conversion from unit to unit default response
func (o *GetConversionFromUnitToUnitDefault) WithStatusCode(code int) *GetConversionFromUnitToUnitDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get conversion from unit to unit default response
func (o *GetConversionFromUnitToUnitDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get conversion from unit to unit default response
func (o *GetConversionFromUnitToUnitDefault) WithPayload(payload *models.APIError) *GetConversionFromUnitToUnitDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get conversion from unit to unit default response
func (o *GetConversionFromUnitToUnitDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConversionFromUnitToUnitDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetConversionFromUnitToUnitDefault) GetConversionFromUnitToUnitResponder() {}

type GetConversionFromUnitToUnitNotImplementedResponder struct {
	middleware.Responder
}

func (*GetConversionFromUnitToUnitNotImplementedResponder) GetConversionFromUnitToUnitResponder() {}

func GetConversionFromUnitToUnitNotImplemented() GetConversionFromUnitToUnitResponder {
	return &GetConversionFromUnitToUnitNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetConversionFromUnitToUnit has not yet been implemented",
		),
	}
}

type GetConversionFromUnitToUnitResponder interface {
	middleware.Responder
	GetConversionFromUnitToUnitResponder()
}
