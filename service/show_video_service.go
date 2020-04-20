package service

import (
"go_demo_all/model"
"go_demo_all/serializer"
)

// CreateVideoService 视频详情服务
type ShowVideoService struct {
}

// Show 创建视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video,id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}

}
