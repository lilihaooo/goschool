package main

import (
	"aschool/conn"
	"aschool/logger"
	_ "aschool/logger"
	"aschool/models"
	"aschool/routers"
	"aschool/settings"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed,err:%v", err)
	}
	// 日志初始化
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// MySQL
	// 连接数据库
	err := conn.InitMySQL(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer conn.CloseMysql() // 程序退出关闭数据库连接

	//连接redis
	err = conn.InitRedis(settings.Conf.RedisConfig)
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer conn.CloseRedis() // 程序退出关闭数据库连接

	// Models －> 数据表
	// 模型绑定
	//dao.DB.AutoMigrate(new(models.User), new(models.Category),
	//	new(models.Post), new(models.Config), new(models.Comment))
	conn.DB.AutoMigrate(new(models.User), new(models.Student), new(models.Class), new(models.Grade), new(models.Teacher))

	//注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf("127.0.0.1:%d", settings.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

	//wg.Add(10000)
	//for i := 0; i < 1000; i++ {
	//	go add()
	//}
	//wg.Wait()
}

func add() {
	for i := 0; i < 1000; i++ {
		var student = models.Student{
			Name:    "小华" + strconv.Itoa(i),
			Created: time.Now().Unix(),
			ClassId: 2,
		}
		err := conn.DB.Create(&student).Error
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	wg.Done()
}
