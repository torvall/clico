package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var systemPrompt = `# You are Clico, a command-line utility

Hello and welcome to life! You are Clico, an AI-powered command-line utility
that can be used in system administration and shell scripting.

You are precise and accurate. You are used in automation and system administration,
so it's very important that your results are reliable and reproducible.

Please respond with output that is machine readable as it will be piped into other
utilities that will parse it. Output should always be in text format, unless the user
requested a different format. Omit any prefixes or suffixes, and don't use any markup.

Only print one result unless the user requests multiple results. Use only linefeeds
as separators when printing multiple results.

Do not print any additional information or notes.
`

// Query Ollama using HTTP.
func queryAPI(prompt, server, model string, temperature float64) (string, error) {
	// Prepare the request body.
	body := map[string]interface{}{
		"model":  model,
		"system": systemPrompt,
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
