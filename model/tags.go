package model

import (
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	TagId    string `gorm:"column:tag_id;type:char(36);unique;comment:主键;" json:"tag_id"`                 // 业务主键
	TeamId   string `gorm:"column:team_id;type:char(36);comment:团队ID;" json:"team_id"`                    // 团队
	TagName  string `gorm:"column:tag_name;type:varchar(255);comment:标签名称;" json:"tag_name"`              // 标签名称
	TagColor string `gorm:"column:tag_color;type:varchar(16);comment:标签颜色;" json:"tag_color"`             // 标签颜色
	TagType  *int32 `gorm:"column:tag_type;type:tinyint;default:1;comment:标签类型 1应用 2文档;" json:"tag_type"` // 标签类型
}

func (t *Tags) TableName() string {
	return "tags"
}
