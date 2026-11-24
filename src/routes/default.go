package routes

import (
	"fmt"
	"net/http"
)

// DefaultRoute godoc
//
//	@Summary	Healthcheck simple
//	@Tags		default
//	@Produce	plain
//	@Success	200	{string}	string	"OK"
//	@Router		/ [get]
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
