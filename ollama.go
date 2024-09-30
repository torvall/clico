package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Query Ollama using HTTP.
func queryOllama(prompt string) string {
	// Prepare the request body.
	body := map[string]interface{}{
		"model":  "llama3.1",
		"prompt": prompt,
		"options": map[string]float64{
			"temperature": 0,
		},
		"stream": false,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal("Error marshalling JSON body: ", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal("Error creating request to Ollama server:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get a response back.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error sending request to Ollama server:", err)
	}
	defer resp.Body.Close()

	// Read the response body.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response from Ollama server:", err)
	}

	// Parse the response JSON body.
	var respData map[string]interface{}
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		log.Fatal("Error unmarshalling response from Ollama server:", err)
	}

	return respData["response"].(string)
}
