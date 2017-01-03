package splunk

import (
	"crypto/tls"
	"time"
	"bytes"
	"os"
	"net/http"
	"encoding/json"
)

type logevent struct {
	Time 		int64		`json:"time" binding:"required"`	// epoch time in seconds
	Host		string  	`json:"host" binding:"required"`	// hostname
	Source		string  	`json:"source" binding:"required"`	// app name
	SourceType	string 		`json:"sourcetype" binding:"required"`	// Splunk bucket to group logs in
	Index		string		`json:"index" binding:"required"`	// idk what it does..
	Event		map[string]string `json:"event" binding:"required"`	// throw any useful key/val pairs here
}

type SplunkLogger struct {
	Url		string		`json:"url" binding:"required"`
	Token		string		`json:"token" binding:"required"`
	Source 		string		`json:"source" binding:"required"`
	SourceType 	string		`json:"sourcetype" binding:"required"`
	Index		string		`json:"index" binding:"required"`
}

func (sl *SplunkLogger) Log(event map[string]string) (err error){
	hostname, _ := os.Hostname()
	// create Splunk log
	splunklog := logevent{
		Time: time.Now().Unix(),
		Host: hostname,
		Source: sl.Source,
		SourceType: sl.SourceType,
		Index: sl.Index,
		Event: event,
	}

	// Convert requestBody struct to byte slice to prep for http.NewRequest
	b, err := json.Marshal(splunklog)
	if err != nil {
		return err
	}

	//log.Print(string(b[:])) // print what the splunk post body will be for checking/debugging

	// make new request
	url := sl.Url  // Splunk url
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Add(content_type, application_json)
	req.Header.Add("Authorization", "Splunk " + sl.Token)
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} // turn off certificate checking
	client := &http.Client{Transport: tr}

	// receive response
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// If statusCode is not good, return error string
	switch res.StatusCode {
	case 200:
	default:
		// Turn response in to string. I'm not smart enough to use io.Readers effectively TODO: Clean the below logic up
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		responseBody := buf.String()
		err = responseBody
		//log.Print(responseBody)	// print error to screen for checking/debugging
	}
	return err
}