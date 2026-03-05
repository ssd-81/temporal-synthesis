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

func main() {

	url := "https://hackattic.com/challenges/dockerized_solutions/problem?access_token=aaa699dde38ea86a"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	defer resp.Body.Close()

	var problem problemStat
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&problem)
	if err != nil {
		fmt.Println("error encountered while decoding input problem")
	}
	
	// pretty, _ := json.MarshalIndent(problem, "", "  ")
	// fmt.Println(string(pretty))


	trToken := "b3a23158.a53d.44a4.aa62.f324a91dbc60"

	// posting json for trigger
	triggerUrl := "https://hackattic.com/_/push/" + trToken
	triggerPayload := TriggerPost{
		RegistryHost: "https://hackattic-registry-sagar-hackattic-docker-cold-glitter-665.fly.dev/",
	}

	jsonTrigger, _ := json.Marshal(triggerPayload)
	resp, err = http.Post(triggerUrl, "application/json", bytes.NewBuffer(jsonTrigger))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))


	// submitUrl := "https://hackattic.com/challenges/dockerized_solutions/solve?access_token=aaa699dde38ea86a"
	// solution := SolutionPost{Secret: "test"}
	// jsonPayload, _ := json.Marshal(solution)

	// resp, err = http.Post(submitUrl, "application/json", bytes.NewBuffer(jsonPayload))
	// if err != nil {
	// 	return
	// }
	// defer resp.Body.Close()

	// finalBody, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(finalBody))
}