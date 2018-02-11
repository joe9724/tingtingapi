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

// NrOrderSerialNuberHandlerFunc turns a function with the right signature into a order serial nuber handler
type NrOrderSerialNuberHandlerFunc func(NrOrderSerialNuberParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrOrderSerialNuberHandlerFunc) Handle(params NrOrderSerialNuberParams) middleware.Responder {
	return fn(params)
}

// NrOrderSerialNuberHandler interface for that can handle valid order serial nuber params
type NrOrderSerialNuberHandler interface {
	Handle(NrOrderSerialNuberParams) middleware.Responder
}

// NewNrOrderSerialNuber creates a new http.Handler for the order serial nuber operation
func NewNrOrderSerialNuber(ctx *middleware.Context, handler NrOrderSerialNuberHandler) *NrOrderSerialNuber {
	return &NrOrderSerialNuber{Context: ctx, Handler: handler}
}

/*NrOrderSerialNuber swagger:route POST /member/order/serialNuber Member orderSerialNuber

充值订单产生流水号

充值订单产生流水号

*/
type NrOrderSerialNuber struct {
	Context *middleware.Context
	Handler NrOrderSerialNuberHandler
}

func (o *NrOrderSerialNuber) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrOrderSerialNuberParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	//res := o.Handler.Handle(Params) // actually handle the request

	var ok OrderSerialNuberOK
	var response models.InlineResponse200
	var status models.Response

	//收到app端请求支付，生成一个流水号，并向相应支付平台发起请求，同时将记录入库
	var serialNumber string
	serialNumber = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	//发起请求


	//入库
	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var t models.RechargeModel
	t.Type = "支付宝"
	t.Status = 0
	t.Order_no = serialNumber
	t.Value = Params.Body.Value
	t.MemberId = Params.Body.MemberID

	db.Table("recharge").Create(&t)

	status.UnmarshalBinary([]byte(_var.Response200(200,"成功发起支付请求")))
	response.Status = &status


	response.Data.Msg = "成功发起支付请求"
	response.Data.Code = 200
	response.Data.SerialNumber = serialNumber
	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)


}
