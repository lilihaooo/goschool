package controller

import (
	"aschool/conn"
	"aschool/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ClassController struct {
}

// 班级列表
type classList struct {
	Id           int
	Name         string
	Grade        int
	ClassTeacher string
	State        string
	Created      int64
}

func (a *ClassController) List(c *gin.Context) {
	var page = 1
	result, err := models.ClassList(page)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "err": err})
		return
	}
	list := make([]classList, len(result))
	for i, v := range result {
		list[i].Id = v.Id
		list[i].Name = v.Name
		list[i].Grade = v.Grade.Name
		list[i].ClassTeacher = v.Teacher.Name
		if v.State == 1 {
			list[i].State = "正常"
		} else {
			list[i].State = "异常"
		}
		list[i].Created = v.Created
	}
	count, _ := models.Count(models.Class{})

	c.JSON(http.StatusOK, gin.H{"code": 1, "list": list, "total": count})
}

type student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created int64  `json:"created"`
}

type classInfo struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Grade        int       `json:"grade"`
	ClassTeacher string    `json:"class_teacher"`
	State        string    `json:"state"`
	Created      int64     `json:"created"`
	Students     []student `json:"students"`
}

func (a *ClassController) GetClassDetailInfo2(c *gin.Context) {
	//使用缓存
	var id int
	if idstring, ok := c.GetQuery("id"); ok {
		id, _ = strconv.Atoi(idstring)
	}
	//先获该班级的信息
	result, err := models.GetOneClass(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "err": err})
		return
	}

	//获得该班级下的学生
	classId := result.Id
	stu, err := models.GetOneClassStudents(classId)
	students := make([]student, len(stu))
	for i, v := range stu {
		students[i].Id = v.Id
		students[i].Name = v.Name
		students[i].Created = v.Created
	}
	classinfo := classInfo{}
	classinfo.Id = result.Id
	classinfo.Name = result.Name
	classinfo.Grade = result.Grade.Name
	classinfo.ClassTeacher = result.Teacher.Name
	if result.State == 1 {
		classinfo.State = "正常"
	} else {
		classinfo.State = "异常"
	}
	classinfo.Created = result.Created
	classinfo.Students = students
	//将班级的学生信息放入redis的list
	var stustr = make([][]byte, len(students))
	for i, v := range students {
		stustr[i], _ = json.Marshal(v)
	}
	res := conn.RedisDb.SAdd("class::"+strconv.Itoa(result.Id), stustr).Err()
	if res != nil {
		fmt.Println(res)
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "class": classinfo})
}

func (a *ClassController) GetClassDetailInfo(c *gin.Context) {
	//不使用缓存
	var id int
	if idstring, ok := c.GetQuery("id"); ok {
		id, _ = strconv.Atoi(idstring)
	}
	result, err := models.ClassDetailInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "err": err})
	}

	var classdetialinfo classInfo
	classdetialinfo.Id = result.Id
	classdetialinfo.Name = result.Name
	classdetialinfo.Grade = result.Grade.Name
	classdetialinfo.ClassTeacher = result.Teacher.Name
	if result.State == 1 {
		classdetialinfo.State = "正常"
	} else {
		classdetialinfo.State = "异常"
	}
	classdetialinfo.Created = result.Created

	student_count := len(result.Students)
	studentInfo := make([]student, student_count)

	for i, v := range result.Students {
		studentInfo[i].Id = v.Id
		studentInfo[i].Name = v.Name
		studentInfo[i].Created = v.Created
	}
	classdetialinfo.Students = studentInfo

	//c.JSON(http.StatusOK, gin.H{"code": 1, "list": classdetialinfo, "student_count":len(result.Students)})
	c.JSON(http.StatusOK, gin.H{"code": 1, "student_count": student_count})
}

// 添加学生 or 修改学生 (看是否传了id)
func (a *ClassController) Save(c *gin.Context) {
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
func (a *ClassController) StuDel(c *gin.Context) {
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
