package numfin

import (
	"fmt"
	"github.com/DmytroKha/numfin/internal/financialdata"
	"github.com/DmytroKha/numfin/internal/openai"
)

const (
	apiUrl    = "https://api.numfin.com/api/v2/financial-values/code?code=%s"
	openaiUrl = "https://api.openai.com/v1/chat/completions"
	apiKey    = "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"
	modelGPT  = "gpt-3.5-turbo"
	prompt    = "Add 4 sentences to this text giving analytics"
)

func MakeAnalyticsFromTheCompany(code string) (string, error) {

	financialDataService := financialdata.NewService(apiUrl)
	openaiService := openai.NewService(openaiUrl, apiKey)

	data, err := financialDataService.FetchCompanyData(code)
	if err != nil {
		return "", fmt.Errorf("Error fetching data for %s: %v\n", code, err)
	}

	text := financialDataService.GenerateText(data)

	gpt3Response, err := openaiService.GenerateResponse(modelGPT, text+"\n"+prompt)
	if err != nil {
		return "", fmt.Errorf("Error generating GPT-3 response for %s: %v\n", code, err)
	}

	return text + "\n" + gpt3Response, nil
}
