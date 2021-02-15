package service

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/Jeffail/gabs/v2"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("<h1>Welcome to my world!</h1>"))
}

func JSON2CSVHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    request, err := gabs.ParseJSON(b)
    fmt.Println(request)
    if request != nil {
        w.Write([]byte("<h1>Convert JSON to CSV !</h1>"))
        return
    }
    http.Error(w, "Bad request - Go away!", 400)
}
