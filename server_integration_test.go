package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordingWinsAndRetrievingThemInMemory(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func TestRecordingWinsAndRetrievingThemInJson(t *testing.T) {
	store := NewJsonPlayerStoreTest()
	server := PlayerServer{store}
	player := "Pepper"
	attemptRequests := 5

	beforeAttempts := httptest.NewRecorder()
	server.ServeHTTP(beforeAttempts, newGetScoreRequest(player))
	lastScore, _ := strconv.Atoi(beforeAttempts.Body.String())

	for i := 0; i < attemptRequests; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	afterAttempts := httptest.NewRecorder()
	server.ServeHTTP(afterAttempts, newGetScoreRequest(player))
	assertStatus(t, afterAttempts.Code, http.StatusOK)

	assertResponseBody(t, afterAttempts.Body.String(), fmt.Sprint(lastScore+attemptRequests))
}

func NewJsonPlayerStoreTest() PlayerStore {
	return &jsonPlayerStore{[]player{}, "players_test.json"}
}
