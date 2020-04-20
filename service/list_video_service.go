package service

import (
	"go_demo_all/model"
	"go_demo_all/serializer"
)

// ListVideoService 视频列表服务
type ListVideoService struct {
}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code:  5000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}

}
