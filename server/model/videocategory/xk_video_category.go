package videocategory

import (
	"time"

	"github.com/shopspring/decimal"
)

type VideoCategory struct {
	VID         uint            `gorm:"primarykey;comment:主键id" json:"vid"`                              // 主键ID
	CreatedTime time.Time       `gorm:"type:datetime(0);autoCreateTime;comment:创建时间" json:"createdTime"` // 创建时间
	UpdatedTime time.Time       `gorm:"type:datetime(0);autoCreateTime;comment:更新时间" json:"updatedTime"` // 更新时间
	Title       string          `gorm: "type:varchar(100);comment:分类标题" json:"title"`                    // 分类标题
	Description string          `gorm:"type:varchar(400);comment:分类描述" json:"description"`               // 分类描述
	ParentID    int             `gorm:"type:int;default:0;comment:父分类ID" json:"parentID"`                // 父分类ID
	Sorted      int             `gorm:"type:int;default:0;comment:排序" json:"sorted"`                     // 排序
	Status      int             `gorm:"type:tinyint(1);default:1;comment:是否激活" json:"status "`           // 是否激活
	IsDelete    int             `gorm:"type:varchar(255);comment:分类图标" json:"isDelete"`                  // 分类删除
	Price       decimal.Decimal `gorm:"type:decimal(10,2);default:0.00;comment:分类价格" json:"price"`       // 分类价格
}

func (VideoCategory) TableName() string {
	return "xk_video_category"
}
