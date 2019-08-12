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

func getTopStarredReposByN(db, n) {
    //TODO: fetch the top starred repos from DB
        results, err := db.Query("SELECT repoId, name, author, stars, created_at FROM Repos order by starts limit " + n)
            if err != nil {
                panic(err.Error()) // proper error handling instead of panic in your app
            }
        // simple supporting structure
        type Tag struct {
            repoId `json:"repoId"`
            name `json:"name"`
            author `json:"author"`
            stars `json:"stars"`
            created_at `json:"created_at"`
        }
        // following is array of json objects. json object hold repo details
        var topNStarredRepos []Tag
        // TODO: iterate thru all the records, add them to json array & return
        // for each record in results
        // push that record to array of json
        return topNStarredRepos;
}