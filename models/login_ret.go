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
	// avatar
	Avatar string `json:"avatar"`
	Name string `json:"name"`
	// gender
	Gender int64 `json:"gender"`

	// member Id
	MemberID int64 `json:"memberId"`
	// membername
	Membername string `json:"memberName"`
	//
	Area string `json:"area"`
	Birth string `json:"birth"`
	Grade string `json:"grade"`
	Level int64 `json:"id"`
	Phone string `json:"phone"`
	Tags string `json:"tags"`
	Status int64 `json:"memberName"`
	Ts  int64 `json:"ts"`
	Money int64 `json:"money"`
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
