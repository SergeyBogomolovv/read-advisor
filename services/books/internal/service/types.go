package service

import (
	"errors"
	"net/url"
	"strconv"
)

var ErrNotFound = errors.New("not found")

type ApiParams struct {
	LangRestrict string
	OrderBy      string
	MaxResults   int
	StartIndex   int
}

const volumesFields = "totalItems,items(id,volumeInfo(title,description,categories,authors,averageRating,publishedDate,publisher,pageCount,imageLinks,ratingsCount))"

const volumeFields = "id,volumeInfo(title,description,categories,authors,averageRating,publishedDate,publisher,pageCount,imageLinks,ratingsCount)"

func (p ApiParams) ToValues() url.Values {
	values := url.Values{}
	if p.LangRestrict != "" {
		values.Set("langRestrict", p.LangRestrict)
	}
	if p.MaxResults > 0 {
		values.Set("maxResults", strconv.Itoa(p.MaxResults))
	}
	if p.OrderBy != "" {
		values.Set("orderBy", p.OrderBy)
	}
	if p.StartIndex > 0 {
		values.Set("startIndex", strconv.Itoa(p.StartIndex))
	}
	return values
}

type Volumes struct {
	TotalItems int32     `json:"totalItems,omitempty"`
	Items      []*Volume `json:"items,omitempty"`
}

type Volume struct {
	Id         string      `json:"id,omitempty"`
	VolumeInfo *VolumeInfo `json:"volumeInfo,omitempty"`
}

type VolumeInfo struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Authors     []string `json:"authors,omitempty"`

	AverageRating float64 `json:"averageRating,omitempty"`
	RatingsCount  int64   `json:"ratingsCount,omitempty"`

	PublishedDate string `json:"publishedDate,omitempty"`
	Publisher     string `json:"publisher,omitempty"`

	PageCount int64 `json:"pageCount,omitempty"`

	ImageLinks *ImageLinks `json:"imageLinks,omitempty"`
}

type ImageLinks struct {
	ExtraLarge     string `json:"extraLarge,omitempty"`
	Large          string `json:"large,omitempty"`
	Medium         string `json:"medium,omitempty"`
	Small          string `json:"small,omitempty"`
	SmallThumbnail string `json:"smallThumbnail,omitempty"`
	Thumbnail      string `json:"thumbnail,omitempty"`
}
