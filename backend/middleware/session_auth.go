package middleware

import (
	"backend/service"
	"net/http"
)


func SessionAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := service.GetSession(r)
        if err != nil || session.Values["user_id"] == nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}