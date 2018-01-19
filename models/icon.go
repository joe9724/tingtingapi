// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Icon icon
// swagger:model Icon
type Icon struct {

	// cover
	Cover string `json:"cover"`

	// icon
	Icon string `json:"icon,omitempty"`

	// id
	ID int64 `json:"id"`

	// name
	Name string `json:"name"`

	// order
	Order int64 `json:"order,omitempty"`

	// 跳转类型(0=大类 1=小类 2=专辑 3=书本 4=章节 5=站内wap 6=第三方链接)
	Type int64 `json:"type"`

	// web Url
	WebURL string `json:"webUrl"`

	TargetId string `json:"targetId"`
}

// Validate validates this icon
func (m *Icon) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Icon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Icon) UnmarshalBinary(b []byte) error {
	var res Icon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
