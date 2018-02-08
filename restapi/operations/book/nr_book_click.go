// Code generated by go-swagger; DO NOT EDIT.

package book

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
	"strconv"
)

// NrBookClickHandlerFunc turns a function with the right signature into a book click handler
type NrBookClickHandlerFunc func(NrBookClickParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrBookClickHandlerFunc) Handle(params NrBookClickParams) middleware.Responder {
	return fn(params)
}

// NrBookClickHandler interface for that can handle valid book click params
type NrBookClickHandler interface {
	Handle(NrBookClickParams) middleware.Responder
}

// NewNrBookClick creates a new http.Handler for the book click operation
func NewNrBookClick(ctx *middleware.Context, handler NrBookClickHandler) *NrBookClick {
	return &NrBookClick{Context: ctx, Handler: handler}
}

/*NrBookClick swagger:route GET /book/click Book bookClick

点赞/取消点赞 书本

点赞/取消点赞 书本

*/
type NrBookClick struct {
	Context *middleware.Context
	Handler NrBookClickHandler
}

func (o *NrBookClick) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrBookClickParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok BookClickOK
	var response models.InlineResponse200
	//var icons models.AlbumBuyResult

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	var status models.Response

	//如果是点赞
	if(*(Params.Action) == "click") {
		//
		value, _ := strconv.ParseInt(*Params.MemberID, 10, 64)
		db.Table("click_book").FirstOrCreate(&models.Fav_Book{}, models.Fav_Book{BookId: *Params.BookID,MemberId:value})
		status.UnmarshalBinary([]byte(_var.Response200(200,"点赞成功")))
	}else if (*(Params.Action) == "unclick"){
		var favModel interface{}
		db.Table("click_book").Where("album_id = ?", Params.BookID).Where("member_id=?",Params.MemberID).Delete(&favModel)
		status.UnmarshalBinary([]byte(_var.Response200(200,"取消点赞成功")))
	}

	//query

	//data

	//response.Data = icons

	//status

	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
