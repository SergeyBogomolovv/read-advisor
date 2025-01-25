package handler

import (
	pb "github.com/SergeyBogomolovv/read-advisor/lib/api/gen/books"
	svc "github.com/SergeyBogomolovv/read-advisor/services/books/internal/service"
)

func convertVolumes(volumes *svc.Volumes) *pb.BookList {
	result := new(pb.BookList)
	result.Total = volumes.TotalItems
	result.Items = make([]*pb.Book, len(volumes.Items))
	for i, item := range volumes.Items {
		result.Items[i] = convertVolume(item)
	}
	return result
}

func convertVolume(v *svc.Volume) *pb.Book {
	book := new(pb.Book)
	book.Id = v.Id
	if v.VolumeInfo != nil {
		book.Title = v.VolumeInfo.Title
		book.Description = v.VolumeInfo.Description
		book.Categories = v.VolumeInfo.Categories
		book.Authors = v.VolumeInfo.Authors
		book.AverageRating = float32(v.VolumeInfo.AverageRating)
		book.RatingsCount = v.VolumeInfo.RatingsCount
		book.PublishedDate = v.VolumeInfo.PublishedDate
		book.Publisher = v.VolumeInfo.Publisher
		book.PageCount = v.VolumeInfo.PageCount
		if v.VolumeInfo.ImageLinks != nil {
			book.ImageLinks = &pb.ImageLinks{
				ExtraLarge:     v.VolumeInfo.ImageLinks.ExtraLarge,
				Large:          v.VolumeInfo.ImageLinks.Large,
				Medium:         v.VolumeInfo.ImageLinks.Medium,
				Small:          v.VolumeInfo.ImageLinks.Small,
				SmallThumbnail: v.VolumeInfo.ImageLinks.SmallThumbnail,
				Thumbnail:      v.VolumeInfo.ImageLinks.Thumbnail,
			}
		}
	}
	return book
}
