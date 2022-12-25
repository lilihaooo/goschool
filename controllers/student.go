package controller

import (
	"aschool/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type StudentController struct {
}

// 学生列表
func (a *StudentController) List(c *gin.Context) {
	keyword, _ := c.GetQuery("keyword")
	var (
		id        int
		page      int = 1
		sort_type int = 1
	)
	if idstring, ok := c.GetQuery("id"); ok {
		id, _ = strconv.Atoi(idstring)
	}
	if pagestring, ok := c.GetQuery("page"); ok {
		page, _ = strconv.Atoi(pagestring)
	}
	if sorttypestring, ok := c.GetQuery("sort_type"); ok {
		sort_type, _ = strconv.Atoi(sorttypestring)
	}

	result, err := models.StudentList(keyword, id, page, sort_type)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "err": err})
		return
	}

	type info struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Sex       string `json:"sex"`
		Age       int    `json:"age"`
		State     string `json:"state"`
		Class     string `json:"class"`
		Grade     any    `json:"grade"`
		TeacherId int    `json:"teacher_id"`
		Created   int64  `json:"created"`
	}
	var res = make([]info, len(result))
	for i, v := range result {
		res[i].Id = v.Id
		res[i].Name = v.Name
		if v.Sex == 1 {
			res[i].Sex = "男"
		} else {
			res[i].Sex = "女"
		}
		res[i].Age = v.Age
		if v.State == 1 {
			res[i].State = "在线"
		} else {
			res[i].State = "毕业"
		}
		res[i].Class = v.Class.Name
		if v.Grade.Name == 0 {
			res[i].Grade = ""
		} else {
			res[i].Grade = v.Grade.Name
		}

		res[i].TeacherId = v.Class.TeacherId
		res[i].Created = v.Created
	}

	//count, _ := models.Count(models.Student{})
	count, _ := models.Count(new(models.Student))

	c.JSON(http.StatusOK, gin.H{"code": 1, "list": res, "total": count})
}

// 添加学生 or 修改学生 (看是否传了id)
func (a *StudentController) Save(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	sex := c.PostForm("sex")
	age := c.PostForm("age")
	grade := c.PostForm("grade")
	class_id := c.PostForm("class_id")
	state := c.PostForm("state")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "err": "姓名不能为空"})
		return
	}
	student := models.Student{}
	student.Name = name
	student.Sex, _ = strconv.Atoi(sex)
	student.Age, _ = strconv.Atoi(age)
	student.GradeId, _ = strconv.Atoi(grade)
	student.ClassId, _ = strconv.Atoi(class_id)
	student.State, _ = strconv.Atoi(state)
	student.Created = time.Now().Unix()

	if idint, _ := strconv.Atoi(id); idint == 0 {
		//新增
		if err := models.CreateStudent(&student); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "err": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "添加成功"})
	} else {
		//修改
		student.Id = idint
		if err := models.UpdateStudent(&student); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "err": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "修改成功"})
	}

}

// 删除学生
func (a *StudentController) StuDel(c *gin.Context) {
	// 获取参数
	idString, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(idString)
	// 执行删除
	if err := models.DeleteStudent(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "删除成功"})
}

// 学生详情
/*
1,
*/

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
