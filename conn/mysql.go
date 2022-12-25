package conn

import (
	"fmt"

	"aschool/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GORM 来访问MySQL
var (
	DB *gorm.DB
)

func InitMySQL(cfg *settings.MySQLConfig) (err error) {
	//连接mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	if err := DB.DB().Ping(); err != nil {
		fmt.Println("连接MySQL失败")
	}
	//禁用复数形式
	DB.SingularTable(true)

	//为表名添加前缀
	// models 定义的数据库影射struct －> MySQL 表
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "lh_" + defaultTableName
	}
	return nil
}

func CloseMysql() {
	DB.Close()
}
