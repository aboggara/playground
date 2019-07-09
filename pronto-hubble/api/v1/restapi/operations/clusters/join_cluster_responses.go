// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "pronto-hubble/api/v1/models"
)

// JoinClusterOKCode is the HTTP code returned for type JoinClusterOK
const JoinClusterOKCode int = 200

/*JoinClusterOK OK

swagger:response joinClusterOK
*/
type JoinClusterOK struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewJoinClusterOK creates JoinClusterOK with default headers values
func NewJoinClusterOK() *JoinClusterOK {

	return &JoinClusterOK{}
}

// WithPayload adds the payload to the join cluster o k response
func (o *JoinClusterOK) WithPayload(payload *models.Cluster) *JoinClusterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the join cluster o k response
func (o *JoinClusterOK) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *JoinClusterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*JoinClusterDefault error

swagger:response joinClusterDefault
*/
type JoinClusterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewJoinClusterDefault creates JoinClusterDefault with default headers values
func NewJoinClusterDefault(code int) *JoinClusterDefault {
	if code <= 0 {
		code = 500
	}

	return &JoinClusterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the join cluster default response
func (o *JoinClusterDefault) WithStatusCode(code int) *JoinClusterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the join cluster default response
func (o *JoinClusterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the join cluster default response
func (o *JoinClusterDefault) WithPayload(payload *models.Error) *JoinClusterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the join cluster default response
func (o *JoinClusterDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *JoinClusterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
