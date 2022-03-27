// Code generated by go-swagger; DO NOT EDIT.

package classes

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

// PostClassesClassIDMembersCreatedCode is the HTTP code returned for type PostClassesClassIDMembersCreated
const PostClassesClassIDMembersCreatedCode int = 201

/*PostClassesClassIDMembersCreated Class members added.


swagger:response postClassesClassIdMembersCreated
*/
type PostClassesClassIDMembersCreated struct {
	/*The URI of the newly created resource.


	 */
	Location strfmt.URI `json:"Location"`
	/*The ID of the newly created resource.


	 */
	XID int64 `json:"X-ID"`
}

// NewPostClassesClassIDMembersCreated creates PostClassesClassIDMembersCreated with default headers values
func NewPostClassesClassIDMembersCreated() *PostClassesClassIDMembersCreated {

	return &PostClassesClassIDMembersCreated{}
}

// WithLocation adds the location to the post classes class Id members created response
func (o *PostClassesClassIDMembersCreated) WithLocation(location strfmt.URI) *PostClassesClassIDMembersCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the post classes class Id members created response
func (o *PostClassesClassIDMembersCreated) SetLocation(location strfmt.URI) {
	o.Location = location
}

// WithXID adds the xId to the post classes class Id members created response
func (o *PostClassesClassIDMembersCreated) WithXID(xID int64) *PostClassesClassIDMembersCreated {
	o.XID = xID
	return o
}

