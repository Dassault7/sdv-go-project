/*
Copyright © 2024 JAROD GUICHARD
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Dassault7/sdv-go-project/models"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

const URL = "https://v2.jokeapi.dev/joke"

var (
	// Used for flags.
	lang      string
	category  string
	blacklist []string
	typeJoke  string
	amount    uint
	OutputUrl bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "get-joke",
	Short: "CLI to get a joke",
	Long: `CLI to get a joke from the internet.

This CLI will fetch a joke from the internet and display it to the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		parseArgs(cmd)
		url := createUrl()
		if OutputUrl {
			cmd.Println(url)
		}

		// Create a WaitGroup
		var wg sync.WaitGroup
		resultChan := make(chan models.JokeResult, amount)

		for i := uint(0); i < amount; i++ {
			wg.Add(1)
			go fetchJoke(url, &wg, resultChan)
		}

		// Wait for all goroutines to finish
		go func() {
			wg.Wait()
			close(resultChan)
		}()

		index := 1 // Index of the joke
		for result := range resultChan {
			if result.IsError {
				// Format the error message
				cmd.Println(fmt.Sprintf("Joke %d:\nError while fetching joke: %s\n", index, result.Error))
			} else {
				// Print the joke
				cmd.Println(fmt.Sprintf("Joke %d:\n%s\n", index, result.Joke))
			}
			index++
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().SortFlags = true
	rootCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Language of the joke")
	rootCmd.Flags().StringVarP(&category, "category", "c", "any", "Category of the joke")
	rootCmd.Flags().StringSliceVarP(&blacklist, "blacklist", "b", []string{}, "Blacklist of categories")
	rootCmd.Flags().UintVarP(&amount, "amount", "a", 1, "Amount of jokes to fetch")
	rootCmd.Flags().StringVarP(&typeJoke, "type", "t", "", "Type of joke to fetch")
	rootCmd.Flags().BoolVarP(&OutputUrl, "output-url", "o", false, "Output the URL to fetch the joke")
}

// exitIfNotValid checks if the value is in the expected values
// and exits if it is not
func exitIfNotValid(cmd *cobra.Command, value string, expected []string, name string) {
	for _, e := range expected {
		if value == e {
			return
		}
	}
	cmd.Printf("Invalid %s: %s\nExpected values: %v\n", name, value, expected)
	os.Exit(1)
}

// parseArgs checks if the arguments are valid
func parseArgs(cmd *cobra.Command) {
	// expected values
	langExpected := []string{"en", "fr", "de", "es", "pt", "cs", "pt"}
	catExpected := []string{"any", "misc", "programming", "pun", "spooky", "christmas"}
	typeExpected := []string{"single", "twopart", ""}
	blacklistExpected := []string{"nsfw", "religious", "political", "racist", "sexist", "explicit"}

	// check the values of the arguments
	exitIfNotValid(cmd, lang, langExpected, "lang")
	exitIfNotValid(cmd, category, catExpected, "category")
	exitIfNotValid(cmd, typeJoke, typeExpected, "type")
	for _, b := range blacklist {
		exitIfNotValid(cmd, b, blacklistExpected, "blacklist")
	}
}

// createUrl creates the URL to fetch the joke
func createUrl() string {
	url := fmt.Sprintf("%s/%s", URL, category)
	var args []string

	if lang != "en" {
		args = append(args, "lang="+lang)
	}

	if len(blacklist) > 0 {
		blacklistFlags := strings.Join(blacklist, ",")
		args = append(args, "blacklistFlags="+blacklistFlags)
	}

	if typeJoke != "" {
		args = append(args, "type="+typeJoke)
	}

	if len(args) > 0 {
		url += "?" + strings.Join(args, "&")
	}
	return url
}

// fetchJoke fetches a joke from the internet
// and sends the result to the resultChan
func fetchJoke(url string, wg *sync.WaitGroup, resultChan chan<- models.JokeResult) {
	defer wg.Done() // Signale que cette goroutine est terminée

	resp, err := http.Get(url)
	if err != nil {
		resultChan <- models.JokeResult{
			Error:   fmt.Sprintf("Error while fetching joke: %v", err),
			IsError: true,
		}
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	switch resp.StatusCode {
	case http.StatusOK:
		var jokeResponse models.JokeResponse
		if err := json.NewDecoder(resp.Body).Decode(&jokeResponse); err != nil {
			resultChan <- models.JokeResult{
				Error:   fmt.Sprintf("Error decoding joke response: %v", err),
				IsError: true,
			}
			return
		}

		// Construire le message de la blague
		if jokeResponse.Type == "single" {
			resultChan <- models.JokeResult{Joke: jokeResponse.Joke}
		} else if jokeResponse.Type == "twopart" {
			joke := fmt.Sprintf("%s\n\n%s", jokeResponse.Setup, jokeResponse.Delivery)
			resultChan <- models.JokeResult{Joke: joke}
		}

	case http.StatusBadRequest:
		var jokeError models.JokeResponseError
		if err := json.NewDecoder(resp.Body).Decode(&jokeError); err != nil {
			resultChan <- models.JokeResult{
				Error:   fmt.Sprintf("Error decoding error response: %v", err),
				IsError: true,
			}
			return
		}
		cause := strings.Join(jokeError.CausedBy, ", ")
		resultChan <- models.JokeResult{
			Error:   fmt.Sprintf("Error while fetching joke: %s\nCause: %s", jokeError.Message, cause),
			IsError: true,
		}

	default:
		resultChan <- models.JokeResult{
			Error:   fmt.Sprintf("Unexpected status code: %d", resp.StatusCode),
			IsError: true,
		}
	}
}
