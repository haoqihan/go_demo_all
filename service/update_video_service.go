package service

import (
	"go_demo_all/model"
	"go_demo_all/serializer"
)

// UpdateVideoService 更新视频的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=300"`
}

// Create 创建视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video,id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	if err = model.DB.Save(&video).Error;err != nil{
		return serializer.Response{
			Code:  5002,
			Msg:   "视频修改失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
