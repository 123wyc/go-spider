package pprof

import (
	"net/http"
	"net/http/pprof"
	"time"
)

var (
	mux    *http.ServeMux
	server *http.Server
)

func Init() {

	mux = http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)

	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	server = &http.Server{
		ReadTimeout:  time.Duration(5000) * time.Millisecond,
		WriteTimeout: time.Duration(5000) * time.Millisecond,
		Handler:      mux,
	}

	go http.ListenAndServe(":16140", nil)
}

// func main() {

// 	Init()

// 	fmt.Println("Hello world!")
// 	for {
// 		time.Sleep(1 * time.Second)
// 	}

// }
