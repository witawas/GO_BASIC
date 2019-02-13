package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start....")
	http.HandleFunc("/todos", todosHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func todosHandler(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("METHOD ==>",req.Method)
	//result := []byte("hello GET todos")
	//w.Write(result)

	if req.Method == "GET" {
		fmt.Fprintf(w, "hello if 1 %s todos", req.Method)
	}

	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			//.....
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("error==>", err)
			return
		}
		fmt.Printf("body : %s\n", body)

		fmt.Fprintf(w, "hello if 2 %s todos", req.Method)
	}

	//fmt.Fprintf(w,"hello %s todos","GET")
}
