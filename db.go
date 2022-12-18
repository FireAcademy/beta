package main

import (
	"os"
    "fmt"
    "time"
    "strconv"
    "database/sql"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func setupDB() {
    db_conn_string := os.Getenv("DB_CONNECTION_STRING")
    if db_conn_string == "" {
        fmt.Printf("DB_CONNECTION_STRING not specified, exiting :(\n")
        return
    }
    var err error
    DB, err = sql.Open("postgres", db_conn_string)
    if err != nil {
        panic(err)
    }

    err = DB.Ping()
    if err != nil {
        panic(err)
    }

    max_idle_conns := os.Getenv("DB_MAX_IDLE_CONNECTIONS")
    if max_idle_conns == "" {
        panic("DB_MAX_IDLE_CONNECTIONS not set")
    }
    max_open_conns := os.Getenv("DB_MAX_OPEN_CONNECTIONS")
    if max_open_conns == "" {
        panic("DB_MAX_OPEN_CONNECTIONS not set")
    }

    max_idle_conns_i, err := strconv.Atoi(max_idle_conns)
    if err != nil {
        panic(err)
    }
    max_open_conns_i, err := strconv.Atoi(max_open_conns)
    if err != nil {
        panic(err)
    }
    // Maximum Idle Connections
    DB.SetMaxIdleConns(max_idle_conns_i)
    // Maximum Open Connections
    DB.SetMaxOpenConns(max_open_conns_i)
    // Connection Lifetime
    DB.SetConnMaxLifetime(30 * time.Second)
}
