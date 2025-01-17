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

	strfmt "github.com/go-openapi/strfmt"
)

// NewColumnFamilyMetricsLiveSsTableCountGetParams creates a new ColumnFamilyMetricsLiveSsTableCountGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsLiveSsTableCountGetParams() *ColumnFamilyMetricsLiveSsTableCountGetParams {

	return &ColumnFamilyMetricsLiveSsTableCountGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsLiveSsTableCountGetParamsWithTimeout creates a new ColumnFamilyMetricsLiveSsTableCountGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsLiveSsTableCountGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsLiveSsTableCountGetParams {

	return &ColumnFamilyMetricsLiveSsTableCountGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsLiveSsTableCountGetParamsWithContext creates a new ColumnFamilyMetricsLiveSsTableCountGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsLiveSsTableCountGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsLiveSsTableCountGetParams {

	return &ColumnFamilyMetricsLiveSsTableCountGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsLiveSsTableCountGetParamsWithHTTPClient creates a new ColumnFamilyMetricsLiveSsTableCountGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsLiveSsTableCountGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsLiveSsTableCountGetParams {

	return &ColumnFamilyMetricsLiveSsTableCountGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsLiveSsTableCountGetParams contains all the parameters to send to the API endpoint
for the column family metrics live ss table count get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsLiveSsTableCountGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics live ss table count get params
func (o *ColumnFamilyMetricsLiveSsTableCountGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsLiveSsTableCountGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics live ss table count get params
func (o *ColumnFamilyMetricsLiveSsTableCountGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics live ss table count get params
func (o *ColumnFamilyMetricsLiveSsTableCountGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsLiveSsTableCountGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics live ss table count get params
func (o *ColumnFamilyMetricsLiveSsTableCountGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics live ss table count get params
func (o *ColumnFamilyMetricsLiveSsTableCountGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsLiveSsTableCountGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics live ss table count get params
func (o *ColumnFamilyMetricsLiveSsTableCountGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsLiveSsTableCountGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
