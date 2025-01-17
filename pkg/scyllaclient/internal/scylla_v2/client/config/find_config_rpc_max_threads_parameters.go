// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigRPCMaxThreadsParams creates a new FindConfigRPCMaxThreadsParams object
// with the default values initialized.
func NewFindConfigRPCMaxThreadsParams() *FindConfigRPCMaxThreadsParams {

	return &FindConfigRPCMaxThreadsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigRPCMaxThreadsParamsWithTimeout creates a new FindConfigRPCMaxThreadsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigRPCMaxThreadsParamsWithTimeout(timeout time.Duration) *FindConfigRPCMaxThreadsParams {

	return &FindConfigRPCMaxThreadsParams{

		timeout: timeout,
	}
}

// NewFindConfigRPCMaxThreadsParamsWithContext creates a new FindConfigRPCMaxThreadsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigRPCMaxThreadsParamsWithContext(ctx context.Context) *FindConfigRPCMaxThreadsParams {

	return &FindConfigRPCMaxThreadsParams{

		Context: ctx,
	}
}

// NewFindConfigRPCMaxThreadsParamsWithHTTPClient creates a new FindConfigRPCMaxThreadsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigRPCMaxThreadsParamsWithHTTPClient(client *http.Client) *FindConfigRPCMaxThreadsParams {

	return &FindConfigRPCMaxThreadsParams{
		HTTPClient: client,
	}
}

/*
FindConfigRPCMaxThreadsParams contains all the parameters to send to the API endpoint
for the find config rpc max threads operation typically these are written to a http.Request
*/
type FindConfigRPCMaxThreadsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config rpc max threads params
func (o *FindConfigRPCMaxThreadsParams) WithTimeout(timeout time.Duration) *FindConfigRPCMaxThreadsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config rpc max threads params
func (o *FindConfigRPCMaxThreadsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config rpc max threads params
func (o *FindConfigRPCMaxThreadsParams) WithContext(ctx context.Context) *FindConfigRPCMaxThreadsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config rpc max threads params
func (o *FindConfigRPCMaxThreadsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config rpc max threads params
func (o *FindConfigRPCMaxThreadsParams) WithHTTPClient(client *http.Client) *FindConfigRPCMaxThreadsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config rpc max threads params
func (o *FindConfigRPCMaxThreadsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigRPCMaxThreadsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
