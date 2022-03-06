package main

import (
	"blog/resource"
	"blog/router"
	"fmt"
	"net/http"
)

func main() {
	err := resource.Init("")
	if err != nil {
		fmt.Println(err)
		return
	}

	r := router.NewRouter()
	s := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	s.ListenAndServe()
}
