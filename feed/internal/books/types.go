package books

import (
	"net/url"
	"strconv"
)

type ApiParams struct {
	Q            string
	Filter       string
	LangRestrict string
	MaxResults   int
	OrderBy      string
	PrintType    string
	StartIndex   int
}

const fields = "kind,totalItems,items(id,etag,kind,searchInfo/textSnippet,selfLink,volumeInfo(title,subtitle,description,mainCategory,categories,authors,averageRating,publishedDate,publisher,pageCount,printedPageCount,samplePageCount,dimensions,imageLinks,language,comicsContent,maturityRating,ratingsCount,seriesInfo))"

func (p ApiParams) String() string {
	values := url.Values{}
	values.Add("q", p.Q)
	values.Add("fields", fields)
	if p.Filter != "" {
		values.Add("filter", p.Filter)
	}
	if p.LangRestrict != "" {
		values.Add("langRestrict", p.LangRestrict)
	}
	if p.MaxResults > 0 {
		values.Add("maxResults", strconv.Itoa(p.MaxResults))
	}
	if p.OrderBy != "" {
		values.Add("orderBy", p.OrderBy)
	}
	if p.PrintType != "" {
		values.Add("printType", p.PrintType)
	}
	if p.StartIndex > 0 {
		values.Add("startIndex", strconv.Itoa(p.StartIndex))
	}
	return values.Encode()
}

type Volumes struct {
	Kind       string    `json:"kind,omitempty"`
	TotalItems int64     `json:"totalItems,omitempty"`
	Items      []*Volume `json:"items,omitempty"`
}

type Volume struct {
	Id         string            `json:"id,omitempty"`
	Etag       string            `json:"etag,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	SearchInfo *VolumeSearchInfo `json:"searchInfo,omitempty"`
	VolumeInfo *VolumeInfo       `json:"volumeInfo,omitempty"`
	SelfLink   string            `json:"selfLink,omitempty"`
}

type VolumeSearchInfo struct {
	TextSnippet string `json:"textSnippet,omitempty"`
}

type VolumeInfo struct {
	Title       string `json:"title,omitempty"`
	Subtitle    string `json:"subtitle,omitempty"`
	Description string `json:"description,omitempty"`

	MainCategory string   `json:"mainCategory,omitempty"`
	Categories   []string `json:"categories,omitempty"`

	Authors []string `json:"authors,omitempty"`

	AverageRating float64 `json:"averageRating,omitempty"`

	PublishedDate string `json:"publishedDate,omitempty"`
	Publisher     string `json:"publisher,omitempty"`

	PageCount        int64 `json:"pageCount,omitempty"`
	PrintedPageCount int64 `json:"printedPageCount,omitempty"`
	SamplePageCount  int64 `json:"samplePageCount,omitempty"`

	Dimensions *Dimensions `json:"dimensions,omitempty"`
	ImageLinks *ImageLinks `json:"imageLinks,omitempty"`

	Language       string            `json:"language,omitempty"`
	ComicsContent  bool              `json:"comicsContent,omitempty"`
	MaturityRating string            `json:"maturityRating,omitempty"`
	RatingsCount   int64             `json:"ratingsCount,omitempty"`
	SeriesInfo     *VolumeSeriesInfo `json:"seriesInfo,omitempty"`
}

type Dimensions struct {
	Height    string `json:"height,omitempty"`
	Thickness string `json:"thickness,omitempty"`
	Width     string `json:"width,omitempty"`
}

type ImageLinks struct {
	ExtraLarge     string `json:"extraLarge,omitempty"`
	Large          string `json:"large,omitempty"`
	Medium         string `json:"medium,omitempty"`
	Small          string `json:"small,omitempty"`
	SmallThumbnail string `json:"smallThumbnail,omitempty"`
	Thumbnail      string `json:"thumbnail,omitempty"`
}

type VolumeSeriesInfo struct {
	BookDisplayNumber    string `json:"bookDisplayNumber,omitempty"`
	Kind                 string `json:"kind,omitempty"`
	ShortSeriesBookTitle string `json:"shortSeriesBookTitle,omitempty"`

	VolumeSeries []*VolumeSeriesInfoVolumeSeries `json:"volumeSeries,omitempty"`
}

type VolumeSeriesInfoVolumeSeries struct {
	OrderNumber    int64  `json:"orderNumber,omitempty"`
	SeriesBookType string `json:"seriesBookType,omitempty"`
	SeriesId       string `json:"seriesId,omitempty"`
}
