// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronGroup/Account/api/private/gen/models"
)

// SmsSignupOKCode is the HTTP code returned for type SmsSignupOK
const SmsSignupOKCode int = 200

/*SmsSignupOK ok

swagger:response smsSignupOK
*/
type SmsSignupOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewSmsSignupOK creates SmsSignupOK with default headers values
func NewSmsSignupOK() *SmsSignupOK {
	return &SmsSignupOK{}
}

// WithPayload adds the payload to the sms signup o k response
func (o *SmsSignupOK) WithPayload(payload *models.LoginResponse) *SmsSignupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sms signup o k response
func (o *SmsSignupOK) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SmsSignupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SmsSignupDefault Error response

swagger:response smsSignupDefault
*/
type SmsSignupDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.SmsSignupDefaultBody `json:"body,omitempty"`
}

// NewSmsSignupDefault creates SmsSignupDefault with default headers values
func NewSmsSignupDefault(code int) *SmsSignupDefault {
	if code <= 0 {
		code = 500
	}

	return &SmsSignupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the sms signup default response
func (o *SmsSignupDefault) WithStatusCode(code int) *SmsSignupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the sms signup default response
func (o *SmsSignupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the sms signup default response
func (o *SmsSignupDefault) WithPayload(payload *models.SmsSignupDefaultBody) *SmsSignupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sms signup default response
func (o *SmsSignupDefault) SetPayload(payload *models.SmsSignupDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SmsSignupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
