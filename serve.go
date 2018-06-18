/*
Serve is a very simple static file server in go

Usage:
-p="8100": port to serve on
-d=".":    the directory of static files to host
*/
package main

import (
  "flag"
  "log"
  "net/http"
)

func logger(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s requested %s", r.RemoteAddr, r.URL)
    h.ServeHTTP(w, r)
  })
}

func main() {
  port := flag.String("p", "3333", "port to serve on")
  directory := flag.String("d", "files", "the directory of static file to host")
  flag.Parse()

  h := http.NewServeMux()

  h.Handle("/", http.FileServer(http.Dir(*directory)))

  hl := logger(h)


  log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
  err := http.ListenAndServe(":"+*port, hl)

  log.Fatal(err)
}
