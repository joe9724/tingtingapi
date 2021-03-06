// Code generated by go-swagger; DO NOT EDIT.

package album

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
	"strconv"
)

// NrAlbumClickHandlerFunc turns a function with the right signature into a album click handler
type NrAlbumClickHandlerFunc func(NrAlbumClickParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrAlbumClickHandlerFunc) Handle(params NrAlbumClickParams) middleware.Responder {
	return fn(params)
}

// NrAlbumClickHandler interface for that can handle valid album click params
type NrAlbumClickHandler interface {
	Handle(NrAlbumClickParams) middleware.Responder
}

// NewNrAlbumClick creates a new http.Handler for the album click operation
func NewNrAlbumClick(ctx *middleware.Context, handler NrAlbumClickHandler) *NrAlbumClick {
	return &NrAlbumClick{Context: ctx, Handler: handler}
}

/*NrAlbumClick swagger:route GET /album/click Album albumClick

点赞/取消点赞专辑

点赞/取消点赞专辑

*/
type NrAlbumClick struct {
	Context *middleware.Context
	Handler NrAlbumClickHandler
}

func (o *NrAlbumClick) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrAlbumClickParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok AlbumClickOK
	var response models.InlineResponse200
	var msg string
	//var icons models.AlbumBuyResult

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	//如果是点赞
	if(*(Params.Action) == "click") {
		//
		value, _ := strconv.ParseInt(*Params.MemberID, 10, 64)
		db.Table("click_album").FirstOrCreate(&models.Fav_Album{}, models.Fav_Album{AlbumId: *Params.AlbumID,MemberId:value})
		msg = "点赞成功"
		fmt.Println("点赞成功")
	}else if (*(Params.Action) == "unclick"){
		var favModel interface{}
		db.Table("click_album").Where("album_id = ?", Params.AlbumID).Where("member_id=?",Params.MemberID).Delete(&favModel)
		msg = "取消点赞成功"
		fmt.Println("取消点赞成功")
	}

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,msg)))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
