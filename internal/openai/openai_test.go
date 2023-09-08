package openai_test

import (
	"errors"
	"testing"

	"github.com/DmytroKha/numfin/internal/openai/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGenerateResponse_Success(t *testing.T) {
	// Create a new instance of the mock service
	mockService := new(mocks.OpenAiService)

	// Set up expected data and return values
	model := "gpt-3.5-turbo"
	prompt := "Translate the following English text to French: 'Hello, world.'"
	expectedResponse := "Bonjour, le monde."

	mockService.On("GenerateResponse", model, prompt).Return(expectedResponse, nil)

	// Use the mock service in your code
	response, err := mockService.GenerateResponse(model, prompt)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)

	// Assert that the expected method was called
	mockService.AssertExpectations(t)
}

func TestGenerateResponse_Error(t *testing.T) {
	// Create a new instance of the mock service
	mockService := new(mocks.OpenAiService)

	// Set up an error scenario
	model := "gpt-3.5-turbo"
	prompt := "Translate the following English text to French: 'Hello, world.'"
	expectedError := errors.New("API request failed")

	mockService.On("GenerateResponse", model, prompt).Return("", expectedError)

	// Use the mock service in your code
	service := mockService
	response, err := service.GenerateResponse(model, prompt)

	// Assertions
	assert.Error(t, err)
	assert.Empty(t, response)
	assert.Equal(t, expectedError, err)

	// Assert that the expected method was called
	mockService.AssertExpectations(t)
}
