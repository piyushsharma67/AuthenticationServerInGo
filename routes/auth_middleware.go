package routes

import (
	"chat-application/utils"
	"context"
	"net/http"
)

func AuthMiddleWare(next http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check for authentication here (e.g., check a session or token)
        // For demonstration, let's assume a simple check for a header
		authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
		token:=authHeader[7:]

		isValid,claims,err:=utils.ValidateToken(token)

		if(!isValid){
			http.Error(w,err,http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", claims)

        next(w, r.WithContext(ctx))
    })
}