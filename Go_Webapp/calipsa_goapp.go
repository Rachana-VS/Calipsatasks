package main

import (
    "fmt"
    "log"
    "net/http"
    "sort"

)

func main() {

    http.HandleFunc("/", handler)
    http.HandleFunc("/hello",handler_hello)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
    var s []string
    for k, _ := range r.Header {
        s=append(s,k)
    }
    sort.Strings(s)
    for _,value := range s  {
        fmt.Println(value)
        w.Write([]byte(value+"\n"))
    }

    fmt.Println("")

}

func handler_hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello Calipsa!</h1>\n"))
//     fmt.Println("Hello Calipsa")


}


