package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
)

type jobApplicationResult struct {
	Name   string
	Age    int
	Passed bool
	Skills []string
}

const url = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func getPassedApplications(list []jobApplicationResult) map[string]struct{} {
	f, err := os.Create("passed_Go_Java.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	set := map[string]struct{}{}

	for _, applicant := range list {
		if applicant.Passed {
			for _, skill := range applicant.Skills {
				if skill == "Java" || skill == "Go" {
					set[applicant.Name] = struct{}{}
					f.WriteString(fmt.Sprint(applicant.Name) + "\n")
					break
				}
			}
		}
	}

	defer f.Close()
	f.Sync()
	return set
}

func getAllApplications(list []jobApplicationResult, passed map[string]struct{}) {
	f, err := os.Create("skills.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	for _, applicant := range list {
		_, found := passed[applicant.Name]
		if found {
			f.WriteString(fmt.Sprint(applicant.Name) + " - ")
			for i, skill := range applicant.Skills {
				if i+1 == len(applicant.Skills) {
					f.WriteString(fmt.Sprint(skill) + "\n")
				} else {
					f.WriteString(fmt.Sprint(skill) + ", ")
				}
			}
		}
	}

	defer f.Close()
	f.Sync()
}

func main() {
	httpClient := pester.New()

	httpResponse, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards jobApplication API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of jobApplication API response"),
		)
	}

	var decodedContent []jobApplicationResult
	err = json.Unmarshal(bodyContent, &decodedContent)

	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	passed := getPassedApplications(decodedContent)

	getAllApplications(decodedContent, passed)
}
