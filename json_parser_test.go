package main

import (
	"encoding/json"
	"testing"
	"testing/fstest"
)

type StubPlayerFile struct {
	Store []StubPlayer `json:"store"`
}

var player1 = StubPlayer{"Pepper", 20}
var player2 = StubPlayer{"Floyd", 10}

var Scores = StubPlayerFile{
	[]StubPlayer{player1, player2},
}

func TestJsonFile(t *testing.T) {
	dataStore, _ := json.Marshal(Scores)

	expectedFile := fstest.MapFS{
		"players_test.json": {
			Data: dataStore,
		},
	}

	err := ToJsonFile("players_test.json", Scores)
	if err != nil {
		panic(err)
	}

	var actual StubPlayerFile
	err = FromJsonFile("players_test.json", &actual)
	if err != nil {
		panic(err)
	}

	expectedByte, _ := expectedFile.ReadFile("players_test.json")
	actualByte, _ := json.Marshal(actual)

	if string(actualByte) != string(expectedByte) {
		t.Errorf("\nExpected: %s\nActual: %s", string(expectedByte), string(actualByte))
	}
}
