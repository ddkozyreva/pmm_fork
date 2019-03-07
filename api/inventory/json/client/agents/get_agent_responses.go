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

// GetAgentReader is a Reader for the GetAgent structure.
type GetAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAgentOK creates a GetAgentOK with default headers values
func NewGetAgentOK() *GetAgentOK {
	return &GetAgentOK{}
}

/*GetAgentOK handles this case with default header values.

A successful response.
*/
type GetAgentOK struct {
	Payload *GetAgentOKBody
}

func (o *GetAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/Get][%d] getAgentOK  %+v", 200, o.Payload)
}

func (o *GetAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAgentBody get agent body
swagger:model GetAgentBody
*/
type GetAgentBody struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`
}

// Validate validates this get agent body
func (o *GetAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentBody) UnmarshalBinary(b []byte) error {
	var res GetAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBody get agent o k body
swagger:model GetAgentOKBody
*/
type GetAgentOKBody struct {

	// external exporter
	ExternalExporter *GetAgentOKBodyExternalExporter `json:"external_exporter,omitempty"`

	// mongodb exporter
	MongodbExporter *GetAgentOKBodyMongodbExporter `json:"mongodb_exporter,omitempty"`

	// mysqld exporter
	MysqldExporter *GetAgentOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`

	// node exporter
	NodeExporter *GetAgentOKBodyNodeExporter `json:"node_exporter,omitempty"`

	// pmm agent
	PMMAgent *GetAgentOKBodyPMMAgent `json:"pmm_agent,omitempty"`

	// qan mysql perfschema agent
	QANMysqlPerfschemaAgent *GetAgentOKBodyQANMysqlPerfschemaAgent `json:"qan_mysql_perfschema_agent,omitempty"`

	// rds exporter
	RDSExporter *GetAgentOKBodyRDSExporter `json:"rds_exporter,omitempty"`
}

// Validate validates this get agent o k body
func (o *GetAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternalExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongodbExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNodeExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMMAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANMysqlPerfschemaAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRDSExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAgentOKBody) validateExternalExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalExporter) { // not required
		return nil
	}

	if o.ExternalExporter != nil {
		if err := o.ExternalExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "external_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *GetAgentOKBody) validateMongodbExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MongodbExporter) { // not required
		return nil
	}

	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *GetAgentOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *GetAgentOKBody) validateNodeExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeExporter) { // not required
		return nil
	}

	if o.NodeExporter != nil {
		if err := o.NodeExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "node_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *GetAgentOKBody) validatePMMAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMAgent) { // not required
		return nil
	}

	if o.PMMAgent != nil {
		if err := o.PMMAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "pmm_agent")
			}
			return err
		}
	}

	return nil
}

func (o *GetAgentOKBody) validateQANMysqlPerfschemaAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlPerfschemaAgent) { // not required
		return nil
	}

	if o.QANMysqlPerfschemaAgent != nil {
		if err := o.QANMysqlPerfschemaAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "qan_mysql_perfschema_agent")
			}
			return err
		}
	}

	return nil
}

func (o *GetAgentOKBody) validateRDSExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.RDSExporter) { // not required
		return nil
	}

	if o.RDSExporter != nil {
		if err := o.RDSExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "rds_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBody) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyExternalExporter ExternalExporter does not run on any Inventory Node.
swagger:model GetAgentOKBodyExternalExporter
*/
type GetAgentOKBodyExternalExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// URL for scraping metrics.
	MetricsURL string `json:"metrics_url,omitempty"`
}

