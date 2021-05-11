package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

type jobApplicationResult struct {
	Name   string
	Age    int
	Passed bool
	Skills []string
}

func fetchHTTPResponse() (*http.Response, error) {
	httpClient := pester.New()
	return httpClient.Get(url)
}

func getResponseContent(response *http.Response) ([]byte, error) {
	return ioutil.ReadAll(response.Body)
}

const url = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func filterPassedApplications(list []jobApplicationResult) []jobApplicationResult {
	var filteredSlice []jobApplicationResult

	for _, applicant := range list {
		if applicant.Passed && containsSelectedSkills(applicant.Skills) {
			filteredSlice = append(filteredSlice, applicant)
		}
	}

	return filteredSlice
}

func containsSelectedSkills(skillSlice []string) bool {
	for _, skill := range skillSlice {
		if skill == "Java" || skill == "Go" {
			return true
		}
	}
	return false
}

func writeToFile(passed []jobApplicationResult) {
	f, err := os.Create("skills.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	for _, applicant := range passed {
		toWrite := fmt.Sprintf("%v - %v\n", applicant.Name, strings.Join(applicant.Skills, ", "))
		f.WriteString(toWrite)
	}

	defer f.Close()
	f.Sync()
}

func main() {

	httpResponse, err := fetchHTTPResponse()
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "fetching data"),
		)
	}

	bodyContent, err := getResponseContent(httpResponse)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of job application API response"),
		)
	}

	var applicationResults []jobApplicationResult

	err = json.Unmarshal(bodyContent, &applicationResults)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	filteredApplicants := filterPassedApplications(applicationResults)

	writeToFile(filteredApplicants)
}
