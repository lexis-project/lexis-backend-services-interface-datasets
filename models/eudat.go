// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Eudat eudat
//
// swagger:model Eudat
type Eudat struct {

	// e u d a t f i o
	EUDATFIO string `json:"EUDAT/FIO,omitempty"`

	// e u d a t f i x e d c o n t e n t
	EUDATFIXEDCONTENT string `json:"EUDAT/FIXED_CONTENT,omitempty"`

	// e u d a t p a r e n t
	EUDATPARENT string `json:"EUDAT/PARENT,omitempty"`

	// e u d a t r e p l i c a
	EUDATREPLICA string `json:"EUDAT/REPLICA,omitempty"`

	// e u d a t r o r
	EUDATROR string `json:"EUDAT/ROR,omitempty"`

	// p ID
	PID string `json:"PID,omitempty"`
}

// Validate validates this eudat
func (m *Eudat) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Eudat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Eudat) UnmarshalBinary(b []byte) error {
	var res Eudat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
