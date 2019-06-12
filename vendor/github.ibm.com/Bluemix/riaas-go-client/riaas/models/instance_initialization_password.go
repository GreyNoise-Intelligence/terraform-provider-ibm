// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InstanceInitializationPassword InstanceEncryptedPassword
// swagger:model instanceInitializationPassword
type InstanceInitializationPassword struct {

	// The administrator password at initialization, encrypted using the first initialized SSH key
	// Format: byte
	EncryptedPassword strfmt.Base64 `json:"encrypted_password,omitempty"`

	// encryption key
	EncryptionKey *InstanceInitializationPasswordEncryptionKey `json:"encryption_key,omitempty"`
}

// Validate validates this instance initialization password
func (m *InstanceInitializationPassword) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryptedPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceInitializationPassword) validateEncryptedPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.EncryptedPassword) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *InstanceInitializationPassword) validateEncryptionKey(formats strfmt.Registry) error {

	if swag.IsZero(m.EncryptionKey) { // not required
		return nil
	}

	if m.EncryptionKey != nil {
		if err := m.EncryptionKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceInitializationPassword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceInitializationPassword) UnmarshalBinary(b []byte) error {
	var res InstanceInitializationPassword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}