package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8080"
	createEvent := "/event"
	listEvent := "/events"
	fmt.Println("URL:>", url)

	{ //POST JSON to server
		fmt.Println("------ SENDING JSON EVENT TO SERVER ----------")
		var jsonStr = []byte(`{"ID":"2", "Title":"Created Event", "Description":"This is an event added by POST request from GO"}`)
		req, err := http.NewRequest("POST", url+createEvent, bytes.NewBuffer(jsonStr))
		//req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}

	{ //GET function /events from server
		fmt.Println("------ GET ALL EVENTS FROM SERVER ----------")
		resp, err := http.Get(url + listEvent)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
}
