// Code generated by go-swagger; DO NOT EDIT.

package analytics

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
)

// NrAnalyticsAppHandlerFunc turns a function with the right signature into a analytics app handler
type NrAnalyticsAppHandlerFunc func(NrAnalyticsAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrAnalyticsAppHandlerFunc) Handle(params NrAnalyticsAppParams) middleware.Responder {
	return fn(params)
}

// NrAnalyticsAppHandler interface for that can handle valid analytics app params
type NrAnalyticsAppHandler interface {
	Handle(NrAnalyticsAppParams) middleware.Responder
}

// NewNrAnalyticsApp creates a new http.Handler for the analytics app operation
func NewNrAnalyticsApp(ctx *middleware.Context, handler NrAnalyticsAppHandler) *NrAnalyticsApp {
	return &NrAnalyticsApp{Context: ctx, Handler: handler}
}

/*NrAnalyticsApp swagger:route GET /analytics/app Analytics analyticsApp

统计

统计

*/
type NrAnalyticsApp struct {
	Context *middleware.Context
	Handler NrAnalyticsAppHandler
}

func (o *NrAnalyticsApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrAnalyticsAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok AnalyticsAppOK
	var response models.InlineResponse200
	//var icons models.AlbumBuyResult

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	var status models.Response
	if(*(Params.TargetType)==0){
		db.Exec("update chapters set play_count=play_count+1 where id=?",*(Params.TargetID))
	}else if (*(Params.TargetType)==4){//支付宝
		db.Exec("update members set money=money+? where id=?",*(Params.Value),*(Params.MemberID))
		db.Exec("insert into recharge(memberId,type,order_no,time,value) values(?,?,?,?,?)",Params.MemberID,4,Params.OrderNo,time.Now().UnixNano()/1000000,Params.Value)
	}else if (*(Params.TargetType)==5){//微信
		db.Exec("update members set money=money+? where id=?",*(Params.Value),*(Params.MemberID))
		db.Exec("insert into recharge(memberId,type,order_no,time,value) values(?,?,?,?,?)",Params.MemberID,5,Params.OrderNo,time.Now().UnixNano()/1000000,Params.Value)
	}



    var code int64
    var msg string

    code = 200
    msg = "上报成功"
    status.Code = &(code)
    status.Msg = &(msg)

	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
