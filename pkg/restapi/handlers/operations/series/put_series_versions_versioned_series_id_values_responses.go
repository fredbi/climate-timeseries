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

// PutSeriesVersionsVersionedSeriesIDValuesNoContentCode is the HTTP code returned for type PutSeriesVersionsVersionedSeriesIDValuesNoContent
const PutSeriesVersionsVersionedSeriesIDValuesNoContentCode int = 204

/*PutSeriesVersionsVersionedSeriesIDValuesNoContent Time series values successfully updated.

Check the response headers to retrieve this resource.


swagger:response putSeriesVersionsVersionedSeriesIdValuesNoContent
*/
type PutSeriesVersionsVersionedSeriesIDValuesNoContent struct {
	/*The URI of the updated resource.


	 */
	Location strfmt.URI `json:"Location"`
	/*The URI of the updated resource.


	 */
	XLocationVersion strfmt.URI `json:"X-LocationVersion"`
	/*The ID of the updated resource.


	 */
	XVersionID int64 `json:"X-VersionID"`
}

// NewPutSeriesVersionsVersionedSeriesIDValuesNoContent creates PutSeriesVersionsVersionedSeriesIDValuesNoContent with default headers values
func NewPutSeriesVersionsVersionedSeriesIDValuesNoContent() *PutSeriesVersionsVersionedSeriesIDValuesNoContent {

	return &PutSeriesVersionsVersionedSeriesIDValuesNoContent{}
}

// WithLocation adds the location to the put series versions versioned series Id values no content response
func (o *PutSeriesVersionsVersionedSeriesIDValuesNoContent) WithLocation(location strfmt.URI) *PutSeriesVersionsVersionedSeriesIDValuesNoContent {
	o.Location = location
	return o
}

// SetLocation sets the location to the put series versions versioned series Id values no content response
func (o *PutSeriesVersionsVersionedSeriesIDValuesNoContent) SetLocation(location strfmt.URI) {
	o.Location = location
}

// WithXLocationVersion adds the xLocationVersion to the put series versions versioned series Id values no content response
func (o *PutSeriesVersionsVersionedSeriesIDValuesNoContent) WithXLocationVersion(xLocationVersion strfmt.URI) *PutSeriesVersionsVersionedSeriesIDValuesNoContent {
	o.XLocationVersion = xLocationVersion
	return o
}

// SetXLocationVersion sets the xLocationVersion to the put series versions versioned series Id values no content response
func (o *PutSeriesVersionsVersionedSeriesIDValuesNoContent) SetXLocationVersion(xLocationVersion strfmt.URI) {
	o.XLocationVersion = xLocationVersion
}

// WithXVersionID adds the xVersionId to the put series versions versioned series Id values no content response
func (o *PutSeriesVersionsVersionedSeriesIDValuesNoContent) WithXVersionID(xVersionID int64) *PutSeriesVersionsVersionedSeriesIDValuesNoContent {
	o.XVersionID = xVersionID
	return o
}

// SetXVersionID sets the xVersionId to the put series versions versioned series Id values no content response
func (o *PutSeriesVersionsVersionedSeriesIDValuesNoContent) SetXVersionID(xVersionID int64) {
	o.XVersionID = xVersionID
}

