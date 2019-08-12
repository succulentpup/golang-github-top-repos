package db

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
    // TODO: get all this details from environment variables
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "password"
    dbName := "gitrepos"
    dbSever := "tcp(127.0.0.1:3306)"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbServer+"/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func insertRepos(db, values ) {
    insForm, err := db.Prepare("INSERT INTO Repos(repoId, name, author, stars, created_at) VALUES(?,?,?,?)"
    if err != nil {
        panic(err.Error())
    }
    insForm.Exec(values.repoId, values.name, values.author, values.starts, values.created_at)
}

fun getMaxRepoCreatedAt(db) (string) {
    results, err := db.Query("SELECT max(repoCreatedAt) FROM Repos")
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    // simple supporting structure
    type Tag struct {
        created_at string `json:"repoCreated_at"`
    }
    var tag Tag
    err = results.Scan(&tag.created_at)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    return &tag.created_at;
}
