//search github jobs
//https://jobs.github.com/positions.json?description=python&location=new+york

//search linkedin jobs
//https://api.linkedin.com/v1/job-search

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	githubAPI := getActiveGithubJobListings("python", "new+york")
	fmt.Println(githubAPI)
	resp, err := http.Get(githubAPI)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// if resp.StatusCode < 200 || resp.StatusCode > 299 {
	// 	data, err := ioutil.ReadAll(resp.Body)

	// }

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(bodyBytes))

}

const (
	//BaseURLGithubAPI comment
	BaseURLGithubAPI = "https://jobs.github.com/positions.json"
)

func getActiveGithubJobListings(searchQuery, location string) string {
	return fmt.Sprintf("%s?description=%s&location=%s", BaseURLGithubAPI, searchQuery, location)
}
