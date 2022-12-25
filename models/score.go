package models

type Score struct {
	Id        int
	ExamId    int
	SubjectId int
	StudentId int
	Score     int
	Created   int64
}

//添加语文成绩

//添加数学成绩

//添加英语成绩

//获取学生列表
/*
1. 只有考试结束才能够添加成绩
2. 只有该班级的任课老师才能添加自己班级自己科目的学生成绩
3. 显示已添加的人数, 还未添加的人数
*/
