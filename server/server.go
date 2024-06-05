package main

import (
	"fmt"
	"net/http"
)

func main() {

    cssFs := http.FileServer(http.Dir("./css"))
    http.Handle("/css/", cssFs)

    imagesFs := http.FileServer(http.Dir("./images"))
    http.Handle("/images/", imagesFs)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            w.Header().Set("Allow", http.MethodGet)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        http.ServeFile(w, r, "dist/index.html")
    })

    fmt.Println("Starting server on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed to start: ", err)
    }
}
