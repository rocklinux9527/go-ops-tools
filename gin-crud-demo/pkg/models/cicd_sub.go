
package models

// import "time"

// 发布订阅表设计

type CicdSub struct {
	ID       int64      `gorm:"column:id" json:"id"`                                          //  id
	AppName  string   `gorm:"column:appName" json:"appName"`                                          // 应用名称
	RobotKeys  string   `gorm:"column:robotKeys" json:"robotKeys"`                                    // 订阅组id
	LastModifyBy string `gorm:"column:lastModifyBy" json:"lastModifyBy"`                              // 最后修改人
	// CreateTime       *time.Time `gorm:"column:create_time" json:"create_time"`                        //  创建时间
	// LastModifyTime   *time.Time `gorm:"column:lastModifyTime" json:"lastModifyTime"`                  //  最后修改时间
	CreateTime       string `gorm:"column:create_time" json:"create_time"`                        //  创建时间
	LastModifyTime   string `gorm:"column:lastModifyTime" json:"lastModifyTime"`                  //  最后修改时间
}

func (*CicdSub) TableName() string {
	return "cicd_sub"
}
