package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var totalprice int
var totalpages int

//テストデータ
var isbns = []string{"9784873115948", "9784873116402", "9784873114750"}

func main() {
	book := make(chan Book)
	quit := make(chan bool)
	exit := make(chan bool)

	for _, isbn := range isbns {
		go BookInfo(isbn, book, quit)
	}

	go listenBook(book, quit, exit)

	<-exit
	fmt.Println("総ページ数：", totalpages, ",総額：", totalprice)

}

type Book struct {
	Title string `json:"title"`
	Pages int    `json:"pages"`
	Price int    `json:"price"`
}

func BookInfo(isbn string, b chan Book, quit chan bool) {

	resp, _ := http.Get("http://www.oreilly.co.jp/books/" + isbn + "/biblio.json")

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var book Book
	json.Unmarshal(body, &book)

	b <- book
	quit <- true

}

func listenBook(book chan Book, quit chan bool, exit chan bool) {
	var i int
	for {
		select {
		case b := <-book:
			totalprice = totalprice + b.Price
			totalpages = totalpages + b.Pages
		case <-quit:
			i++
			if i >= len(isbns) {
				exit <- true
			}
		}
	}
}
