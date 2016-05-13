package line

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

func CallLineApi(endpoint, serviceUri string, accessToken string, req , rsp interface{}) error {
	url := endpoint + serviceUri

	headers := make(http.Header, 0)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}
	headers["X-LINE-ChannelToken"] = []string{accessToken}

	payload, _ := json.MarshalIndent(req, "", "  ")

	response, err := PostContent(url, headers, payload)


	if response != nil {
		json.Unmarshal(response, rsp)
	}

	if err != nil {
		return fmt.Errorf("line_api: Unable to call '%v' API. %v", url, err.Error())
	}

	return nil
}


func PostContent(url string, headers http.Header, payload []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	setHttpHeaders(req, headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return body, HttpError{Status: resp.Status, StatusCode: resp.StatusCode}
	}

	return body, err
}


func setHttpHeaders(request *http.Request, headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			request.Header.Set(key, value)
		}
	}
}

type HttpError struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
}

func (err HttpError) Error() string {
	return fmt.Sprintf("serve_http: %v code: %v", err.Status, err.StatusCode)
}