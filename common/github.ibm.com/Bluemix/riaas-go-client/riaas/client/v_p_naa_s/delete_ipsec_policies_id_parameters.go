// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteIpsecPoliciesIDParams creates a new DeleteIpsecPoliciesIDParams object
// with the default values initialized.
func NewDeleteIpsecPoliciesIDParams() *DeleteIpsecPoliciesIDParams {
	var ()
	return &DeleteIpsecPoliciesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIpsecPoliciesIDParamsWithTimeout creates a new DeleteIpsecPoliciesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIpsecPoliciesIDParamsWithTimeout(timeout time.Duration) *DeleteIpsecPoliciesIDParams {
	var ()
	return &DeleteIpsecPoliciesIDParams{

		timeout: timeout,
	}
}

// NewDeleteIpsecPoliciesIDParamsWithContext creates a new DeleteIpsecPoliciesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIpsecPoliciesIDParamsWithContext(ctx context.Context) *DeleteIpsecPoliciesIDParams {
	var ()
	return &DeleteIpsecPoliciesIDParams{

		Context: ctx,
	}
}

// NewDeleteIpsecPoliciesIDParamsWithHTTPClient creates a new DeleteIpsecPoliciesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIpsecPoliciesIDParamsWithHTTPClient(client *http.Client) *DeleteIpsecPoliciesIDParams {
	var ()
	return &DeleteIpsecPoliciesIDParams{
		HTTPClient: client,
	}
}

/*DeleteIpsecPoliciesIDParams contains all the parameters to send to the API endpoint
for the delete ipsec policies ID operation typically these are written to a http.Request
*/
type DeleteIpsecPoliciesIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The IPsec policy identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) WithTimeout(timeout time.Duration) *DeleteIpsecPoliciesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) WithContext(ctx context.Context) *DeleteIpsecPoliciesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) WithHTTPClient(client *http.Client) *DeleteIpsecPoliciesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) WithGeneration(generation int64) *DeleteIpsecPoliciesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) WithID(id string) *DeleteIpsecPoliciesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) WithVersion(version string) *DeleteIpsecPoliciesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete ipsec policies ID params
func (o *DeleteIpsecPoliciesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIpsecPoliciesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}