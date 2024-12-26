package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var API_URL = GetEnv("API_URL", "http://localhost:8080")

func MakeAPIRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %v", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", API_URL, endpoint), reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %v", err)
	}
	defer res.Body.Close()

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return responseBytes, nil
}
