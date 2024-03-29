// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "pronto-hubble/api/v1/models"
)

// RegisterDeviceCreatedCode is the HTTP code returned for type RegisterDeviceCreated
const RegisterDeviceCreatedCode int = 201

/*RegisterDeviceCreated Created

swagger:response registerDeviceCreated
*/
type RegisterDeviceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ID `json:"body,omitempty"`
}

// NewRegisterDeviceCreated creates RegisterDeviceCreated with default headers values
func NewRegisterDeviceCreated() *RegisterDeviceCreated {

	return &RegisterDeviceCreated{}
}

// WithPayload adds the payload to the register device created response
func (o *RegisterDeviceCreated) WithPayload(payload *models.ID) *RegisterDeviceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register device created response
func (o *RegisterDeviceCreated) SetPayload(payload *models.ID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterDeviceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*RegisterDeviceDefault error

swagger:response registerDeviceDefault
*/
type RegisterDeviceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterDeviceDefault creates RegisterDeviceDefault with default headers values
func NewRegisterDeviceDefault(code int) *RegisterDeviceDefault {
	if code <= 0 {
		code = 500
	}

	return &RegisterDeviceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the register device default response
func (o *RegisterDeviceDefault) WithStatusCode(code int) *RegisterDeviceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the register device default response
func (o *RegisterDeviceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the register device default response
func (o *RegisterDeviceDefault) WithPayload(payload *models.Error) *RegisterDeviceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register device default response
func (o *RegisterDeviceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterDeviceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
