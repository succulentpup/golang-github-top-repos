package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    "encoding/json"
    "topRepos/db"
    "github.com/graphql-go/graphql"
)

type formattedData struct {
   repoId string
   repoName string
   repoAuthor string
   repoStars string
   repoCreated_at string
}

func getTopStarredReposByN(db, n) {
// TODO: parse/marshal/unMarshal the output as per golang & return
    return db.getTopStarredReposByN(n)
}

// TODO: read the URI from the environment variable
func main () {
    dbConnection := db.dbConn()
    fmt.Println("Connection to DB is SUCCESSFUL")
    topStarredRepos := getTopStarredReposByN(dbConnection, 10) //TODO: get it thru environment
    fmt.Println("Top 10 Repos based on the number of stars")

    // TODO: I've already filtered a ton of attributes in getRepos
    // while saving the repo details into the DB.
    // If that's not the case, then we would have seen overFetching issue
    // However, to simulate that lets assume 5 fields returning to end user itself is over overFetching
    // Hence we decided to take advante of API query language graphQL

    // following code is an example found in internet on how to create simple graphQL server
    // that returns only required fields
    fields := graphql.Fields{
       "repoName": &graphql.Field{
           Type: graphql.String,
           Resolve: func(p graphql.ResolveParams) (interface{}, error) {
              return "world", nil
           },
       },
    }
    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, err := graphql.NewSchema(schemaConfig)
        if err != nil {
           log.Fatalf("failed to create new schema, error: %v", err)
        }
        // Query
        query := `
            {
                repoName
            }
        `
        params := graphql.Params{Schema: schema, RequestString: query}
        r := graphql.Do(params)
        if len(r.Errors) > 0 {
            log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
        }
        rJSON, _ := json.Marshal(r)
        fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
