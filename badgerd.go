package main

import (
	"encoding/json"
	"fmt"
	"github.com/IDMWORKS/badgerd/badge"
	"github.com/IDMWORKS/badgerd/status"
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
	Port  string `json:port`
}

func main() {
	config = readConfig()
	http.HandleFunc("/badge/", badgeHandler)

	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatal(err)
	}
}

func readConfig() *Config {
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
	return conf
}

func badgeHandler(writer http.ResponseWriter, req *http.Request) {
	urlBits := strings.Split(req.URL.Path, "/")
	project := urlBits[2]

	var verb string
	if len(urlBits) >= 4 {
		verb = urlBits[3]
	} else {
		verb = "build-status"
	}

	status, err := getStatus(project)
	if err != nil {
		log.Println("Error - " + err.Error())
		http.ServeFile(writer, req, "badges/"+badge.BuildErrorBadge)
		return
	}

	badgeFile := badge.BuildErrorBadge
	switch verb {
	case "build-status":
		badgeFile, err = badge.ForBuildStatus(status)
	case "rcov":
		badgeFile, err = badge.ForRCov(status)
	default:
		err = fmt.Errorf("Unknown verb: '%s'", verb)
	}

	if err != nil {
		log.Println("Error - " + err.Error())
	}

	log.Printf("%s/%s - %s", project, verb, badgeFile)
	http.ServeFile(writer, req, "badges/"+badgeFile)
}

func getStatus(project string) (*status.BuildStatus, error) {
	url := "http://" + config.Host + "/job/" + project + "/api/json"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.User, config.Token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	status := status.BuildStatus{}
	err = json.Unmarshal(body, &status)

	if err != nil {
		return nil, err
	}

	return &status, nil
}
