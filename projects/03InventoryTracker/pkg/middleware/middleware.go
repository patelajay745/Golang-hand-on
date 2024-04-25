package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/patelajay745/projects/03InventoryTracker/pkg/utils"
)

func CheckUserRole(handler http.HandlerFunc, allowedRoles ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)
		if !ok {
			http.Error(w, "Role not found in context", http.StatusInternalServerError)
			return
		}

		allowed := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				allowed = true
				break
			}
		}

		if !allowed {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		handler(w, r)
	}
}

// func CheckUserRole(next http.Handler, allowedRoles ...string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		role, ok := r.Context().Value("role").(string)
// 		if !ok {
// 			http.Error(w, "Role not found in context", http.StatusInternalServerError)
// 			return
// 		}

// 		allowed := false
// 		for _, allowedRole := range allowedRoles {
// 			if role == allowedRole {
// 				allowed = true
// 				break
// 			}
// 		}

// 		if !allowed {
// 			http.Error(w, "Unauthorized", http.StatusForbidden)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			//utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header is required")
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.Id)
		ctx = context.WithValue(ctx, "role", claims.Role)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request details
		log.Printf(
			"Method: %s, Path: %s, Time: %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
