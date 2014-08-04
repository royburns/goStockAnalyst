package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("...")
	r, err := http.Get("http://ichart.yahoo.com/table.csv?s=600000.SS&d=09&e=8&f=2010&g=d")
	if err != nil {
		log.Fatal(err)
		return
	}

	file, err := os.OpenFile("c:\\test.csv", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	// var buf []byte
	// n, err := r.Body.Read(buf)
	buf, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	// fmt.Println(n)
	n, err := file.Write(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(n)

	// // res, err := http.Get("http://www.google.com/robots.txt")
	// res, err := http.Get("http://ichart.yahoo.com/table.csv?s=600000.SS&a=08&b=25&c=2010&d=09&e=8&f=2010&g=d")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)
}
