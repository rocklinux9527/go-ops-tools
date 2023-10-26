package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"github.com/spf13/viper"
	"eagle-cicd-sub/pkg/models"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func LoadDbConfig() (*gorm.DB, error)  {
	// 使用Viper加载配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// 从配置文件中获取数据库连接信息
	dbConfig := viper.Sub("mysql")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.GetString("user"),
		dbConfig.GetString("password"),
		dbConfig.GetString("host"),
		dbConfig.GetInt("port"),
		dbConfig.GetString("dbname"),
		dbConfig.GetString("charset"),
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 连接池配置
		PrepareStmt: true, // 预编译语句
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 设置连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	// 设置最大打开的连接数
	sqlDB.SetMaxOpenConns(dbConfig.GetInt("max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(dbConfig.GetInt("max_idle_connections"))
	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.GetInt("max_life_seconds")) * time.Second)

	// defer db.Close()  v2 版本不需要这个了

	if err := db.AutoMigrate(&models.CicdSub{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	return db,nil
}
