# Go-Splunk-HTTP
A simple and lightweight HTTP Splunk logging package for Go. Instantiates a logging connection object to your Splunk server and allows you to submit log events as desired. [Uses HTTP event collection on a Splunk server](http://docs.splunk.com/Documentation/Splunk/latest/Data/UsetheHTTPEventCollector).

[![GoDoc](https://godoc.org/github.com/ZachtimusPrime/Go-Splunk-HTTP/splunk?status.svg)](https://godoc.org/github.com/ZachtimusPrime/Go-Splunk-HTTP/splunk)
[![Build Status](https://travis-ci.org/ZachtimusPrime/Go-Splunk-HTTP.svg?branch=master)](https://travis-ci.org/ZachtimusPrime/Go-Splunk-HTTP) 
[![Coverage Status](https://coveralls.io/repos/github/ZachtimusPrime/Go-Splunk-HTTP/badge.svg?branch=master)](https://coveralls.io/github/ZachtimusPrime/Go-Splunk-HTTP?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ZachtimusPrime/Go-Splunk-HTTP)](https://goreportcard.com/report/github.com/ZachtimusPrime/Go-Splunk-HTTP) 

## Table of Contents ##

* [Installation](#installation)
* [Usage](#usage)

## Installation ##

```bash
go get "github.com/ZachtimusPrime/Go-Splunk-HTTP/splunk"
```

## Usage ##

Construct a new Splunk HTTP client, then send log events as desired. 

For example:

```go
package main

import (
        "github.com/ZachtimusPrime/Go-Splunk-HTTP/splunk"
)

func main() {

		// Create new Splunk client
		splunk := splunk.NewClient(
		        nil, 
				"https://{your-splunk-URL}:8088/services/collector",
				"{your-token}",
				"{your-source}",
				"{your-sourcetype}",
				"{your-index}"
		)
		
		// Send a log event with the client
		err := splunk.Log(interface{"msg": "send key/val pairs or json objects here", "msg2": "anything that is useful to you in the log event"})
		if err != nil {
        		return err
        }
}

```
