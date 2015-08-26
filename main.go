package main

import (
    "log"
    "net/http"
    "encoding/json"
    "fmt"
)

func main() {
    resp, err := http.Get("http://reddit.com/r/golang.json")
    if err != nil {
        log.Fatal(err)
    }
    if resp.StatusCode != http.StatusOK {
        log.Fatal(resp.Status)
    }
    if err != nil {
        log.Fatal(err)
    }

    type Item struct {
        Title string
        URL   string
    }

    type Response struct {
        Data struct {
            Children []struct {
                Data Item
            }
        }
    }

    r := new(Response)
    err = json.NewDecoder(resp.Body).Decode(r)

    for _, child := range r.Data.Children {
        fmt.Println(child.Data.Title)
    }

}


