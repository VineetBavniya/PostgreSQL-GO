package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VineetBavniya/PostgreSQL-GO/router"
)

func main(){
    r := router.Router()

    fmt.Println("Starting Server on the Port 9001")

    log.Fatal(http.ListenAndServe(":9001", r))

}