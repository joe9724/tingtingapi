// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// AlbumListFavURL generates an URL for the album list fav operation
type AlbumListFavURL struct {
	Client      *string
	EndTs       *int64
	Grade       *string
	Imei        *string
	IsRecommend *int64
	MemberID    *string
	PageIndex   *int64
	PageSize    *int64
	StartTs     *int64
	Tags        *string
	Ts          *int64
	Val         *string
	Version     *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AlbumListFavURL) WithBasePath(bp string) *AlbumListFavURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AlbumListFavURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *AlbumListFavURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/album/list/fav"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/TingtingApi/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var client string
	if o.Client != nil {
		client = *o.Client
	}
	if client != "" {
		qs.Set("client", client)
	}

	var endTs string
	if o.EndTs != nil {
		endTs = swag.FormatInt64(*o.EndTs)
	}
	if endTs != "" {
		qs.Set("endTs", endTs)
	}

	var grade string
	if o.Grade != nil {
		grade = *o.Grade
	}
	if grade != "" {
		qs.Set("grade", grade)
	}

	var imei string
	if o.Imei != nil {
		imei = *o.Imei
	}
	if imei != "" {
		qs.Set("imei", imei)
	}

	var isRecommend string
	if o.IsRecommend != nil {
		isRecommend = swag.FormatInt64(*o.IsRecommend)
	}
	if isRecommend != "" {
		qs.Set("isRecommend", isRecommend)
	}

	var memberID string
	if o.MemberID != nil {
		memberID = *o.MemberID
	}
	if memberID != "" {
		qs.Set("memberId", memberID)
	}

	var pageIndex string
	if o.PageIndex != nil {
		pageIndex = swag.FormatInt64(*o.PageIndex)
	}
	if pageIndex != "" {
		qs.Set("pageIndex", pageIndex)
	}

	var pageSize string
	if o.PageSize != nil {
		pageSize = swag.FormatInt64(*o.PageSize)
	}
	if pageSize != "" {
		qs.Set("pageSize", pageSize)
	}

	var startTs string
	if o.StartTs != nil {
		startTs = swag.FormatInt64(*o.StartTs)
	}
	if startTs != "" {
		qs.Set("startTs", startTs)
	}

	var tags string
	if o.Tags != nil {
		tags = *o.Tags
	}
	if tags != "" {
		qs.Set("tags", tags)
	}

	var ts string
	if o.Ts != nil {
		ts = swag.FormatInt64(*o.Ts)
	}
	if ts != "" {
		qs.Set("ts", ts)
	}

	var val string
	if o.Val != nil {
		val = *o.Val
	}
	if val != "" {
		qs.Set("val", val)
	}

	var version string
	if o.Version != nil {
		version = *o.Version
	}
	if version != "" {
		qs.Set("version", version)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *AlbumListFavURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *AlbumListFavURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *AlbumListFavURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on AlbumListFavURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on AlbumListFavURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *AlbumListFavURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
