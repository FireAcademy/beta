package main

import (
	"os"
    "time"
    "strconv"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
    db_conn_string := os.Getenv("DB_CONNECTION_STRING")
    if db_conn_string == "" {
        panic("DB_CONNECTION_STRING not specified, exiting :(\n")
        return
    }
    var err error
    DB, err = gorm.Open(postgres.Open(db_conn_string))
    if err != nil {
        panic(err)
    }

    sqlDB, err := DB.DB()
    if err != nil {
        panic(err)
    }

    err = sqlDB.Ping()
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
    sqlDB.SetMaxIdleConns(max_idle_conns_i)
    // Maximum Open Connections
    sqlDB.SetMaxOpenConns(max_open_conns_i)
    // Connection Lifetime
    sqlDB.SetConnMaxLifetime(30 * time.Second)
}
