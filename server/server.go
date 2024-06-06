package main

import (
	"fmt"
	"net/http"
)

func main() {
    // Serve static files from the dist directory
	fs := http.FileServer(http.Dir("./dist/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	fs = http.FileServer(http.Dir("./dist/images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

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
