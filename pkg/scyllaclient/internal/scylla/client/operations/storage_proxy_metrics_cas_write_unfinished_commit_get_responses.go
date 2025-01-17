// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// StorageProxyMetricsCasWriteUnfinishedCommitGetReader is a Reader for the StorageProxyMetricsCasWriteUnfinishedCommitGet structure.
type StorageProxyMetricsCasWriteUnfinishedCommitGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsCasWriteUnfinishedCommitGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsCasWriteUnfinishedCommitGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsCasWriteUnfinishedCommitGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsCasWriteUnfinishedCommitGetOK creates a StorageProxyMetricsCasWriteUnfinishedCommitGetOK with default headers values
func NewStorageProxyMetricsCasWriteUnfinishedCommitGetOK() *StorageProxyMetricsCasWriteUnfinishedCommitGetOK {
	return &StorageProxyMetricsCasWriteUnfinishedCommitGetOK{}
}

/*
StorageProxyMetricsCasWriteUnfinishedCommitGetOK handles this case with default header values.

StorageProxyMetricsCasWriteUnfinishedCommitGetOK storage proxy metrics cas write unfinished commit get o k
*/
type StorageProxyMetricsCasWriteUnfinishedCommitGetOK struct {
	Payload int32
}

func (o *StorageProxyMetricsCasWriteUnfinishedCommitGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyMetricsCasWriteUnfinishedCommitGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsCasWriteUnfinishedCommitGetDefault creates a StorageProxyMetricsCasWriteUnfinishedCommitGetDefault with default headers values
func NewStorageProxyMetricsCasWriteUnfinishedCommitGetDefault(code int) *StorageProxyMetricsCasWriteUnfinishedCommitGetDefault {
	return &StorageProxyMetricsCasWriteUnfinishedCommitGetDefault{
		_statusCode: code,
	}
}

/*
StorageProxyMetricsCasWriteUnfinishedCommitGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsCasWriteUnfinishedCommitGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics cas write unfinished commit get default response
func (o *StorageProxyMetricsCasWriteUnfinishedCommitGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsCasWriteUnfinishedCommitGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsCasWriteUnfinishedCommitGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsCasWriteUnfinishedCommitGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
