// Code generated by go-swagger; DO NOT EDIT.

package search

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

// SearchHandlerFunc turns a function with the right signature into a search handler
type SearchHandlerFunc func(SearchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchHandlerFunc) Handle(params SearchParams) middleware.Responder {
	return fn(params)
}

// SearchHandler interface for that can handle valid search params
type SearchHandler interface {
	Handle(SearchParams) middleware.Responder
}

// NewSearch creates a new http.Handler for the search operation
func NewSearch(ctx *middleware.Context, handler SearchHandler) *Search {
	return &Search{Context: ctx, Handler: handler}
}

/*Search swagger:route GET /search Search search

搜索

搜索

*/
type Search struct {
	Context *middleware.Context
	Handler SearchHandler
}

func (o *Search) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSearchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok SearchOK
	var response models.InlineResponse20010
	var albumList models.InlineResponse20010AlbumList
	var tagList []models.Tag

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	// var test []interface{}
	//db.Table("sub_category_items").Select("sub_category_items.name, category_album_relation.albumId").Joins("left join category_album_relation on category_album_relation.categoryId = sub_category_items.id and sub_category_items.id=?",1).Scan(&test)
	//db.Joins("JOIN sub_category_items ON sub_category_items.id = category_album_relation.albumId AND sub_category_items.id = ?",1).Where("credit_cards.number = ?", "411111111111").Find(&test)

	if (*(Params.Action) == "getContent"){
		rows, err := db.Table("albums").Select("albums.name, albums.id,albums.author_avatar,albums.author_name,albums.books_number,albums.icon,albums.play_count,albums.sub_title,albums.time").Joins("left join category_album_relation on category_album_relation.albumId = albums.id").Where("albums.name like '%南%'").Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Rows()
		//var temp []models.Album
		for rows.Next() {
			var name string
			var albumId int64
			var author_avatar string
			var author_name string
			var books_number int64
			var icon string
			var play_count int64
			var sub_title string
			var time int64

			err = rows.Scan(&name, &albumId, &author_avatar, &author_name, &books_number, &icon, &play_count, &sub_title, &time)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(name, albumId)
			var t models.Album
			t.AlbumID = albumId
			t.AlbumName = name
			t.Value = 1
			t.Icon = icon
			t.Author_Name = author_name
			t.Author_Avatar = author_avatar
			t.Books_Number = books_number
			t.Icon = icon
			t.Play_Count = play_count
			t.Sub_Title = sub_title
			t.Time = time
			//temp = append(temp,t)
			albumList = append(albumList, &t)
		}
		response.AlbumList = albumList
	}else if(*(Params.Action) == "getTags"){
            db.Table("tags").Where("status=?",0).Find(&tagList)
            response.TagList = tagList
	}

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
