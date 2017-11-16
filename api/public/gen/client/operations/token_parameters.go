// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTokenParams creates a new TokenParams object
// with the default values initialized.
func NewTokenParams() *TokenParams {
	var ()
	return &TokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenParamsWithTimeout creates a new TokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenParamsWithTimeout(timeout time.Duration) *TokenParams {
	var ()
	return &TokenParams{

		timeout: timeout,
	}
}

// NewTokenParamsWithContext creates a new TokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenParamsWithContext(ctx context.Context) *TokenParams {
	var ()
	return &TokenParams{

		Context: ctx,
	}
}

// NewTokenParamsWithHTTPClient creates a new TokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenParamsWithHTTPClient(client *http.Client) *TokenParams {
	var ()
	return &TokenParams{
		HTTPClient: client,
	}
}

/*TokenParams contains all the parameters to send to the API endpoint
for the token operation typically these are written to a http.Request
*/
type TokenParams struct {

	/*Jwt*/
	Jwt string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the token params
func (o *TokenParams) WithTimeout(timeout time.Duration) *TokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token params
func (o *TokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token params
func (o *TokenParams) WithContext(ctx context.Context) *TokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token params
func (o *TokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token params
func (o *TokenParams) WithHTTPClient(client *http.Client) *TokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token params
func (o *TokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJwt adds the jwt to the token params
func (o *TokenParams) WithJwt(jwt string) *TokenParams {
	o.SetJwt(jwt)
	return o
}

// SetJwt adds the jwt to the token params
func (o *TokenParams) SetJwt(jwt string) {
	o.Jwt = jwt
}

// WriteToRequest writes these params to a swagger request
func (o *TokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param jwt
	qrJwt := o.Jwt
	qJwt := qrJwt
	if qJwt != "" {
		if err := r.SetQueryParam("jwt", qJwt); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}