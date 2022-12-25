package controller

import (
	"aschool/models"
	"aschool/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AdminController struct {
}

// 登录
func (a *AdminController) Login(c *gin.Context) {

	// 1、 接收参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 2、 数据库验证OK
	user, _ := models.Login(username, password)

	if len(user) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "用户名或者密码错误"})
		return
	}
	// 3、返回token
	//fmt.Println(user[0])
	token := util.GenerateToken(user[0])
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "登陆成功", "token": token})
	return

}

// 注册
func (a *AdminController) Register(c *gin.Context) {
	password := c.PostForm("password")
	username := c.PostForm("username")
	if password == "" || username == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "err": "用户名或者密码不能为空"})
		return
	}
	user := models.User{}
	password = models.EncryptPassword([]byte(password))
	user.Username = username
	user.Password = password
	user.Created = time.Now().Unix()
	user.State = 1
	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "注册成功"})
}

// 主页
func (a *AdminController) Main(c *gin.Context) {
	claims, _ := c.Get("claims")
	//类型断言
	value, _ := claims.(util.Claims)
	c.JSON(http.StatusOK, gin.H{"code:": 1, "user": value.Username, "id": value.UserId})
	return
}

//系统配置信息展示
//func (a *AdminController) Config(c *gin.Context) {
//	// load config list
//	result, _ := models.ConfigList()
//
//	// name -> value
//	options := make(map[string]string)
//	for _, v := range result {
//		options[v.Name] = v.Value
//	}
//	// config.html
//	dataList := gin.H{}
//	dataList["config"] = options
//	c.HTML(http.StatusOK, "config.html", dataList)
//}

//系统配置信息更新
//func (a *AdminController) AddConfig(c *gin.Context) {
//	options := make(map[string]string)
//	mp := make(map[string]*models.Config)
//
//	// 获取所有的配置
//	result, _ := models.ConfigList()
//	for _, v := range result {
//		options[v.Name] = v.Value
//		mp[v.Name] = v
//	}
//	// 按照每个字段更新KV config表数据
//	// 更新 (ID)
//	// 插入
//
//	if c.Request.Method == "POST" {
//		keys := []string{"url", "title", "keywords", "description", "email", "timezone", "start", "qq"}
//
//		for _, key := range keys {
//			val := c.PostForm(key) // form 表单数据
//			if _, ok := mp[key]; !ok {
//				options[key] = val
//				models.UpdateConfig(&models.Config{Name: key, Value: val})
//			} else {
//				opt := mp[key]
//				if err := models.UpdateConfig(&models.Config{Id: opt.Id, Name: key, Value: val}); err != nil {
//					continue
//				}
//			}
//
//		}
//	}
//
//	// 给个提示语
//	msg := "数据保存成功"
//	info := "<script> alert('" + msg + "');window.history.go(-1); </script>"
//	c.Writer.WriteString(info)
//	// 跳转
//	c.Redirect(http.StatusMovedPermanently, "/admin/config")
//}

func AdminList(c *gin.Context) {

}

//后台首页
//func (a *AdminController) Index(c *gin.Context) {
//	dataList := gin.H{}
//
//	// 1. 类目
//	categoryList, _ := models.CategoryList()
//	dataList["categorys"] = categoryList
//	// 2. 文章信息
//	var (
//		page     int = 1
//		pagesize int = 5
//		list     []*models.Post
//		offset   int
//	)
//
//	if pageStr, ok := c.GetQuery("page"); ok {
//		page, _ = strconv.Atoi(pageStr)
//	}
//	offset = (page - 1) * pagesize
//	// 3. 总的记录数据
//	count, _ := models.Count(new(models.Post))
//
//	if count > 0 {
//		list, _ = models.GetArtileList(offset, pagesize)
//	}
//
//	// 分页
//	pagebar := util.NewPager(page, int(count), pagesize, "/admin/index", true).ToString()
//	dataList["pagebar"] = template.HTML(pagebar)
//	dataList["list"] = list
//	// 4. list.html
//	c.HTML(http.StatusOK, "list.html", dataList)
//}

