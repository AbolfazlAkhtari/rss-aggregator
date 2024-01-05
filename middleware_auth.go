package main

import (
	"fmt"
	"gituhb.com/AbolfazlAkhtari/RSS/internal/auth"
	"gituhb.com/AbolfazlAkhtari/RSS/internal/database"
	"net/http"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apikey, err := auth.GetApiKey(request.Header)
		if err != nil {
			respondWithError(writer, 400, fmt.Sprintf("Auth err: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(request.Context(), apikey)
		if err != nil {
			respondWithError(writer, 404, fmt.Sprintf("Couldn't get User: %v", err))
			return
		}

		handler(writer, request, user)
	}
}
