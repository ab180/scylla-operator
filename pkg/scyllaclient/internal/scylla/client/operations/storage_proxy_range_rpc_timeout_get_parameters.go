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

// NewStorageProxyRangeRPCTimeoutGetParams creates a new StorageProxyRangeRPCTimeoutGetParams object
// with the default values initialized.
func NewStorageProxyRangeRPCTimeoutGetParams() *StorageProxyRangeRPCTimeoutGetParams {

	return &StorageProxyRangeRPCTimeoutGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyRangeRPCTimeoutGetParamsWithTimeout creates a new StorageProxyRangeRPCTimeoutGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyRangeRPCTimeoutGetParamsWithTimeout(timeout time.Duration) *StorageProxyRangeRPCTimeoutGetParams {

	return &StorageProxyRangeRPCTimeoutGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyRangeRPCTimeoutGetParamsWithContext creates a new StorageProxyRangeRPCTimeoutGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyRangeRPCTimeoutGetParamsWithContext(ctx context.Context) *StorageProxyRangeRPCTimeoutGetParams {

	return &StorageProxyRangeRPCTimeoutGetParams{

		Context: ctx,
	}
}

// NewStorageProxyRangeRPCTimeoutGetParamsWithHTTPClient creates a new StorageProxyRangeRPCTimeoutGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyRangeRPCTimeoutGetParamsWithHTTPClient(client *http.Client) *StorageProxyRangeRPCTimeoutGetParams {

	return &StorageProxyRangeRPCTimeoutGetParams{
		HTTPClient: client,
	}
}

/*
StorageProxyRangeRPCTimeoutGetParams contains all the parameters to send to the API endpoint
for the storage proxy range Rpc timeout get operation typically these are written to a http.Request
*/
type StorageProxyRangeRPCTimeoutGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy range Rpc timeout get params
func (o *StorageProxyRangeRPCTimeoutGetParams) WithTimeout(timeout time.Duration) *StorageProxyRangeRPCTimeoutGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy range Rpc timeout get params
func (o *StorageProxyRangeRPCTimeoutGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy range Rpc timeout get params
func (o *StorageProxyRangeRPCTimeoutGetParams) WithContext(ctx context.Context) *StorageProxyRangeRPCTimeoutGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy range Rpc timeout get params
func (o *StorageProxyRangeRPCTimeoutGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy range Rpc timeout get params
func (o *StorageProxyRangeRPCTimeoutGetParams) WithHTTPClient(client *http.Client) *StorageProxyRangeRPCTimeoutGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy range Rpc timeout get params
func (o *StorageProxyRangeRPCTimeoutGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyRangeRPCTimeoutGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
