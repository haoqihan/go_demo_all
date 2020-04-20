package api

import (
	"github.com/gin-gonic/gin"
	service2 "go_demo_all/service"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	service := service2.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	service := service2.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)

}

// ListVideo 视频列表接口
func ListVideo(c *gin.Context) {
	service := service2.ListVideoService{}
	res := service.List()
	c.JSON(200, res)
}

//  UpdateVideo 更新视频的接口
func UpdateVideo(c *gin.Context) {
	service := service2.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// DeleteVideo 删除视频接口
func DeleteVideo(c *gin.Context) {
	service := service2.DeleteVideoService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)

}
