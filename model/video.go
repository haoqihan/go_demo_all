package model

import "github.com/jinzhu/gorm"

// Video 的模型
type Video struct {
	gorm.Model
	Title string
	Info  string
}
