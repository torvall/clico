package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Query Ollama using HTTP.
func queryAPI(prompt, server, model string, temperature float64) (string, error) {
	// Prepare the request body.
	body := map[string]interface{}{
		"model":  model,
		"prompt": prompt,
		"options": map[string]float64{
			"temperature": float64(temperature),
		},
		"stream": false,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON body: %s", err.Error())
	}

	req, err := http.NewRequest("POST", server+"/api/generate", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get a response back.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to server: %s", err.Error())
	}
	defer resp.Body.Close()

	// Read the response body.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %s", err.Error())
	}

	// Parse the response JSON body.
	var respData map[string]interface{}
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %s", err.Error())
	}

	return respData["response"].(string), nil
}
