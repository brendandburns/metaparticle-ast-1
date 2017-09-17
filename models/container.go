// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Container container
// swagger:model container

type Container struct {

	// env
	Env []*EnvVar `json:"env"`

	// image
	// Required: true
	Image *string `json:"image"`
}

/* polymorph container env false */

/* polymorph container image false */

// Validate validates this container
func (m *Container) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Container) validateEnv(formats strfmt.Registry) error {

	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {

		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {

			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Container) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Container) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Container) UnmarshalBinary(b []byte) error {
	var res Container
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
