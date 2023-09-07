package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//go:generate mockery --dir . --name OpenAiService --output ./mocks
type OpenAiService interface {
	GenerateResponse(model, prompt string) (string, error)
}

type Service struct {
	BaseURL string
	APIKey  string
}

func NewService(baseURL, apiKey string) *Service {
	return &Service{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (s *Service) GenerateResponse(model, prompt string) (string, error) {

	var attempt int

	resp, err := CheckRateLimit(s, attempt, model, prompt)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return prompt, err
	}

	if _, exists := data["error"]; exists {
		return prompt, fmt.Errorf(data["error"].(map[string]interface{})["message"].(string))
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return content, nil
}

func CheckRateLimit(s *Service, attempt int, model, prompt string) (*http.Response, error) {

	if attempt > 3 {
		return nil, fmt.Errorf("API request failed with status code: %d", http.StatusTooManyRequests)
	}
	attempt += 1

	requestBody := map[string]interface{}{
		"model":    model,
		"messages": []interface{}{map[string]interface{}{"role": "system", "content": prompt}},
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", s.BaseURL, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		time.Sleep(time.Minute)
		return CheckRateLimit(s, attempt, model, prompt)
	}

	return resp, nil
}
