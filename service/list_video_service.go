package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// ShowVideoService 视频投稿服务
type ListVideoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	total:=0

	if service.Limit ==0 {
		service.Limit=8
	}

	if err:=model.DB.Model(model.Video{}).Count(&total).Error; err!=nil {
		return serializer.Response{
			Status:50000,
			Msg: "视频列表查询错误",
			Error: err.Error(),
		}
	}

	if err:=model.DB.Limit(service.Limit).Offset(service.Start).Find(&videos).Error;err!=nil{
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库链接错误",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(
		serializer.BuildVideos(videos),
		uint(total))
}
