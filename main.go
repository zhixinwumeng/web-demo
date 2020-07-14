package main

import (
	"fmt"
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"<h1> hello,this is my first page! </h1>")
}

func main() {
	fmt.Println("服务端运行中。。。")
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":8000",nil)

}
