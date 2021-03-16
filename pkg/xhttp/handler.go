package xhttp

import (
    "net/http"
)

type RESTfulHandler interface {
    GET(http.ResponseWriter, *http.Request)
    HEAD(http.ResponseWriter, *http.Request)
    POST(http.ResponseWriter, *http.Request)
    DELETE(http.ResponseWriter, *http.Request)
    PATCH(http.ResponseWriter, *http.Request)
    PUT(http.ResponseWriter, *http.Request)
    OPTIONS(http.ResponseWriter, *http.Request)
    ANY(http.ResponseWriter, *http.Request)
}
