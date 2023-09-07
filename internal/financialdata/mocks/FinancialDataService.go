// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	financialdata "github.com/DmytroKha/numfin/internal/financialdata"

	mock "github.com/stretchr/testify/mock"
)

// FinancialDataService is an autogenerated mock type for the FinancialDataService type
type FinancialDataService struct {
	mock.Mock
}

// FetchCompanyData provides a mock function with given fields: companyCode
func (_m *FinancialDataService) FetchCompanyData(companyCode string) (*financialdata.FinancialData, error) {
	ret := _m.Called(companyCode)

	var r0 *financialdata.FinancialData
	if rf, ok := ret.Get(0).(func(string) *financialdata.FinancialData); ok {
		r0 = rf(companyCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*financialdata.FinancialData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(companyCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateText provides a mock function with given fields: data
func (_m *FinancialDataService) GenerateText(data *financialdata.FinancialData) string {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func(*financialdata.FinancialData) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewFinancialDataService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFinancialDataService creates a new instance of FinancialDataService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFinancialDataService(t mockConstructorTestingTNewFinancialDataService) *FinancialDataService {
	mock := &FinancialDataService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
