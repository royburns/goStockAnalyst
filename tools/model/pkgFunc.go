package model

type PkgFunc struct {
	Id    int64  `xorm:"BIGINT(20)"`
	Pid   int64  `xorm:"index BIGINT(20)"`
	Path  string `xorm:"VARCHAR(150)"`
	Name  string `xorm:"index VARCHAR(100)"`
	Doc   string `xorm:"TEXT"`
	IsOld int    `xorm:"TINYINT(1)"`
}
