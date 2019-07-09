// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "pronto-hubble/api/v1/models"
)

// DeleteDeviceNoContentCode is the HTTP code returned for type DeleteDeviceNoContent
const DeleteDeviceNoContentCode int = 204

/*DeleteDeviceNoContent Deleted

swagger:response deleteDeviceNoContent
*/
type DeleteDeviceNoContent struct {
}

// NewDeleteDeviceNoContent creates DeleteDeviceNoContent with default headers values
func NewDeleteDeviceNoContent() *DeleteDeviceNoContent {

	return &DeleteDeviceNoContent{}
}

// WriteResponse to the client
func (o *DeleteDeviceNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteDeviceDefault error

swagger:response deleteDeviceDefault
*/
type DeleteDeviceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDeviceDefault creates DeleteDeviceDefault with default headers values
func NewDeleteDeviceDefault(code int) *DeleteDeviceDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteDeviceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete device default response
func (o *DeleteDeviceDefault) WithStatusCode(code int) *DeleteDeviceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete device default response
func (o *DeleteDeviceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete device default response
func (o *DeleteDeviceDefault) WithPayload(payload *models.Error) *DeleteDeviceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete device default response
func (o *DeleteDeviceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeviceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}