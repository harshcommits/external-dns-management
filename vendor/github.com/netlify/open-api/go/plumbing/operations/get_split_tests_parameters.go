// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetSplitTestsParams creates a new GetSplitTestsParams object
// with the default values initialized.
func NewGetSplitTestsParams() *GetSplitTestsParams {
	var ()
	return &GetSplitTestsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSplitTestsParamsWithTimeout creates a new GetSplitTestsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSplitTestsParamsWithTimeout(timeout time.Duration) *GetSplitTestsParams {
	var ()
	return &GetSplitTestsParams{

		timeout: timeout,
	}
}

// NewGetSplitTestsParamsWithContext creates a new GetSplitTestsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSplitTestsParamsWithContext(ctx context.Context) *GetSplitTestsParams {
	var ()
	return &GetSplitTestsParams{

		Context: ctx,
	}
}

// NewGetSplitTestsParamsWithHTTPClient creates a new GetSplitTestsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSplitTestsParamsWithHTTPClient(client *http.Client) *GetSplitTestsParams {
	var ()
	return &GetSplitTestsParams{
		HTTPClient: client,
	}
}

/*GetSplitTestsParams contains all the parameters to send to the API endpoint
for the get split tests operation typically these are written to a http.Request
*/
type GetSplitTestsParams struct {

	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get split tests params
func (o *GetSplitTestsParams) WithTimeout(timeout time.Duration) *GetSplitTestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get split tests params
func (o *GetSplitTestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get split tests params
func (o *GetSplitTestsParams) WithContext(ctx context.Context) *GetSplitTestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get split tests params
func (o *GetSplitTestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get split tests params
func (o *GetSplitTestsParams) WithHTTPClient(client *http.Client) *GetSplitTestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get split tests params
func (o *GetSplitTestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the get split tests params
func (o *GetSplitTestsParams) WithSiteID(siteID string) *GetSplitTestsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get split tests params
func (o *GetSplitTestsParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSplitTestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
