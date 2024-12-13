package cmd

import (
	"testing"
)

func TestCreateUrl(t *testing.T) {
	lang = "fr"
	category = "programming"
	blacklist = []string{"nsfw", "racist"}
	typeJoke = "single"

	expected := "https://v2.jokeapi.dev/joke/programming?lang=fr&blacklistFlags=nsfw,racist&type=single"
	actual := createUrl()

	if actual != expected {
		t.Errorf("Expected URL %s, got %s", expected, actual)
	}
}
