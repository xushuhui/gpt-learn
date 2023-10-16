package lesson2

import (
	"context"
	"fmt"
	"testing"

	"github.com/sashabaranov/go-openai"
)

// 用 GPT 写一段用户评语判断的程序，输入用户评论，输出评论是正向还是反向的，分别用 Y和 N 表示
func TestSentiment(t *testing.T) {
	response, err := Chat("这件衣服我爱的不行")
	if err != nil {
		t.Errorf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response.Choices[0].Message.Content)
	// Y

	response, err = Chat("不理不睬")
	if err != nil {
		t.Errorf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(response.Choices[0].Message.Content)
	// N
	response, err = Chat("用餐环境差，等待时间长")
	if err != nil {
		t.Errorf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(response.Choices[0].Message.Content)
	// N
}

func Chat(userprompt string) (response openai.ChatCompletionResponse, err error) {
	client := openai.NewClient("your openai key")
	prompt := "I input user comments and you output  by Y and N if the comments are positive or negative "
	response, err = client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userprompt,
				},
			},
			Temperature: 0,
			MaxTokens:   500,
		},
	)
	if err != nil {
		return response, err
	}
	return response, nil
}
