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
	"tingtingbackend/var"
)

// AlbumDetailHandlerFunc turns a function with the right signature into a album detail handler
type AlbumDetailHandlerFunc func(AlbumDetailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AlbumDetailHandlerFunc) Handle(params AlbumDetailParams) middleware.Responder {
	return fn(params)
}

// AlbumDetailHandler interface for that can handle valid album detail params
type AlbumDetailHandler interface {
	Handle(AlbumDetailParams) middleware.Responder
}

// NewAlbumDetail creates a new http.Handler for the album detail operation
func NewAlbumDetail(ctx *middleware.Context, handler AlbumDetailHandler) *AlbumDetail {
	return &AlbumDetail{Context: ctx, Handler: handler}
}

/*AlbumDetail swagger:route GET /album/detail Album albumDetail

专辑详情

专辑详情

*/
type AlbumDetail struct {
	Context *middleware.Context
	Handler AlbumDetailHandler
}

func (o *AlbumDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAlbumDetailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok AlbumDetailOK
	var response models.InlineResponse2005
	var album models.Album

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	db.Table("albums").Where(map[string]interface{}{"status":0}).Where("id=?",Params.AlbumID).Find(&album)

	//response.Data.AlbumID = album.ID
	//response.Data.AlbumName = album.Name
	response.IsFav = true
	response.Data = &album
	response.Data.AlbumID = album.ID
    response.Data.AlbumName = album.Name

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
