package main

import (
	"main/types"
	"net/http"
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

	http.ListenAndServe(":8080", mux)
}
