// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// PostSeriesSeriesIDCreatedCode is the HTTP code returned for type PostSeriesSeriesIDCreated
const PostSeriesSeriesIDCreatedCode int = 201

/*PostSeriesSeriesIDCreated Series version successfully created.

Check the response headers to retrieve this resource.


swagger:response postSeriesSeriesIdCreated
*/
type PostSeriesSeriesIDCreated struct {
	/*The URI of the newly created versioned resource.


	 */
	Location strfmt.URI `json:"Location"`
	/*The URI of the newly created versioned resource.


	 */
	XLocationVersion strfmt.URI `json:"X-LocationVersion"`
	/*The UUID of the newly created versioned resource.


	 */
	XVersionID strfmt.UUID `json:"X-VersionID"`
}

// NewPostSeriesSeriesIDCreated creates PostSeriesSeriesIDCreated with default headers values
func NewPostSeriesSeriesIDCreated() *PostSeriesSeriesIDCreated {

	return &PostSeriesSeriesIDCreated{}
}

// WithLocation adds the location to the post series series Id created response
func (o *PostSeriesSeriesIDCreated) WithLocation(location strfmt.URI) *PostSeriesSeriesIDCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the post series series Id created response
func (o *PostSeriesSeriesIDCreated) SetLocation(location strfmt.URI) {
	o.Location = location
}

// WithXLocationVersion adds the xLocationVersion to the post series series Id created response
func (o *PostSeriesSeriesIDCreated) WithXLocationVersion(xLocationVersion strfmt.URI) *PostSeriesSeriesIDCreated {
	o.XLocationVersion = xLocationVersion
	return o
}

// SetXLocationVersion sets the xLocationVersion to the post series series Id created response
func (o *PostSeriesSeriesIDCreated) SetXLocationVersion(xLocationVersion strfmt.URI) {
	o.XLocationVersion = xLocationVersion
}

// WithXVersionID adds the xVersionId to the post series series Id created response
func (o *PostSeriesSeriesIDCreated) WithXVersionID(xVersionID strfmt.UUID) *PostSeriesSeriesIDCreated {
	o.XVersionID = xVersionID
	return o
}

// SetXVersionID sets the xVersionId to the post series series Id created response
func (o *PostSeriesSeriesIDCreated) SetXVersionID(xVersionID strfmt.UUID) {
	o.XVersionID = xVersionID
}

