package service

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("<h1>Welcome to my world!</h1>"))
}

func JSON2CSVHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Convert JSON to CSV !</h1>"))
}
