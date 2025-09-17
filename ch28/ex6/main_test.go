package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert2.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student) // 결과 변환
	assert.Nil(err)
	assert.Equal("aaa", student.Name)
	assert.Equal(16, student.Age)
	assert.Equal(87, student.Score)
}