// WriteResponse to the client
func (o *PostSeriesSeriesIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location.String()
	if location != "" {
		rw.Header().Set("Location", location)
	}

	// response header X-LocationVersion

	xLocationVersion := o.XLocationVersion.String()
	if xLocationVersion != "" {
		rw.Header().Set("X-LocationVersion", xLocationVersion)
	}

	// response header X-VersionID

	xVersionID := o.XVersionID.String()
	if xVersionID != "" {
		rw.Header().Set("X-VersionID", xVersionID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

func (o *PostSeriesSeriesIDCreated) PostSeriesSeriesIDResponder() {}

// PostSeriesSeriesIDBadRequestCode is the HTTP code returned for type PostSeriesSeriesIDBadRequest
const PostSeriesSeriesIDBadRequestCode int = 400

/*PostSeriesSeriesIDBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response postSeriesSeriesIdBadRequest
*/
type PostSeriesSeriesIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostSeriesSeriesIDBadRequest creates PostSeriesSeriesIDBadRequest with default headers values
func NewPostSeriesSeriesIDBadRequest() *PostSeriesSeriesIDBadRequest {

	return &PostSeriesSeriesIDBadRequest{}
}

// WithPayload adds the payload to the post series series Id bad request response
func (o *PostSeriesSeriesIDBadRequest) WithPayload(payload *models.APIError) *PostSeriesSeriesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post series series Id bad request response
func (o *PostSeriesSeriesIDBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSeriesSeriesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostSeriesSeriesIDBadRequest) PostSeriesSeriesIDResponder() {}

// PostSeriesSeriesIDUnauthorizedCode is the HTTP code returned for type PostSeriesSeriesIDUnauthorized
const PostSeriesSeriesIDUnauthorizedCode int = 401

/*PostSeriesSeriesIDUnauthorized Unauthorized access for a lack of authentication


swagger:response postSeriesSeriesIdUnauthorized
*/
type PostSeriesSeriesIDUnauthorized struct {
}

// NewPostSeriesSeriesIDUnauthorized creates PostSeriesSeriesIDUnauthorized with default headers values
func NewPostSeriesSeriesIDUnauthorized() *PostSeriesSeriesIDUnauthorized {

	return &PostSeriesSeriesIDUnauthorized{}
}

// WriteResponse to the client
func (o *PostSeriesSeriesIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *PostSeriesSeriesIDUnauthorized) PostSeriesSeriesIDResponder() {}

// PostSeriesSeriesIDForbiddenCode is the HTTP code returned for type PostSeriesSeriesIDForbidden
const PostSeriesSeriesIDForbiddenCode int = 403

/*PostSeriesSeriesIDForbidden Forbidden access for a lack of sufficient privileges


swagger:response postSeriesSeriesIdForbidden
*/
type PostSeriesSeriesIDForbidden struct {
}

// NewPostSeriesSeriesIDForbidden creates PostSeriesSeriesIDForbidden with default headers values
func NewPostSeriesSeriesIDForbidden() *PostSeriesSeriesIDForbidden {

	return &PostSeriesSeriesIDForbidden{}
}

// WriteResponse to the client
func (o *PostSeriesSeriesIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *PostSeriesSeriesIDForbidden) PostSeriesSeriesIDResponder() {}

// PostSeriesSeriesIDNotFoundCode is the HTTP code returned for type PostSeriesSeriesIDNotFound
const PostSeriesSeriesIDNotFoundCode int = 404

/*PostSeriesSeriesIDNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response postSeriesSeriesIdNotFound
*/
type PostSeriesSeriesIDNotFound struct {
}

// NewPostSeriesSeriesIDNotFound creates PostSeriesSeriesIDNotFound with default headers values
func NewPostSeriesSeriesIDNotFound() *PostSeriesSeriesIDNotFound {

	return &PostSeriesSeriesIDNotFound{}
}

// WriteResponse to the client
func (o *PostSeriesSeriesIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *PostSeriesSeriesIDNotFound) PostSeriesSeriesIDResponder() {}

// PostSeriesSeriesIDMethodNotAllowedCode is the HTTP code returned for type PostSeriesSeriesIDMethodNotAllowed
const PostSeriesSeriesIDMethodNotAllowedCode int = 405

/*PostSeriesSeriesIDMethodNotAllowed MethodNotAllowed


swagger:response postSeriesSeriesIdMethodNotAllowed
*/
type PostSeriesSeriesIDMethodNotAllowed struct {
}

// NewPostSeriesSeriesIDMethodNotAllowed creates PostSeriesSeriesIDMethodNotAllowed with default headers values
func NewPostSeriesSeriesIDMethodNotAllowed() *PostSeriesSeriesIDMethodNotAllowed {

	return &PostSeriesSeriesIDMethodNotAllowed{}
}

// WriteResponse to the client
func (o *PostSeriesSeriesIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *PostSeriesSeriesIDMethodNotAllowed) PostSeriesSeriesIDResponder() {}

// PostSeriesSeriesIDInternalServerErrorCode is the HTTP code returned for type PostSeriesSeriesIDInternalServerError
const PostSeriesSeriesIDInternalServerErrorCode int = 500

/*PostSeriesSeriesIDInternalServerError An internal error has occured. See error details.


swagger:response postSeriesSeriesIdInternalServerError
*/
type PostSeriesSeriesIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostSeriesSeriesIDInternalServerError creates PostSeriesSeriesIDInternalServerError with default headers values
func NewPostSeriesSeriesIDInternalServerError() *PostSeriesSeriesIDInternalServerError {

	return &PostSeriesSeriesIDInternalServerError{}
}

// WithPayload adds the payload to the post series series Id internal server error response
func (o *PostSeriesSeriesIDInternalServerError) WithPayload(payload *models.APIError) *PostSeriesSeriesIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post series series Id internal server error response
func (o *PostSeriesSeriesIDInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSeriesSeriesIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostSeriesSeriesIDInternalServerError) PostSeriesSeriesIDResponder() {}

/*PostSeriesSeriesIDDefault Other error. See error details.


swagger:response postSeriesSeriesIdDefault
*/
type PostSeriesSeriesIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostSeriesSeriesIDDefault creates PostSeriesSeriesIDDefault with default headers values
func NewPostSeriesSeriesIDDefault(code int) *PostSeriesSeriesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PostSeriesSeriesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post series series ID default response
func (o *PostSeriesSeriesIDDefault) WithStatusCode(code int) *PostSeriesSeriesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post series series ID default response
func (o *PostSeriesSeriesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post series series ID default response
func (o *PostSeriesSeriesIDDefault) WithPayload(payload *models.APIError) *PostSeriesSeriesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post series series ID default response
func (o *PostSeriesSeriesIDDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSeriesSeriesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PostSeriesSeriesIDDefault) PostSeriesSeriesIDResponder() {}

type PostSeriesSeriesIDNotImplementedResponder struct {
	middleware.Responder
}

func (*PostSeriesSeriesIDNotImplementedResponder) PostSeriesSeriesIDResponder() {}

func PostSeriesSeriesIDNotImplemented() PostSeriesSeriesIDResponder {
	return &PostSeriesSeriesIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.PostSeriesSeriesID has not yet been implemented",
		),
	}
}

type PostSeriesSeriesIDResponder interface {
	middleware.Responder
	PostSeriesSeriesIDResponder()
}