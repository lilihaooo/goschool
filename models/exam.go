package models

type Exam struct {
	Id            int
	Name          string
	SubjectId     string
	ClassIds      string
	ExamTeacherId int
	State         int
	Created       string
}

//添加考试
