// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "pronto-hubble/api/v1/models"
)

// ListOKCode is the HTTP code returned for type ListOK
const ListOKCode int = 200

/*ListOK OK

swagger:response listOK
*/
type ListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.User `json:"body,omitempty"`
}

// NewListOK creates ListOK with default headers values
func NewListOK() *ListOK {

	return &ListOK{}
}

// WithPayload adds the payload to the list o k response
func (o *ListOK) WithPayload(payload []*models.User) *ListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list o k response
func (o *ListOK) SetPayload(payload []*models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.User, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*ListDefault error

swagger:response listDefault
*/
type ListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListDefault creates ListDefault with default headers values
func NewListDefault(code int) *ListDefault {
	if code <= 0 {
		code = 500
	}

	return &ListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list default response
func (o *ListDefault) WithStatusCode(code int) *ListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list default response
func (o *ListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list default response
func (o *ListDefault) WithPayload(payload *models.Error) *ListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list default response
func (o *ListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
