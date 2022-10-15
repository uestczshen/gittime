package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
)

var (
	logger *log.Logger
)

func InitLog() {
	logger = log.New(os.Stdout, "TRACE:", log.Ltime|log.Lshortfile)
}

func main() {
	InitLog()
	mux := http.NewServeMux()
	mux.HandleFunc("/", safeHandler(handleRoot))
	mux.HandleFunc("/healthz", handleHealthz)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type MyHandler func(http.ResponseWriter, *http.Request)

func safeHandler(handler MyHandler) MyHandler {
	return func(rsp http.ResponseWriter, req *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				trace := make([]byte, 0)
				runtime.Stack(trace, false)
				logger.Printf("handler panic trace:%v", string(trace))
				rsp.WriteHeader(http.StatusInternalServerError)
			}
		}()
		handler(rsp, req)
	}
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	retCode := int(http.StatusOK)
	if version := os.Getenv("VERSION"); version == "" {
		w.Header().Add("VERSION", "not found")
	} else {
		w.Header().Add("VERSION", version)
	}
	remoteAddr := req.RemoteAddr
	for k, vals := range req.Header {
		for _, val := range vals {
			w.Header().Add(k, val)
		}
	}

	logger.Printf("[handleRoot] clientAddr[%s] retCode:%v", remoteAddr, retCode)
	w.WriteHeader(retCode)
	w.Write([]byte("忽如一夜春风来, 千树万树梨花开"))
}

func handleHealthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
