package models

//
//import (
//	"aschool/dao"
//	"aschool/settings"
//)
//
//type Student struct {
//	Id      int    `json:"id"`
//	Name    string `json:"name"`
//	Sex     int    `json:"sex"`
//	Age     int    `json:"age"`
//	GradeId int    `json:"grade_id"`
//	ClassId int    `json:"class_id"`
//	State   int    `json:"state"`
//	Created int64  `json:"created"`
//}
//
//// 添加
//func CreateStudent(student *Student) (err error) {
//	err = dao.DB.Create(&student).Error
//	return
//}
//
//// 删除
//func DeleteStudent(id int) (err error) {
//	err = dao.DB.Where("id=?", id).Delete(&Student{}).Error
//	return
//}
//
//// 修改
//func UpdateStudent(student *Student) (err error) {
//	err = dao.DB.Save(student).Error
//	return
//}
//
//type studentInfo struct {
//	Id        int    `json:"id"`
//	Name      string `json:"name"`
//	Sex       int    `json:"sex"`
//	Age       int    `json:"age"`
//	GradeName string `json:"grade_name"`
//	ClassName string `json:"class_name"`
//	State     int    `json:"state"`
//	Created   int64  `json:"created"`
//}
//
//// 查找
//func StudentList(keyword string, id int, page int, sort_type int) (studentInfo []*studentInfo, err error) {
//	if page < 1 {
//		page = 1
//	}
//	db := dao.DB
//	db = db.Debug().Table("lh_student").Select("lh_student.id, lh_student.name ,sex, age, lh_grade.name as grade_name, lh_class.name as class_name ,lh_student.state, lh_student.created as created").
//		Joins("left join lh_class on class_id = lh_class.id").
//		Joins("left join lh_grade on lh_student.grade_id = lh_grade.id")
//	// 设置过滤字段
//	if id > 0 {
//		db = db.Where("lh_student.id = ?", id)
//	}
//	// 搜索关键词（采用模糊匹配）
//	if len(keyword) > 0 {
//		db = db.Where("lh_student.name like ?", "%"+keyword+"%")
//	}
//	// 设置分页
//	pagesize := settings.Conf.Pagesize
//	offset := (page - 1) * pagesize
//	db = db.Offset(offset).Limit(pagesize)
//
//	if sort_type == 1 {
//		db = db.Order("lh_student.id desc")
//	} else {
//		db = db.Order("lh_student.id asc")
//	}
//	if err = db.Scan(&studentInfo).Error; err != nil {
//		return nil, err
//	}
//	return
//}
