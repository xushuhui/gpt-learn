package lesson3

import (
	"context"
	"fmt"
	"testing"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/prompts"
)

func TestLlmPromptTemplate(t *testing.T) {
	prompt := prompts.NewPromptTemplate(
		"What is a good name for a company that makes {{.product}}? And only return the best one",
		[]string{"product"},
	)
	chain := chains.NewLLMChain(llm, prompt)

	result, err := chains.Run(context.Background(), chain,
		map[string]any{
			"product": "colorful socks",
		})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
	// ChromaSock
}

func TestLlmPromptTemplateApply(t *testing.T) {
	prompt := prompts.NewPromptTemplate(
		"What is a good name for a company that makes {{.product}}? And only return the best one",
		[]string{"product"},
	)
	chain := chains.NewLLMChain(llm, prompt)
	// 并发调用，注意控制goroutine数量，openai call limit
	result, err := chains.Apply(context.Background(), chain,
		[]map[string]any{
			{"product": "colorful socks"},
			{"product": "cloudnative devops platform"},
			{"product": "Noise cancellation headphone"},
		}, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
	// [map[text:Socks in Color] map[text:CloudCraft] map[text:SilenceSync]]
}

func TestLLMRequest(t *testing.T) {
	result, err := QueryBaidu("今天北京天气？")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

func QueryBaidu(question string) (string, error) {
	return "", nil
}
//FIXME:
func TestSummaryChain(t *testing.T) {
	summarizingPrompt := prompts.NewPromptTemplate(`
Summarize the following content into a sentence less than 20 words:
---
{{.content}}
`, []string{"content"})
	summarizingChain := chains.NewLLMChain(llm, summarizingPrompt)
	summarizingChain.OutputKey = "summary"
	translatingPrompt := prompts.NewPromptTemplate(`translate "{{.summary}}" into Chinese:`, []string{"summary"})
	translatingChain := chains.NewLLMChain(llm, translatingPrompt)
	translatingChain.OutputKey = "translated"
	chs := []chains.Chain{summarizingChain, translatingChain}
	overallChain, err := chains.NewSequentialChain(chs, []string{"content"},
		[]string{"summary", "translated"})
	if err != nil {
		t.Error(err)
		return
	}
	result,err := overallChain.Call(context.Background(),map[string]any{
		"content":`LangChain is a framework for developing applications powered by language models. It enables applications that are:

		Data-aware: connect a language model to other sources of data
		Agentic: allow a language model to interact with its environment
		The main value props of LangChain are:
		
		Components: abstractions for working with language models, along with a collection of implementations for each abstraction. Components are modular and easy-to-use, whether you are using the rest of the LangChain framework or not
		Off-the-shelf chains: a structured assembly of components for accomplishing specific higher-level tasks
		Off-the-shelf chains make it easy to get started. For more complex applications and nuanced use-cases, components make it easy to customize existing chains or build new ones.`,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("summary: ",result["summary"])
	fmt.Println("translated: ",result["translated"])
	//translated:  LangChain 是一个框架，可用于使用语言模型开发数据驱动和交互式应用。它提供模块化组件和预构建的链条，方便定制和实施
}
