# jobs_discovery

Package jobs_discovery is a wrapper for the github jobs api
More info on github jobs api: https://jobs.github.com/api

## Installation

Import the library with

```golang
import "github.com/schairez/jobs-discovery/githubjobs"
```

## Usage

```golang
package main

import (
    "fmt"
    "github.com/schairez/jobs-discovery/githubjobs"
    "net/http"
    "time"
)

func main() {

    paramsMap := map[string]string{
        "description": "python",
        "location":    "sf",
        "full_time":   "true"}

    var c = &githubjobs.JobsClient{
        Client: &http.Client{Timeout: 5 * time.Second}}
    c.SetGithubJobsURL(paramsMap)
    var jobsData []JobsListing
    jobsData = c.GetPositionsResultStruct()
    fmt.Println(jobsData[0].Description)
    fmt.Println(stripHTML(jobsData[0].Description))

}
```