// SetXID sets the xId to the post classes class Id members created response
func (o *PostClassesClassIDMembersCreated) SetXID(xID int64) {
	o.XID = xID
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

func (o *PostClassesClassIDMembersCreated) PostClassesClassIDMembersResponder() {}

// PostClassesClassIDMembersBadRequestCode is the HTTP code returned for type PostClassesClassIDMembersBadRequest
const PostClassesClassIDMembersBadRequestCode int = 400

/*PostClassesClassIDMembersBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response postClassesClassIdMembersBadRequest
*/
type PostClassesClassIDMembersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostClassesClassIDMembersBadRequest creates PostClassesClassIDMembersBadRequest with default headers values
func NewPostClassesClassIDMembersBadRequest() *PostClassesClassIDMembersBadRequest {

	return &PostClassesClassIDMembersBadRequest{}
}

// WithPayload adds the payload to the post classes class Id members bad request response
func (o *PostClassesClassIDMembersBadRequest) WithPayload(payload *models.APIError) *PostClassesClassIDMembersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post classes class Id members bad request response
func (o *PostClassesClassIDMembersBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostClassesClassIDMembersBadRequest) PostClassesClassIDMembersResponder() {}

// PostClassesClassIDMembersUnauthorizedCode is the HTTP code returned for type PostClassesClassIDMembersUnauthorized
const PostClassesClassIDMembersUnauthorizedCode int = 401

/*PostClassesClassIDMembersUnauthorized Unauthorized access for a lack of authentication


swagger:response postClassesClassIdMembersUnauthorized
*/
type PostClassesClassIDMembersUnauthorized struct {
}

// NewPostClassesClassIDMembersUnauthorized creates PostClassesClassIDMembersUnauthorized with default headers values
func NewPostClassesClassIDMembersUnauthorized() *PostClassesClassIDMembersUnauthorized {

	return &PostClassesClassIDMembersUnauthorized{}
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *PostClassesClassIDMembersUnauthorized) PostClassesClassIDMembersResponder() {}

// PostClassesClassIDMembersForbiddenCode is the HTTP code returned for type PostClassesClassIDMembersForbidden
const PostClassesClassIDMembersForbiddenCode int = 403

/*PostClassesClassIDMembersForbidden Forbidden access for a lack of sufficient privileges


swagger:response postClassesClassIdMembersForbidden
*/
type PostClassesClassIDMembersForbidden struct {
}

// NewPostClassesClassIDMembersForbidden creates PostClassesClassIDMembersForbidden with default headers values
func NewPostClassesClassIDMembersForbidden() *PostClassesClassIDMembersForbidden {

	return &PostClassesClassIDMembersForbidden{}
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *PostClassesClassIDMembersForbidden) PostClassesClassIDMembersResponder() {}

// PostClassesClassIDMembersNotFoundCode is the HTTP code returned for type PostClassesClassIDMembersNotFound
const PostClassesClassIDMembersNotFoundCode int = 404

/*PostClassesClassIDMembersNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response postClassesClassIdMembersNotFound
*/
type PostClassesClassIDMembersNotFound struct {
}

// NewPostClassesClassIDMembersNotFound creates PostClassesClassIDMembersNotFound with default headers values
func NewPostClassesClassIDMembersNotFound() *PostClassesClassIDMembersNotFound {

	return &PostClassesClassIDMembersNotFound{}
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *PostClassesClassIDMembersNotFound) PostClassesClassIDMembersResponder() {}

// PostClassesClassIDMembersMethodNotAllowedCode is the HTTP code returned for type PostClassesClassIDMembersMethodNotAllowed
const PostClassesClassIDMembersMethodNotAllowedCode int = 405

/*PostClassesClassIDMembersMethodNotAllowed MethodNotAllowed


swagger:response postClassesClassIdMembersMethodNotAllowed
*/
type PostClassesClassIDMembersMethodNotAllowed struct {
}

// NewPostClassesClassIDMembersMethodNotAllowed creates PostClassesClassIDMembersMethodNotAllowed with default headers values
func NewPostClassesClassIDMembersMethodNotAllowed() *PostClassesClassIDMembersMethodNotAllowed {

	return &PostClassesClassIDMembersMethodNotAllowed{}
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *PostClassesClassIDMembersMethodNotAllowed) PostClassesClassIDMembersResponder() {}

// PostClassesClassIDMembersConflictCode is the HTTP code returned for type PostClassesClassIDMembersConflict
const PostClassesClassIDMembersConflictCode int = 409

/*PostClassesClassIDMembersConflict Resource already exists. An object creation was requested, but this object was already existing.


swagger:response postClassesClassIdMembersConflict
*/
type PostClassesClassIDMembersConflict struct {
}

// NewPostClassesClassIDMembersConflict creates PostClassesClassIDMembersConflict with default headers values
func NewPostClassesClassIDMembersConflict() *PostClassesClassIDMembersConflict {

	return &PostClassesClassIDMembersConflict{}
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

func (o *PostClassesClassIDMembersConflict) PostClassesClassIDMembersResponder() {}

// PostClassesClassIDMembersInternalServerErrorCode is the HTTP code returned for type PostClassesClassIDMembersInternalServerError
const PostClassesClassIDMembersInternalServerErrorCode int = 500

/*PostClassesClassIDMembersInternalServerError An internal error has occured. See error details.


swagger:response postClassesClassIdMembersInternalServerError
*/
type PostClassesClassIDMembersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostClassesClassIDMembersInternalServerError creates PostClassesClassIDMembersInternalServerError with default headers values
func NewPostClassesClassIDMembersInternalServerError() *PostClassesClassIDMembersInternalServerError {

	return &PostClassesClassIDMembersInternalServerError{}
}

// WithPayload adds the payload to the post classes class Id members internal server error response
func (o *PostClassesClassIDMembersInternalServerError) WithPayload(payload *models.APIError) *PostClassesClassIDMembersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post classes class Id members internal server error response
func (o *PostClassesClassIDMembersInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostClassesClassIDMembersInternalServerError) PostClassesClassIDMembersResponder() {}

/*PostClassesClassIDMembersDefault Other error. See error details.


swagger:response postClassesClassIdMembersDefault
*/
type PostClassesClassIDMembersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostClassesClassIDMembersDefault creates PostClassesClassIDMembersDefault with default headers values
func NewPostClassesClassIDMembersDefault(code int) *PostClassesClassIDMembersDefault {
	if code <= 0 {
		code = 500
	}

	return &PostClassesClassIDMembersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post classes class ID members default response
func (o *PostClassesClassIDMembersDefault) WithStatusCode(code int) *PostClassesClassIDMembersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post classes class ID members default response
func (o *PostClassesClassIDMembersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post classes class ID members default response
func (o *PostClassesClassIDMembersDefault) WithPayload(payload *models.APIError) *PostClassesClassIDMembersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post classes class ID members default response
func (o *PostClassesClassIDMembersDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostClassesClassIDMembersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostClassesClassIDMembersDefault) PostClassesClassIDMembersResponder() {}

type PostClassesClassIDMembersNotImplementedResponder struct {
	middleware.Responder
}

func (*PostClassesClassIDMembersNotImplementedResponder) PostClassesClassIDMembersResponder() {}

func PostClassesClassIDMembersNotImplemented() PostClassesClassIDMembersResponder {
	return &PostClassesClassIDMembersNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.PostClassesClassIDMembers has not yet been implemented",
		),
	}
}

type PostClassesClassIDMembersResponder interface {
	middleware.Responder
	PostClassesClassIDMembersResponder()
}
