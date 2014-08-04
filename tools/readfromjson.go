package main

import (
	"encoding/json"
	"fmt"
	"github.com/royburns/goStockAnalyst/models"
	"os"
	// "strings"
)

type JsonData struct {
	Result []StockInfo `json:"result"`
}

type StockInfo struct {
	ID          int64  `json:"NUM,string"`
	FullName    string `json:"FULLNAME"`
	ProductName string `json:"PRODUCTNAME"`
	ProductID   string `json:"PRODUCTID"`
}

func ReadJson(path string, j interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return json.NewDecoder(f).Decode(j)
}

// func (j *Json) WriteTo(w io.Writer) (int64, error) {
// 	b, err := json.MarshalIndent(j, "", "\t")
// 	if err != nil {
// 		return 0, err
// 	}
// 	n, err := w.Write(append(b, '\n'))
// 	return int64(n), err
// }

func main() {

	models.InitDB()
	fmt.Println("Read Begin")
	// filename := "../doc/ss_a_page1.txt"
	filename := "E:\\GoProjects\\src\\github.com\\royburns\\goStockAnalyst\\doc\\ss_a_page4.txt"
	j := new(JsonData)
	err := ReadJson(filename, j)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(len(j.Result))
	}
	fmt.Println("Read End")

	fmt.Println("Insert Begin")
	sc := make([]models.StockCompany, 0)
	for i := 0; i < len(j.Result); i++ {
		var tmp models.StockCompany
		tmp.Id = j.Result[i].ID
		tmp.Name = j.Result[i].ProductName
		tmp.FullName = j.Result[i].FullName

		tmp.ACodeID = j.Result[i].ProductID
		tmp.AName = j.Result[i].ProductName

		sc = append(sc, tmp)
	}
	orm := models.GetEngine()
	orm.Insert(sc)
	fmt.Println("Insert End")
}
