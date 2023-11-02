package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_atc/lib/utils"
	"io"
	"net/http"
)

type DeepInfra struct {
	apiKey  string
	apiBase string
}

type GenerationResult struct {
	Text string `json:"generated_text"`
}

type DeepInfraCompletionResult struct {
	RequestID string             `json:"request_id"`
	Results   []GenerationResult `json:"results"`
}

func NewDeepInfra(config utils.AppConfig) (DeepInfra, error) {
	return DeepInfra{
		apiKey:  config.DeepInfraApiKey,
		apiBase: "https://api.deepinfra.com/v1/inference/meta-llama/Llama-2-13b-chat-hf",
	}, nil
}

func (d *DeepInfra) Complete(text string) (DeepInfraCompletionResult, error) {
	post_body, err := json.Marshal(map[string]string{
		"input": text,
	})
	utils.HandleError(err)
	requestBody := bytes.NewBuffer(post_body)

	req, err := http.NewRequest("POST", d.apiBase, requestBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", d.apiKey))
	utils.HandleError(err)

	res, err := http.DefaultClient.Do(req)
	utils.HandleError(err)

	defer res.Body.Close()
	responseData, err := io.ReadAll(res.Body)
	utils.HandleError(err)

	var result DeepInfraCompletionResult
	if err := json.Unmarshal(responseData, &result); err != nil {
		utils.HandleError(err)
	}

	return result, nil
}
