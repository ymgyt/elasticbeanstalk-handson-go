package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"sort"
	"strings"
)

const (
	addr = ":5000"
	v = "v0.0.1"
)

func main() {
	mux := 	http.NewServeMux()
	mux.HandleFunc("/", dump)
	mux.HandleFunc("/dump", dump)
	mux.HandleFunc("/env/name", envName)
	mux.HandleFunc("/env", env)
	mux.HandleFunc("/version", version)

	s := &http.Server{Addr: addr,  Handler: mux}
	fmt.Println("running on ", addr)
	fmt.Fprintln(os.Stderr, s.ListenAndServe())
}

func dump(w http.ResponseWriter, r *http.Request){
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	w.Write(dump)
}

func envName(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,os.Getenv("ENV"))
}

func env(w http.ResponseWriter, r *http.Request) {
	envs := os.Environ()
	sort.Slice(envs, func(i,j int) bool {
		return envs[i] < envs[j]
	})
	io.WriteString(w, strings.Join(envs, "\n"))
}

func version(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,v)
}
