// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// AddRDSExporterReader is a Reader for the AddRDSExporter structure.
type AddRDSExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRDSExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddRDSExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddRDSExporterOK creates a AddRDSExporterOK with default headers values
func NewAddRDSExporterOK() *AddRDSExporterOK {
	return &AddRDSExporterOK{}
}

/*AddRDSExporterOK handles this case with default header values.

A successful response.
*/
type AddRDSExporterOK struct {
	Payload *AddRDSExporterOKBody
}

func (o *AddRDSExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddRDSExporter][%d] addRdsExporterOK  %+v", 200, o.Payload)
}

func (o *AddRDSExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRDSExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddRDSExporterBody add RDS exporter body
swagger:model AddRDSExporterBody
*/
type AddRDSExporterBody struct {

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// A list of Service identifiers (Node identifiers are extracted from Services).
	ServiceIds []string `json:"service_ids"`
}

// Validate validates this add RDS exporter body
func (o *AddRDSExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSExporterBody) UnmarshalBinary(b []byte) error {
	var res AddRDSExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSExporterOKBody add RDS exporter o k body
swagger:model AddRDSExporterOKBody
*/
type AddRDSExporterOKBody struct {

	// rds exporter
	RDSExporter *AddRDSExporterOKBodyRDSExporter `json:"rds_exporter,omitempty"`
}

// Validate validates this add RDS exporter o k body
func (o *AddRDSExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRDSExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRDSExporterOKBody) validateRDSExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.RDSExporter) { // not required
		return nil
	}

	if o.RDSExporter != nil {
		if err := o.RDSExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRdsExporterOK" + "." + "rds_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddRDSExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSExporterOKBodyRDSExporter RDSExporter runs on Generic or Container Node and exposes RemoteAmazonRDS Node and AmazonRDSMySQL Service metrics.
swagger:model AddRDSExporterOKBodyRDSExporter
*/
type AddRDSExporterOKBodyRDSExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// A list of Service identifiers (Node identifiers are extracted from Services).
	ServiceIds []string `json:"service_ids"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this add RDS exporter o k body RDS exporter
func (o *AddRDSExporterOKBodyRDSExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addRdsExporterOKBodyRdsExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addRdsExporterOKBodyRdsExporterTypeStatusPropEnum = append(addRdsExporterOKBodyRdsExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddRDSExporterOKBodyRDSExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddRDSExporterOKBodyRDSExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddRDSExporterOKBodyRDSExporterStatusSTARTING captures enum value "STARTING"
	AddRDSExporterOKBodyRDSExporterStatusSTARTING string = "STARTING"

	// AddRDSExporterOKBodyRDSExporterStatusRUNNING captures enum value "RUNNING"
	AddRDSExporterOKBodyRDSExporterStatusRUNNING string = "RUNNING"

	// AddRDSExporterOKBodyRDSExporterStatusWAITING captures enum value "WAITING"
	AddRDSExporterOKBodyRDSExporterStatusWAITING string = "WAITING"

	// AddRDSExporterOKBodyRDSExporterStatusSTOPPING captures enum value "STOPPING"
	AddRDSExporterOKBodyRDSExporterStatusSTOPPING string = "STOPPING"

	// AddRDSExporterOKBodyRDSExporterStatusDONE captures enum value "DONE"
	AddRDSExporterOKBodyRDSExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddRDSExporterOKBodyRDSExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addRdsExporterOKBodyRdsExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddRDSExporterOKBodyRDSExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addRdsExporterOK"+"."+"rds_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSExporterOKBodyRDSExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSExporterOKBodyRDSExporter) UnmarshalBinary(b []byte) error {
	var res AddRDSExporterOKBodyRDSExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
