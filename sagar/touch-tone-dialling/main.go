package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)



type problemStatement struct {
	WavUrl string `json:"wav_url"`

}

type solution struct {
	
}

// send: https://hackattic.com/challenges/touch_tone_dialing/solve?access_token=aaa699dde38ea86a

func main() {

	resp, err := http.Get("https://hackattic.com/challenges/touch_tone_dialing/problem?access_token=aaa699dde38ea86a")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var problem problemStatement
	json.Unmarshal(body, &problem)
	fmt.Println(problem)

	currDir, err := os.Getwd()
	fileName := "challenge.wav"
	fullPath := filepath.Join(currDir, fileName)

	err = DownloadFile(problem.WavUrl, fullPath)
	if err != nil {
		fmt.Println("couldn't not get the current working directory")
		return
	}
	fmt.Println("current working directory: ", currDir)
	
	err = DownloadFile(problem.WavUrl, currDir)
	if err != nil {
		fmt.Println("file could not be downloaded")
		return 
	}
	


	

	data := solution{
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
	panic(err)
	}


	// Make POST request with JSON data
	resp, err = http.Post("https://hackattic.com/challenges/password_hashing/solve?access_token=aaa699dde38ea86a", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
	panic(err)
	}
	defer resp.Body.Close()
	resBytes, _ := io.ReadAll(resp.Body)
	fmt.Println()
	fmt.Println(string(resBytes))

}

func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}