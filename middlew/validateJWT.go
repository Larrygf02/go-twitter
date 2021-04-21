package middlew

import "net/http"

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token"+err.Error(), 400)
			return
		}
		next.ServeHTTP(w, r)
	}
}
