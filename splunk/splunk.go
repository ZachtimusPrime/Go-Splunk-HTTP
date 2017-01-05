package splunk

import (
	"crypto/tls"
	"time"
	"bytes"
	"os"
	"net/http"
	"encoding/json"
	"errors"
)

// Event represents the log event object that is sent to Splunk when *HTTPCollector.Log is called.
type Event struct {
	Time 		int64		`json:"time" binding:"required"`	// epoch time in seconds
	Host		string  	`json:"host" binding:"required"`	// hostname
	Source		string  	`json:"source" binding:"required"`	// app name
	SourceType	string 		`json:"sourcetype" binding:"required"`	// Splunk bucket to group logs in
	Index		string		`json:"index" binding:"required"`	// idk what it does..
	Event		map[string]string `json:"event" binding:"required"`	// throw any useful key/val pairs here
}

type Client struct {
	HTTPClient *http.Client	 // HTTP client used to communicate with the API
	URL string
	Token string
	Source string
	SourceType string
	Index string
}

func NewClient(httpClient *http.Client, URL string, Token string, Source string, SourceType string, Index string) (*Client) {
	// Create a new client
	if httpClient == nil {
		tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} // turn off certificate checking
		httpClient = &http.Client{Timeout: time.Second * 20, Transport: tr}
	}

	c := &Client{HTTPClient: httpClient, URL: URL, Token: Token, Source: Source, SourceType: SourceType, Index: Index}

	return c
}

func NewEvent(event map[string]string, source string, sourcetype string, index string) (Event) {
	hostname, _ := os.Hostname()
	e := Event{Time: time.Now().Unix(), Host: hostname, Source: source, SourceType: sourcetype, Index: index, Event: event}

	return e
}

func (c *Client) Log(event map[string]string) (error) {
	// create Splunk log
	log := NewEvent(event, c.Source, c.SourceType, c.Index)

	// Convert requestBody struct to byte slice to prep for http.NewRequest
	b, err := json.Marshal(log)
	if err != nil {
		return err
	}

	//log.Print(string(b[:])) // print what the splunk post body will be for checking/debugging

	// make new request
	url := c.URL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Splunk " + c.Token)

	// receive response
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	// If statusCode is not good, return error string
	switch res.StatusCode {
	case 200:
	default:
		// Turn response into string and return it
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		responseBody := buf.String()
		err = errors.New(responseBody)

	//log.Print(responseBody)	// print error to screen for checking/debugging
	}

	return err
}
