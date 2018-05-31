package routing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/johananl/gophercon-eu-workshop/pkg/routing"
)

func TestBaseRouter(t *testing.T) {
	handler := routing.BaseRouter()

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}
}
