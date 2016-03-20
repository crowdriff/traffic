# Traffic [![Build Status](https://travis-ci.org/crowdriff/traffic.svg?branch=master)](https://travis-ci.org/crowdriff/traffic) [![Go Report Card](https://goreportcard.com/badge/github.com/crowdriff/traffic)](https://goreportcard.com/report/github.com/crowdriff/traffic)
===

A traffic pattern generator for all your Go testing/benchmarking needs.  
  
## Quick Start Examples

### Example 1: HTTP Server
Say I have an HTTP Server that I want to test and my simulated traffic pattern should be 25% `GET` requests to the `/hello` endpoint and 75% `GET` requests to the `/bye` endpoint.
  
```go
// Create a new traffic generator that will execute 1000 requests
gen := traffic.NewGenerator(1000)
// Add a traffic pattern that'll hit the /hello endpoint with
// 25% probability
gen.AddPattern(&traffic.Pattern{
	Probability: 25,
	Fn: func() {
		URL := fmt.Sprintf("http://%s/hello/world", serverURL)
		http.Get(URL)
	},
})
// Add a second traffic pattern that'll hit the /bye endpoint with
// 75% probability
gen.AddPattern(&traffic.Pattern{
	Probability: 75,
	Fn: func() {
		URL := fmt.Sprintf("http://%s/bye/world", serverURL)
		http.Get(URL)
	},
})
// Execute the traffic generator
gen.Execute()
```
  
This will execute 1000 requests sequentially and make an attempt to do a 25/75 split between `/hello` and `/bye`.