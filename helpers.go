package aoc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	FileFormat = "day%d/input" 
	SessionKey = "SESSION"
	URLFormat  = "https://adventofcode.com/2019/day/%d/input"
)

func GetDayInput(day int) (string, error) {
	filePath := fmt.Sprintf(FileFormat, day)
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return inputFromFile(filePath)
	}

	session := os.Getenv(SessionKey)
	if session != "" {
		url := fmt.Sprintf(URLFormat, day)
		return inputFromRemote(url, session)
	}

	return "", fmt.Errorf("input for day %d not found", day)
}

func inputFromFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	out, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(out), "\n"), nil
}

func inputFromRemote(url, session string) (string, error) {
	client := new(http.Client)
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))
	
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(out), "\n"), nil
}