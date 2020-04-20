package serializer

import "go_demo_all/model"

// User 用户序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildVideo 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		videos = append(videos, BuildVideo(item))
	}
	return videos
}
