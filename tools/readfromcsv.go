package main

import (
	"encoding/csv"
	"fmt"
	"github.com/royburns/goStockAnalyst/models"
	"os"
	"strconv"
	"strings"
)

// 0 公司代码,1 公司简称,2 公司全称,3 英文名称,4 注册地址,
// 5 A股代码, 6 A股简称, 7 A股上市日期, 8 A股总股本, 9 A股流通股本,
// 10 B股代码, 11 B股简称, 12 B股上市日期, 13 B股总股本, 14 B股流通股本,
// 15 地区, 16 省份, 17 城市, 18 所属行业, 19 公司网址
func main() {
	models.InitDB()
	fmt.Println("Read Begin")
	f, err := os.Open("../doc/sz_上市公司列表.csv")
	if err != nil {
		fmt.Println("Error.")
	}
	defer f.Close()
	r := csv.NewReader(f)

	lines, err := r.ReadAll()
	fmt.Println("Read End")

	fmt.Println("Insert Begin")
	orm := models.GetEngine()
	// sc := make([]models.StockCompany, 0)
	fmt.Println(len(lines))
	for i := 1; i < len(lines); i++ {
		if lines[i][0] != "" {
			var tmp models.StockCompany

			// fmt.Println(lines[i][0])
			number, _ := strconv.Atoi(lines[i][0])
			tmp.Id = int64(number)
			tmp.Name = lines[i][1]
			tmp.FullName = lines[i][2]
			tmp.EnName = lines[i][3]
			tmp.Address = lines[i][4]

			if lines[i][5] != "" {
				tmp.ACodeID = fmt.Sprintf("%06s", lines[i][5])
			}
			tmp.AName = lines[i][6]
			tmp.AData = lines[i][7]
			number, _ = strconv.Atoi(strings.Join(strings.Split(lines[i][8], ","), ""))
			tmp.ACapital = int64(number)
			number, _ = strconv.Atoi(strings.Join(strings.Split(lines[i][9], ","), ""))
			tmp.AOutCapital = int64(number)

			if lines[i][10] != "" {
				tmp.BCodeID = fmt.Sprintf("%06s", lines[i][10])
			}
			tmp.BName = lines[i][11]
			tmp.BData = lines[i][12]
			number, _ = strconv.Atoi(strings.Join(strings.Split(lines[i][13], ","), ""))
			tmp.BCapital = int64(number)
			number, _ = strconv.Atoi(strings.Join(strings.Split(lines[i][14], ","), ""))
			tmp.BOutCapital = int64(number)

			tmp.Region = lines[i][15]
			tmp.Province = lines[i][16]
			tmp.City = lines[i][17]
			tmp.Industry = lines[i][18]
			tmp.Site = lines[i][19]

			// sc = append(sc, tmp)
			orm.Insert(tmp)
		}
	}

	fmt.Println("Insert End")
}
