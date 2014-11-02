package handler

import (
	"github.com/gorilla/mux"
	"github.com/squiidz/ferretShack/module/db"
	"log"
	"net/http"
)

func DefaultHandler(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/home", http.StatusFound)
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("SUCCESS"))
}

func ViewHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("VIEW GOES HERE"))
}

func ArticleHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("ARTICLE GOES HERE"))
	artId := mux.Vars(req)["id"]
	log.Println(artId)
	session, err := db.NewSession()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
	log.Println(session.Host)
}
