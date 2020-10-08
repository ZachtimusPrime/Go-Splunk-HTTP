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

import "github.com/ZachtimusPrime/Go-Splunk-HTTP/splunk"

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
		
	// Use the client to send a log with the go host's current time
	err := splunk.Log(
		interface{"msg": "send key/val pairs or json objects here", "msg2": "anything that is useful to you in the log event"}
	)
	if err != nil {
        	return err
        }
	
	// Use the client to send a log with a provided timestamp
	err = splunk.LogWithTime(
		time.Now(),
		interface{"msg": "send key/val pairs or json objects here", "msg2": "anything that is useful to you in the log event"}
	)
	if err != nil {
		return err
	}
	
	// Use the client to send a batch of log events
	var events []splunk.Event
	events = append(
		events,
		splunk.NewEvent(
			interface{"msg": "event1"},
			"{desired-source}",
			"{desired-sourcetype}",
			"{desired-index}"
		)
	)
	events = append(
		events,
		splunk.NewEvent(
			interface{"msg": "event2"},
			"{desired-source}",
			"{desired-sourcetype}",
			"{desired-index}"
		)
	)
	err = splunk.LogEvents(events)
	if err != nil {
		return err
	}
}

```

## Splunk Writer  ##
To support logging libraries, and other output, we've added an asynchronous Writer. It supports retries, and different intervals for flushing messages & max log messages in its buffer

The easiest way to get access to the writer with an existing client is to do:

```go
writer := splunkClient.Writer()
```

This will give you an io.Writer you can use to direct output to splunk. However, since the io.Writer() is asynchronous, it will never return an error from its Write() function. To access errors generated from the Client,
Instantiate your Writer this way:

```go
splunk.Writer{
  Client: splunkClient
}
```
Since the type will now be splunk.Writer(), you can access the `Errors()` function, which returns a channel of errors. You can then spin up a goroutine to listen on this channel and report errors, or you can handle however you like. 

Optionally, you can add more configuration to the writer.

```go
splunk.Writer {
  Client: splunkClient,
  FlushInterval: 10 *time.Second, // How often we'll flush our buffer
  FlushThreshold: 25, // Max messages we'll keep in our buffer, regardless of FlushInterval
  MaxRetries: 2, // Number of times we'll retry a failed send
}
```

