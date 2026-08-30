package service

import (
	"net/url"
	"sync"

	"fizz-buzz.com/internal/model"
)

var (
	statsMu     sync.Mutex
	statsCounts = make(map[string]int)
	statsParams = make(map[string]map[string]string)
)

func RecordRequest(values url.Values) {
	params := make(map[string]string, len(values))
	for name := range values {
		params[name] = values.Get(name)
	}
	key := values.Encode()

	// if we want a persistent storage, we can use a database or a file to store the counts and params
	statsMu.Lock()
	defer statsMu.Unlock()
	statsCounts[key]++
	statsParams[key] = params
}

func MostUsedRequest() model.RequestStats {
	statsMu.Lock()
	defer statsMu.Unlock()

	var topKey string
	var topCount int
	for key, count := range statsCounts {
		if count > topCount {
			topKey = key
			topCount = count
		}
	}

	return model.RequestStats{
		Params: statsParams[topKey],
		Count:  topCount,
	}
}
