package app

import (
    "net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    f(w, r)
}

func Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "text/html")

    w.Write([]byte("<h1>Hello World!</h1>"))
}
