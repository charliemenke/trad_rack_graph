package router

import (
    "net/http"
)

// build and return server router
func New() {
    router := http.NewServeMux()

    router.HandleFunc()
}
