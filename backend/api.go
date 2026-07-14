package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type APIServer struct {
	port      string
	externals map[string]any
}

// function type
type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string
}

// ----
// Initialize
func NewAPIServer(port string) *APIServer {
	return &APIServer{
		port:      port,
		externals: make(map[string]any),
	}
}

// Add external API
func (apiServer *APIServer) AddExternalAPI(name string, address any) {
	apiServer.externals[name] = address
}

// Run the server
func (apiServer *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/api/filter/{city}", makeHTTPHandleFunc(apiServer.FilterCity)).Methods("GET")
	router.HandleFunc("/api/search/{city}", makeHTTPHandleFunc(apiServer.SearchCity)).Methods("GET")
	router.HandleFunc("/api/forecast/{city}", makeHTTPHandleFunc(apiServer.Forecast)).Methods("GET")

	//fs := http.FileServer(http.Dir("../frontend/dist"))
	router.PathPrefix("/").Handler(spaHandler("../frontend/dist"))

	fmt.Println("Server starting on port:", apiServer.port[1:])
	log.Fatal(http.ListenAndServe(apiServer.port, router))
}

// Wrapper to simplify the error response handling
func makeHTTPHandleFunc(f apiFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		EnableCORS(&w)

		if err := f(w, r); err != nil {
			ResponseJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

// JSON response function
func ResponseJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}

// Allow CORS
func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true") // Required for credentials
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

// Fallback Handler
/*
	React Router DOM or any client side routing will now work by just
	using a plain file server: http.FileServer(http.Dir("../frontend/dist")),
	since there's no actual dist/dashboard file - only index.html knows how to
	route it client-side.
*/
func spaHandler(staticPath string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(staticPath))
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(staticPath, filepath.Clean(r.URL.Path))
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	}
}
