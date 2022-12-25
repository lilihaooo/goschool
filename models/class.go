package models

import (
	"aschool/conn"
	//"gorm.io/gorm/clause"
)

type Class struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	GradeId   int           `json:"grade_id"`
	TeacherId int           `json:"teacher_id"`
	State     int           `json:"state"`
	Created   int64         `json:"created"`
	Students  []StudentInfo `json:"students"`
	Grade     Grade
	Teacher   TeacherInfo
}

type StudentInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	State   int    `json:"state"`
	ClassId int    `json:"class_id"`
	Created int64  `json:"created"`
}

type TeacherInfo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}

func (StudentInfo) TableName() string {
	return "student"
}
func (TeacherInfo) TableName() string {
	return "teacher"
}

// 添加
func CreateClass(student *Student) (err error) {
	err = conn.DB.Create(&student).Error
	return
}

// 删除
func DeleteClass(id int) (err error) {
	err = conn.DB.Where("id=?", id).Delete(&Student{}).Error
	return
}

// 修改
func UpdateClass(student *Student) (err error) {
	err = conn.DB.Save(student).Error
	return
}

// 查找
func ClassList(page int) ([]Class, error) {
	if page < 1 {
		page = 1
	}
	db := conn.DB
	var class []Class
	err := db.Model(&Class{}).Preload("Grade").Preload("Teacher").Find(&class).Error
	return class, err
}

//班级信息详情
/*
1, 该班级的学生列表   ---  以班级id为键名缓存每个班级的学生信息
2, 年级名称
3, 状态
4, 班主任
*/
func ClassDetailInfo(classId int) (Class, error) {
	db := conn.DB
	var class Class
	err := db.Debug().Model(&Class{}).Preload("Grade").Preload("Teacher").Preload("Students").Find(&class, classId).Error
	return class, err
}

// 获得某个班级的信息
func GetOneClass(classId int) (Class, error) {
	db := conn.DB
	var class Class
	err := db.Model(&Class{}).Preload("Grade").Preload("Teacher").Find(&class, classId).Error
	return class, err
}
