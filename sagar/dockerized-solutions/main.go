package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type problemStat struct {
	Credentials struct {
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"credentials"`
	IgnitionKey  string `json:"ignition_key"`
	TriggerToken string `json:"trigger_token"`
}

type TriggerPost struct {
	RegistryHost string `json:"registry_host"`
}

type SolutionPost struct {
	Secret string `json:"secret"`
}

// testing for flow
// caching the problem is invalidates it
const (
	accessToken  = "aaa699dde38ea86a"
	registryHost = "hackattic-registry-sagar-hackattic-docker-damp-glade-7985.fly.dev"
)

func main() {
	var problem problemStat
	// not caching the problem
	url := "https://hackattic.com/challenges/dockerized_solutions/problem?access_token=" + accessToken
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching problem:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &problem); err != nil {
		fmt.Println("error unmarshaling:", err)
		return
	}
	prettyProblem, _ := json.MarshalIndent(problem, "", "  ")
	fmt.Println(string(prettyProblem))

	triggerUrl := "https://hackattic.com/_/push/" + problem.TriggerToken
	triggerPayload := TriggerPost{
		RegistryHost: "hackattic-registry-sagar-hackattic-docker-damp-glade-7985.fly.dev",
	}

	jsonTrigger, _ := json.Marshal(triggerPayload)
	resp, err = http.Post(triggerUrl, "application/json", bytes.NewBuffer(jsonTrigger))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
