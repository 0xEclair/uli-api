package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// ShowVideoService 视频投稿服务
type ShowVideoService struct {
}

// Show 创建视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err:= model.DB.First(&video,id).Error

	if err!=nil{
		return serializer.Response{
			Status:404,
			Msg: "视频信息获取失败",
			Error: err.Error(),
		}
		//c.JSON(404,serializer.Response{
		//	Status:404,
		//	Msg:"视频信息获取失败",
		//	Error:err.Error(),
		//})
	}
	return serializer.Response{
		Data:   serializer.BuildVideo(video),
	}
}
