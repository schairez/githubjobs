package positions

import (
	"net/http"
)

type (

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
	//JobsClient is a Github Jobs REST API Client
	JobsClient struct {
		Client     *http.Client
		JobsURLApi string
		// Timeout          time.Time time.Second * 30
		JSONBodyResponse []JobsListing
	}
)

//ListJobsResponse struct
// ListJobsResponse struct {
// 	Responses []JobsListing
// }
