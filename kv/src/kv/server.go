package main

import (
    "time"
    "strings"
    "log"
    "net/http"
    "fmt"
)

func RunServer(store Store) {
    log.Println("Starting server at :9119")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            key := strings.TrimPrefix(r.URL.Path, "/")
            if len(key) > 0 {
                value := store.Get(key)
                if len(value) > 0 {
                    fmt.Fprintf(w, "%#v", value)
                } else {
                    w.WriteHeader(http.StatusNotFound)
                }
            } else {
                fmt.Fprintf(w, "%#v", store.GetAll())
            }
        case "POST":
            key := r.FormValue("key")
            value := r.FormValue("value")
            store.Set(key, value)
        default:
            fmt.Fprintf(w, "Method not supported")
        }
    })

    ticker := time.NewTicker(time.Millisecond * 1000)
    go func() {
        for _ = range ticker.C {
            store.Sync()
        }
    }()

    panic(http.ListenAndServe(":9119", nil))
}
