package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	config *Config
)

type Config struct {
	Host  string `json:host`
	User  string `json:user`
	Token string `json:token`
}

type buildStatus struct {
	DisplayName string `json:displayName`
	Url         string `json:url`
	Color       string `json:color`
}

func main() {
	config, _ = readConfig()

	http.HandleFunc("/badger/", badgeHandler)
	http.ListenAndServe(":8080", nil)
}

func readConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	conf := &Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf, nil
}

func getBadge(status string) string {
	switch {
	case status == "blue":
		return "build-passing.svg"
	case status == "red":
		return "build-failing.svg"
	case strings.Index(status, "_anime") > 0:
		return "build-building.svg"
	}
	return "build-error.svg"
}

func badgeHandler(writer http.ResponseWriter, req *http.Request) {
	project := strings.Split(req.URL.Path, "/")[2]
	status, _ := getStatus(project)
	badgeFile := getBadge(status)
	http.ServeFile(writer, req, "badges/"+badgeFile)
}

func getStatus(project string) (string, error) {
	url := "http://" + config.Host + "/job/" + project + "/api/json"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(config.User, config.Token)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	status := buildStatus{}
	err = json.Unmarshal(body, &status)

	if err != nil {
		return "", err
	}

	return status.Color, nil
}
