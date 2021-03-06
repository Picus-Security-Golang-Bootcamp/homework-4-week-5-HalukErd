// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaginationResponse pagination response
//
// swagger:model PaginationResponse
type PaginationResponse struct {

	// Current page index
	Page int32 `json:"page,omitempty"`

	// Page Length
	PageLength int32 `json:"pageLength,omitempty"`

	// Page sort type
	Sort string `json:"sort,omitempty"`

	// Total page size
	TotalPages int32 `json:"totalPages,omitempty"`
}

// Validate validates this pagination response
func (m *PaginationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pagination response based on context it is used
func (m *PaginationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationResponse) UnmarshalBinary(b []byte) error {
	var res PaginationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
