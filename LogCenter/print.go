package LogCenter

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
)

type Options struct {
	Path string
	Port string
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {

	op := &Options{Path: "./", Port: "8001"}

	data, _ := ioutil.ReadFile("./config.json")

	json.Unmarshal(data, op)

	http.Handle("/", http.FileServer(http.Dir(op.Path)))
	err := http.ListenAndServe(":" + op.Port, Log(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
