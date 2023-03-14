package sql

import (
	"time"
)

/**
create table departments
(
    id   int auto_increment primary key,
    name varchar(255) not null,
    constraint name unique (name)
);
*/

type Department struct {
	DepartmentID uint64 `gorm:"column:id;primaryKey;type:int(11);autoIncrement"`
	Name         string `gorm:"column:name;type:varchar(255);not null;unique"`
}

/**
create table employees
(
    user_id       int auto_increment primary key,
    name          varchar(255)                        not null,
    sex           enum ('男', '女')                   not null,
    age           int                                 not null,
    phone         varchar(255)                        not null,
    email         varchar(255)                        not null,
    position      varchar(255)                        not null,
    marry         enum ('已婚', '未婚', '离婚')        not null,
    education     varchar(255)                        not null,
    join_time     timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    department_id int                                 not null,
    constraint fk_employees_department foreign key (department_id) references departments (id) on update cascade
);
*/

type Employee struct {
	UserID       uint64     `gorm:"column:user_id;primaryKey;type:int(11);autoIncrement"`
	Name         string     `gorm:"column:name;type:varchar(255);not null"`
	Sex          string     `gorm:"column:sex;type:enum('男','女');not null"`
	Age          uint64     `gorm:"column:age;type:int(11);not null"`
	Phone        string     `gorm:"column:phone;type:varchar(255);not null"`
	Email        string     `gorm:"column:email;type:varchar(255);not null"`
	Position     string     `gorm:"column:position;type:varchar(255);not null"`
	Marry        string     `gorm:"column:marry;type:enum('已婚','未婚','离婚');not null;"`
	Education    string     `gorm:"column:education;type:varchar(255);not null"`
	JoinTime     time.Time  `gorm:"column:join_time;type:timestamp;not null"`
	Department   Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;foreignKey:department_id;references:id"`
	DepartmentID uint64     `gorm:"column:department_id;type:int(11);not null"`
}

const (
	AddDefaultDepartment = "INSERT IGNORE INTO departments (id, name) VALUES (1000000, 'Admin')"
	AddDefaultEmployee   = "INSERT IGNORE INTO employees (user_id, name, sex, age, phone, email, position, marry, education, join_time, department_id) VALUES (1000000, 'Admin', '男', 18, '00000000000', 'admin@example.com', 'Admin', '未婚', '本科', '1970-01-01 00:00:00', 1000000)"
)
