package service

import (
    "bytes"
    "encoding/json"
	"net/http"
	"testing"
	"net/http/httptest"
)

func TestJSON2CSVHandler(t *testing.T) {
    requestBody, err := json.Marshal(map[string]string{
        "request": "post",
        "url": "https://www.google.com",
    })
	req, err := http.NewRequest("POST", "/request", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(JSON2CSVHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
    // Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
