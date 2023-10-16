package lesson2

import (
	"context"
	"fmt"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestRecognizeIntent(t *testing.T) {
	client := openai.NewClient("your openai key")
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Recognize the intent from the user's input",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "提醒我明早8点有会议",
				},
			},
			Temperature: 0,
			MaxTokens:   500,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response.Choices[0].Message.Content)
	//intent: set_reminder
	//entities:
	//- reminder_type: meeting
	//- date: tomorrow
	//- time: 8:00 AM
}

func TestGenerateSQL(t *testing.T) {
	client := openai.NewClient("your openai key")
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `You are a software engineer, you can anwser the user request based on the given tables:
					table “students“ with the columns [id, name, course_id, score] 
					table "courses" with the columns [id, name]`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "计算所有学生英语课程的平均成绩",
				},
			},
			Temperature: 0,
			MaxTokens:   500,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response.Choices[0].Message.Content)
	/*
		To calculate the average score for all students in the English course, you need to join the tables "students" and "courses" based on the course_id column. Then you can filter the rows where the course name is "English" and calculate the average score.

		Here is an example SQL query to achieve this:

		```
		SELECT AVG(score) AS average_score
		FROM students s
		JOIN courses c ON s.course_id = c.id
		WHERE c.name = 'English'
		```

		This query will return the average score for all students who are enrolled in the English course.
	*/
}

func TestRecognizeIntentOutputJson(t *testing.T) {
	client := openai.NewClient("your openai key")
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `Recognize the intent from the user's input and format output as JSON string. 
					The output JSON string includes: "intention", "paramters"`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "提醒我明早8点有会议",
				},
			},
			Temperature: 0,
			MaxTokens:   500,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response.Choices[0].Message.Content)
	/*
			{
		  		"intention": "reminder",
		  		"parameters": {
		    		"time": "明早8点",
		    		"title": "会议"
		  		}
			}
	*/
}

func TestGenerateSQLOutputSQL(t *testing.T) {
	client := openai.NewClient("your openai key")
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `You are a software engineer, you can write a SQL string as the anwser according to the user request 
					The user's requirement is based on the given tables:
					   table “students“ with the columns [id, name, course_id, score];
					   table "courses" with the columns [id, name].`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "列出英语课程成绩大于80分的学生",
				},
			},
			Temperature: 0,
			MaxTokens:   500,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response.Choices[0].Message.Content)
	/*
		SELECT students.name
		FROM students
		JOIN courses ON students.course_id = courses.id
		WHERE courses.name = '英语' AND students.score > 80;
	*/
}

func TestThrowErrorOutput(t *testing.T) {
	client := openai.NewClient("your openai key")
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `You are a software engineer, you can write a SQL string as the anwser according to the user request.
					Also, when you cannot create the SQL query for the user's request based on the given tables, please, only return "invalid request"
								   The user's requirement is based on the given tables:
									  table “students“ with the columns [id, name, course_id, score];
									  table "courses" with the columns [id, name].`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "列出年龄大于13的学生",
				},
			},
			Temperature: 0,
			MaxTokens:   500,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response.Choices[0].Message.Content)
	// invalid request
}
