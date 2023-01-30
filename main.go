package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
	"os"
    _ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
  value := os.Getenv(key)
  if len(value) == 0 {
      return defaultValue
  }
  return value
}

func dbConnection() string{
  host := getEnv("POSTGRES_HOST", "postgres")
  port := getEnv("POSTGRES_PORT", "5432")
  user := getEnv("POSTGRES_USER", "postgres")
  password := getEnv("POSTGRES_PASSWORD", "postgres")
  dbname := getEnv("POSTGRES_DB_NAME", "postgres")
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  fmt.Println(psqlInfo)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    // panic(err)
    fmt.Println(err)
    return "Failure"
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    // panic(err)
    fmt.Println(err)
    return "Failure"
  }

  fmt.Println("Successfully connected!")
  return "Connected"
}

func homePage(w http.ResponseWriter, r *http.Request){
  tenant := getEnv("TENANT_NAME", "default")
  status := dbConnection()
  fmt.Fprintf(w, "Welcome to the HomePage: %s", tenant)
  fmt.Println("Endpoint Hit: homePage", tenant)
  fmt.Fprintf(w, "Database Status: %s", status)
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
  handleRequests()
}