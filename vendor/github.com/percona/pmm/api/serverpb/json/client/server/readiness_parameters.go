// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewReadinessParams creates a new ReadinessParams object
// with the default values initialized.
func NewReadinessParams() *ReadinessParams {

	return &ReadinessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadinessParamsWithTimeout creates a new ReadinessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadinessParamsWithTimeout(timeout time.Duration) *ReadinessParams {

	return &ReadinessParams{

		timeout: timeout,
	}
}

// NewReadinessParamsWithContext creates a new ReadinessParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadinessParamsWithContext(ctx context.Context) *ReadinessParams {

	return &ReadinessParams{

		Context: ctx,
	}
}

// NewReadinessParamsWithHTTPClient creates a new ReadinessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadinessParamsWithHTTPClient(client *http.Client) *ReadinessParams {

	return &ReadinessParams{
		HTTPClient: client,
	}
}

/*ReadinessParams contains all the parameters to send to the API endpoint
for the readiness operation typically these are written to a http.Request
*/
type ReadinessParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the readiness params
func (o *ReadinessParams) WithTimeout(timeout time.Duration) *ReadinessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the readiness params
func (o *ReadinessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the readiness params
func (o *ReadinessParams) WithContext(ctx context.Context) *ReadinessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the readiness params
func (o *ReadinessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the readiness params
func (o *ReadinessParams) WithHTTPClient(client *http.Client) *ReadinessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the readiness params
func (o *ReadinessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadinessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
