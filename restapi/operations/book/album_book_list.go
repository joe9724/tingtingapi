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
	"tingtingapi/var"
)

// AlbumBookListHandlerFunc turns a function with the right signature into a album book list handler
type AlbumBookListHandlerFunc func(AlbumBookListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AlbumBookListHandlerFunc) Handle(params AlbumBookListParams) middleware.Responder {
	return fn(params)
}

// AlbumBookListHandler interface for that can handle valid album book list params
type AlbumBookListHandler interface {
	Handle(AlbumBookListParams) middleware.Responder
}

// NewAlbumBookList creates a new http.Handler for the album book list operation
func NewAlbumBookList(ctx *middleware.Context, handler AlbumBookListHandler) *AlbumBookList {
	return &AlbumBookList{Context: ctx, Handler: handler}
}

/*AlbumBookList swagger:route GET /album/bookList Book albumBookList

专辑下的书列表

登录接口

*/
type AlbumBookList struct {
	Context *middleware.Context
	Handler AlbumBookListHandler
}

func (o *AlbumBookList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAlbumBookListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok AlbumBookListOK
	var response models.InlineResponse20019
	var books models.InlineResponse20019Books

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	//db.Table("albums").Where(map[string]interface{}{"status":0}).Find(&categoryList).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
	//query

	//var test []interface{}
	//db.Table("sub_category_items").Select("sub_category_items.name, category_album_relation.albumId").Joins("left join category_album_relation on category_album_relation.categoryId = sub_category_items.id and sub_category_items.id=?",1).Scan(&test)
	//db.Joins("JOIN sub_category_items ON sub_category_items.id = category_album_relation.albumId AND sub_category_items.id = ?",1).Where("credit_cards.number = ?", "411111111111").Find(&test)

	rows, err := db.Table("books").Select("books.name, album_book_relation.bookId,books.author_name,books.play_count,books.clips_number,books.icon,books.sub_title").Joins("left join album_book_relation on album_book_relation.bookId = books.id").Where("books.status=?",0).Where("album_book_relation.status=?",0).Where("album_book_relation.albumId=?",Params.AlbumID).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Rows()
	if err !=nil{
		fmt.Println("err is",err.Error())
	}
	//var temp []models.Album
	for rows.Next() {
		var name string
		var icon string
		var authorAvatar string
		var authorName string
		var bookId int64
		var playCount int64
		var clipsNumber int64
		var subTitle string

		err = rows.Scan(&name,&bookId,&authorName,&playCount,&clipsNumber,&icon,subTitle)
		if err != nil{
			fmt.Println(err.Error())
		}

		var t models.Book
		t.BookID = bookId
		t.Name = name
		t.Icon = icon
		t.AuthorAvatar = authorAvatar
		t.AuthorName = authorName
		t.PlayCount = playCount
		t.ClipsNumber = clipsNumber
		t.SubTitle = subTitle
		//temp = append(temp,t)
		books = append(books,&t)
	}

	response.Books = books

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
