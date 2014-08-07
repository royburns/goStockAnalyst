package models

import (
// "fmt"
// "github.com/astaxie/beego"
// "time"
// "strings"
)

const (
	Shanghai_Prefix  = "ss"
	Shanghai_Postfix = ".ss"
	Shenzhen_Prefix  = "sz"
	Shenzhen_Postfix = ".sz"
)

// 0 公司代码,1 公司简称,2 公司全称,3 英文名称,4 注册地址,
// 5 A股代码, 6 A股简称, 7 A股上市日期, 8 A股总股本, 9 A股流通股本,
// 10 B股代码, 11 B股简称, 12 B股上市日期, 13 B股总股本, 14 B股流通股本,
// 15 地区, 16 省份, 17 城市, 18 所属行业, 19 公司网址
type StockCompany struct {
	Id       int64 `xorm:"pk"`
	Name     string
	FullName string
	EnName   string
	Address  string

	ACodeID     string
	AName       string
	AData       string
	ACapital    int64
	AOutCapital int64

	BCodeID     string
	BName       string
	BData       string
	BCapital    int64
	BOutCapital int64

	Region   string
	Province string
	City     string
	Industry string
	Site     string
}
type StockDay struct {
	//
	Code string
	// Data uint32
	Date string

	Open  float32
	Close float32
	Hight float32
	Low   float32

	Amount   float32 //
	Vol      int32   //gu
	Reserved int32
}

type StockWeek struct {
	Code string
	Data uint32

	Open  float32
	Close float32
	Hight float32
	Low   float32

	Amount   float32
	Vol      int32
	Reserved int32
}

type StockMonth struct {
	Code string
	Data uint32

	Open  float32
	Close float32
	Hight float32
	Low   float32

	Amount   float32
	Vol      int32
	Reserved int32
}

type StockYear struct {
	Code string
	Data uint32

	Open  float32
	Close float32
	Hight float32
	Low   float32

	Amount   float32
	Vol      int32
	Reserved int32
}

func (this *StockDay) GetOpen() float32 {
	return 0.01 * float32(this.Open)
}

func (this *StockDay) GetClose() float32 {
	return 0.01 * float32(this.Close)
}

func (this *StockDay) GetHight() float32 {
	return 0.01 * float32(this.Hight)
}

func (this *StockDay) GetLow() float32 {
	return 0.01 * float32(this.Low)
}
