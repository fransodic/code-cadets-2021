package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

type pokemonAPIResponse struct {
	Name                   string
	LocationAreaEncounters string `json:"location_area_encounters"`
}

type pokemonEncounters struct {
	LocationArea struct {
		Name string
	} `json:"location_area"`
}

type pokemonStruct struct {
	Name     string
	Location []string
}

func parseProgramArguments(pokemon *string) {
	flag.StringVar(pokemon, "pokemon", "", "desired pokemon's name or number")

	flag.Parse()
}

func fetchHTTPResponse(url string) (*http.Response, error) {
	httpClient := pester.New()
	return httpClient.Get(url)
}

func getResponseContent(response *http.Response) ([]byte, error) {
	return ioutil.ReadAll(response.Body)
}

func getDataFromURL(url string) ([]byte, error) {

	httpResponse, err := fetchHTTPResponse(url)
	if err != nil {
		return nil, errors.WithMessage(err, "fetching data from pokemon API")
	}

	bodyContent, err := getResponseContent(httpResponse)
	if err != nil {
		return nil, errors.WithMessage(err, "reading body of pokemon API response")
	}

	return bodyContent, nil
}

func marshalPokemonStruct(pokemon string, locations []pokemonEncounters) pokemonStruct {
	var resultPokemon pokemonStruct

	resultPokemon.Name = pokemon
	for _, locationArea := range locations {
		resultPokemon.Location = append(resultPokemon.Location, locationArea.LocationArea.Name)
	}

	return resultPokemon
}

func printResults(pokemon pokemonStruct) {
	jsonToPrint, err := json.Marshal(pokemon)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "marshaling the result structure"),
		)
	}

	fmt.Println(string(jsonToPrint))
}

const url = "https://pokeapi.co/api/v2/pokemon/"

func main() {

	var pokemon string
	var pokemonResult pokemonAPIResponse
	var pokemonLocations []pokemonEncounters

	parseProgramArguments(&pokemon)

	if len(pokemon) == 0 {
		log.Fatal(
			errors.New("pokemon not specified"),
		)
	}

	bodyContent, err := getDataFromURL(url + pokemon)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bodyContent, &pokemonResult)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	pokemon = pokemonResult.Name

	bodyContent, err = getDataFromURL(pokemonResult.LocationAreaEncounters)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "fetching data from pokemon API"),
		)
	}

	err = json.Unmarshal(bodyContent, &pokemonLocations)
	if err != nil {
		log.Fatal(err)
	}
	finalJson := marshalPokemonStruct(pokemon, pokemonLocations)
	printResults(finalJson)
}