// WriteResponse to the client
func (o *PutSeriesVersionsVersionedSeriesIDValuesNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

	xVersionID := swag.FormatInt64(o.XVersionID)
	if xVersionID != "" {
		rw.Header().Set("X-VersionID", xVersionID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *PutSeriesVersionsVersionedSeriesIDValuesNoContent) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// PutSeriesVersionsVersionedSeriesIDValuesBadRequestCode is the HTTP code returned for type PutSeriesVersionsVersionedSeriesIDValuesBadRequest
const PutSeriesVersionsVersionedSeriesIDValuesBadRequestCode int = 400

/*PutSeriesVersionsVersionedSeriesIDValuesBadRequest Client error in request. Input did not pass validations. See error details.


swagger:response putSeriesVersionsVersionedSeriesIdValuesBadRequest
*/
type PutSeriesVersionsVersionedSeriesIDValuesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPutSeriesVersionsVersionedSeriesIDValuesBadRequest creates PutSeriesVersionsVersionedSeriesIDValuesBadRequest with default headers values
func NewPutSeriesVersionsVersionedSeriesIDValuesBadRequest() *PutSeriesVersionsVersionedSeriesIDValuesBadRequest {

	return &PutSeriesVersionsVersionedSeriesIDValuesBadRequest{}
}

// WithPayload adds the payload to the put series versions versioned series Id values bad request response
func (o *PutSeriesVersionsVersionedSeriesIDValuesBadRequest) WithPayload(payload *models.APIError) *PutSeriesVersionsVersionedSeriesIDValuesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put series versions versioned series Id values bad request response
func (o *PutSeriesVersionsVersionedSeriesIDValuesBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSeriesVersionsVersionedSeriesIDValuesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PutSeriesVersionsVersionedSeriesIDValuesBadRequest) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// PutSeriesVersionsVersionedSeriesIDValuesUnauthorizedCode is the HTTP code returned for type PutSeriesVersionsVersionedSeriesIDValuesUnauthorized
const PutSeriesVersionsVersionedSeriesIDValuesUnauthorizedCode int = 401

/*PutSeriesVersionsVersionedSeriesIDValuesUnauthorized Unauthorized access for a lack of authentication


swagger:response putSeriesVersionsVersionedSeriesIdValuesUnauthorized
*/
type PutSeriesVersionsVersionedSeriesIDValuesUnauthorized struct {
}

// NewPutSeriesVersionsVersionedSeriesIDValuesUnauthorized creates PutSeriesVersionsVersionedSeriesIDValuesUnauthorized with default headers values
func NewPutSeriesVersionsVersionedSeriesIDValuesUnauthorized() *PutSeriesVersionsVersionedSeriesIDValuesUnauthorized {

	return &PutSeriesVersionsVersionedSeriesIDValuesUnauthorized{}
}

// WriteResponse to the client
func (o *PutSeriesVersionsVersionedSeriesIDValuesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

func (o *PutSeriesVersionsVersionedSeriesIDValuesUnauthorized) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// PutSeriesVersionsVersionedSeriesIDValuesForbiddenCode is the HTTP code returned for type PutSeriesVersionsVersionedSeriesIDValuesForbidden
const PutSeriesVersionsVersionedSeriesIDValuesForbiddenCode int = 403

/*PutSeriesVersionsVersionedSeriesIDValuesForbidden Forbidden access for a lack of sufficient privileges


swagger:response putSeriesVersionsVersionedSeriesIdValuesForbidden
*/
type PutSeriesVersionsVersionedSeriesIDValuesForbidden struct {
}

// NewPutSeriesVersionsVersionedSeriesIDValuesForbidden creates PutSeriesVersionsVersionedSeriesIDValuesForbidden with default headers values
func NewPutSeriesVersionsVersionedSeriesIDValuesForbidden() *PutSeriesVersionsVersionedSeriesIDValuesForbidden {

	return &PutSeriesVersionsVersionedSeriesIDValuesForbidden{}
}

// WriteResponse to the client
func (o *PutSeriesVersionsVersionedSeriesIDValuesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *PutSeriesVersionsVersionedSeriesIDValuesForbidden) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// PutSeriesVersionsVersionedSeriesIDValuesNotFoundCode is the HTTP code returned for type PutSeriesVersionsVersionedSeriesIDValuesNotFound
const PutSeriesVersionsVersionedSeriesIDValuesNotFoundCode int = 404

/*PutSeriesVersionsVersionedSeriesIDValuesNotFound Resource not found. The object requested does not exist or is not visible to the user.


swagger:response putSeriesVersionsVersionedSeriesIdValuesNotFound
*/
type PutSeriesVersionsVersionedSeriesIDValuesNotFound struct {
}

// NewPutSeriesVersionsVersionedSeriesIDValuesNotFound creates PutSeriesVersionsVersionedSeriesIDValuesNotFound with default headers values
func NewPutSeriesVersionsVersionedSeriesIDValuesNotFound() *PutSeriesVersionsVersionedSeriesIDValuesNotFound {

	return &PutSeriesVersionsVersionedSeriesIDValuesNotFound{}
}

// WriteResponse to the client
func (o *PutSeriesVersionsVersionedSeriesIDValuesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *PutSeriesVersionsVersionedSeriesIDValuesNotFound) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowedCode is the HTTP code returned for type PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed
const PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowedCode int = 405

/*PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed MethodNotAllowed


swagger:response putSeriesVersionsVersionedSeriesIdValuesMethodNotAllowed
*/
type PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed struct {
}

// NewPutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed creates PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed with default headers values
func NewPutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed() *PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed {

	return &PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed{}
}

// WriteResponse to the client
func (o *PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

func (o *PutSeriesVersionsVersionedSeriesIDValuesMethodNotAllowed) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

// PutSeriesVersionsVersionedSeriesIDValuesInternalServerErrorCode is the HTTP code returned for type PutSeriesVersionsVersionedSeriesIDValuesInternalServerError
const PutSeriesVersionsVersionedSeriesIDValuesInternalServerErrorCode int = 500

/*PutSeriesVersionsVersionedSeriesIDValuesInternalServerError An internal error has occured. See error details.


swagger:response putSeriesVersionsVersionedSeriesIdValuesInternalServerError
*/
type PutSeriesVersionsVersionedSeriesIDValuesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPutSeriesVersionsVersionedSeriesIDValuesInternalServerError creates PutSeriesVersionsVersionedSeriesIDValuesInternalServerError with default headers values
func NewPutSeriesVersionsVersionedSeriesIDValuesInternalServerError() *PutSeriesVersionsVersionedSeriesIDValuesInternalServerError {

	return &PutSeriesVersionsVersionedSeriesIDValuesInternalServerError{}
}

// WithPayload adds the payload to the put series versions versioned series Id values internal server error response
func (o *PutSeriesVersionsVersionedSeriesIDValuesInternalServerError) WithPayload(payload *models.APIError) *PutSeriesVersionsVersionedSeriesIDValuesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put series versions versioned series Id values internal server error response
func (o *PutSeriesVersionsVersionedSeriesIDValuesInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSeriesVersionsVersionedSeriesIDValuesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PutSeriesVersionsVersionedSeriesIDValuesInternalServerError) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

/*PutSeriesVersionsVersionedSeriesIDValuesDefault Other error. See error details.


swagger:response putSeriesVersionsVersionedSeriesIdValuesDefault
*/
type PutSeriesVersionsVersionedSeriesIDValuesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPutSeriesVersionsVersionedSeriesIDValuesDefault creates PutSeriesVersionsVersionedSeriesIDValuesDefault with default headers values
func NewPutSeriesVersionsVersionedSeriesIDValuesDefault(code int) *PutSeriesVersionsVersionedSeriesIDValuesDefault {
	if code <= 0 {
		code = 500
	}

	return &PutSeriesVersionsVersionedSeriesIDValuesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put series versions versioned series ID values default response
func (o *PutSeriesVersionsVersionedSeriesIDValuesDefault) WithStatusCode(code int) *PutSeriesVersionsVersionedSeriesIDValuesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put series versions versioned series ID values default response
func (o *PutSeriesVersionsVersionedSeriesIDValuesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put series versions versioned series ID values default response
func (o *PutSeriesVersionsVersionedSeriesIDValuesDefault) WithPayload(payload *models.APIError) *PutSeriesVersionsVersionedSeriesIDValuesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put series versions versioned series ID values default response
func (o *PutSeriesVersionsVersionedSeriesIDValuesDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSeriesVersionsVersionedSeriesIDValuesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PutSeriesVersionsVersionedSeriesIDValuesDefault) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

type PutSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder struct {
	middleware.Responder
}

func (*PutSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder) PutSeriesVersionsVersionedSeriesIDValuesResponder() {
}

func PutSeriesVersionsVersionedSeriesIDValuesNotImplemented() PutSeriesVersionsVersionedSeriesIDValuesResponder {
	return &PutSeriesVersionsVersionedSeriesIDValuesNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.PutSeriesVersionsVersionedSeriesIDValues has not yet been implemented",
		),
	}
}

type PutSeriesVersionsVersionedSeriesIDValuesResponder interface {
	middleware.Responder
	PutSeriesVersionsVersionedSeriesIDValuesResponder()
}