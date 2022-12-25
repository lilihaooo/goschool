package models

import "aschool/conn"

func Count(value interface{}) (int, error) {
	var total int = 0
	conn.DB.Model(value).Count(&total)
	return total, nil
}
