package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/royburns/goStockAnalyst/models"
	"strconv"
	"strings"
	// "time"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {

	this.Data["Website"] = "goTestLinkReport.org"
	this.Data["Email"] = "roy.burns@163.com"

	this.TplNames = "index.tpl"
}

// type StockDay struct {
// 	//
// 	Code string
// 	Data uint32

// 	Open  float32
// 	Close float32
// 	Hight float32
// 	Low   float32

// 	Amount   float32 //
// 	Vol      int32   //gu
// 	Reserved int32
// }
func (this *IndexController) Stock() {
	fmt.Println("In GetStock()...")
	req := httplib.Get("http://xueqiu.com/S/SH601166/historical.csv")
	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}

	days := make([]*models.StockDay, 0)
	lines := strings.Split(str, "\n")
	for i := 1; i < len(lines); i++ {
		column := strings.Split(lines[i], ",")
		temp := new(models.StockDay)

		temp.Code = column[0]

		// t := time.Parse("2014-08-07", column[1])
		// temp.Data = t.
		temp.Date = column[1]

		f, _ := strconv.ParseFloat(column[2], 2)
		temp.Open = float32(f)
		f, _ = strconv.ParseFloat(column[3], 2)
		temp.Hight = float32(f)
		f, _ = strconv.ParseFloat(column[4], 2)
		temp.Low = float32(f)
		f, _ = strconv.ParseFloat(column[5], 2)
		temp.Close = float32(f)

		i, _ = strconv.Atoi(column[6])
		temp.Vol = int32(i)

		days = append(days, temp)
	}

	this.Data["json"] = days
	this.ServeJson()

	// var v interface{}
	// err := req.ToJson(v)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(v)
}

type page struct {
	IsActive  bool
	TestPlans string
	PageNum   int
}

// calPageList returns page lists.
func calPageList(p, maxPageNum int) []*page {
	listSize := 15
	hls := listSize / 2
	pl := make([]*page, 0, listSize)

	start, end := p-hls, p+hls
	if p < hls+1 {
		start, end = 1, listSize
	}

	if end > maxPageNum {
		end = maxPageNum
	}

	for i := start; i <= end; i++ {
		pl = append(pl, &page{
			IsActive: i == p,
			PageNum:  i,
		})
	}
	return pl
}
