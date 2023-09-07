package financialdata

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
)

//go:generate mockery --dir . --name FinancialDataService --output ./mocks
type FinancialDataService interface {
	FetchCompanyData(companyCode string) (*FinancialData, error)
	GenerateText(data *FinancialData) string
}

type Service struct {
	APIURL    string
	Validator *validator.Validate
}

func NewService(apiURL string) *Service {
	return &Service{
		APIURL:    apiURL,
		Validator: validator.New(),
	}
}

type FinancialData struct {
	ChartDataFinancials ChartDataFinancials `json:"chartDataFinancials"`
	BalanceSheet        BalanceSheet        `json:"balanceSheet"`
}

type ChartDataFinancials struct {
	CompanyName string `json:"company_name" validate:"required"`
	Ticker      string `json:"ticker" validate:"required"`
}

type BalanceSheet struct {
	AverageYearlyAssetGrowth    string `json:"average_yearly_asset_growth" validate:"required"`
	QuarterOverQuarterChange    string `json:"quarter_over_quarter_asset_change_percent" validate:"required"`
	AssetsGrowthRate            string `json:"assets_growth_rate" validate:"required"`
	TotalCurrentAssets          string `json:"total_current_assets" validate:"required"`
	CashAndShortTermInvestments string `json:"cash_and_short_term_investments" validate:"required"`
	CashInvestmentsGrowth       string `json:"cash_and_short_term_investments_growth" validate:"required"`
}

func (s *Service) FetchCompanyData(companyCode string) (*FinancialData, error) {
	url := fmt.Sprintf(s.APIURL, companyCode)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var financialData FinancialData
	if err = json.Unmarshal(body, &financialData); err != nil {
		return nil, err
	}

	if err = s.Validator.Struct(&financialData); err != nil {
		return nil, err
	}

	return &financialData, nil
}

func (s *Service) GenerateText(data *FinancialData) string {
	template := `
Diving into the fiscal trajectory of %s, we observe an average asset growth. This rate, interestingly, stands at %s, reflecting both the company's highs and lows. When compared quarter-over-quarter, this figure adjusts to %s. A look back at the past year reveals a total asset change of %s.

In the realm of current assets, %s clocks in at %s in the reporting currency. A significant portion of these assets, precisely %s, is held in cash and short-term investments. This segment shows a change of %s%% when juxtaposed with last year's data.
`

	return fmt.Sprintf(template, data.ChartDataFinancials.CompanyName, data.BalanceSheet.AverageYearlyAssetGrowth, data.BalanceSheet.QuarterOverQuarterChange, data.BalanceSheet.AssetsGrowthRate, data.ChartDataFinancials.Ticker, data.BalanceSheet.TotalCurrentAssets, data.BalanceSheet.CashAndShortTermInvestments, data.BalanceSheet.CashInvestmentsGrowth)
}
