// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (

)

// Icon icon
// swagger:model Icon
type Fav_Album struct {
	ID int64 `json:"id,omitempty"`
	// cover
	AlbumId int64 `json:"albumId,omitempty"`

	MemberId int64 `json:"memberId,omitempty"`
	// id
	Status int64 `json:"status,omitempty"`

}

