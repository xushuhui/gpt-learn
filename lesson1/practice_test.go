package lesson1

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cast"
)

// 1. 写一句中文，通过 API 让 GPT 输出英文译文
// 2. 分别计算问题 1 中英语句子和对应中文译文的token数
func TestTranslate(t *testing.T) {
	client := openai.NewClient("your openai key")
	chinese := "在未来还没有到来的时候，总要有人把它创造出来，那个人应该是我们。"
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a translator,I input a Chinese sentence, and you output English",
				},

				{
					Role:    openai.ChatMessageRoleUser,
					Content: chinese,
				},
			},
			Temperature: 0,
			MaxTokens:   200,
		},
	)
	if err != nil {
		t.Errorf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response.Choices[0].Message.Content)
	// When the future has not yet arrived, someone must create it, and that person should be us.
	english := response.Choices[0].Message.Content

	encoding, err := tiktoken.EncodingForModel("gpt-3.5-turbo")
	if err != nil {
		t.Errorf("getEncoding: %v", err)
		return
	}
	cnTokens := encoding.Encode(chinese, nil, nil)
	enTokens := encoding.Encode(english, nil, nil)
	fmt.Println("chinese: " + chinese + " ; " + cast.ToString(len(cnTokens)) + " tokens")
	fmt.Println("english: " + english + " ; " + cast.ToString(len(enTokens)) + " tokens")
	// chinese: 在未来还没有到来的时候，总要有人把它创造出来，那个人应该是我们。 ; 35 tokens
	// english: When the future has not yet arrived, someone has to create it, and that person should be us. ; 21 tokens
}
