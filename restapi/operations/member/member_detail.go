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
)

// MemberDetailHandlerFunc turns a function with the right signature into a member detail handler
type MemberDetailHandlerFunc func(MemberDetailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MemberDetailHandlerFunc) Handle(params MemberDetailParams) middleware.Responder {
	return fn(params)
}

// MemberDetailHandler interface for that can handle valid member detail params
type MemberDetailHandler interface {
	Handle(MemberDetailParams) middleware.Responder
}

// NewMemberDetail creates a new http.Handler for the member detail operation
func NewMemberDetail(ctx *middleware.Context, handler MemberDetailHandler) *MemberDetail {
	return &MemberDetail{Context: ctx, Handler: handler}
}

/*MemberDetail swagger:route GET /member/detail Member memberDetail
会员详情
会员详情
*/
type MemberDetail struct {
	Context *middleware.Context
	Handler MemberDetailHandler
}

func (o *MemberDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMemberDetailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok MemberDetailOK
	var response models.InlineResponse200155
	var detail models.Member

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	db.Table("members").Where(map[string]interface{}{"status":0}).Where("id=?",Params.MemberID).Last(&detail)
	//detail.BuyedAlbumsCount = 2
	detail.Password = ""
	response.Data = &detail

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}