// Code generated by go-swagger; DO NOT EDIT.

package debugrequest

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

// NewGetDebugRequestParams creates a new GetDebugRequestParams object
// with the default values initialized.
func NewGetDebugRequestParams() *GetDebugRequestParams {
	var ()
	return &GetDebugRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDebugRequestParamsWithTimeout creates a new GetDebugRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDebugRequestParamsWithTimeout(timeout time.Duration) *GetDebugRequestParams {
	var ()
	return &GetDebugRequestParams{

		timeout: timeout,
	}
}

// NewGetDebugRequestParamsWithContext creates a new GetDebugRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDebugRequestParamsWithContext(ctx context.Context) *GetDebugRequestParams {
	var ()
	return &GetDebugRequestParams{

		Context: ctx,
	}
}

// NewGetDebugRequestParamsWithHTTPClient creates a new GetDebugRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDebugRequestParamsWithHTTPClient(client *http.Client) *GetDebugRequestParams {
	var ()
	return &GetDebugRequestParams{
		HTTPClient: client,
	}
}

/*GetDebugRequestParams contains all the parameters to send to the API endpoint
for the get debug request operation typically these are written to a http.Request
*/
type GetDebugRequestParams struct {

	/*DebugRequestID
	  ID of config to return

	*/
	DebugRequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get debug request params
func (o *GetDebugRequestParams) WithTimeout(timeout time.Duration) *GetDebugRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get debug request params
func (o *GetDebugRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get debug request params
func (o *GetDebugRequestParams) WithContext(ctx context.Context) *GetDebugRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get debug request params
func (o *GetDebugRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get debug request params
func (o *GetDebugRequestParams) WithHTTPClient(client *http.Client) *GetDebugRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get debug request params
func (o *GetDebugRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDebugRequestID adds the debugRequestID to the get debug request params
func (o *GetDebugRequestParams) WithDebugRequestID(debugRequestID string) *GetDebugRequestParams {
	o.SetDebugRequestID(debugRequestID)
	return o
}

// SetDebugRequestID adds the debugRequestId to the get debug request params
func (o *GetDebugRequestParams) SetDebugRequestID(debugRequestID string) {
	o.DebugRequestID = debugRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDebugRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param debugRequestId
	if err := r.SetPathParam("debugRequestId", o.DebugRequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
