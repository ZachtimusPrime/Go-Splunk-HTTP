# Go-Splunk-HTTP
A simple and lightweight HTTP Splunk logging package for Go. Instantiates a logging connection object to your Splunk server and allows you to submit log events as desired. [Uses HTTP event collection on a Splunk server](http://docs.splunk.com/Documentation/Splunk/latest/Data/UsetheHTTPEventCollector).

[![GoDoc](https://godoc.org/github.com/ZachtimusPrime/Go-Splunk-HTTP?status.svg)](https://godoc.org/github.com/ZachtimusPrime/Go-Splunk-HTTP) 
[![Build Status](https://travis-ci.org/ZachtimusPrime/Go-Splunk-HTTP.svg?branch=master)](https://travis-ci.org/ZachtimusPrime/Go-Splunk-HTTP) 
[![Coverage Status](https://coveralls.io/repos/github/ZachtimusPrime/Go-Splunk-HTTP/badge.svg?branch=master)](https://coveralls.io/github/ZachtimusPrime/Go-Splunk-HTTP?branch=master) 
[![Go Report Card](https://goreportcard.com/badge/github.com/ZachtimusPrime/Go-Splunk-HTTP)](https://goreportcard.com/report/github.com/ZachtimusPrime/Go-Splunk-HTTP) 

## Table of Contents ##

* [Installation](#installation)
* [Usage](#usage)

## Installation ##

```bash
go get "github.com/ZachtimusPrime/Go-Splunk-HTTP"
```

## Usage ##

Construct a new Splunk HTTP connection, then send log events as desired. 

For example:

```go
package main

import (
        "github.com/ZachtimusPrime/Go-Splunk-HTTP"
)

func main() {

		// Create new connection to splunk
		sl := splunk.HTTPCollector{
				Url: "https://{splunk-URL}:8088/services/collector",
				Token: "{your-token}",
				Source: "{your-source}",
				SourceType: "{your-sourcetype}",
		}
		
		// Send a log event
		sl.Log(map[string]string{"msg": "send key/val pairs here", "msg2": "anything that is useful to you in the log event"})

```
