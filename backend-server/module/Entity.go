package module

import (
	"SimpleDouyin/util"
	"time"
)

// Entity 所有表的module都应由以下字段 并继承以下字段
type Entity struct {
	ID       int64     `json:"id,omitempty" gorm:"primaryKey"`                                                            // 主键
	IsDelete int32     `json:"is_delete,omitempty"`                                                                       // 逻辑删除
	CreateAt time.Time `json:"create_at,omitempty" gorm:"not null;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdateAt time.Time `json:"update_at,omitempty" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
}

const (
	BeforeDel = 0
	Deleted   = 1
)

// CreateEntity 快速创建一个实体
func CreateEntity() Entity {
	return Entity{
		ID:       util.SF.NextID(),
		IsDelete: BeforeDel,
	}
}
