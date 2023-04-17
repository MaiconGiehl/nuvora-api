package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func TravelCompanyMiddleware(h http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("AAAAAAAAAA")
		token := strings.Split(r.Header.Values("Authorization")[0], "Bearer ")[1]
		fmt.Print("BBBBBBBBBBBBBBB")
		permissionLevel, err := extractPermissionLevel(token)
		if err != nil {
			json.NewEncoder(w).Encode("invalid token")
			return 
		}
		if permissionLevel != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("wrong permission level")
			return
		}
		h.ServeHTTP(w, r)
	})
}