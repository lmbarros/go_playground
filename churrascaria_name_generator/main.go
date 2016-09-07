package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handleGimmeAName(w http.ResponseWriter, r *http.Request) {
	nouns := [...][2]string{
		{"Galpão", "M"},
		{"Fogo", "M"},
		{"Churrasco", "M"},
		{"Braseiro", "M"},
		{"Chama", "F"},
		{"Brasa", "F"},
		{"Estância", "F"},
		{"Estrela", "F"},
		{"Tradição", "F"},
	}

	adjectivesM := [...]string{
		"Gaúcho",
		"do Sul",
		"Crioulo",
		"do Rio Grande",
		"Nativo",
		"Nativista",
		"da Fronteira",
		"Gauchesco",
		"Gaudério",
		"Farrapo",
	}

	adjectivesF := [...]string{
		"Gaúcha",
		"do Sul",
		"Crioula",
		"do Rio Grande",
		"Nativa",
		"Nativista",
		"da Fronteira",
		"Gauchesca",
		"Gaudéria",
	}

	i := rand.Intn(len(nouns))
	noun := nouns[i][0]

	var adjective string

	if nouns[i][1] == "M" {
		adjective = adjectivesM[rand.Intn(len(adjectivesM))]
	} else {
		adjective = adjectivesF[rand.Intn(len(adjectivesF))]
	}

	fmt.Fprintf(w, "%s %s", noun, adjective)
}

func handleEverythingElse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Please access /gimmeaname to get a churrascaria name.")
}

func main() {
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/gimmeaname", handleGimmeAName)
	http.HandleFunc("/", handleEverythingElse)

	http.ListenAndServe(":8080", nil)
}
