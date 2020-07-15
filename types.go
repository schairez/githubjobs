package githubjobs

import (
	"net/http"
)

const (
	//BaseURLGithubPositionsAPI base of github jobs rest api
	BaseURLGithubPositionsAPI = "https://jobs.github.com/positions.json"
)

type (
	//Client is a Github Jobs REST API Client
	Client struct {
		Client     *http.Client
		JobsURLApi string
	}
	//ArrOfJobListingStructs is an arr of Jobslisting structs
	ArrOfJobListingStructs []JobsListing

	//JobsListing struct of interested fields of json response
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
	//OptParams is the query params for the github jobs client
	OptParams struct {
		Description *string `json:"description,omitempty"`
		Location    *string `json:"location,omitempty"`
		Lat         *string `json:"lat,omitempty"`
		Long        *string `json:"long,omitempty"`
		FullTime    *string `json:"full_time,omitempty"`
	}
)

/*

description — A search term, such as "ruby" or "java". This parameter is aliased to search.
location — A city name, zip code, or other location search term.
lat — A specific latitude. If used, you must also send long and must not send location.
long — A specific longitude. If used, you must also send lat and must not send location.
full_time — If you want to limit results to full time positions set this parameter to 'true'.

*/

// type OptParams struct {
// 	description string
// 	location    string
// 	lat         string
// 	long        string
// 	full_time   bool
// }
