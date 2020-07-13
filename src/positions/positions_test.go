package positions

import (
	"net/http"
	"testing"
	"time"
)

func TestBuildGithubJobsURL(t *testing.T) {
	paramsMap := map[string]string{"description": "python", "location": "sf", "full_time": "true"}
	jobsURL := buildGithubJobsURL(paramsMap)
	t.Log(jobsURL)
	if jobsURL != "https://jobs.github.com/positions.json?description=python&location=sf&full_time=true" {
		t.Error()
	}

}

func TestGetJSONresponse(t *testing.T) {

	paramsMap := map[string]string{
		"description": "python",
		"location":    "sf",
		"full_time":   "true"}

	var c = &JobsClient{
		Client: &http.Client{Timeout: 5 * time.Second}}
	c.SetGithubJobsURL(paramsMap)
	var jobsData []JobsListing
	jobsData = c.GetJSONResponse()
	t.Log(jobsData)
	// t.Log(c.GetJSONResponse())

}

func TestStringHTMLSplitter(t *testing.T) {

	var body = `
<p>We are looking for a tenacious, passionate, creative, software-driven individual to guide the design and development of our payments software.</p> <p><strong>OBJECTIVE #1 Take on a project and see it to its end.</strong></p> <ul> <li>Potential Obstacles: We work by 6 week cycles (similar to sprints) where it can be easy to lose focus or forget to plan ahead to ensure the successful completion of your projects.</li> <li>Actions: Understand the value of a successful cycle completion and use your tenacity to carry each project to the very end.</li> <li>Results: Your project completion numbers in Jira are stellar and you have concentration available for the handful of projects on your plate.</li> </ul> <p><strong>OBJECTIVE #2 Work as a synchronous team in a remote environment.</strong></p> <ul> <li>Potential Obstacles: Working remote has many personal and team-based challenges from timezone differences, to communication and trust of your colleagues.</li> <li>Actions: Develop deep working relationships with your colleagues to know what they’re working on, and how to support each other.</li> <li>Results: The t\
\ulture is strong, Rebilly’s bigger projects efficiently move forward.</li> </ul> <p><strong>OBJECTIVE #3 /*lp your colleagues raise their skills with thoughtful code reviews and feedback on ideas.</strong></p> <ul> <li>Potential Obstacles: Your day will be filled with your own projects and priorities, and offering feedback to a remote team from afar can be intimidating.</li> <li>Actions: Prioritize your schedule and share your unique knowledge and experience to help make your team stronger.</li> <li>Results: Your team is more precise, creative, and their knowledge is deepened with your help.</li> </ul> <p><strong>KEY FUNCTIONAL AREAS OF RESPONSIBILITY</strong></p> <ul> <li> <p>What will you manage? Each teammate receives a role within a small team per cycle that can include a handful of small projects, one large project, or staying outside of the cycle to support clients, bugs, and cycle planning. Along with that, you’ll be responsible for ensuring adherence to programming and documentation policies, code standards, testing, release, and reporting updates on your projects.</p> </li> <li> <p>What will you contribute to? You will contribute to code reviews, generating projects for each cycle, feedback for ideas (including the ideas for other areas of the company), and hiring new teammates in the Product Team.</p> </li> <li> <p>What will you support? Outside of generally supporting your teammates in the Product Circle, Rebilly is an open environment where everyone is welcome in every area of the company. You are welcome to support any project you feel called to participate in.</p> </li> <li> <p>How much support per project? Teams working in a cycle can be 1, 2, or sometimes 4 people. You will lean on your cycle teammates to support the completion of the cycle, as well as receive support from teammates supporting the cycle. You can ask as many teammates to join your project as you need to complete the mission.</p> </li> </ul> <p><strong>SOME OF THE PROJECTS WE ARE WORKING ON</strong></p> <ul> <li>Customer Payments Portal - create a secure and highly customizable customer payments portal</li> <li>Proficiency in Information Architecture - support usability in our database and more</li> <li>Payment Gateway Integrations - expand our library of integrations</li> </ul> <p><strong>APPLY!</strong></p> <p>Please send your resume through Apply By API: <a href="https://app.applybyapi.com/posting/7/">https://app.applybyapi.com/posting/7/</a></p> <p><strong>HIRING PROCESS</strong></p> <ul> <li>Apply By API* -<a href="https://app.applybyapi.com/posting/7/">https://app.applybyapi.com/posting/7/</a> </li> <li>Small code challenge (unless you have extensive open source contributions)</li> <li>Interview with the Product Team</li> <li>Interview with the Product Team Lead</li> <li>Offer and Hiring</li> <li>The total process should take less than 2 weeks. *We are only accepting applications through ApplyByAPI at this time.</li> </ul> <p><strong>WORKING AT REBILLY</strong></p> <p>Rebilly’s Purpose: Find and create freedom through your work.</p> <p>Rebilly offers a competitive salary, training and development, birthday lunches, and provides the computer of your choosing. Further perks and benefits are dependent on your work location of choice.</p> <p>Rebilly is a Teal Organization, meaning our organizational structure may differ than what you’re used to, but we think you’ll be pleasantly surprised. In a very tiny nutshell, this means that as an organization we value self-management, self-organization, as well as the wholeness of the individuals that make up our team (meaning you should be yourself at work, and do the work that inspires you.)</p> <p>Come check us out at <a href="https://www.rebilly.com/careers/">https://www.rebilly.com/careers/</a></p> <p><strong>ABOUT REBILLY</strong></p> <p>Rebilly is now a payments software after pivoting from a focus on subscription billing in October 2019. Rebilly’s comprehensive feature set is built to help our merchants get more of their customers from the order form to the thank you page, and more renewals paid. We took lessons learned from years of experience in the subscription business and millions of subscribers and applied them to make our payment system clever, flexible, and easy to use for our clients.</p> ",
how_to_apply: "<p>Please send your resume through Apply By API: <a href="https://app.applybyapi.com/posting/7/">https://app.applybyapi.com/posting/7/</a></p>
	 `
	t.Log(stripHTML(body))

}
