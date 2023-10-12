package lesson1

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"

	"github.com/spf13/cast"
)

func TestHelloGpt(t *testing.T) {
	client := openai.NewClient("your openai key")
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "你是一个AI助理",
				},

				{
					Role:    openai.ChatMessageRoleUser,
					Content: "你好！你叫什么名字？",
				},
			},
			Temperature: 0.9,
			MaxTokens:   200,
			N:           3,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response.Choices[0].Message.Content)
	// 你好！我是一个AI助理，还没有名字。你可以为我取一个名字吗？
	for _, v := range response.Choices {
		fmt.Println(v.Message.Content)
	}
	// 你好！我是一个AI助理，所以没有一个具体的名字。你可以随意给我取一个名字，或者直接称呼我为助理。有什么我可以帮助你的吗？
	// 你好！我是AI助理，没有具体的名字。您可以随意称呼我为助理或者任何您喜欢的名字。有什么我可以帮助您的吗？
	// 你好！我是一个AI助理，没有具体的名字。你可以随意称呼我。有什么我可以帮助你的吗？
}

func TestEncodeToken(t *testing.T) {
	chinese := "在未来还没有到来的时候，总要有人把它创造出来，那个人应该是我们。"
	encoding, err := tiktoken.EncodingForModel("gpt-3.5-turbo")
	if err != nil {
		t.Errorf("getEncoding: %v", err)
		return
	}
	tokens := encoding.Encode(chinese, nil, nil)
	fmt.Println(tokens)
	//[19000 39442 37507 98806 81543 28037 37507 9554 13646 20022 247 3922 60843 31634 19361 17792 24326 232 8676 225 6701 249 67178 20834 37507 3922 45932 96 19483 17792 51611 76982 21043 98739 1811]

	num_of_tokens_in_chinese := len(encoding.Encode(chinese, nil, nil))
	fmt.Println("chinese: " + chinese + " ; " + cast.ToString(num_of_tokens_in_chinese) + " tokens")
	// chinese: 在未来还没有到来的时候，总要有人把它创造出来，那个人应该是我们。 ; 35 tokens
}
