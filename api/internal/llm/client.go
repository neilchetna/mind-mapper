package llm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	llmDTO "github.com/neilchetna/mind-mapper/internal/llm/dto"
)

type LLMClient struct {
}

func (c *LLMClient) MakeLLMCall(prompt string) (string, error) {
	httpClient := &http.Client{}
	model := os.Getenv("LLM_MODEL")
	payload := llmDTO.LLMRequest{
		Model:  model,
		Prompt: prompt,
	}
	llmUrl := os.Getenv("LLM_HTTP_URL")
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", llmUrl, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var fullResponse strings.Builder

	for decoder.More() {
		var chunk llmDTO.LLMResponse
		if err := decoder.Decode(&chunk); err != nil {
			return "", err
		}

		fullResponse.WriteString(chunk.Response)

		if chunk.Done {
			break
		}
	}

	return fullResponse.String(), nil
}
