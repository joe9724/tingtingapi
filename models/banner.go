// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Banner 大图
// swagger:model Banner
type Banner struct {

	// cover
	Cover string `json:"cover,omitempty"`
	// type
	Type int64 `json:"type,omitempty"`
	// 跳转类型(0=大类 1=小类 2=专辑 3=书本 4=章节 5=站内wap 6=第三方链接)
	Name string `json:"name,omitempty"`
	// 备用字段
	ExtraInfo1 int64 `json:"extraInfo1,omitempty"`

	// 备用字段
	ExtraInfo2 string `json:"extraInfo2,omitempty"`

	// 备用字段
	ExtraInfo3 string `json:"extraInfo3,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

    Jumpid int64 `json:"jumpid"`

	// order
	Order int64 `json:"order,omitempty"`
	


}

// Validate validates this banner
func (m *Banner) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Banner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Banner) UnmarshalBinary(b []byte) error {
	var res Banner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
