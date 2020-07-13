# jobs_discovery

Package jobs_discovery is a wrapper for the github jobs api
https://jobs.github.com/api

## Installation

Import the library with

```golang
import "github.com/denisbrodbeck/jobs-discovery"
```

## Usage

```golang
package main

import (
	"fmt"
	"github.com/schairez/jobs-discovery"
)

func main() {

    paramsMap := map[string]string{
        "description": "python",
        "location":    "sf",
        "full_time":   "true"}

    var c = &JobsClient{
        Client: &http.Client{Timeout: 5 * time.Second}}
    c.SetGithubJobsURL(paramsMap)
    var jobsData []JobsListing
    jobsData = c.GetJSONResponse()
    fmt.Println(jobsData[0].Description)
    fmt.Println(stripHTML(jobsData[0].Description))

}
```
