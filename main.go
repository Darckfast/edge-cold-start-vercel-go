package main

import (
	"main/types"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mailru/easyjson"
)

func main() {
	mux := httprouter.New()
	mux.HandlerFunc("POST", "/", func(w http.ResponseWriter, r *http.Request) {
		easyjson.MarshalToWriter(types.Payload{
			Time: time.Now().UnixMilli(),
		}, w)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, mux)
}
