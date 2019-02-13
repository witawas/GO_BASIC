package todo

import (
	"fmt"
	"net/http"
)

func todosHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello %s todos", req.Method)
}

func main() {

}
