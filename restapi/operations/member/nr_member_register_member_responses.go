// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"tingtingapi/models"
)

// MemberRegisterMemberOKCode is the HTTP code returned for type MemberRegisterMemberOK
const MemberRegisterMemberOKCode int = 200

/*MemberRegisterMemberOK 登录成功，返回登录信息

swagger:response memberRegisterMemberOK
*/
type MemberRegisterMemberOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20016 `json:"body,omitempty"`
}

// NewMemberRegisterMemberOK creates MemberRegisterMemberOK with default headers values
func NewMemberRegisterMemberOK() *MemberRegisterMemberOK {
	return &MemberRegisterMemberOK{}
}

// WithPayload adds the payload to the member register member o k response
func (o *MemberRegisterMemberOK) WithPayload(payload *models.InlineResponse20016) *MemberRegisterMemberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member register member o k response
func (o *MemberRegisterMemberOK) SetPayload(payload *models.InlineResponse20016) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberRegisterMemberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrMemberRegisterMemberDefault error

swagger:response memberRegisterMemberDefault
*/
type NrMemberRegisterMemberDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrMemberRegisterMemberDefault creates NrMemberRegisterMemberDefault with default headers values
func NewNrMemberRegisterMemberDefault(code int) *NrMemberRegisterMemberDefault {
	if code <= 0 {
		code = 500
	}

	return &NrMemberRegisterMemberDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member register member default response
func (o *NrMemberRegisterMemberDefault) WithStatusCode(code int) *NrMemberRegisterMemberDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member register member default response
func (o *NrMemberRegisterMemberDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member register member default response
func (o *NrMemberRegisterMemberDefault) WithPayload(payload *models.Error) *NrMemberRegisterMemberDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member register member default response
func (o *NrMemberRegisterMemberDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrMemberRegisterMemberDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}