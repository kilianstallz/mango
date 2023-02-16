package check

import (
	"encoding/json"
	"log"
	"net/http"
)

func QueryLatestGoVersion() (string, error) {
	// get the lastest version from the Go site
	response, err := http.Get("https://go.dev/dl/?mode=json")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// read the body
	type data struct {
		Version string `json:"version"`
		Stable  bool   `json:"stable"`
		Files   []struct {
			OS       string `json:"os"`
			Arch     string `json:"arch"`
			Kind     string `json:"kind"`
			Filename string `json:"filename"`
		} `json:"files"`
	}

	var body []data

	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}

	return body[0].Version, nil
}
