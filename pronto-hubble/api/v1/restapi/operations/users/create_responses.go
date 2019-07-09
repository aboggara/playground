// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "pronto-hubble/api/v1/models"
)

// CreateCreatedCode is the HTTP code returned for type CreateCreated
const CreateCreatedCode int = 201

/*CreateCreated Created

swagger:response createCreated
*/
type CreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ID `json:"body,omitempty"`
}

// NewCreateCreated creates CreateCreated with default headers values
func NewCreateCreated() *CreateCreated {

	return &CreateCreated{}
}

// WithPayload adds the payload to the create created response
func (o *CreateCreated) WithPayload(payload *models.ID) *CreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create created response
func (o *CreateCreated) SetPayload(payload *models.ID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateDefault error

swagger:response createDefault
*/
type CreateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDefault creates CreateDefault with default headers values
func NewCreateDefault(code int) *CreateDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create default response
func (o *CreateDefault) WithStatusCode(code int) *CreateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create default response
func (o *CreateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create default response
func (o *CreateDefault) WithPayload(payload *models.Error) *CreateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create default response
func (o *CreateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
