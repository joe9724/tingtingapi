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
	"tingtingapi/var"
	"fmt"
	"time"
)

// NrMemberLoginByThirdPartyHandlerFunc turns a function with the right signature into a member login by third party handler
type NrMemberLoginByThirdPartyHandlerFunc func(NrMemberLoginByThirdPartyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrMemberLoginByThirdPartyHandlerFunc) Handle(params NrMemberLoginByThirdPartyParams) middleware.Responder {
	return fn(params)
}

// NrMemberLoginByThirdPartyHandler interface for that can handle valid member login by third party params
type NrMemberLoginByThirdPartyHandler interface {
	Handle(NrMemberLoginByThirdPartyParams) middleware.Responder
}

// NewNrMemberLoginByThirdParty creates a new http.Handler for the member login by third party operation
func NewNrMemberLoginByThirdParty(ctx *middleware.Context, handler NrMemberLoginByThirdPartyHandler) *NrMemberLoginByThirdParty {
	return &NrMemberLoginByThirdParty{Context: ctx, Handler: handler}
}

/*NrMemberLoginByThirdParty swagger:route POST /member/loginByThirdParty Member memberLoginByThirdParty

第三方登录

第三方登录

*/
type NrMemberLoginByThirdParty struct {
	Context *middleware.Context
	Handler NrMemberLoginByThirdPartyHandler
}

func (o *NrMemberLoginByThirdParty) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrMemberLoginByThirdPartyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	//res := o.Handler.Handle(Params) // actually handle the request

	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var ok LoginOK
	var res models.InlineResponse20021
	var state models.Response
	var loginRet models.LoginRet

	db.Table("members").Where("openid=?", Params.Body.Openid).Where("platform=?", Params.Body.Type).First(&loginRet)
	// 判断是否已经存在用户信息，存在则返回这个用户信息
	if loginRet.ID != 0 {
		// 修改最后一次登录时间
		sql := "UPDATE members SET ts = ? WHERE id = ? AND status = 0"
		db.Exec(sql, time.Now().Unix(), loginRet.ID)
	} else {
		sql := "INSERT INTO members(name, avatar, open_id, platform, ts) VALUES (?,?,?,?,?)"
		db.Exec(sql, Params.Body.Name, Params.Body.Avatar, Params.Body.Openid, Params.Body.Type, time.Now().Unix())
		// 写完之后再查询一次，保证用户存在
		db.Table("members").Where("openid=?", Params.Body.Openid).Where("platform=?", Params.Body.Type).First(&loginRet)
	}

	if loginRet.ID != 0 {
		loginRet.Password = ""
		res.Data = &loginRet
		state.UnmarshalBinary([]byte(_var.Response200(200, "登录成功")))
		res.Status = &state
	} else {
		state.UnmarshalBinary([]byte(_var.Response200(301, "用户不存在")))
		res.Status = &state
	}

	ok.SetPayload(&res)

	o.Context.Respond(rw, r, route.Produces, route, res)

}
