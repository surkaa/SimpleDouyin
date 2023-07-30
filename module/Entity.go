package module

import "time"

// Entity 所有表的module都应由以下字段 并继承以下字段
type Entity struct {
	ID       int64     `json:"id"`                  // 主键
	IsDelete int32     `json:"is_delete,omitempty"` // 逻辑删除
	CreateAt time.Time `json:"create_at,omitempty"` // 创建时间
	UpdateAt time.Time `json:"update_at,omitempty"` // 更新时间
}
