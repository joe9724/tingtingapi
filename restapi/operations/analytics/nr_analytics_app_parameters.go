// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrAnalyticsAppParams creates a new NrAnalyticsAppParams object
// with the default values initialized.
func NewNrAnalyticsAppParams() NrAnalyticsAppParams {
	var ()
	return NrAnalyticsAppParams{}
}

// NrAnalyticsAppParams contains all the bound params for the analytics app operation
// typically these are obtained from a http.Request
//
// swagger:parameters /analytics/app
type NrAnalyticsAppParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*点赞/取消点赞 click/unclick
	  In: query
	*/
	Action *string
	/*客户端类型
	  In: query
	*/
	Client *string
	/*唯一识别号
	  In: query
	*/
	Imei *string
	/*用户id
	  In: query
	*/
	MemberID *string
	/*被点赞内容id
	  In: query
	*/
	TargetID *int64
	/*被点赞内容类型
	  In: query
	*/
	TargetType *int64
	/*版本号
	  In: query
	*/
	Version *string

	Value *float64

	OrderNo *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrAnalyticsAppParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAction, qhkAction, _ := qs.GetOK("action")
	if err := o.bindAction(qAction, qhkAction, route.Formats); err != nil {
		res = append(res, err)
	}

	qClient, qhkClient, _ := qs.GetOK("client")
	if err := o.bindClient(qClient, qhkClient, route.Formats); err != nil {
		res = append(res, err)
	}

	qImei, qhkImei, _ := qs.GetOK("imei")
	if err := o.bindImei(qImei, qhkImei, route.Formats); err != nil {
		res = append(res, err)
	}

	qMemberID, qhkMemberID, _ := qs.GetOK("memberId")
	if err := o.bindMemberID(qMemberID, qhkMemberID, route.Formats); err != nil {
		res = append(res, err)
	}

	qValue, qhkValue, _ := qs.GetOK("value")
	if err := o.bindValue(qValue, qhkValue, route.Formats); err != nil {
		res = append(res, err)
	}

	qTargetID, qhkTargetID, _ := qs.GetOK("targetId")
	if err := o.bindTargetID(qTargetID, qhkTargetID, route.Formats); err != nil {
		res = append(res, err)
	}

	qTargetType, qhkTargetType, _ := qs.GetOK("targetType")
	if err := o.bindTargetType(qTargetType, qhkTargetType, route.Formats); err != nil {
		res = append(res, err)
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrderNo, qhkOrderNo, _ := qs.GetOK("order_no")
	if err := o.bindOrderNo(qOrderNo, qhkOrderNo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrAnalyticsAppParams) bindAction(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Action = &raw

	return nil
}

func (o *NrAnalyticsAppParams) bindClient(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrAnalyticsAppParams) bindImei(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrAnalyticsAppParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.MemberID = &raw

	return nil
}

func (o *NrAnalyticsAppParams) bindValue(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("value", "query", "int64", raw)
	}
	o.Value = &value

	return nil
}

func (o *NrAnalyticsAppParams) bindTargetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("targetId", "query", "int64", raw)
	}
	o.TargetID = &value

	return nil
}

func (o *NrAnalyticsAppParams) bindTargetType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("targetType", "query", "int64", raw)
	}
	o.TargetType = &value

	return nil
}

func (o *NrAnalyticsAppParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrAnalyticsAppParams) bindOrderNo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.OrderNo = &raw

	return nil
}
