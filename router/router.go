package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/qimpl/housing/models"
	"github.com/qimpl/housing/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

var unProtectedRoutes = []string{"/swagger/"}

// CreateRouter create housing API routes
func CreateRouter() {
	router := mux.NewRouter()
	APIRouter := router.PathPrefix("/api/v1").Subrouter()

	APIRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Use(jwtVerifyTokenMiddleware)

	createHousingRouter(APIRouter)
	createHousingTypeRouter(APIRouter)
	createHousingStatusRouter(APIRouter)
	createVisitRouter(APIRouter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	handler := c.Handler(router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		log.Print(err)
	}
}

func jwtVerifyTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestURI := strings.TrimPrefix(r.RequestURI, "/api/v1")
		for _, route := range unProtectedRoutes {
			if strings.Contains(requestURI, route) {
				next.ServeHTTP(w, r)
				return
			}
		}

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
		if len(authHeader) != 2 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			var badRequest *models.BadRequest
			json.NewEncoder(w).Encode(badRequest.GetError("Malformed Authorization HTTP header"))

			return
		}

		if _, err := services.ValidateJwtToken(authHeader[1]); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			var Unauthorized *models.Unauthorized
			json.NewEncoder(w).Encode(Unauthorized.GetError(fmt.Sprintf("Invalid jwt token : %s", err)))

			return
		}

		next.ServeHTTP(w, r)
	})
}
