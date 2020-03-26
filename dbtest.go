// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

package main

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)
const (
    host    = "localhost"
    port    = "54321"
    user    = "postgres"
    password    = "postgres"
    dbname  = "postgres"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host,port,user,password,dbname )
    
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)

    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    } else {
        fmt.Println("Successfully Connected to database",host,"on port", port)
    }
    sqlStatement:= `
    INSERT INTO users ( age,email,first_name,last_name)
    VALUES ($1,$2,$3,$4) RETURNING id`

    id := 0
    err = db.QueryRow(sqlStatement, 100, "spam.justinong@gmail.com","Justin","Ong").Scan(&id)
    if err != nil {
        panic(err)
    }
    fmt.Println ("New record id is: ", id)
}

