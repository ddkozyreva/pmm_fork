// Code generated by go-swagger; DO NOT EDIT.

package server_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetReadOnlySettingsReader is a Reader for the GetReadOnlySettings structure.
type GetReadOnlySettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReadOnlySettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReadOnlySettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetReadOnlySettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReadOnlySettingsOK creates a GetReadOnlySettingsOK with default headers values
func NewGetReadOnlySettingsOK() *GetReadOnlySettingsOK {
	return &GetReadOnlySettingsOK{}
}

/*
GetReadOnlySettingsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetReadOnlySettingsOK struct {
	Payload *GetReadOnlySettingsOKBody
}

// IsSuccess returns true when this get read only settings Ok response has a 2xx status code
func (o *GetReadOnlySettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get read only settings Ok response has a 3xx status code
func (o *GetReadOnlySettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get read only settings Ok response has a 4xx status code
func (o *GetReadOnlySettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get read only settings Ok response has a 5xx status code
func (o *GetReadOnlySettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get read only settings Ok response a status code equal to that given
func (o *GetReadOnlySettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get read only settings Ok response
func (o *GetReadOnlySettingsOK) Code() int {
	return 200
}

func (o *GetReadOnlySettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/server/settings/readonly][%d] getReadOnlySettingsOk %s", 200, payload)
}

func (o *GetReadOnlySettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/server/settings/readonly][%d] getReadOnlySettingsOk %s", 200, payload)
}

func (o *GetReadOnlySettingsOK) GetPayload() *GetReadOnlySettingsOKBody {
	return o.Payload
}

func (o *GetReadOnlySettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetReadOnlySettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReadOnlySettingsDefault creates a GetReadOnlySettingsDefault with default headers values
func NewGetReadOnlySettingsDefault(code int) *GetReadOnlySettingsDefault {
	return &GetReadOnlySettingsDefault{
		_statusCode: code,
	}
}

/*
GetReadOnlySettingsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetReadOnlySettingsDefault struct {
	_statusCode int

	Payload *GetReadOnlySettingsDefaultBody
}

// IsSuccess returns true when this get read only settings default response has a 2xx status code
func (o *GetReadOnlySettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get read only settings default response has a 3xx status code
func (o *GetReadOnlySettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get read only settings default response has a 4xx status code
func (o *GetReadOnlySettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get read only settings default response has a 5xx status code
func (o *GetReadOnlySettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get read only settings default response a status code equal to that given
func (o *GetReadOnlySettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get read only settings default response
func (o *GetReadOnlySettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetReadOnlySettingsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/server/settings/readonly][%d] GetReadOnlySettings default %s", o._statusCode, payload)
}

func (o *GetReadOnlySettingsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/server/settings/readonly][%d] GetReadOnlySettings default %s", o._statusCode, payload)
}

func (o *GetReadOnlySettingsDefault) GetPayload() *GetReadOnlySettingsDefaultBody {
	return o.Payload
}

func (o *GetReadOnlySettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetReadOnlySettingsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetReadOnlySettingsDefaultBody get read only settings default body
swagger:model GetReadOnlySettingsDefaultBody
*/
type GetReadOnlySettingsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetReadOnlySettingsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get read only settings default body
func (o *GetReadOnlySettingsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReadOnlySettingsDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetReadOnlySettings default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetReadOnlySettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get read only settings default body based on the context it is used
func (o *GetReadOnlySettingsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReadOnlySettingsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetReadOnlySettings default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetReadOnlySettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetReadOnlySettingsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReadOnlySettingsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetReadOnlySettingsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReadOnlySettingsDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//     // or ...
//     if (any.isSameTypeAs(Foo.getDefaultInstance())) {
//       foo = any.unpack(Foo.getDefaultInstance());
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := anypb.New(foo)
//      if err != nil {
//        ...
//      }
//      ...
//      foo := &pb.Foo{}
//      if err := any.UnmarshalTo(foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model GetReadOnlySettingsDefaultBodyDetailsItems0
*/
type GetReadOnlySettingsDefaultBodyDetailsItems0 struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com. As of May 2023, there are no widely used type server
	// implementations and no plans to implement one.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	AtType string `json:"@type,omitempty"`

	// get read only settings default body details items0
	GetReadOnlySettingsDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *GetReadOnlySettingsDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// A URL/resource name that uniquely identifies the type of the serialized
		// protocol buffer message. This string must contain at least
		// one "/" character. The last segment of the URL's path must represent
		// the fully qualified name of the type (as in
		// `path/google.protobuf.Duration`). The name should be in a canonical form
		// (e.g., leading "." is not accepted).
		//
		// In practice, teams usually precompile into the binary all types that they
		// expect it to use in the context of Any. However, for URLs which use the
		// scheme `http`, `https`, or no scheme, one can optionally set up a type
		// server that maps type URLs to message definitions as follows:
		//
		// * If no scheme is provided, `https` is assumed.
		// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
		//   value in binary format, or produce an error.
		// * Applications are allowed to cache lookup results based on the
		//   URL, or have them precompiled into a binary to avoid any
		//   lookup. Therefore, binary compatibility needs to be preserved
		//   on changes to types. (Use versioned type names to manage
		//   breaking changes.)
		//
		// Note: this functionality is not currently available in the official
		// protobuf release, and it is not used for type URLs beginning with
		// type.googleapis.com. As of May 2023, there are no widely used type server
		// implementations and no plans to implement one.
		//
		// Schemes other than `http`, `https` (or the empty scheme) might be
		// used with implementation specific semantics.
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv GetReadOnlySettingsDefaultBodyDetailsItems0

	rcv.AtType = stage1.AtType
	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "@type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		o.GetReadOnlySettingsDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o GetReadOnlySettingsDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {
		// A URL/resource name that uniquely identifies the type of the serialized
		// protocol buffer message. This string must contain at least
		// one "/" character. The last segment of the URL's path must represent
		// the fully qualified name of the type (as in
		// `path/google.protobuf.Duration`). The name should be in a canonical form
		// (e.g., leading "." is not accepted).
		//
		// In practice, teams usually precompile into the binary all types that they
		// expect it to use in the context of Any. However, for URLs which use the
		// scheme `http`, `https`, or no scheme, one can optionally set up a type
		// server that maps type URLs to message definitions as follows:
		//
		// * If no scheme is provided, `https` is assumed.
		// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
		//   value in binary format, or produce an error.
		// * Applications are allowed to cache lookup results based on the
		//   URL, or have them precompiled into a binary to avoid any
		//   lookup. Therefore, binary compatibility needs to be preserved
		//   on changes to types. (Use versioned type names to manage
		//   breaking changes.)
		//
		// Note: this functionality is not currently available in the official
		// protobuf release, and it is not used for type URLs beginning with
		// type.googleapis.com. As of May 2023, there are no widely used type server
		// implementations and no plans to implement one.
		//
		// Schemes other than `http`, `https` (or the empty scheme) might be
		// used with implementation specific semantics.
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = o.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.GetReadOnlySettingsDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.GetReadOnlySettingsDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this get read only settings default body details items0
func (o *GetReadOnlySettingsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get read only settings default body details items0 based on context it is used
func (o *GetReadOnlySettingsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetReadOnlySettingsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReadOnlySettingsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetReadOnlySettingsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReadOnlySettingsOKBody get read only settings OK body
swagger:model GetReadOnlySettingsOKBody
*/
type GetReadOnlySettingsOKBody struct {
	// settings
	Settings *GetReadOnlySettingsOKBodySettings `json:"settings,omitempty"`
}

// Validate validates this get read only settings OK body
func (o *GetReadOnlySettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReadOnlySettingsOKBody) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if o.Settings != nil {
		if err := o.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getReadOnlySettingsOk" + "." + "settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getReadOnlySettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get read only settings OK body based on the context it is used
func (o *GetReadOnlySettingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReadOnlySettingsOKBody) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {
	if o.Settings != nil {

		if swag.IsZero(o.Settings) { // not required
			return nil
		}

		if err := o.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getReadOnlySettingsOk" + "." + "settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getReadOnlySettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetReadOnlySettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReadOnlySettingsOKBody) UnmarshalBinary(b []byte) error {
	var res GetReadOnlySettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReadOnlySettingsOKBodySettings ReadOnlySettings represents a stripped-down version of PMM Server settings that can be accessed by users of all roles.
swagger:model GetReadOnlySettingsOKBodySettings
*/
type GetReadOnlySettingsOKBodySettings struct {
	// True if updates are enabled.
	UpdatesEnabled bool `json:"updates_enabled,omitempty"`

	// True if telemetry is enabled.
	TelemetryEnabled bool `json:"telemetry_enabled,omitempty"`

	// True if Advisor is enabled.
	AdvisorEnabled bool `json:"advisor_enabled,omitempty"`

	// True if Alerting is enabled.
	AlertingEnabled bool `json:"alerting_enabled,omitempty"`

	// PMM Server public address.
	PMMPublicAddress string `json:"pmm_public_address,omitempty"`

	// True if Backup Management is enabled.
	BackupManagementEnabled bool `json:"backup_management_enabled,omitempty"`

	// True if Azure Discover is enabled.
	AzurediscoverEnabled bool `json:"azurediscover_enabled,omitempty"`

	// True if Access Control is enabled.
	EnableAccessControl bool `json:"enable_access_control,omitempty"`
}

// Validate validates this get read only settings OK body settings
func (o *GetReadOnlySettingsOKBodySettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get read only settings OK body settings based on context it is used
func (o *GetReadOnlySettingsOKBodySettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetReadOnlySettingsOKBodySettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReadOnlySettingsOKBodySettings) UnmarshalBinary(b []byte) error {
	var res GetReadOnlySettingsOKBodySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
