package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type APIResponse struct {
	TotalSize          float64
	Start              float64
	NextCollectionLink string
	PrevCollectionLink string
	Entries            []Entry
}

func (this *APIResponse) Iter() <-chan Entry {
	ch := make(chan Entry)

	go func() {
		iter(this, ch)
		close(ch)
	}()

	return ch
}

func iter(data *APIResponse, ch chan Entry) {
	for _, entry := range data.Entries {
		ch <- entry
	}

	if data.NextCollectionLink != "" {
		next, err := fetchBugList(data.NextCollectionLink)
		if err == nil {
			iter(&next, ch)
		}
	}
}

type Entry struct {
	Title    string
	URL      string
	Assignee string
	Status string
}

func GetBugs(date time.Time) (APIResponse, error) {
	url, _ := buildURL(date.UTC().String())
	return fetchBugList(url)
}

func fetchBugList(url string) (apiResp APIResponse, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	var dataMap map[string]interface{}

	err = json.Unmarshal(responseBody, &dataMap)
	if err != nil {
		return
	}

	apiResp, err = parseResponse(dataMap)
	return
}

func getTime() string {
	date := time.Date(2012, time.January, 1, 0, 0, 0, 0, time.UTC)

	return date.String()
}

func buildURL(modifiedSince string) (uri string, err error) {
	baseUrl, err := url.Parse("https://api.launchpad.net/1.0/elementary")
	if err != nil {
		return
	}

	params := url.Values{}
	params.Add("ws.op", "searchTasks")
	params.Add("modified_since", modifiedSince)

	baseUrl.RawQuery = params.Encode()

	uri = baseUrl.String()
	return
}

func parseResponse(dataMap map[string]interface{}) (apiResp APIResponse, err error) {

	if totalSize, ok := dataMap["total_size"]; ok {
		if apiResp.TotalSize, ok = totalSize.(float64); !ok {
			err = errors.New("Could not get `total_size` from response")
			return
		}
	}

	if start, ok := dataMap["start"]; ok {
		if apiResp.Start, ok = start.(float64); !ok {
			err = errors.New("Could not get `start` from response")
			return
		}
	}

	if nextColl, ok := dataMap["next_collection_link"]; ok {
		if apiResp.NextCollectionLink, ok = nextColl.(string); !ok {
			fmt.Println("Could not get `next_collection_link` from response")
		}
	} else if prevColl, ok := dataMap["prev_collection_link"]; ok {
		if apiResp.PrevCollectionLink, ok = prevColl.(string); !ok {
			fmt.Println("Could not get `prev_collection_link` from response")
		}
	}

	if entries, ok := dataMap["entries"]; ok {
		if entriesArray, ok := entries.([]interface{}); !ok {
			err = errors.New("Could not parse `entries` from response")
			return
		} else {

			apiResp.Entries, err = parseEntries(entriesArray)
			if err != nil {
				return
			}
		}

	} else {
		err = errors.New("Could not get `entries` from response")
		return
	}
	return
}

func parseEntries(entriesArray []interface{}) (entries []Entry, err error) {
	entries = make([]Entry, len(entriesArray))

	for i, rawEntry := range entriesArray {
		var entryMap map[string]interface{}
		entry := Entry{}
		ok := true

		if entryMap, ok = rawEntry.(map[string]interface{}); !ok {
			err = errors.New("Could not parse entry number " + strconv.Itoa(i))
			return
		}

		entry.Title, _ = entryMap["title"].(string)

		entry.URL, _ = entryMap["web_link"].(string)

		if entry.Assignee, ok = entryMap["assignee_link"].(string); !ok {
			entry.Assignee = "elementary Devs"
		} else {
			entry.Assignee = entry.Assignee[31:]
		}

		entry.Status, _ = entryMap["status"].(string)
		entries[i] = entry
	}
	return
}
