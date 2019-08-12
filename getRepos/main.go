package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    "encoding/json"
    "topRepos/db"
)

type formattedData struct {
   repoId string
   repoName string
   repoAuthor string
   repoStars string
   repoCreated_at string
}

func getEarlierDateByNdays(n int) (string) {
    now := time.Now()
    earlierTime := now.AddDate(0, 0, -n)
    earlierDate := earlierTime.Format("2006-01-02")
    return earlierDate
}

func getFetchReposApiUrl(queryParams string) (string) {
    // TODO: get the restAPI as a constant imported from a file
    restAPI := "https://api.github.com/search/repositories?q=created:>"
    return restAPI + queryParams
}

func pluckReqAttributesAndFormat(listOfReposHavingUnwantedAttributesAswell) ([]formattedData) {
   // TODO: process the array of json objects to pluck only required keys
   // for this challenge, I would pick the following attributes only.
   // repoId, repoName, repoAuthor, repoStars and repoCreated_at


    // Following is the dummy object that I've created
    // plucking the attributes needed for insertion would be prepared
    // by iterating thru listOfReposHavingUnwantedAttributesAswell
    var pluckedAndFormattedData []formattedData
    pluckedAndFormattedData := [
    {
        "repoId" : "1",
        "repoName": "fruitsCart",
        "repoAuthor": "ganesh",
        "repoStars": "122",
        "repoCreated_at": "2019-08-10"
    },
    {
        "repoId" : "2",
        "repoName": "nodeServiceTemplate",
        "repoAuthor": "Kyle",
        "repoStars": "28",
        "repoCreated_at": "2019-08-11"
    }
   ]
   return pluckedAndFormattedData;
}

func getOnlyQualifiedRepos(dbConnection, listOfRepos map[string]interface{}) ([]listOfRepos)  {
    // TODO: iterate through this formattedData, check on repoCreated_at
    // only return the entries that are greater than max of repoCreated_at that is in db

    recentRepoFromDB := db.getMaxRepoCreatedAt(dbConnection)
    created_at := time.now(recentRepoFromDB)

    // for each repo in listOfRepos, filter only the repo > created_at
    return qualifiedReposToInsert
}

// TODO: read the URI from the environment variable
func main () {
    earlierDate := getEarlierDateByNdays(7) //TODO: get it thru environment
    response, err := http.Get(getFetchReposApiUrl(earlierDate))
    if err != nil {
       fmt.Println("Fetching the repositories is failed");
    }
    data, _ := ioutil.ReadAll(response.Body)
    var result map[string]interface{}
    json.Unmarshal([]byte(data), &result)
        if err != nil {
            panic(err)
        }
    dbConnection := db.dbConn()
    fmt.Println("Connection to DB is SUCCESSFUL")
    latestReposToInsert := getOnlyQualifiedRepos(dbConnection, result);
    dataToInsert := pluckReqAttributesAndFormat(latestReposToInsert)
    fmt.Println("DB Connection successful");
    db.insertRepos(dbConnection, dataToInsert)
    fmt.Println("Repos got successfully inserted into DB")
}