package service

import (
	"eagle-cicd-sub/api/v1"
	"eagle-cicd-sub/logger"
	"eagle-cicd-sub/pkg/database"
	"eagle-cicd-sub/pkg/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

// 日志记录

func setupLogger() {
	logRotation := logger.LogBck()
	log.SetOutput(logRotation)
}

// sql 连接器函数

func initializeDatabase() (*gorm.DB, error) {
	db, err := database.LoadDbConfig()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 函数字段检验 checks for empty fields.

func validateCreateData(appName, robotkeys, modifyBy string) error {
	if appName == "" || robotkeys == "" || modifyBy == "" {
		return errors.New("please fill in all required fields")
	}
	return nil
}



 // 处理sub 数据返回数据结构逻辑

func mapSubListData(cicdSubList []models.CicdSub) []v1.GetSubData {
	var subDataList []v1.GetSubData
	for _, cicdSub := range cicdSubList {
		if cicdSub.AppName != "" {
			subData := v1.GetSubData{
				ID:            cicdSub.ID,
				AppName:       cicdSub.AppName,
				LastModifyBy:  cicdSub.LastModifyBy,
				LastModifyTime: cicdSub.LastModifyTime,
				CreateTime:    cicdSub.CreateTime,
				RobotKeys:     cicdSub.RobotKeys,
			}
			subDataList = append(subDataList, subData)
		}
	}
	return subDataList
}


// Create 创建 sub 表记录

func CreateDataCicdSub(appName, robotkeys, modifyBy string) (int, error) {
	if err := validateCreateData(appName, robotkeys, modifyBy); err != nil {
		return 1, err
	}
	db, err := initializeDatabase()
	setupLogger()
	log.Printf("create cicd imput xargs appName: %s robotkeys: %s robotkeys: %s",appName,robotkeys, modifyBy)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return 1, err
	}
	currentTime := time.Now()
	timestamp := currentTime.Format("2006-01-02 15:04:05")  // 格式化时间为字符串
	newCicdSub := &models.CicdSub{
		AppName:        appName,
		RobotKeys:      robotkeys,
		LastModifyBy:   modifyBy,
		CreateTime:     timestamp,
		LastModifyTime: timestamp,
	}
	code, err := createCicdSub(db, newCicdSub)
	if err != nil {
		log.Printf("create cicd return code: %d, error: %v", code, err)
	}
	return code,err
}

func createCicdSub(db *gorm.DB, cicdSub *models.CicdSub) (int, error) {
	// 检查是否 已存在相同的 AppName
	setupLogger()
	var existingCicdSub models.CicdSub

	if err := db.Where("appName = ?", cicdSub.AppName).First(&existingCicdSub).Error; err == nil {
		// 存在相同的 AppName 以存在返回 2
		log.Printf("create cicd 应用已经存在 App:%s... , error: %v",cicdSub.AppName, err)
		return 2, errors.New("record with the same appName already exists")
	}
	if err := db.Create(cicdSub).Error; err != nil {
		// 创建失败 返回1
		log.Printf("create cicd 创建订阅失败 AppName:%s... , error: %v",cicdSub.AppName, err)
		return 1, err
	}
	// 创建成功 返回0
	return 0, nil
}


// Get SubList 分页查询 Sub记录

func GetDataCicdSubList( page, pageSize int, appName string )(*v1.GetCicdSubListResponse, error) {
	setupLogger()
	// fmt.Println("查询参数进来了",page,pageSize,appName)

	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return nil, err
	}
	var subDataList []v1.GetSubData

	cicdSubList, err := getCicdSubList(db, page, pageSize, appName)

	if err != nil {
		fmt.Printf("Error getting CicdSub list: %v\n", err)
		return nil,err
	}
	// 映射查询结果到 GetSubData 结构体
	if cicdSubList != nil && len(cicdSubList) > 0 {
		for _, cicdSub := range cicdSubList {
			if cicdSub.AppName != "" {
				subData := v1.GetSubData{
					ID:             cicdSub.ID,
					AppName:        cicdSub.AppName,
					LastModifyBy:   cicdSub.LastModifyBy,
					LastModifyTime: cicdSub.LastModifyTime,
					CreateTime:     cicdSub.CreateTime,
					RobotKeys:      cicdSub.RobotKeys,
				}
				subDataList = append(subDataList, subData)
			}
		}
	}else {
		fmt.Printf("查询输入的应用不存在: AppName %s", appName)
		subDataList = []v1.GetSubData{}  // Set it to an empty slice
	}

	// 创建 GetCicdSubListResponse 结构体

	response := v1.GetCicdSubListResponse{
		Code: 0, // 设置状态码
		Msg:  "", // 设置状态信息
	}

	// 设置 Data.List 为映射后的 GetSubData 结构体切片
	response.Data.List = subDataList
	response.Data.Total = len(subDataList) // 设置总数
	return &response, nil
}


