// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LoginRet login ret
// swagger:model LoginRet
type LoginRet struct {
	ID int64 `json:"id"`
	// cover
	Area string `json:"area"`

	// icon
	Avatar string `json:"avatar"`

	// id
	Gender int64 `json:"gender"`

	// name
	Grade string `json:"grade"`

	// order
	Level int64 `json:"level"`

	Name string `json:"name"`

	Phone string `json:"phone"`

	Tags string `json:"tags"`

	// web Url
	Password string `json:"password"`

	Money int64 `json:"money"`

	Ts  int64 `json:"ts"`

	Status  int64 `json:"status"`

	Ts1  string `json:"ts1"`

	Birth string `json:"birth"`
}

// Validate validates this login ret
func (m *LoginRet) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *LoginRet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginRet) UnmarshalBinary(b []byte) error {
	var res LoginRet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
