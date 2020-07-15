# githubjobs

Package githubjobs is a wrapper for the github jobs api
More info on github jobs api: https://jobs.github.com/api

## Installation

Import the library with

```golang
import "github.com/schairez/githubjobs"
```

## Usage

```golang
package main

import (
    "fmt"
    "github.com/schairez/githubjobs"
    "net/http"
    "time"
)

func main() {

    paramsMap := map[string]string{
        "description": "python",
        "location":    "sf",
        "full_time":   "true"}

    c := NewClient()
    c.SetGithubJobsURL(paramsMap)
    jobsData,err := c.GetPositionsResultStruct()
    if err != nil {
        panic(err)

    }
    fmt.Println(jobsData[0].Description)
    fmt.Println(stripHTML(jobsData[0].Description))

}
```
