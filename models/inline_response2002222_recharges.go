// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2002Recharges inline response2002 recharges
// swagger:model inlineResponse2002Recharges
type InlineResponse2002222Recharges []*Recharge

// Validate validates this inline response2002 recharges
func (m InlineResponse2002222Recharges) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
