// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	middleware "github.com/go-openapi/runtime/middleware"
	"tingtingapi/models"
	"fmt"
	"tingtingbackend/var"
)

// MemberRechargeListHandlerFunc turns a function with the right signature into a member recharge list handler
type MemberRechargeListHandlerFunc func(MemberRechargeListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MemberRechargeListHandlerFunc) Handle(params MemberRechargeListParams) middleware.Responder {
	return fn(params)
}

// MemberRechargeListHandler interface for that can handle valid member recharge list params
type MemberRechargeListHandler interface {
	Handle(MemberRechargeListParams) middleware.Responder
}

// NewMemberRechargeList creates a new http.Handler for the member recharge list operation
func NewMemberRechargeList(ctx *middleware.Context, handler MemberRechargeListHandler) *MemberRechargeList {
	return &MemberRechargeList{Context: ctx, Handler: handler}
}

/*MemberRechargeList swagger:route POST /member/recharge/list Member memberRechargeList

获取充值列表(含条件检索)

获取充值列表

*/
type MemberRechargeList struct {
	Context *middleware.Context
	Handler MemberRechargeListHandler
}

func (o *MemberRechargeList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMemberRechargeListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok MemberRechargeListOK
	var response models.InlineResponse2002222
	var rechargeList models.InlineResponse2002222Recharges

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	db.Table("recharge").Where(map[string]interface{}{"status":0}).Where("memberId=?",Params.MemberID).Find(&rechargeList)
	//query

	//data

	response.Recharges = rechargeList

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