//博文添加
//func (a *AdminController) Article(c *gin.Context) {
//	dataList := gin.H{}
//	// 1. 获取所有类目
//	categoryList, _ := models.CategoryList()
//	dataList["categorys"] = categoryList
//
//	idString, _ := c.GetQuery("id")
//	if idString != "" { // 修改展示操作
//		id, _ := strconv.Atoi(idString)
//		article, _ := models.GetDetailById(id)
//		dataList["post"] = article
//	}
//	// 2. c.HTML
//	c.HTML(http.StatusOK, "_form.html", dataList)
//}

//保存
//func (a *AdminController) Save(c *gin.Context) {
//	post := models.Post{}
//	post.UserId = 1
//	post.Title = c.PostForm("title")
//	post.Content = c.PostForm("content")
//
//	is_top := c.PostForm("is_top")
//	if is_top == "" {
//		post.IsTop = 0
//	} else {
//		post.IsTop, _ = strconv.Atoi(is_top)
//	}
//
//	post.Types, _ = strconv.Atoi(c.PostForm("types"))
//	post.Tags = c.PostForm("tags")
//	post.Url = c.PostForm("url")
//	post.CategoryId, _ = strconv.Atoi(c.PostForm("cate_id"))
//	post.Info = c.PostForm("info")
//	post.Image = c.PostForm("Image")
//	post.Created = time.Now()
//	post.Updated = time.Now()
//
//	id, _ := strconv.Atoi(c.PostForm("id"))
//	if id == 0 {
//		// 新增
//		models.CreatePost(&post)
//	} else {
//		// 更新
//		post.Id = id
//		models.UpdatePost(&post)
//	}
//	c.Redirect(http.StatusMovedPermanently, "/admin/index")
//}

// 文章删除
//func (a *AdminController) PostDel(c *gin.Context) {
//
//	// 获取参数
//	idString, _ := c.GetQuery("id")
//	// 执行删除
//	models.DeletePost(idString)
//	// 直接跳转
//	c.Redirect(http.StatusMovedPermanently, "/admin/index")
//}

//类目
//func (a *AdminController) Category(c *gin.Context) {
//
//	categoryList, _ := models.CategoryList()
//
//	dataList := gin.H{
//		"categorys": categoryList,
//	}
//
//	c.HTML(http.StatusOK, "category.html", dataList)
//}

//添加修改类目
//func (a *AdminController) Categoryadd(c *gin.Context) {
//	dataList := gin.H{} // 空的时候 －>  添加添加界面
//	idString, _ := c.GetQuery("id")
//	if idString != "" { // 编辑
//		id, _ := strconv.Atoi(idString)
//		categories, _ := models.GetCategoryById(id)
//		dataList["cate"] = categories[0]
//	}
//	c.HTML(http.StatusOK, "category_add.html", dataList)
//}

//保存类目
//func (a *AdminController) CategorySave(c *gin.Context) {
//	//  获取参数POST
//	name := c.PostForm("name")
//	idString := c.PostForm("id")
//	fmt.Println("name = ", name)
//	fmt.Println("idString = ", idString)
//
//	var category models.Category
//	category.Name = name
//	category.Updated = time.Now()
//	if idString == "" {
//		// insert
//		category.Created = time.Now()
//	} else {
//		//update
//		nilTime := time.Time{}
//		id, _ := strconv.Atoi(idString)
//
//		categories, _ := models.GetCategoryById(id)
//
//		if categories[0].Created == nilTime {
//			category.Created = time.Now()
//		} else {
//			category.Created = categories[0].Created
//		}
//		category.Id = id
//	}
//
//	models.UpdateCategory(&category)
//
//	c.Redirect(http.StatusMovedPermanently, "/admin/category")
//}

// 类目删除
//func (a *AdminController) CategoryDel(c *gin.Context) {
//
//	idString, _ := c.GetQuery("id")
//	models.DeleteCategory(idString)
//
//	c.Redirect(http.StatusMovedPermanently, "/admin/category")
//}
