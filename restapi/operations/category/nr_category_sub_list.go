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
	"tingtingapi/var"

	"fmt"
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
	//var categoryList models.InlineResponse20023Data

	//根据categoryId找到对应的子类
	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var CategoryList []models.SubCategoryItem

	db.Table("sub_category_items").Where("category_id=?",*(Params.CategoryID)).Where("status=?",0).Find(&CategoryList)
	for j:=0; j<len(CategoryList);j++  {
		//根据subCategoryId请求出前6条albumList
		var AlbumList []models.Album
		db.Table("albums").Select("albums.name, albums.id,albums.author_avatar,albums.author_name,albums.books_number,albums.icon,albums.play_count,albums.sub_title,albums.time").Joins("left join category_album_relation on category_album_relation.albumId = albums.id").Where("albums.status=? and category_album_relation.status=? and category_album_relation.categoryId=?",0,0,CategoryList[j].ID).Find(&AlbumList)

		for k:=0; k<len(AlbumList);k++  {
			AlbumList[k].AlbumName = AlbumList[k].Name
			AlbumList[k].AlbumID = AlbumList[k].ID
		}
		CategoryList[j].AlbumList = AlbumList
	}

	var BookList []models.Book
	rows, err := db.Table("books").Select("books.name, category_book_relation.bookId,books.author_name,books.play_count,books.clips_number,books.icon,books.sub_title").Joins("left join category_book_relation on category_book_relation.bookId = books.id").Where("category_book_relation.categoryId=?",Params.CategoryID).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize))).Rows()
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
		BookList = append(BookList,t)
	}
	fmt.Println("bookList is",BookList)

	/*response.Data = categoryList

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
			album.Play_Count = 123
			album.Books_Number = 34
			album.Value = 118
			album.Icon = "http://tingting-resource.bitekun.xin/resource/image/cover.png"
			album.Cover = "http://tingting-resource.bitekun.xin/resource/image/cover.png"
			albums = append(albums,album)
		}
		sub.AlbumList = albums
		//
		SubCategories = append(SubCategories,sub)
	}*/

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status
	response.DataList = CategoryList
	response.BookList = BookList

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
