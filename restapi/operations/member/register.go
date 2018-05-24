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
    "io/ioutil"
	"time"
	"strings"
	"runtime"
	"strconv"
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
	var msg string
	var code int64
	//var loginRet models.LoginRet
	//var msg string
	//var code int64
	msg = "注册成功"
	var status models.Response
	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	//query
	fmt.Println("phone=",Params.PhoneNumber)
	//判断验证码是否正确
	validateCode := *(Params.ValidateCode)
	var sms models.SendSms
	db.Table("sms").Where("phone=?",Params.PhoneNumber).Where("code=?",validateCode).Last(&sms)
	if (sms.Id == 0){
		status.UnmarshalBinary([]byte(_var.Response200(401,"验证码不正确")))
		response.Status = &status
		ok.SetPayload(&response)
		o.Context.Respond(rw, r, route.Produces, route, ok)

		return
	}

	//判断member里是有已有phone对应的记录，没有的话插入，有的话更新
	var filename string
	var birth string
	var grade string

	filename = strconv.FormatInt((time.Now().Unix()),10)
	var member models.Member
	db.Table("members").Where("phone=?",Params.PhoneNumber).Find(&member)

	if(Params.Membername!=nil) {
		member.Name = *Params.Membername
	}
	member.Password = *Params.Password
	member.Phone = *Params.PhoneNumber
	member.Ts = time.Now().Unix()
	member.Ts1 = "2018-02-24 18:00:15.322041"
	if (Params.BirthYear != nil) {
		birth = strconv.FormatInt(*(Params.BirthYear),10) + "-" + strconv.FormatInt(*(Params.BirthMonth),10) + "-" + strconv.FormatInt(*(Params.BirthDay),10)
		member.Birth = birth
	}
	if (Params.Grade != nil) {
		grade = strconv.FormatInt(*(Params.Grade),10)
		member.Grade = grade
	}
	if(member.ID==int64(0)){
		//查找是否已有用户名被占用了
		var temp models.Member
		db.Table("members").Where("name=?",Params.Membername).Find(&temp)
		if (temp.ID==0){
			//insert
			code = 200
			msg = "注册成功"
			has,_ := HasAvatar(Params,filename)
            if has == true{
            	member.Avatar = filename
			}
			db.Save(&member)
		}else{
			code = 202
			msg = "用户名已被注册,请换个用户名"
		}

	}else{
		code = 203
		msg = "手机号已经被注册"
	}

	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}



func HasAvatar(Params  RegisterParams , filename string)  (r bool,saveFileName string){
	//如果有cover
	if (Params.Avatar!=nil) {
		avatar, err := ioutil.ReadAll(Params.Avatar)
		if err != nil {
			fmt.Println("err upload:", err.Error())
		}
		contentType := http.DetectContentType(avatar)
		var lower string
		lower = strings.ToLower(contentType)
		if(strings.Contains(lower,"jp")||(strings.Contains(lower,"pn"))) {
			if(runtime.GOOS == "windows") {
				err1 := ioutil.WriteFile(filename+".jpg", avatar, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			}else{
				err1 := ioutil.WriteFile("/root/go/src/resource/image/avatar/"+filename+".jpg", avatar, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			}
			return true,filename+".jpg"
		}else{
			return false,""
		}


	}else{
		return false,""
	}
}