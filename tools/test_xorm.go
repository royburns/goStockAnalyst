package main

import (
	"fmt"
	"github.com/Unknwon/gowalker/utils"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var (
	orm    *xorm.Engine
	tables []interface{}
)

type Test struct {
	Id   int64  `xorm:"ID pk"`
	Name string `xorm:"Name"`
}

type PkgFunc struct {
	Id    int64
	Pid   int64  `xorm:"index"` // Id of package documentation it belongs to.
	Path  string `xorm:"VARCHAR(150)"`
	Name  string `xorm:"index VARCHAR(100)"`
	Doc   string `xorm:"TEXT"`
	IsOld bool   // Indicates if the function no longer exists.
}

func main() {
	//
	InitDB()
}

func InitDB() (err error) {
	utils.LoadConfig("../conf/app.conf")
	dbtype := utils.Cfg.MustValue("db", "dbtype")
	host := utils.Cfg.MustValue("db", "host")
	port := utils.Cfg.MustValue("db", "port")
	name := utils.Cfg.MustValue("db", "name")
	user := utils.Cfg.MustValue("db", "user")
	pwd := utils.Cfg.MustValue("db", "pwd")
	ssl := utils.Cfg.MustValue("db", "ssl")

	switch dbtype {
	case "mysql":
		// connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		// 	user, pwd, host, port, name)
		connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			user, pwd, host, port, name)
		fmt.Errorf(connstr)
		orm, err = xorm.NewEngine("mysql", connstr)
		if err != nil {
			beego.Debug(fmt.Sprintf("Failed to create new engine : %v\n", err))
		}
	case "postgres":
		cnnstr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
			user, pwd, host, port, name, ssl)
		orm, err = xorm.NewEngine("postgres", cnnstr)
	default:
		beego.Debug(fmt.Sprintf("Unknown db type: %s\n", dbtype))
		return fmt.Errorf("Unknown db type: %s\n", dbtype)
	}

	if err != nil {
		beego.Debug(fmt.Sprintf("models.init(failed to connect database): %v\n", err))
		return fmt.Errorf("models.init(failed to connect database): %v\n", err)
	}

	// orm.ShowSQL = true
	// orm.ShowInfo = true
	// orm.ShowDebug = true
	// orm.ShowWarn = true
	orm.ShowErr = true

	orm.SetMapper(core.SameMapper{})

	fmt.Println("success!")
	// return orm.Sync(tables...)
	err = orm.Sync(new(Test), new(PkgFunc))
	if err != nil {
		fmt.Println("Failed to Sync()!")
	}
	return
}
