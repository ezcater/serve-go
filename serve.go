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

func main() {
  port := flag.String("p", "3333", "port to serve on")
  directory := flag.String("d", "files", "the directory of static file to host")
  flag.Parse()

  http.Handle("/", http.FileServer(http.Dir(*directory)))

  log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
  log.Fatal(http.ListenAndServe(":"+*port, nil))
}
