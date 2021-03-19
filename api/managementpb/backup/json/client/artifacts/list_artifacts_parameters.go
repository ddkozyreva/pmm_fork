// Code generated by go-swagger; DO NOT EDIT.

package artifacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListArtifactsParams creates a new ListArtifactsParams object
// with the default values initialized.
func NewListArtifactsParams() *ListArtifactsParams {
	var ()
	return &ListArtifactsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListArtifactsParamsWithTimeout creates a new ListArtifactsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListArtifactsParamsWithTimeout(timeout time.Duration) *ListArtifactsParams {
	var ()
	return &ListArtifactsParams{

		timeout: timeout,
	}
}

// NewListArtifactsParamsWithContext creates a new ListArtifactsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListArtifactsParamsWithContext(ctx context.Context) *ListArtifactsParams {
	var ()
	return &ListArtifactsParams{

		Context: ctx,
	}
}

// NewListArtifactsParamsWithHTTPClient creates a new ListArtifactsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListArtifactsParamsWithHTTPClient(client *http.Client) *ListArtifactsParams {
	var ()
	return &ListArtifactsParams{
		HTTPClient: client,
	}
}

/*ListArtifactsParams contains all the parameters to send to the API endpoint
for the list artifacts operation typically these are written to a http.Request
*/
type ListArtifactsParams struct {

	/*Body*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list artifacts params
func (o *ListArtifactsParams) WithTimeout(timeout time.Duration) *ListArtifactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list artifacts params
func (o *ListArtifactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list artifacts params
func (o *ListArtifactsParams) WithContext(ctx context.Context) *ListArtifactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list artifacts params
func (o *ListArtifactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list artifacts params
func (o *ListArtifactsParams) WithHTTPClient(client *http.Client) *ListArtifactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list artifacts params
func (o *ListArtifactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list artifacts params
func (o *ListArtifactsParams) WithBody(body interface{}) *ListArtifactsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list artifacts params
func (o *ListArtifactsParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListArtifactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
