package api

import "net/http"

type Item struct{
    Path     string   `json:"path"`
    Methods  []string `json:"methods"`
    Response Response `json:"response"`
}

type Response struct {
    ContentType string `json:"contentType"`
    Headers map[string][]string `json:"headers"`
    Cookies []http.Cookie `json:"cookies"`
    Content string `json:"content"`
}
