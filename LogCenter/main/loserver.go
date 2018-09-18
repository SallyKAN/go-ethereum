package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	httpPort := 4001
	//openLogFile(logPath)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/p2p",logHandler)
	http.HandleFunc("/pow",mineHandler)
	fmt.Printf("listening on %v\n", httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort),nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1><div>Welcome to whereever you are</div>")
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func logHandler(w http.ResponseWriter, r *http.Request) {
		//log.SetOutput(w)
		//fmt.Fprintf(w,"<h1>hello</h1>")
		//fmt.Fprintf(log.)
	fmt.Printf("Logging to %v\n", "output")
		dat, err := ioutil.ReadFile("output")
		check(err)
		fmt.Fprintf(w,string(dat))

		}
func mineHandler(w http.ResponseWriter, r *http.Request)	{
	fmt.Printf("Logging to %v\n", "output2")
	dat, err := ioutil.ReadFile("output2")
	check(err)
	fmt.Fprintf(w,string(dat))
}


//func openLogFile(logfile string) {
//	if logfile != "" {
//		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
//
//		if err != nil {
//			log.Fatal("OpenLogfile: os.OpenFile:", err)
//		}
//
//		log.SetOutput(lf)
//	}
//}
