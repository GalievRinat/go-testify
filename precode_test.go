package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)

	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	status := responseRecorder.Code
	require.NotEmpty(t, status, http.StatusOK)

	require.NotEmpty(t, responseRecorder.Body)

	count := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, count, totalCount)
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	require.NotEmpty(t, status, http.StatusOK)

	require.NotEmpty(t, responseRecorder.Body)

	expected := `count missing`
	assert.Equal(t, responseRecorder.Body.String(), expected)
}

func TestMainHandlerWhenCityWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=orenburg", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	require.NotEmpty(t, status, http.StatusOK)

	require.NotEmpty(t, responseRecorder.Body)

	expected := `wrong city value`
	assert.Equal(t, responseRecorder.Body.String(), expected)
}

func TestMainHandlerWhenCountWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=AAA&city=moskow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	require.NotEmpty(t, status, http.StatusOK)

	require.NotEmpty(t, responseRecorder.Body)

	expected := `wrong count value`
	assert.Equal(t, responseRecorder.Body.String(), expected)
}
