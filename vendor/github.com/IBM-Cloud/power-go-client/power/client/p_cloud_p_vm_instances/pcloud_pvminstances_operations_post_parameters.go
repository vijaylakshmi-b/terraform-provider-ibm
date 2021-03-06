// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPvminstancesOperationsPostParams creates a new PcloudPvminstancesOperationsPostParams object
// with the default values initialized.
func NewPcloudPvminstancesOperationsPostParams() *PcloudPvminstancesOperationsPostParams {
	var ()
	return &PcloudPvminstancesOperationsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesOperationsPostParamsWithTimeout creates a new PcloudPvminstancesOperationsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPvminstancesOperationsPostParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesOperationsPostParams {
	var ()
	return &PcloudPvminstancesOperationsPostParams{

		timeout: timeout,
	}
}

// NewPcloudPvminstancesOperationsPostParamsWithContext creates a new PcloudPvminstancesOperationsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPvminstancesOperationsPostParamsWithContext(ctx context.Context) *PcloudPvminstancesOperationsPostParams {
	var ()
	return &PcloudPvminstancesOperationsPostParams{

		Context: ctx,
	}
}

// NewPcloudPvminstancesOperationsPostParamsWithHTTPClient creates a new PcloudPvminstancesOperationsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPvminstancesOperationsPostParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesOperationsPostParams {
	var ()
	return &PcloudPvminstancesOperationsPostParams{
		HTTPClient: client,
	}
}

/*PcloudPvminstancesOperationsPostParams contains all the parameters to send to the API endpoint
for the pcloud pvminstances operations post operation typically these are written to a http.Request
*/
type PcloudPvminstancesOperationsPostParams struct {

	/*Body
	  Parameters for the desired operations

	*/
	Body *models.PVMInstanceOperation
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PvmInstanceID
	  PCloud PVM Instance ID

	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesOperationsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) WithContext(ctx context.Context) *PcloudPvminstancesOperationsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesOperationsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) WithBody(body *models.PVMInstanceOperation) *PcloudPvminstancesOperationsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) SetBody(body *models.PVMInstanceOperation) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesOperationsPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesOperationsPostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances operations post params
func (o *PcloudPvminstancesOperationsPostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesOperationsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
