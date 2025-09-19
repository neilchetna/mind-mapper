package dto

type LLMResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}
