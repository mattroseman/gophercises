package story

import (
	"io"
	"io/ioutil"
	"testing"
)

var testStoryPath = "testStory.json"

var arcNameTestCases = []struct {
	arcName       string
	expectedTitle string
}{
	{"intro", "The Little Blue Gopher"},
	{"new-york", "Visiting New York"},
	{"debate", "The Great Debate"},
	{"sean-kelly", "Exit Stage Left"},
	{"mark-bates", "Costume Time"},
	{"denver", "Hockey and Ski Slopes"},
	{"home", "Home Sweet Home"},
}

func TestReadStory(t *testing.T) {
	jsonBytes, err := ioutil.ReadFile(testStoryPath)
	if err != nil && err != io.EOF {
		panic(err)
	}

	testStory, err := ReadStory(jsonBytes)
	if err != nil {
		panic(err)
	}

	for _, testCase := range arcNameTestCases {
		if arc, ok := testStory[testCase.arcName]; ok {
			if arc.Title != testCase.expectedTitle {
				t.Fatalf("arc %s has incorrect title. got %s, want %s",
					testCase.arcName, arc.Title, testCase.expectedTitle)
			}
		} else {
			t.Fatalf("arc %s was expected but was not found", testCase.arcName)
		}
	}
}
