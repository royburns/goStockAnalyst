package model

type Test struct {
	Id   int64  `xorm:"BIGINT(20)"`
	Name string `xorm:"VARCHAR(255)"`
}
