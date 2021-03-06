// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// PostSeriesVersionsVersionedSeriesIDCreatedCode is the HTTP code returned for type PostSeriesVersionsVersionedSeriesIDCreated
const PostSeriesVersionsVersionedSeriesIDCreatedCode int = 201

/*PostSeriesVersionsVersionedSeriesIDCreated Time series version successfully created.

Check the response headers to retrieve this resource.


swagger:response postSeriesVersionsVersionedSeriesIdCreated
*/
type PostSeriesVersionsVersionedSeriesIDCreated struct {
	/*The URI of the newly created resource.


	 */
	Location strfmt.URI `json:"Location"`
	/*The ID of the updated resource.


	 */
	XID int64 `json:"X-ID"`
}

// NewPostSeriesVersionsVersionedSeriesIDCreated creates PostSeriesVersionsVersionedSeriesIDCreated with default headers values
func NewPostSeriesVersionsVersionedSeriesIDCreated() *PostSeriesVersionsVersionedSeriesIDCreated {

	return &PostSeriesVersionsVersionedSeriesIDCreated{}
}

// WithLocation adds the location to the post series versions versioned series Id created response
func (o *PostSeriesVersionsVersionedSeriesIDCreated) WithLocation(location strfmt.URI) *PostSeriesVersionsVersionedSeriesIDCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the post series versions versioned series Id created response
func (o *PostSeriesVersionsVersionedSeriesIDCreated) SetLocation(location strfmt.URI) {
	o.Location = location
}

// WithXID adds the xId to the post series versions versioned series Id created response
func (o *PostSeriesVersionsVersionedSeriesIDCreated) WithXID(xID int64) *PostSeriesVersionsVersionedSeriesIDCreated {
	o.XID = xID
	return o
}

