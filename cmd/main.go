package main

import (
	"fmt"
	"net/http"
	"pr-trigger-go/internal/router"
)

func main() {
	fmt.Println("Pr-Trigger-GO")

	MainRouter := router.Router()

	http.ListenAndServe(":3000", MainRouter)
}
