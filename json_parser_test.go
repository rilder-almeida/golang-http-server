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
	testFile := fstest.MapFS{
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

	expected, _ := testFile.ReadFile("players_test.json")
	if actualByte, _ := json.Marshal(actual); string(actualByte) != string(expected) {
		t.Errorf("\nExpected: %s\nActual: %s", string(expected), string(actualByte))
	}

}
