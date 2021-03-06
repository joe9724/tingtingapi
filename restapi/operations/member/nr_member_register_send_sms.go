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
	"tingtingapi/var"
	"time"
	"math/rand"
	_"strconv"
)

// NrMemberRegisterSendSmsHandlerFunc turns a function with the right signature into a member register send sms handler
type NrMemberRegisterSendSmsHandlerFunc func(NrMemberRegisterSendSmsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrMemberRegisterSendSmsHandlerFunc) Handle(params NrMemberRegisterSendSmsParams) middleware.Responder {
	return fn(params)
}

// NrMemberRegisterSendSmsHandler interface for that can handle valid member register send sms params
type NrMemberRegisterSendSmsHandler interface {
	Handle(NrMemberRegisterSendSmsParams) middleware.Responder
}

// NewNrMemberRegisterSendSms creates a new http.Handler for the member register send sms operation
func NewNrMemberRegisterSendSms(ctx *middleware.Context, handler NrMemberRegisterSendSmsHandler) *NrMemberRegisterSendSms {
	return &NrMemberRegisterSendSms{Context: ctx, Handler: handler}
}

/*NrMemberRegisterSendSms swagger:route GET /member/register/sendSms Member memberRegisterSendSms

下发注册时的验证码

下发注册时的验证码

*/
type NrMemberRegisterSendSms struct {
	Context *middleware.Context
	Handler NrMemberRegisterSendSmsHandler
}

func (o *NrMemberRegisterSendSms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrMemberRegisterSendSmsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok MemberRegisterSendSmsOK
	var response models.InlineResponse20016

	var status models.Response

	var msg string
	var code string

	msg = "操作成功"
	code = "200"

	//(1)产生验证码
	//var code string
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(vcode)
	code = vcode
	//code = "654321"
	// (2)查询数据库内是否5分钟内已经有验证码下发记录 没有的话请求第三方下发验证码
    var findRecord models.SendSms


	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()


	//query
    // type =0 正常注册      =1 第三方登录后绑定手机号     =2 快捷登录
		db.Table("sms").Where("type=?", 0).Where(map[string]interface{}{"phone": Params.PhoneNumber}).Where("ts>?", time.Now().Unix()-5*60).Last(&findRecord)
		fmt.Println(findRecord)
		if findRecord.Id == 0 {
			code = code
			fmt.Println("5分钟内没有delay1=", time.Now().Unix()-300)
			//调用第三方sp平台下发短信
			_var.SendMsg(*(Params.PhoneNumber), code,*(Params.Type))
		} else {
			fmt.Println("5分钟内有delay2=", time.Now().Unix()-300)
			code = findRecord.Code
			//调用第三方sp平台下发短信
			_var.SendMsg(*(Params.PhoneNumber), code,*(Params.Type))
		}
		var send bool
		send = true
		//(3)第二步成功后回调后入库
		if (send == true) {
			findRecord.Code = code
			findRecord.Phone = *(Params.PhoneNumber)
			findRecord.Type = *(Params.Type)
			findRecord.Ts = int64(time.Now().Unix())
			db.Table("sms").Create(&findRecord)
		}

		// 如果是快捷登录，将验证码作为临时密码存放到member里
		/*if(*(Params.Type) ==1 ){  //
			fmt.Println("register.sendsms.phone=?",Params.PhoneNumber)
			var member models.Member
			db.Table("members").Where("phone=?",Params.PhoneNumber).Find(&member)
			member.Password = code
			member.Phone = *Params.PhoneNumber
			member.Ts = time.Now().Unix()
			member.Ts1 = "2018-02-24 18:00:15.322041"
			db.Save(&member)
			fmt.Println("register.sendsms.code=?",code)
		}*/

		/*codevalue,err := strconv.ParseInt(code,10,64)
		if err!=nil{

		}*/
	status.UnmarshalBinary([]byte(_var.Response200(200,msg)))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
