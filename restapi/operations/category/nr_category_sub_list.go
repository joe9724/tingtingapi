// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	middleware "github.com/go-openapi/runtime/middleware"
	"tingtingapi/models"
	_"fmt"
	"tingtingbackend/var"

)

// NrCategorySubListHandlerFunc turns a function with the right signature into a category sub list handler
type NrCategorySubListHandlerFunc func(NrCategorySubListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrCategorySubListHandlerFunc) Handle(params NrCategorySubListParams) middleware.Responder {
	return fn(params)
}

// NrCategorySubListHandler interface for that can handle valid category sub list params
type NrCategorySubListHandler interface {
	Handle(NrCategorySubListParams) middleware.Responder
}

// NewNrCategorySubList creates a new http.Handler for the category sub list operation
func NewNrCategorySubList(ctx *middleware.Context, handler NrCategorySubListHandler) *NrCategorySubList {
	return &NrCategorySubList{Context: ctx, Handler: handler}
}

/*NrCategorySubList swagger:route GET /category/sub/list Category categorySubList

获取子类列表

获取子类列表

*/
type NrCategorySubList struct {
	Context *middleware.Context
	Handler NrCategorySubListHandler
}

func (o *NrCategorySubList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrCategorySubListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok CategorySubListOK
	var response models.InlineResponse20023
	var categoryList models.InlineResponse20023Data

	response.Data = categoryList

	var SubCategories []models.SubCategoryItem

	for k:=0; k<6;k++  {
		var sub models.SubCategoryItem
		sub.CategoryID = 1
		sub.Name = "猜你喜欢"
		sub.Icon = "http://tingting-resource.bitekun.xin/resource/image/avatar.jpg"
		//
		var albums []models.Album
		for i:=0;i<6 ;i++  {
			var album models.Album
			album.Name = "现代文学"
			album.ID = int64(i)
			album.AlbumName = "现代文学"
			album.PlayCount = 123
			album.BooksNumber = 34
			album.Value = 118
			album.Icon = "http://tingting-resource.bitekun.xin/resource/image/cover.png"
			album.Cover = "http://tingting-resource.bitekun.xin/resource/image/cover.png"
			albums = append(albums,album)
		}
		sub.AlbumList = albums
		//
		SubCategories = append(SubCategories,sub)
	}

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status
	response.DataList = SubCategories

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
