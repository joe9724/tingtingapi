// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-openapi/runtime/middleware"
	"tingtingapi/models"
	"fmt"
	"tingtingbackend/var"
)

// BookDefaultHandlerFunc turns a function with the right signature into a book default handler
type BookDefaultHandlerFunc func(BookDefaultParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookDefaultHandlerFunc) Handle(params BookDefaultParams) middleware.Responder {
	return fn(params)
}

// BookDefaultHandler interface for that can handle valid book default params
type BookDefaultHandler interface {
	Handle(BookDefaultParams) middleware.Responder
}

// NewBookDefault creates a new http.Handler for the book default operation
func NewBookDefault(ctx *middleware.Context, handler BookDefaultHandler) *BookDefault {
	return &BookDefault{Context: ctx, Handler: handler}
}

/*BookDefault swagger:route GET /book/default Book bookDefault

根据用户信息匹配书本

根据用户信息匹配书本

*/
type BookDefault struct {
	Context *middleware.Context
	Handler BookDefaultHandler
}

func (o *BookDefault) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookDefaultParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok BookDefaultOK
	var response models.InlineResponse20077
	var books models.InlineResponse2007Books

	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	//db.Table("albums").Where(map[string]interface{}{"status":0}).Find(&categoryList).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))

	//db.Table("sub_category_items").Select("sub_category_items.name, category_album_relation.albumId").Joins("left join category_album_relation on category_album_relation.categoryId = sub_category_items.id and sub_category_items.id=?",1).Scan(&test)
	//db.Joins("JOIN sub_category_items ON sub_category_items.id = category_album_relation.albumId AND sub_category_items.id = ?",1).Where("credit_cards.number = ?", "411111111111").Find(&test)

	if (2>1) {rows, err := db.Table("books").Select("books.name, tag_book_relation.bookId,books.author_name,books.play_count,books.clips_number").Joins("left join tag_book_relation on tag_book_relation.bookId = books.id").Limit(3).Offset(0).Rows()
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

			err = rows.Scan(&name,&bookId,&authorName,&playCount,&clipsNumber)
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

}