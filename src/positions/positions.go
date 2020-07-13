package positions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	//ListJobsResponse struct
	ListJobsResponse struct {
		Responses []JobsListing
	}
	//JobsListing struct
	JobsListing struct {
		ListingID   string `json:"id"`
		ListingType string `json:"type"`
		URL         string `json:"url"`
		CreatedAt   string `json:"created_at"`
		Company     string `json:"company"`
		CompanyURL  string `json:"company_url"`
		Location    string `json:"location"`
		Title       string `json:"title"`
		Description string `json:"description"`
		HowToApply  string `json:"how_to_apply"`
	}
)

// listMapType list of objs
// type listMapType []map[string]string

const (
	//BaseURLGithubPositionsAPI comment
	BaseURLGithubPositionsAPI = "https://jobs.github.com/positions.json?"
)

func buildGithubJobsURL(uriParams map[string]string) string {
	size := len(uriParams)
	cnt := 0
	urlparams := ""
	for k, v := range uriParams {
		urlparams += k + "=" + v
		cnt++
		if size != cnt {
			urlparams += "&"
		}

	}
	return BaseURLGithubPositionsAPI + urlparams

}

/*

description — A search term, such as "ruby" or "java". This parameter is aliased to search.
location — A city name, zip code, or other location search term.
lat — A specific latitude. If used, you must also send long and must not send location.
long — A specific longitude. If used, you must also send lat and must not send location.
full_time — If you want to limit results to full time positions set this parameter to 'true'.

*/

//GET /positions.json
func getPositions(description, location string) {

	githubURL := getActiveGithubJobListingsURI(description, location)
	fmt.Println(githubURL)
	resp, err := http.Get(githubURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(bodyBytes))
	var ResponseStruct ListJobsResponse
	json.Unmarshal(bodyBytes, &ResponseStruct)
	ResponseJSON, _ := json.MarshalIndent(ResponseStruct, "", "\t")
	err = ioutil.WriteFile("output.json", ResponseJSON, 0644)

}
