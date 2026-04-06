package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

const credsFile = "credentials.json"

func main() {
	var problem problemStat
	url := "https://hackattic.com/challenges/dockerized_solutions/problem?access_token=aaa699dde38ea86a"

	if data, err := os.ReadFile(credsFile); err == nil {
		json.Unmarshal(data, &problem)
	} else {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("error: Get url")
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error: reading response body")
			return
		}

		if err := json.Unmarshal(body, &problem); err != nil {
			fmt.Println("error: unmarshaling problem")
			return
		}

		os.WriteFile(credsFile, body, 0644)
	}

	pretty, _ := json.MarshalIndent(problem, "", "  ")
	fmt.Println(string(pretty))

	fmt.Println("waiting for manual execution: ")
	fmt.Scanln()

	triggerUrl := "https://hackattic.com/_/push/" + problem.TriggerToken
	triggerPayload := TriggerPost{
		RegistryHost: "hackattic-registry-sagar-hackattic-docker-cold-glitter-665.fly.dev",
	}

	jsonTrigger, _ := json.Marshal(triggerPayload)
	resp, err := http.Post(triggerUrl, "application/json", bytes.NewBuffer(jsonTrigger))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
