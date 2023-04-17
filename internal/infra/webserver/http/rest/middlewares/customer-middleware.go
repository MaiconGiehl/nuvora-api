package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
)

func CustomerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Values("Authorization")[0], "Bearer ")[1]
		permissionLevel, err := extractPermissionLevel(token)
		if err != nil {
			json.NewEncoder(w).Encode("invalid token")
			return
		}
		if permissionLevel != 3 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("wrong permission level")
			return
		}
		h.ServeHTTP(w, r)
	})
}
