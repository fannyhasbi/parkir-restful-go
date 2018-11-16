package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

const (
  DB_USER     = "postgres"
  DB_PASSWORD = "hasbi12345"
  DB_NAME     = "parking"
)

func main(){
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=5433 sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres", dbinfo)
  checkErr(err)
  defer db.Close()

  fmt.Println("# SELECT Database")
  rows, err := db.Query("SELECT * FROM owner")
  checkErr(err)

  fmt.Println("oid | nama | nim")

  for rows.Next(){
    var oid int
    var nama string
    var nim string

    err = rows.Scan(&oid, &nama, &nim)
    checkErr(err)

    fmt.Printf("%v | %11v | %v\n", oid, nama, nim)
  }

}

func checkErr(err error){
  if err != nil {
    panic(err)
  }
}