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

// LoginHandlerFunc turns a function with the right signature into a login handler
type LoginHandlerFunc func(LoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginHandlerFunc) Handle(params LoginParams) middleware.Responder {
	return fn(params)
}

// LoginHandler interface for that can handle valid login params
type LoginHandler interface {
	Handle(LoginParams) middleware.Responder
}

// NewLogin creates a new http.Handler for the login operation
func NewLogin(ctx *middleware.Context, handler LoginHandler) *Login {
	return &Login{Context: ctx, Handler: handler}
}

/*Login swagger:route GET /member/login Member login

登录接口

登录接口

*/
type Login struct {
	Context *middleware.Context
	Handler LoginHandler
}

func (o *Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoginParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok LoginOK
	var response models.InlineResponse20021
	var loginRet models.LoginRet

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	fmt.Println("phone=",Params.Phone)
	fmt.Println("password=",Params.Password)
	fmt.Println("loginType=",Params.LoginType)

	db.Table("members").Where("phone=?",Params.Phone).Where("password=?",Params.Password).Find(&loginRet).Limit(1)

	fmt.Println(loginRet)
	response.Data = &loginRet
	/*response.Data.MemberID = response.Data.Id
	response.Data.Membername = response.Data.Name

	response.Data.MemberID = 1
	response.Data.Membername = "bf"
	response.Data.Gender=1
	response.Data.Avatar = "http://tingting-resource.bitekun.xin/resource/image/avatar.jpg"
	response.Data.Name = "pf"
	response.Data.Area = "江苏 南京"
	response.Data.Birth = "1991-08-17"
	response.Data.Grade = "1-3"
	response.Data.Id = 1
	response.Data.Level = 1
	response.Data.Phone = "18963602871"
	response.Data.Money = 118
	response.Data.Ts = 1787868685*/

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)



}
