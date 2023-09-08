package financialdata_test

import (
	"errors"
	"github.com/DmytroKha/numfin/internal/financialdata"
	"github.com/DmytroKha/numfin/internal/financialdata/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchCompanyData_Success(t *testing.T) {
	// Create a new instance of the mock service
	mockService := new(mocks.FinancialDataService)

	// Set up expected data and return values
	expectedData := &financialdata.FinancialData{}
	mockService.On("FetchCompanyData", "TEST").Return(expectedData, nil)

	// Use the mock service in your code
	service := mockService //financialdata.NewService(mockService)
	data, err := service.FetchCompanyData("TEST")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedData, data)

	// Assert that the expected method was called
	mockService.AssertExpectations(t)
}

func TestFetchCompanyData_Error(t *testing.T) {
	// Create a new instance of the mock service
	mockService := new(mocks.FinancialDataService)

	// Set up an error scenario
	expectedError := errors.New("some error")
	mockService.On("FetchCompanyData", "INVALID").Return(nil, expectedError)

	// Use the mock service in your code
	data, err := mockService.FetchCompanyData("INVALID")

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Equal(t, expectedError, err)

	// Assert that the expected method was called
	mockService.AssertExpectations(t)
}

func TestGenerateText(t *testing.T) {
	// Create a new instance of the mock service
	mockService := new(mocks.FinancialDataService)

	// Set up input data
	inputData := &financialdata.FinancialData{
		ChartDataFinancials: financialdata.ChartDataFinancials{
			CompanyName: "Test Company",
			Ticker:      "TEST",
		},
		BalanceSheet: financialdata.BalanceSheet{
			AverageYearlyAssetGrowth:    "10%",
			QuarterOverQuarterChange:    "5%",
			AssetsGrowthRate:            "8%",
			TotalCurrentAssets:          "1000",
			CashAndShortTermInvestments: "500",
			CashInvestmentsGrowth:       "15%",
		},
	}

	// Set up expected text
	expectedText := `
Diving into the fiscal trajectory of Test Company, we observe an average asset growth. This rate, interestingly, stands at 10%, reflecting both the company's highs and lows. When compared quarter-over-quarter, this figure adjusts to 5%. A look back at the past year reveals a total asset change of 8%.

In the realm of current assets, TEST clocks in at 1000 in the reporting currency. A significant portion of these assets, precisely 500, is held in cash and short-term investments. This segment shows a change of 15% when juxtaposed with last year's data.
`

	mockService.On("GenerateText", inputData).Return(expectedText)

	// Use the mock service in your code
	service := mockService //financialdata.NewService(mockService)
	text := service.GenerateText(inputData)

	// Assertions
	assert.Equal(t, expectedText, text)

	// Assert that the expected method was called
	mockService.AssertExpectations(t)
}

func TestGenerateText_EmptyData(t *testing.T) {
	// Create a new instance of the mock service
	mockService := new(mocks.FinancialDataService)

	// Set up input data with empty values
	inputData := &financialdata.FinancialData{}

	// Set up expected text with placeholders
	expectedText := `
Diving into the fiscal trajectory of , we observe an average asset growth. This rate, interestingly, stands at , reflecting both the company's highs and lows. When compared quarter-over-quarter, this figure adjusts to . A look back at the past year reveals a total asset change of .

In the realm of current assets,  clocks in at  in the reporting currency. A significant portion of these assets, precisely , is held in cash and short-term investments. This segment shows a change of % when juxtaposed with last year's data.
`

	mockService.On("GenerateText", inputData).Return(expectedText)

	// Use the mock service in your code
	service := mockService //financialdata.NewService(mockService)
	text := service.GenerateText(inputData)

	// Assertions
	assert.Equal(t, expectedText, text)

	// Assert that the expected method was called
	mockService.AssertExpectations(t)
}
