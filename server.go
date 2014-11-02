package main

import (
	"flag"
	"github.com/gorilla/mux"
	hand "github.com/squiidz/ferretShack/module/handler"
	"log"
	"net/http"
)

var (
	port   *string
	router = mux.NewRouter()
)

func init() {
	port = flag.String("port", "8080", "-port [your port]")
	flag.Parse()
	*port = ":" + *port
	log.Printf("[+] Server Running on port %s", *port)
}

func main() {
	router.Handle("/home", http.HandlerFunc(hand.HomeHandler))
	router.Handle("/view", http.HandlerFunc(hand.ViewHandler))
	router.Handle("/article/{id}", http.HandlerFunc(hand.ArticleHandler))
	router.Handle("/", http.HandlerFunc(hand.DefaultHandler))
	log.Panic(http.ListenAndServe(*port, router))
}
