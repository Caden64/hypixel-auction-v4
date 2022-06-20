package MojangRequests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const baseURL = "https://sessionserver.mojang.com/session/minecraft/profile/"

type name struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Properties []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"properties"`
}

func UUIDToUser(uuid string, client *http.Client) (string, error) {

	req, err := http.NewRequest(http.MethodGet, baseURL+uuid, nil)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	req.Header.Set("user-agent", "golang uuid to name conversion")

	fmt.Println(req.URL)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer func(Body io.ReadCloser) {

		err = Body.Close()
		if err != nil {
			log.Panicf("unable to close body: %v\n", err)
		}

	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err

	}

	var data name
	err = json.Unmarshal(body, &data)
	if err != nil {

		f, err := os.Create("test.json")

		if err != nil {
			log.Fatalf("error: %v \n", err)
		}

		f.Write(body)
		f.Sync()

		defer f.Close()

		return "", err
	}

	return data.Name, nil

}
