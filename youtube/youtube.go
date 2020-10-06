package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	ID    string `json:"id"`
	Stats Stats  `json:"statistics"`
}
type Stats struct {
	Views       string `json:"views"`
	Subscribers string `json:"subscriberCount"`
}

func GetSubscribers() (Item, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		log.Println(err)
		return Item{}, err
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	var res Response

	if err = json.Unmarshal(body, &res); err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	return res.Items[0], nil
}
