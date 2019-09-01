package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// ShowVideoService 视频投稿服务
type ListVideoService struct {
}

// Show 创建视频
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err:= model.DB.Find(&videos).Error

	if err!=nil{
		return serializer.Response{
			Status:50000,
			Msg: "视频列表查询错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data:   serializer.BuildVideos(videos),
	}
}
