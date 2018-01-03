// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrMemberUploadRecordParams creates a new NrMemberUploadRecordParams object
// with the default values initialized.
func NewNrMemberUploadRecordParams() NrMemberUploadRecordParams {
	var ()
	return NrMemberUploadRecordParams{}
}

// NrMemberUploadRecordParams contains all the bound params for the member upload record operation
// typically these are obtained from a http.Request
//
// swagger:parameters /member/uploadRecord
type NrMemberUploadRecordParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*备用字段
	  In: formData
	*/
	CategoryID *int64
	/*
	  In: formData
	*/
	Content *string
	/*封面.
	  In: formData
	*/
	Cover io.ReadCloser
	/*录音文件.
	  Required: true
	  In: formData
	*/
	Icon io.ReadCloser
	/*会员id
	  In: formData
	*/
	MemberID *int64
	/*
	  Required: true
	  In: formData
	*/
	SubTitle string
	/*
	  In: formData
	*/
	Summary *string
	/*
	  In: formData
	*/
	Tag *string
	/*Description of file contents.
	  Required: true
	  In: formData
	*/
	Title string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrMemberUploadRecordParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdCategoryID, fdhkCategoryID, _ := fds.GetOK("categoryId")
	if err := o.bindCategoryID(fdCategoryID, fdhkCategoryID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdContent, fdhkContent, _ := fds.GetOK("content")
	if err := o.bindContent(fdContent, fdhkContent, route.Formats); err != nil {
		res = append(res, err)
	}

	cover, coverHeader, err := r.FormFile("cover")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "cover", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindCover(cover, coverHeader); err != nil {
		res = append(res, err)
	} else {
		o.Cover = &runtime.File{Data: cover, Header: coverHeader}
	}

	icon, iconHeader, err := r.FormFile("icon")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "icon", err))
	} else if err := o.bindIcon(icon, iconHeader); err != nil {
		res = append(res, err)
	} else {
		o.Icon = &runtime.File{Data: icon, Header: iconHeader}
	}

	fdMemberID, fdhkMemberID, _ := fds.GetOK("memberId")
	if err := o.bindMemberID(fdMemberID, fdhkMemberID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdSubTitle, fdhkSubTitle, _ := fds.GetOK("subTitle")
	if err := o.bindSubTitle(fdSubTitle, fdhkSubTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	fdSummary, fdhkSummary, _ := fds.GetOK("summary")
	if err := o.bindSummary(fdSummary, fdhkSummary, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTag, fdhkTag, _ := fds.GetOK("tag")
	if err := o.bindTag(fdTag, fdhkTag, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTitle, fdhkTitle, _ := fds.GetOK("title")
	if err := o.bindTitle(fdTitle, fdhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrMemberUploadRecordParams) bindCategoryID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("categoryId", "formData", "int64", raw)
	}
	o.CategoryID = &value

	return nil
}

func (o *NrMemberUploadRecordParams) bindContent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Content = &raw

	return nil
}

func (o *NrMemberUploadRecordParams) bindCover(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *NrMemberUploadRecordParams) bindIcon(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *NrMemberUploadRecordParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("memberId", "formData", "int64", raw)
	}
	o.MemberID = &value

	return nil
}

func (o *NrMemberUploadRecordParams) bindSubTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("subTitle", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("subTitle", "formData", raw); err != nil {
		return err
	}

	o.SubTitle = raw

	return nil
}

func (o *NrMemberUploadRecordParams) bindSummary(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Summary = &raw

	return nil
}

func (o *NrMemberUploadRecordParams) bindTag(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Tag = &raw

	return nil
}

func (o *NrMemberUploadRecordParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("title", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("title", "formData", raw); err != nil {
		return err
	}

	o.Title = raw

	return nil
}