func getCicdSubList(db *gorm.DB, page, pageSize int, appName string) ([]models.CicdSub, error) {
	setupLogger()
	offset := (page - 1) * pageSize
	var cicdSubList []models.CicdSub
	query := db.Offset(offset).Limit(pageSize)

	if appName != "" {
		query = query.Where("appName = ?", appName)
	}

	if err := query.Find(&cicdSubList).Error; err != nil {
		log.Printf("query cicd 出现问题 App:%s... , error: %v",appName, err)
		return nil, err
	}
	return cicdSubList, nil
}


// Update Sub 更新记录

func UpdateDataCicdSub(Id int64, appName, robotkeys, modifyBy string) (int, error) {
	setupLogger()
	if err := validateCreateData(appName, robotkeys, modifyBy); err != nil {
		return 1, err
	}
	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return 1, err
	}
	currentTime := time.Now()
	// 格式化时间为字符串
	timestamp := currentTime.Format("2006-01-02 15:04:05")
	updateSub := &models.CicdSub{
		AppName:        appName,
		RobotKeys:      robotkeys,
		LastModifyBy:   modifyBy,
		CreateTime:     "",
		LastModifyTime: timestamp,
	}
	code, err := updateCicdSub(db,Id,updateSub )
	if err != nil {
		fmt.Printf("Error Update CicdSub: %v\n", err)
		return 1, err
	}
	return code, err

}

func updateCicdSub(db *gorm.DB, id int64, cicdSub *models.CicdSub) (int, error) {
	setupLogger()
	var existingCicdSub models.CicdSub
	if err := db.Where("id = ?", id).First(&existingCicdSub).Error; err != nil {
		// 记录不存在，返回 2
		log.Printf("update  Cicd Sub 数据不存在, Id: %d , error: %v", id, err)
		return 2, fmt.Errorf("Record not found ")
	}
	// 更新记录
	if err := db.Model(&existingCicdSub).Clauses(clause.OnConflict{UpdateAll: true}).Updates(cicdSub).Error; err != nil {
		// 更新失败，返回 1
		log.Printf("Update  Cicd Sub 出现问题 App:%s... , error: %v", err)
		return 1, err
	}

	// 更新成功，返回 0
	return 0, nil
}



// Delete Sub 删除 Sub记录

func DeleteCicdSub(Id int64, appName string) (int, error) {
	setupLogger()
	if appName == "" || Id == 0 {
		// 至少一个字段为空的情况下执行的操作
		fmt.Println("请填写所有必填字段,缺少应用名称和Id")
		return 1, nil
	}
	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return 1, err
	}
	code, err := deleteCicdSub(db,Id,appName)
	if err != nil {
		fmt.Printf("Error delere CicdSub Record: %v\n", err)
	}
	return code, err
}


func deleteCicdSub(db *gorm.DB, id int64, appName string)( int, error) {
	setupLogger()
	// 检查记录是否存在 id 和应用名称是否存在 存在删除不存在给返回

	var existingCicdSub models.CicdSub

	if err := db.Where("id = ? AND appName = ?", id, appName).First(&existingCicdSub).Error; err != nil {
		return 2, fmt.Errorf(" Record not found")
	}

	if err := db.Delete(&existingCicdSub).Error; err != nil {
		// 删除失败，返回 1
		return 1, err
	}
	// 删除成功，返回 0
	return 0, nil
}
