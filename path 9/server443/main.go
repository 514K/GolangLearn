package main

import (
	"fmt"
	"net/http"
	// "golang.org/x/crypto/acme/autocert"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS user! Your config: %+v", r.TLS)
	})

	// log.Fatal(http.Serve(autocert.NewListener("mydomainname.com"), mux))
}
