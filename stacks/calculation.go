package stacks

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

// CalculateBestHost returns the host which has the best latency
func (stack *Stack) CalculateBestHost() (string, error) {
	// Define the host with the best latency
	bestHostLatency := int64(-1)
	bestHost := ""
	for _, host := range stack.Hosts {
		latency := pingHost(host)
		if latency == -1 {
			continue
		}
		if latency < bestHostLatency || bestHostLatency == -1 {
			bestHostLatency = latency
			bestHost = host
		}
	}

	// Return the host with the best latency
	if bestHost == "" {
		return "", ErrAllHostsOffline
	}
	return bestHost, nil
}

// pingHost returns the latency to a HTTP(s) host in ms or -1 if the host is offline
func pingHost(host string) int64 {
	// Create a request to the HTTP(s) host
	request, err := http.NewRequest("GET", host, nil)
	if err != nil {
		return -1
	}

	// Inject the go-httpstat context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(request.Context(), &result)
	request = request.WithContext(ctx)

	// Execute the request
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return -1
	}
	_, err = io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return -1
	}
	response.Body.Close()
	result.End(time.Now())

	// Return the latency
	return result.Connect.Milliseconds()
}
