package main

import (
    "fmt"
    "flag"
    "strings"
)

var serverArg = flag.Bool("server", false, "start server on port 9119")

var setArg = flag.String("set", "", "set key:value pair")
var getArg = flag.String("get", "", "get value by key")
var getallArg = flag.Bool("getall", false, "get all key:value pairs")
var fileArg = flag.String("file", "/tmp/KvStore", "specify store file")

func main() {
    flag.Parse()

    store := NewFileStore(*fileArg)
    defer store.Close()

    if *serverArg {
        //TODO graceful shutdown
        RunServer(store)
        return
    }

    if len(*setArg) > 0 {
        kv := strings.SplitN(*setArg, ":", 2)
        if len(kv) == 2 {
            store.Set(kv[0], kv[1])
        }
    }

    if len(*getArg) > 0 {
        fmt.Println(store.Get(*getArg))
    }

    if *getallArg {
        fmt.Printf("%#v", store.GetAll())
    }
}
