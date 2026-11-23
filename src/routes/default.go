package routes

import (
	"fmt"
	"net/http"
)

func DefaultRoute() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "OK")
	})

}
