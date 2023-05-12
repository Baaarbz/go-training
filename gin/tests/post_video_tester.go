package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/v1/videos"
	method := "POST"

	payload := strings.NewReader(`{
    "title": "Fight Club",
    "description": "Test description",
    "url": "https://www.youtube.com/watch?v=O1nDozs-LxI&pp=ygUVZmlnaHRpbmcgY2x1YiB0cmFpbGVy"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic YmFyYno6cHdk")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
