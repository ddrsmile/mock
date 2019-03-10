package common

import (
    "github.com/gorilla/mux"
)

func SetRoutes(r *mux.Router) {
    r.HandleFunc("/health", GetHealthHandler())
}
