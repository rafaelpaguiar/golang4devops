package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	ResponseOutput *http.Response
}

func (m MockClient) Get(url string) (resp *http.Response, err error) {
	return m.ResponseOutput, nil
}

func TestDoGetRequest(t *testing.T) {
	words := WordsPage{
		Page: Page{"Words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}

	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Error("marshal error: %s", err)
	}
	apiInstance := API{
		Options: Options{},
		Client: MockClient{
			ResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}

	response, err := apiInstance.DoGetRequest("http://localhost/words")
	if err != nil {
		t.Error("DoGetRequesterror: %s", err)
	}

	if response == nil {

		t.Fatalf("Response is empty")
	}
	if response.GetResponse() != strings.Join([]string{"a", "b"}, ", ") {
		t.Error("Unexected response: %s", response.GetResponse())
	}

}
