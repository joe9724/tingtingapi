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

// NrMemberInitHandlerFunc turns a function with the right signature into a member init handler
type NrMemberInitHandlerFunc func(NrMemberInitParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrMemberInitHandlerFunc) Handle(params NrMemberInitParams) middleware.Responder {
	return fn(params)
}

// NrMemberInitHandler interface for that can handle valid member init params
type NrMemberInitHandler interface {
	Handle(NrMemberInitParams) middleware.Responder
}

// NewNrMemberInit creates a new http.Handler for the member init operation
func NewNrMemberInit(ctx *middleware.Context, handler NrMemberInitHandler) *NrMemberInit {
	return &NrMemberInit{Context: ctx, Handler: handler}
}

/*NrMemberInit swagger:route GET /member/init Member memberInit

启动页接口

启动页接口

*/
type NrMemberInit struct {
	Context *middleware.Context
	Handler NrMemberInitHandler
}

type Web struct{
	Url string `json:"url"`
}


func (o *NrMemberInit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrMemberInitParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok MemberInitOK
	var response models.InitModel
	var initData models.InitData

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	//query
	db.Table("init").Where(map[string]interface{}{"status":0}).Find(&initData).Limit(1)

	fmt.Println(initData)

	initData.Msg = initData.Msg
	initData.DownloadURL = "http://www.baidu.com"
	initData.Force = initData.Force
	initData.Number = initData.Number
	//var aboutus string
	//aboutus = "http://www.baidu.com"*/

	var webs []*Web

	var intraInfo models.InitDataExtraInfo

	initData.ExtraInfo = intraInfo

	db.Raw("select url from web").Find(&webs)

	intraInfo.AQURL = (webs[0].Url)

	intraInfo.SpecialURL = (webs[1].Url)

	intraInfo.AboutUsURL = (webs[2].Url)

	initData.ExtraInfo = intraInfo

	//var s interface{}
	//db.Table("init").Where(map[string]interface{}{"status":0}).Find(&s).Limit(1)
	//fmt.Println(s)
	//data
	response.Data = &initData
	response.Data.Force = "0"

//	response.Data.ExtraInfo.AboutUsURL = &(webs[0].Url)

	//status
	var status models.Status
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)
}
