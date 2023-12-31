// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OpenAiService is an autogenerated mock type for the OpenAiService type
type OpenAiService struct {
	mock.Mock
}

// GenerateResponse provides a mock function with given fields: model, prompt
func (_m *OpenAiService) GenerateResponse(model string, prompt string) (string, error) {
	ret := _m.Called(model, prompt)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(model, prompt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(model, prompt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOpenAiService interface {
	mock.TestingT
	Cleanup(func())
}

// NewOpenAiService creates a new instance of OpenAiService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOpenAiService(t mockConstructorTestingTNewOpenAiService) *OpenAiService {
	mock := &OpenAiService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
