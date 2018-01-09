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

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrMemberEditParams creates a new NrMemberEditParams object
// with the default values initialized.
func NewNrMemberEditParams() NrMemberEditParams {
	var ()
	return NrMemberEditParams{}
}

// NrMemberEditParams contains all the bound params for the member edit operation
// typically these are obtained from a http.Request
//
// swagger:parameters /member/edit
type NrMemberEditParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
	/*出生日
		  In: formData
		*/
	MemberID *int64
	/*登录名
	  In: formData
	*/
	Membername *string
	/*头像
	  In: formData
	*/
	Avatar io.ReadCloser
	/*出生日
	  In: formData
	*/
	BirthDay *int64
	/*出生月
	  In: formData
	*/
	BirthMonth *int64
	/*出生年
	  In: formData
	*/
	BirthYear *int64
	/*客户端类型
	  In: formData
	*/
	Client *string
	/*年级
	  In: formData
	*/
	Grade *string
	/*唯一识别号
	  In: formData
	*/
	Imei *string
	/*密码
	  In: formData
	*/
	Password *string
	/*手机号
	  In: formData
	*/
	PhoneNumber *string
	/*时间戳
	  In: formData
	*/
	Ts *int64
	/*加密串
	  In: formData
	*/
	Val *string
	/*短信验证码
	  In: formData
	*/
	ValidateCode *string
	/*版本号
	  In: formData
	*/
	Version *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrMemberEditParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	fdMembername, fdhkMembername, _ := fds.GetOK("Membername")
	if err := o.bindMembername(fdMembername, fdhkMembername, route.Formats); err != nil {
		res = append(res, err)
	}

	avatar, avatarHeader, err := r.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "avatar", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindAvatar(avatar, avatarHeader); err != nil {
		res = append(res, err)
	} else {
		o.Avatar = &runtime.File{Data: avatar, Header: avatarHeader}
	}

	fdBirthDay, fdhkBirthDay, _ := fds.GetOK("birth-day")
	if err := o.bindBirthDay(fdBirthDay, fdhkBirthDay, route.Formats); err != nil {
		res = append(res, err)
	}

	fdBirthMonth, fdhkBirthMonth, _ := fds.GetOK("birth-month")
	if err := o.bindBirthMonth(fdBirthMonth, fdhkBirthMonth, route.Formats); err != nil {
		res = append(res, err)
	}

	fdBirthYear, fdhkBirthYear, _ := fds.GetOK("birth-year")
	if err := o.bindBirthYear(fdBirthYear, fdhkBirthYear, route.Formats); err != nil {
		res = append(res, err)
	}

	fdClient, fdhkClient, _ := fds.GetOK("client")
	if err := o.bindClient(fdClient, fdhkClient, route.Formats); err != nil {
		res = append(res, err)
	}

	fdGrade, fdhkGrade, _ := fds.GetOK("grade")
	if err := o.bindGrade(fdGrade, fdhkGrade, route.Formats); err != nil {
		res = append(res, err)
	}

	fdImei, fdhkImei, _ := fds.GetOK("imei")
	if err := o.bindImei(fdImei, fdhkImei, route.Formats); err != nil {
		res = append(res, err)
	}

	fdPassword, fdhkPassword, _ := fds.GetOK("password")
	if err := o.bindPassword(fdPassword, fdhkPassword, route.Formats); err != nil {
		res = append(res, err)
	}

	fdPhoneNumber, fdhkPhoneNumber, _ := fds.GetOK("phoneNumber")
	if err := o.bindPhoneNumber(fdPhoneNumber, fdhkPhoneNumber, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTs, fdhkTs, _ := fds.GetOK("ts")
	if err := o.bindTs(fdTs, fdhkTs, route.Formats); err != nil {
		res = append(res, err)
	}

	fdMemberID, fdhkMemberID, _ := fds.GetOK("memberId")
	if err := o.bindMemberID(fdMemberID, fdhkMemberID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdVal, fdhkVal, _ := fds.GetOK("val")
	if err := o.bindVal(fdVal, fdhkVal, route.Formats); err != nil {
		res = append(res, err)
	}

	fdValidateCode, fdhkValidateCode, _ := fds.GetOK("validateCode")
	if err := o.bindValidateCode(fdValidateCode, fdhkValidateCode, route.Formats); err != nil {
		res = append(res, err)
	}

	fdVersion, fdhkVersion, _ := fds.GetOK("version")
	if err := o.bindVersion(fdVersion, fdhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrMemberEditParams) bindMembername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Membername = &raw

	return nil
}

func (o *NrMemberEditParams) bindAvatar(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *NrMemberEditParams) bindBirthDay(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("birth-day", "formData", "int64", raw)
	}
	o.BirthDay = &value

	return nil
}

func (o *NrMemberEditParams) bindBirthMonth(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("birth-month", "formData", "int64", raw)
	}
	o.BirthMonth = &value

	return nil
}

func (o *NrMemberEditParams) bindBirthYear(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("birth-year", "formData", "int64", raw)
	}
	o.BirthYear = &value

	return nil
}

func (o *NrMemberEditParams) bindClient(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Client = &raw

	return nil
}

func (o *NrMemberEditParams) bindGrade(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value := raw
	/*if err != nil {
		return errors.InvalidType("grade", "formData", "int64", raw)
	}*/
	o.Grade = &value

	return nil
}

func (o *NrMemberEditParams) bindImei(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Imei = &raw

	return nil
}

func (o *NrMemberEditParams) bindPassword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Password = &raw

	return nil
}

func (o *NrMemberEditParams) bindPhoneNumber(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.PhoneNumber = &raw

	return nil
}

func (o *NrMemberEditParams) bindTs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ts", "formData", "int64", raw)
	}
	o.Ts = &value

	return nil
}

func (o *NrMemberEditParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrMemberEditParams) bindVal(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Val = &raw

	return nil
}

func (o *NrMemberEditParams) bindValidateCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ValidateCode = &raw

	return nil
}

func (o *NrMemberEditParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Version = &raw

	return nil
}
