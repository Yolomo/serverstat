package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	/*"io"*/
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getJsonVar(js *simplejson.Json, str string) string {
	result, err := js.Get("hostname").String()
	if err != nil {
		log.Println(err)
	}
	return result
}

func main() {

	dat, err := ioutil.ReadFile("hosts/xps.json")
	check(err)

	fmt.Print(string(dat))

	js, err := simplejson.NewJson(dat)
	if err != nil {
		log.Fatalln(err)
	}

	total, err := js.Get("hostname").String()
	kak := total

	fmt.Println("Hostname: " + getJsonVar(js, "hostname"))
	fmt.Println("" + kak)
}