// SetXID sets the xId to the post series versions versioned series Id created response
func (o *PostSeriesVersionsVersionedSeriesIDCreated) SetXID(xID int64) {
	o.XID = xID
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location.String()
	if location != "" {
		rw.Header().Set("Location", location)
	}

	// response header X-ID

	xID := swag.FormatInt64(o.XID)
	if xID != "" {
		rw.Header().Set("X-ID", xID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

func (o *PostSeriesVersionsVersionedSeriesIDCreated) PostSeriesVersionsVersionedSeriesIDResponder() {}

// PostSeriesVersionsVersionedSeriesIDBadRequestCode is the HTTP code returned for type PostSeriesVersionsVersionedSeriesIDBadRequest
const PostSeriesVersionsVersionedSeriesIDBadRequestCode int = 400

/*PostSeriesVersionsVersionedSeriesIDBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response postSeriesVersionsVersionedSeriesIdBadRequest
*/
type PostSeriesVersionsVersionedSeriesIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostSeriesVersionsVersionedSeriesIDBadRequest creates PostSeriesVersionsVersionedSeriesIDBadRequest with default headers values
func NewPostSeriesVersionsVersionedSeriesIDBadRequest() *PostSeriesVersionsVersionedSeriesIDBadRequest {

	return &PostSeriesVersionsVersionedSeriesIDBadRequest{}
}

// WithPayload adds the payload to the post series versions versioned series Id bad request response
func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) WithPayload(payload *models.APIError) *PostSeriesVersionsVersionedSeriesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post series versions versioned series Id bad request response
func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostSeriesVersionsVersionedSeriesIDBadRequest) PostSeriesVersionsVersionedSeriesIDResponder() {
}

// PostSeriesVersionsVersionedSeriesIDUnauthorizedCode is the HTTP code returned for type PostSeriesVersionsVersionedSeriesIDUnauthorized
const PostSeriesVersionsVersionedSeriesIDUnauthorizedCode int = 401

/*PostSeriesVersionsVersionedSeriesIDUnauthorized Unauthorized access for a lack of authentication


swagger:response postSeriesVersionsVersionedSeriesIdUnauthorized
*/
type PostSeriesVersionsVersionedSeriesIDUnauthorized struct {
}

// NewPostSeriesVersionsVersionedSeriesIDUnauthorized creates PostSeriesVersionsVersionedSeriesIDUnauthorized with default headers values
func NewPostSeriesVersionsVersionedSeriesIDUnauthorized() *PostSeriesVersionsVersionedSeriesIDUnauthorized {

	return &PostSeriesVersionsVersionedSeriesIDUnauthorized{}
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *PostSeriesVersionsVersionedSeriesIDUnauthorized) PostSeriesVersionsVersionedSeriesIDResponder() {
}

// PostSeriesVersionsVersionedSeriesIDForbiddenCode is the HTTP code returned for type PostSeriesVersionsVersionedSeriesIDForbidden
const PostSeriesVersionsVersionedSeriesIDForbiddenCode int = 403

/*PostSeriesVersionsVersionedSeriesIDForbidden Forbidden access for a lack of sufficient privileges


swagger:response postSeriesVersionsVersionedSeriesIdForbidden
*/
type PostSeriesVersionsVersionedSeriesIDForbidden struct {
}

// NewPostSeriesVersionsVersionedSeriesIDForbidden creates PostSeriesVersionsVersionedSeriesIDForbidden with default headers values
func NewPostSeriesVersionsVersionedSeriesIDForbidden() *PostSeriesVersionsVersionedSeriesIDForbidden {

	return &PostSeriesVersionsVersionedSeriesIDForbidden{}
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *PostSeriesVersionsVersionedSeriesIDForbidden) PostSeriesVersionsVersionedSeriesIDResponder() {
}

// PostSeriesVersionsVersionedSeriesIDNotFoundCode is the HTTP code returned for type PostSeriesVersionsVersionedSeriesIDNotFound
const PostSeriesVersionsVersionedSeriesIDNotFoundCode int = 404

/*PostSeriesVersionsVersionedSeriesIDNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response postSeriesVersionsVersionedSeriesIdNotFound
*/
type PostSeriesVersionsVersionedSeriesIDNotFound struct {
}

// NewPostSeriesVersionsVersionedSeriesIDNotFound creates PostSeriesVersionsVersionedSeriesIDNotFound with default headers values
func NewPostSeriesVersionsVersionedSeriesIDNotFound() *PostSeriesVersionsVersionedSeriesIDNotFound {

	return &PostSeriesVersionsVersionedSeriesIDNotFound{}
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *PostSeriesVersionsVersionedSeriesIDNotFound) PostSeriesVersionsVersionedSeriesIDResponder() {
}

// PostSeriesVersionsVersionedSeriesIDMethodNotAllowedCode is the HTTP code returned for type PostSeriesVersionsVersionedSeriesIDMethodNotAllowed
const PostSeriesVersionsVersionedSeriesIDMethodNotAllowedCode int = 405

/*PostSeriesVersionsVersionedSeriesIDMethodNotAllowed MethodNotAllowed


swagger:response postSeriesVersionsVersionedSeriesIdMethodNotAllowed
*/
type PostSeriesVersionsVersionedSeriesIDMethodNotAllowed struct {
}

// NewPostSeriesVersionsVersionedSeriesIDMethodNotAllowed creates PostSeriesVersionsVersionedSeriesIDMethodNotAllowed with default headers values
func NewPostSeriesVersionsVersionedSeriesIDMethodNotAllowed() *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed {

	return &PostSeriesVersionsVersionedSeriesIDMethodNotAllowed{}
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *PostSeriesVersionsVersionedSeriesIDMethodNotAllowed) PostSeriesVersionsVersionedSeriesIDResponder() {
}

// PostSeriesVersionsVersionedSeriesIDConflictCode is the HTTP code returned for type PostSeriesVersionsVersionedSeriesIDConflict
const PostSeriesVersionsVersionedSeriesIDConflictCode int = 409

/*PostSeriesVersionsVersionedSeriesIDConflict Resource already exists. An object creation was requested, but this object was already existing.


swagger:response postSeriesVersionsVersionedSeriesIdConflict
*/
type PostSeriesVersionsVersionedSeriesIDConflict struct {
}

// NewPostSeriesVersionsVersionedSeriesIDConflict creates PostSeriesVersionsVersionedSeriesIDConflict with default headers values
func NewPostSeriesVersionsVersionedSeriesIDConflict() *PostSeriesVersionsVersionedSeriesIDConflict {

	return &PostSeriesVersionsVersionedSeriesIDConflict{}
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

func (o *PostSeriesVersionsVersionedSeriesIDConflict) PostSeriesVersionsVersionedSeriesIDResponder() {
}

// PostSeriesVersionsVersionedSeriesIDInternalServerErrorCode is the HTTP code returned for type PostSeriesVersionsVersionedSeriesIDInternalServerError
const PostSeriesVersionsVersionedSeriesIDInternalServerErrorCode int = 500

/*PostSeriesVersionsVersionedSeriesIDInternalServerError An internal error has occured. See error details.


swagger:response postSeriesVersionsVersionedSeriesIdInternalServerError
*/
type PostSeriesVersionsVersionedSeriesIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostSeriesVersionsVersionedSeriesIDInternalServerError creates PostSeriesVersionsVersionedSeriesIDInternalServerError with default headers values
func NewPostSeriesVersionsVersionedSeriesIDInternalServerError() *PostSeriesVersionsVersionedSeriesIDInternalServerError {

	return &PostSeriesVersionsVersionedSeriesIDInternalServerError{}
}

// WithPayload adds the payload to the post series versions versioned series Id internal server error response
func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) WithPayload(payload *models.APIError) *PostSeriesVersionsVersionedSeriesIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post series versions versioned series Id internal server error response
func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostSeriesVersionsVersionedSeriesIDInternalServerError) PostSeriesVersionsVersionedSeriesIDResponder() {
}

