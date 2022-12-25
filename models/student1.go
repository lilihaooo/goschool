package models

import (
	"aschool/conn"
	"aschool/settings"
)

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Sex     int    `json:"sex"`
	Age     int    `json:"age"`
	State   int    `json:"state"`
	ClassId int    `json:"class_id"`
	GradeId int    `json:"grade_id"`
	Created int64  `json:"created"`
	Class   ClassInfo
	Grade   Grade
}

type ClassInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GradeId   int    `json:"grade_id"`
	TeacherId int    `json:"teacher_id"`
}

func (ClassInfo) TableName() string {
	return "class"
}

//

// 查找分页
func StudentList(keyword string, id int, page int, sort_type int) ([]Student, error) {
	if page < 1 {
		page = 1
	}
	db := conn.DB
	db = db.Model(&Student{}).Preload("Class").Preload("Grade")

	// 设置过滤字段
	if id > 0 {
		db = db.Where("lh_student.id = ?", id)
	}
	// 搜索关键词（采用模糊匹配）
	if len(keyword) > 0 {
		db = db.Where("lh_student.name like ?", "%"+keyword+"%")
	}
	// 设置分页
	pagesize := settings.Conf.Pagesize
	offset := (page - 1) * pagesize
	db = db.Offset(offset).Limit(pagesize)

	if sort_type == 1 {
		db = db.Order("lh_student.id desc")
	} else {
		db = db.Order("lh_student.id asc")
	}

	var student []Student
	//err := db.Model(&Student{}).Related(&class, "ClassId").Find(&student).Error
	err := db.Find(&student).Error
	return student, err
}

// 添加
func CreateStudent(student *Student) (err error) {
	err = conn.DB.Create(&student).Error
	return
}

// 删除
func DeleteStudent(id int) (err error) {
	err = conn.DB.Where("id=?", id).Delete(&Student{}).Error
	return
}

// 修改
func UpdateStudent(student *Student) (err error) {
	err = conn.DB.Save(student).Error
	return
}

// 获取某个班级的全部学生
func GetOneClassStudents(id int) ([]Student, error) {
	db := conn.DB
	var students []Student
	err := db.Debug().Where("class_id = ?", id).Preload("Grade").Limit(50).Find(&students).Error
	return students, err
}
