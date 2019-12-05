package aoc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	FileFormat = "day%d/input"
	SessionKey = "AOC_SESSION_KEY"
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

func InputToSlice(input string, separator string) []string {
	return strings.Split(input, separator)
}

func SliceToDigits(slice []string) ([]int, error) {
	digits := make([]int, len(slice))
	for idx, item := range slice {
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		digits[idx] = num
	}
	return digits, nil
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
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