/*PostSeriesVersionsVersionedSeriesIDDefault Other error. See error details.


swagger:response postSeriesVersionsVersionedSeriesIdDefault
*/
type PostSeriesVersionsVersionedSeriesIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostSeriesVersionsVersionedSeriesIDDefault creates PostSeriesVersionsVersionedSeriesIDDefault with default headers values
func NewPostSeriesVersionsVersionedSeriesIDDefault(code int) *PostSeriesVersionsVersionedSeriesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PostSeriesVersionsVersionedSeriesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post series versions versioned series ID default response
func (o *PostSeriesVersionsVersionedSeriesIDDefault) WithStatusCode(code int) *PostSeriesVersionsVersionedSeriesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post series versions versioned series ID default response
func (o *PostSeriesVersionsVersionedSeriesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post series versions versioned series ID default response
func (o *PostSeriesVersionsVersionedSeriesIDDefault) WithPayload(payload *models.APIError) *PostSeriesVersionsVersionedSeriesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post series versions versioned series ID default response
func (o *PostSeriesVersionsVersionedSeriesIDDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSeriesVersionsVersionedSeriesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostSeriesVersionsVersionedSeriesIDDefault) PostSeriesVersionsVersionedSeriesIDResponder() {}

type PostSeriesVersionsVersionedSeriesIDNotImplementedResponder struct {
	middleware.Responder
}

func (*PostSeriesVersionsVersionedSeriesIDNotImplementedResponder) PostSeriesVersionsVersionedSeriesIDResponder() {
}

func PostSeriesVersionsVersionedSeriesIDNotImplemented() PostSeriesVersionsVersionedSeriesIDResponder {
	return &PostSeriesVersionsVersionedSeriesIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.PostSeriesVersionsVersionedSeriesID has not yet been implemented",
		),
	}
}

type PostSeriesVersionsVersionedSeriesIDResponder interface {
	middleware.Responder
	PostSeriesVersionsVersionedSeriesIDResponder()
}
