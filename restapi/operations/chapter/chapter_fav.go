// Code generated by go-swagger; DO NOT EDIT.

package chapter

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

// ChapterFavHandlerFunc turns a function with the right signature into a chapter fav handler
type ChapterFavHandlerFunc func(ChapterFavParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ChapterFavHandlerFunc) Handle(params ChapterFavParams) middleware.Responder {
	return fn(params)
}

// ChapterFavHandler interface for that can handle valid chapter fav params
type ChapterFavHandler interface {
	Handle(ChapterFavParams) middleware.Responder
}

// NewChapterFav creates a new http.Handler for the chapter fav operation
func NewChapterFav(ctx *middleware.Context, handler ChapterFavHandler) *ChapterFav {
	return &ChapterFav{Context: ctx, Handler: handler}
}

/*ChapterFav swagger:route GET /chapter/fav Chapter chapterFav

收藏/取消收藏章节

收藏/取消收藏章节

*/
type ChapterFav struct {
	Context *middleware.Context
	Handler ChapterFavHandler
}

func (o *ChapterFav) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewChapterFavParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok ChapterFavOK
	var response models.InlineResponse20016
	//var icons models.AlbumBuyResult

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	var status models.Response
	value, _ := strconv.ParseInt(*Params.ChapterID, 10, 64)
	//如果是收藏
	if(*(Params.Action) == "fav") {
		db.Table("fav_chapter").FirstOrCreate(&models.Fav_Chapter{}, models.Fav_Chapter{ChapterId: value,MemberId:*Params.Userid})
		status.UnmarshalBinary([]byte(_var.Response200(200,"收藏成功")))
	}else if (*(Params.Action) == "unfav"){
		var favModel interface{}
		db.Table("fav_chapter").Where("chapter_id = ?", value).Where("member_id=?",Params.Userid).Delete(&favModel)
		status.UnmarshalBinary([]byte(_var.Response200(200,"取消收藏成功")))
	}

	//query

	//data

	//response.Data = icons

	//status

	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)
}
