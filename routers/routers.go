package routers

import (
	"aschool/controllers"
	"aschool/logger"
	"aschool/middlewares"
	"aschool/settings"
	"github.com/gin-gonic/gin"
)

func helloHandler(ctx *gin.Context) {
	// hello 函数
	ctx.JSON(200, gin.H{
		"mode": settings.Conf.Mode,
		"host": settings.Conf.MySQLConfig.Host,
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 中间件
	r.Use(logger.GinLogger())

	r.GET("/hello", helloHandler)

	////前端系统
	//blog := controller.BlogController{}
	//v1Group := r.Group("v1")
	//{
	//	// 主页
	//	v1Group.GET("/home", blog.GetHome)
	//	// 文章列表
	//	v1Group.GET("/article", blog.GetArticleList)
	//	// 文章详情
	//	v1Group.GET("/detail/:id", blog.GetArticleDetail)
	//	// 创建评价
	//	v1Group.POST("/comment", blog.CreateComment)
	//	// 资源
	//	v1Group.GET("/resource", blog.Resource)
	//	// 关于我们
	//	v1Group.GET("/about", blog.GetAbout)
	//}

	//后端系统
	v2Group := r.Group("admin")
	v2Group.Use(middlewares.JWTMiddleware())
	admin := controller.AdminController{}
	student := controller.StudentController{}
	class := controller.ClassController{}
	{
		// 主页面登录
		// localhost:9002/admin/login
		v2Group.POST("/login", admin.Login)

		//注册
		// localhost:9002/admin/register
		v2Group.POST("/register", admin.Register)
		//v2Group.GET("/logout", admin.Logout)

		// 主页
		// localhost:9002/admin/main
		v2Group.GET("/main", admin.Main)

		//学生列表
		v2Group.GET("/stulist", student.List)
		//添加or修改
		v2Group.POST("/stusave", student.Save)
		//删除学生
		v2Group.GET("/studel", student.StuDel)

		//班级列表
		v2Group.GET("/classlist", class.List)
		//班级详情
		v2Group.GET("/classdetail", class.GetClassDetailInfo)
		v2Group.GET("/classdetail2", class.GetClassDetailInfo2)

		//////提交更新
		//v2Group.POST("/addconfig", admin.AddConfig)
		//// 博文列表
		//v2Group.GET("/index", admin.Index)
		//
		//// 博文添加
		////// 显示
		//v2Group.GET("/article", admin.Article)
		////// 文章保存
		//v2Group.POST("/save", admin.Save)
		////// 文章删除
		//v2Group.GET("/delete", admin.PostDel)
		//
		//// 类目主页
		//v2Group.GET("/category", admin.Category)
		//// 类目增加
		//v2Group.GET("/categoryadd", admin.Categoryadd)
		//// 类目保存
		//v2Group.POST("/categorysave", admin.CategorySave)
		//// 类目删除
		//v2Group.GET("/categorydel", admin.CategoryDel)
	}
	return r
}
