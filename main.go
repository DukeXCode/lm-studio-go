package lm

import (
	"fmt"
)

// Example usage
func main() {
	url := "http://localhost:1234/v1/chat/completions"
	model := "google/gemma-3-12b"
	messages := []Message{
		{Role: "system", Content: "Always answer in rhymes."},
		{Role: "user", Content: "How are you doing today?"},
	}
	temperature := 0.7

	response, err := SendChatCompletion(url, model, messages, temperature)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(response.Choices) > 0 {
		fmt.Println("Assistant:", response.Choices[0].Message.Content)
	} else {
		fmt.Println("No choices in response.")
	}
}
