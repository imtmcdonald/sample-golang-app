package main

import (
    "fmt"
    "log"
    "net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request){
	tenant := os.Getenv("TENANT_NAME")
    fmt.Fprintf(w, "Welcome to the HomePage: %s", tenant)
    fmt.Println("Endpoint Hit: homePage", tenant)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}