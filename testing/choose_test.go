package testing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEndpoint(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1/choose/1/0")
	if err != nil {
		t.Error("Choose endpoint is not accessible. :(")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Unable to read body of the response. :(")
	}
	fmt.Println(string(body))
}
