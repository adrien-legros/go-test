package greeting

import (
	"net/http"
	"testing"
	"io/ioutil"
)

func TestHttp(t *testing.T){
	resp, err := http.Get("http://localhost:8090/hello")
	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                t.Error("expected", nil, "got", err.Error())
        }
	bodyString := string(bodyBytes)
	if bodyString != "hello\n" {
		t.Error("expected status success got:", bodyString)
	}
}