// Validate validates this get agent o k body external exporter
func (o *GetAgentOKBodyExternalExporter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyExternalExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyExternalExporter) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyExternalExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyMongodbExporter MongoDBExporter runs on Generic or Container Node and exposes MongoDB Service metrics.
swagger:model GetAgentOKBodyMongodbExporter
*/
type GetAgentOKBodyMongodbExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// MongoDB URI for scraping metrics. (See https://docs.mongodb.com/manual/reference/connection-string/)
	ConnectionString string `json:"connection_string,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this get agent o k body mongodb exporter
func (o *GetAgentOKBodyMongodbExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getAgentOKBodyMongodbExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAgentOKBodyMongodbExporterTypeStatusPropEnum = append(getAgentOKBodyMongodbExporterTypeStatusPropEnum, v)
	}
}

const (

	// GetAgentOKBodyMongodbExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	GetAgentOKBodyMongodbExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// GetAgentOKBodyMongodbExporterStatusSTARTING captures enum value "STARTING"
	GetAgentOKBodyMongodbExporterStatusSTARTING string = "STARTING"

	// GetAgentOKBodyMongodbExporterStatusRUNNING captures enum value "RUNNING"
	GetAgentOKBodyMongodbExporterStatusRUNNING string = "RUNNING"

	// GetAgentOKBodyMongodbExporterStatusWAITING captures enum value "WAITING"
	GetAgentOKBodyMongodbExporterStatusWAITING string = "WAITING"

	// GetAgentOKBodyMongodbExporterStatusSTOPPING captures enum value "STOPPING"
	GetAgentOKBodyMongodbExporterStatusSTOPPING string = "STOPPING"

	// GetAgentOKBodyMongodbExporterStatusDONE captures enum value "DONE"
	GetAgentOKBodyMongodbExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *GetAgentOKBodyMongodbExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getAgentOKBodyMongodbExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetAgentOKBodyMongodbExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("getAgentOK"+"."+"mongodb_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyMongodbExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyMongodbExporter) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyMongodbExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyMysqldExporter MySQLdExporter runs on Generic or Container Node and exposes MySQL and AmazonRDSMySQL Service metrics.
swagger:model GetAgentOKBodyMysqldExporter
*/
type GetAgentOKBodyMysqldExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this get agent o k body mysqld exporter
func (o *GetAgentOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getAgentOKBodyMysqldExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAgentOKBodyMysqldExporterTypeStatusPropEnum = append(getAgentOKBodyMysqldExporterTypeStatusPropEnum, v)
	}
}

const (

	// GetAgentOKBodyMysqldExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	GetAgentOKBodyMysqldExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// GetAgentOKBodyMysqldExporterStatusSTARTING captures enum value "STARTING"
	GetAgentOKBodyMysqldExporterStatusSTARTING string = "STARTING"

	// GetAgentOKBodyMysqldExporterStatusRUNNING captures enum value "RUNNING"
	GetAgentOKBodyMysqldExporterStatusRUNNING string = "RUNNING"

	// GetAgentOKBodyMysqldExporterStatusWAITING captures enum value "WAITING"
	GetAgentOKBodyMysqldExporterStatusWAITING string = "WAITING"

	// GetAgentOKBodyMysqldExporterStatusSTOPPING captures enum value "STOPPING"
	GetAgentOKBodyMysqldExporterStatusSTOPPING string = "STOPPING"

	// GetAgentOKBodyMysqldExporterStatusDONE captures enum value "DONE"
	GetAgentOKBodyMysqldExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *GetAgentOKBodyMysqldExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getAgentOKBodyMysqldExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetAgentOKBodyMysqldExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("getAgentOK"+"."+"mysqld_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyNodeExporter NodeExporter runs on Generic on Container Node and exposes its metrics.
swagger:model GetAgentOKBodyNodeExporter
*/
type GetAgentOKBodyNodeExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this get agent o k body node exporter
func (o *GetAgentOKBodyNodeExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getAgentOKBodyNodeExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAgentOKBodyNodeExporterTypeStatusPropEnum = append(getAgentOKBodyNodeExporterTypeStatusPropEnum, v)
	}
}

const (

	// GetAgentOKBodyNodeExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	GetAgentOKBodyNodeExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// GetAgentOKBodyNodeExporterStatusSTARTING captures enum value "STARTING"
	GetAgentOKBodyNodeExporterStatusSTARTING string = "STARTING"

	// GetAgentOKBodyNodeExporterStatusRUNNING captures enum value "RUNNING"
	GetAgentOKBodyNodeExporterStatusRUNNING string = "RUNNING"

	// GetAgentOKBodyNodeExporterStatusWAITING captures enum value "WAITING"
	GetAgentOKBodyNodeExporterStatusWAITING string = "WAITING"

	// GetAgentOKBodyNodeExporterStatusSTOPPING captures enum value "STOPPING"
	GetAgentOKBodyNodeExporterStatusSTOPPING string = "STOPPING"

	// GetAgentOKBodyNodeExporterStatusDONE captures enum value "DONE"
	GetAgentOKBodyNodeExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *GetAgentOKBodyNodeExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getAgentOKBodyNodeExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetAgentOKBodyNodeExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("getAgentOK"+"."+"node_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyNodeExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyNodeExporter) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyNodeExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyPMMAgent PMMAgent runs on Generic on Container Node.
swagger:model GetAgentOKBodyPMMAgent
*/
type GetAgentOKBodyPMMAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// True if Agent is running and connected to pmm-managed.
	Connected bool `json:"connected,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`
}

// Validate validates this get agent o k body PMM agent
func (o *GetAgentOKBodyPMMAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyPMMAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyPMMAgent) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyPMMAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyQANMysqlPerfschemaAgent QANMySQLPerfSchemaAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model GetAgentOKBodyQANMysqlPerfschemaAgent
*/
type GetAgentOKBodyQANMysqlPerfschemaAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this get agent o k body QAN mysql perfschema agent
func (o *GetAgentOKBodyQANMysqlPerfschemaAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyQANMysqlPerfschemaAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyQANMysqlPerfschemaAgent) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyQANMysqlPerfschemaAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyRDSExporter RDSExporter runs on Generic or Container Node and exposes RemoteAmazonRDS Node and AmazonRDSMySQL Service metrics.
swagger:model GetAgentOKBodyRDSExporter
*/
type GetAgentOKBodyRDSExporter struct {

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

// Validate validates this get agent o k body RDS exporter
func (o *GetAgentOKBodyRDSExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getAgentOKBodyRdsExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAgentOKBodyRdsExporterTypeStatusPropEnum = append(getAgentOKBodyRdsExporterTypeStatusPropEnum, v)
	}
}

const (

	// GetAgentOKBodyRDSExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	GetAgentOKBodyRDSExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// GetAgentOKBodyRDSExporterStatusSTARTING captures enum value "STARTING"
	GetAgentOKBodyRDSExporterStatusSTARTING string = "STARTING"

	// GetAgentOKBodyRDSExporterStatusRUNNING captures enum value "RUNNING"
	GetAgentOKBodyRDSExporterStatusRUNNING string = "RUNNING"

	// GetAgentOKBodyRDSExporterStatusWAITING captures enum value "WAITING"
	GetAgentOKBodyRDSExporterStatusWAITING string = "WAITING"

	// GetAgentOKBodyRDSExporterStatusSTOPPING captures enum value "STOPPING"
	GetAgentOKBodyRDSExporterStatusSTOPPING string = "STOPPING"

	// GetAgentOKBodyRDSExporterStatusDONE captures enum value "DONE"
	GetAgentOKBodyRDSExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *GetAgentOKBodyRDSExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getAgentOKBodyRdsExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetAgentOKBodyRDSExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("getAgentOK"+"."+"rds_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyRDSExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyRDSExporter) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyRDSExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
