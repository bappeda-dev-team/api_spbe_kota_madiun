package main

import (
        "encoding/json"
        "log"
        "net/http"
)

func main() {
        addr := ":3001"
        mux := http.NewServeMux()

        mux.HandleFunc("/",
        func(w http.ResponseWriter, r *http.Request) {
                enc := json.NewEncoder(w)
                w.Header().
                Set("Content-Type",
                "application/json; charset=utf-8")

                resp := Resp{
                        Message: "SPBE KOTA MADIUN",
                }

                if err := enc.Encode(resp); err != nil {
                        panic(err)
                }
        })

        log.Printf("listening on %s\n", addr)

        log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
        Message string `json:"message"`
}

