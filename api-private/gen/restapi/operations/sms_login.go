// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

// SmsLoginHandlerFunc turns a function with the right signature into a sms login handler
type SmsLoginHandlerFunc func(SmsLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SmsLoginHandlerFunc) Handle(params SmsLoginParams) middleware.Responder {
	return fn(params)
}

// SmsLoginHandler interface for that can handle valid sms login params
type SmsLoginHandler interface {
	Handle(SmsLoginParams) middleware.Responder
}

// NewSmsLogin creates a new http.Handler for the sms login operation
func NewSmsLogin(ctx *middleware.Context, handler SmsLoginHandler) *SmsLogin {
	return &SmsLogin{Context: ctx, Handler: handler}
}

/*SmsLogin swagger:route POST /smsLogin smsLogin

sms login

*/
type SmsLogin struct {
	Context *middleware.Context
	Handler SmsLoginHandler
}

func (o *SmsLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSmsLoginParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	zap.L().Named("api").Info("SmsLogin", zap.Any("request", &Params))

	res := o.Handler.Handle(Params) // actually handle the request

	zap.L().Named("api").Info("SmsLogin", zap.Any("response", &res))

	o.Context.Respond(rw, r, route.Produces, route, res)

}
