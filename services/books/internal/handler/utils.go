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
	return &pb.Book{
		Id:            v.Id,
		Title:         v.VolumeInfo.Title,
		Description:   v.VolumeInfo.Description,
		Categories:    v.VolumeInfo.Categories,
		Authors:       v.VolumeInfo.Authors,
		AverageRating: float32(v.VolumeInfo.AverageRating),
		RatingsCount:  v.VolumeInfo.RatingsCount,
		PublishedDate: v.VolumeInfo.PublishedDate,
		Publisher:     v.VolumeInfo.Publisher,
		PageCount:     int32(v.VolumeInfo.PageCount),
		ImageLinks: &pb.ImageLinks{
			ExtraLarge:     v.VolumeInfo.ImageLinks.ExtraLarge,
			Large:          v.VolumeInfo.ImageLinks.Large,
			Medium:         v.VolumeInfo.ImageLinks.Medium,
			Small:          v.VolumeInfo.ImageLinks.Small,
			SmallThumbnail: v.VolumeInfo.ImageLinks.SmallThumbnail,
			Thumbnail:      v.VolumeInfo.ImageLinks.Thumbnail,
		},
	}
}
