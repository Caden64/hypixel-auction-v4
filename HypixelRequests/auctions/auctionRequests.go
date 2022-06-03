package auctions

import (
	"encoding/json"
	"errors"
	"fmt"
	"hypixel-auction-v4/HypixelRequests"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const url = "https://api.hypixel.net/skyblock/auctions"

// AuctionRequest to send request and then return unmarshalled data
func AuctionRequest(page int, client *http.Client) (AuctionData, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("error with new http request %v\n", err)
	}

	req.Header.Set("user-agent", "auction parser golang")
	req.URL.RawQuery = "page=" + strconv.Itoa(page)

	fmt.Println(req.URL)
	resp, err := client.Do(req)

	if err != nil {

		fmt.Printf("Error with request: %v\n", err)
		if resp == nil {

			fmt.Printf("Nothing returned\n")

		} else {
			fmt.Printf("This is what was returned %v\n", resp.Status)
		}

		return AuctionData{}, errors.New("error doing request")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadGateway {
		fmt.Println(resp.Status)

		return AuctionData{}, errors.New("request is bad ")

	} else if resp.StatusCode == http.StatusBadGateway {
		req, err = http.NewRequest(http.MethodGet, "https://api.hypixel.net/skyblock/auctions", nil)
		if err != nil {
			fmt.Printf("error with new http request %v\n", err)
		}

		req.Header.Set("user-agent", "auction parser golang")
		req.URL.RawQuery = "page=" + strconv.Itoa(page)

		fmt.Println(req.URL)
		resp, err = client.Do(req)

		if err != nil {

			fmt.Printf("Error with request: %v\n", err)
			if resp == nil {

				fmt.Printf("Nothing returned\n")

			} else {
				fmt.Printf("This is what was returned %v\n", resp.Status)
			}

			return AuctionData{}, errors.New("error doing request")
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.Status)

			return AuctionData{}, errors.New("request is bad")
		}
	}
	//fmt.Printf("Request succeeded, page: %v \n", req.URL)

	defer func(Body io.ReadCloser) {

		err = Body.Close()
		if err != nil {
			log.Panicf("unable to close body: %v\n", err)
		}

	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AuctionData{}, err

	}

	var data AuctionData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return AuctionData{}, err

	}

	return data, nil
}

func AllPagesAuctions() *AllAuctionData {
	var wg sync.WaitGroup

	client := HypixelRequests.NewClient()

	var auctions AllAuctionData
	wg.Add(1)

	err := auctions.AddData(&wg, 0, client)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for i := 1; i < auctions.Pages; i++ {
		wg.Add(1)

		err = auctions.AddData(&wg, i, client)

		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
	wg.Wait()

	return &auctions

}

func (c *AllAuctionData) AddData(wg *sync.WaitGroup, page int, client *http.Client) error {

	current, err := AuctionRequest(page, client)
	defer wg.Done()
	if err != nil {
		return err
	}

	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.LastUpdated = current.LastUpdated
	c.Pages = current.TotalPages
	for _, i := range current.Auctions {

		c.Auctions = append(c.Auctions, i)
	}

	return nil

}
