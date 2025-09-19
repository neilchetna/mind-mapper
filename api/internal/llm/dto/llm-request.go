package dto

type LLMRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}
