package handler

import "net/http"

func Router() {
	http.HandleFunc("/", getLogin)
	http.HandleFunc("/login", postLogint)

}
