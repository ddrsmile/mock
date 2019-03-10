package main

import (
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "log"
    "mock/api"
    "net/http"
    "os"
)

func health(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("OK"))
}

func getRoutes() http.Handler {
    router := mux.NewRouter()
    router.HandleFunc("/health", health)
    api.SetRoutes(router)
    return handlers.LoggingHandler(os.Stdout, router)
}

func main() {
    log.Println("server start.")
    log.Fatal(http.ListenAndServe(":8080", getRoutes()))

}