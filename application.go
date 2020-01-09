package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

const addr = ":5000"

func main() {
	mux := 	http.NewServeMux()
	mux.HandleFunc("/dump", dump)
	mux.HandleFunc("/", dump)

	s := &http.Server{Addr: addr,  Handler: mux}
	fmt.Println("running on ", addr)
	fmt.Fprintln(os.Stderr, s.ListenAndServe())
}

func dump(w http.ResponseWriter, r *http.Request){
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(dump)
	w.WriteHeader(http.StatusOK)
}
