package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var todos = make(map[string]todo)

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
		//fmt.Fprintf(w, "hello if 1 %s todos", req.Method)
		b, err1 := json.Marshal(todos)
		if err1 != nil {
			//.....
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("error==>", err1)
			return
		}
		fmt.Printf("%T => \n %v \n %s \n", b, b, b)
		//fmt.Fprintf("%T => \n %v \n %s \n", b, b, b)
		w.Header().Set("content-Type", "application/json")
		w.Write(b)
	}

	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			//.....
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("error==>", err)
			return
		}

		fmt.Printf("%T =>  %v \n %s \n", body, body, body)

		fmt.Printf("body : %s\n", body)

		//fmt.Fprintf(w, "hello if 2 %s todos", req.Method)

		fmt.Println("-----------------------------------------------")
		var t todo
		err1 := json.Unmarshal(body, &t)

		fmt.Printf("% #v\n", t)
		fmt.Println("error ==>", err1)

		//if todos == nil {
		//	todos = make(map[string]todo)
		//}

		fmt.Println("request ID=", t.ID)

		//item, ok := todos[t.ID] ==> for check map
		//	if todos[t.ID] != nil { // check for address
		todos[t.ID] = t
		//	}

		fmt.Println(" map size", len(todos))
		fmt.Printf(" map % #v\n", todos)

		//b, err1 := json.Marshal(t)

		//for i, _ := range todos {
		b, err1 := json.Marshal(todos)
		if err1 != nil {
			//.....
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("error==>", err1)
			return
		}
		fmt.Printf("%T => \n %v \n %s \n", b, b, b)
		//fmt.Fprintf("%T => \n %v \n %s \n", b, b, b)
		w.Header().Set("content-Type", "application/json")
		w.Write(b)

		//fmt.Fprintf(w, "hello if 1 %s todos", req.Method)
		//	}

		//fmt.Println(err1)

		//fmt.Fprintf(w, "% #v\n", todos)
	}

	//fmt.Fprintf(w,"hello %s todos","GET")

}
