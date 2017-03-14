package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.google.com")
	//if err != nil {
	//	log.Fatal(err)
	//	}
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	fmt.Printf("%s", page)
}

//blank identifier: should be like to get the log error
//res, err :=
