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

	"time"
)

// RegisterHandlerFunc turns a function with the right signature into a register handler
type RegisterHandlerFunc func(RegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterHandlerFunc) Handle(params RegisterParams) middleware.Responder {
	return fn(params)
}

// RegisterHandler interface for that can handle valid register params
type RegisterHandler interface {
	Handle(RegisterParams) middleware.Responder
}

// NewRegister creates a new http.Handler for the register operation
func NewRegister(ctx *middleware.Context, handler RegisterHandler) *Register {
	return &Register{Context: ctx, Handler: handler}
}

/*Register swagger:route POST /member/register Member register

注册接口

注册接口

*/
type Register struct {
	Context *middleware.Context
	Handler RegisterHandler
}

func (o *Register) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegisterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok RegisterOK
	var response models.InlineResponse20016
	//var loginRet models.LoginRet
	//var msg string
	//var code int64
	var status models.Response
	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	fmt.Println("phone=",Params.PhoneNumber)

	//判断member里是有已有phone对应的记录，没有的话插入，有的话更新

	var member models.Member
	db.Table("members").Where("phone=?",Params.PhoneNumber).Find(&member)

	if(member.ID==int64(0)){
		//查找是否已有用户名被占用了
		var temp models.Member
		db.Table("members").Where("name=?",Params.Membername).Find(&temp)
		if (temp.ID==0){
			//insert
			status.UnmarshalBinary([]byte(_var.Response200(200,"注册成功")))
			fmt.Println("注册成功")
			db.Table("members").FirstOrCreate(&models.Member{}, models.Member{Phone: *Params.PhoneNumber,Name: *Params.Membername,Password:*Params.Password,Ts:time.Now().Unix()})
		}else{
			status.UnmarshalBinary([]byte(_var.Response200(202,"用户名已经被占用")))
		}

	}else{

		//update
		status.UnmarshalBinary([]byte(_var.Response200(201,"更新成功")))
		if(Params.Membername!=nil) {
			member.Name = *Params.Membername
		}
		member.Password = *Params.Password
		member.Phone = *Params.PhoneNumber
		member.Ts = time.Now().Unix()
		db.Save(&member)
		//db.Table("members").FirstOrCreate(&models.Member{}, models.Member{Phone: *Params.PhoneNumber,Name: *Params.Membername,Password:*Params.Password})
		fmt.Println("更新成功")
	}

	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
