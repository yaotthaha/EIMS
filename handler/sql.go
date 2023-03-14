package handler

import "time"

type EmployeeInfo struct {
	UserID         uint64    `gorm:"column:user_id;primary_key" json:"user_id,omitempty"`
	Name           string    `gorm:"column:name" json:"name,omitempty"`
	Sex            string    `gorm:"column:sex" json:"sex,omitempty"`
	Age            uint8     `gorm:"column:age" json:"age,omitempty"`
	Phone          string    `gorm:"column:phone" json:"phone,omitempty"`
	Email          string    `gorm:"column:email" json:"email,omitempty"`
	Position       string    `gorm:"column:position" json:"position,omitempty"`
	Marry          string    `gorm:"column:marry" json:"marry,omitempty"`
	Education      string    `gorm:"column:education" json:"education,omitempty"`
	JoinTime       time.Time `gorm:"column:join_time;autoCreateTime" json:"join_time,omitempty"`
	DepartmentID   uint64    `gorm:"column:department_id;foreignKey:department_id;references:id" json:"department_id,omitempty"`
	DepartmentName string    `gorm:"column:department_name;many2many:department;joinReferences:name" json:"department_name,omitempty"`
}
