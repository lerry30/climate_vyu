# Handling API Endpoint

This struct handles api request and holds properties like port and other external structs that can be use internally.

```go
type APIServer struct {
    port string
    externals mapp[string]any
}
```

## Types

The <code>apiFunc</code> function type is use by wrapper for adding extra functionalities, especially handling errors instead of relaying with the plain function (<code>http.HandleFunc</code> function).

<code>APIError</code> is a struct for error messages so defining error message doesn't required to specified the type all the time. The actual purpose is to make the format consistent.

```go
// function type
type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string
}
```

## Constructor

Initializing and assinging property values.

```go
func NewAPIServer(port string) *APIServer {
	return &APIServer{
		port:      port,
		externals: make(map[string]any),
	}
}
```

## Methods

For a small app like this creating a map of type <code>any</code> to hold external structs, that can be use inside the <code>APIServer</code> methods, is fine but for large applications interface is the must have and preferred pattern for actually any case.

With proper key name for a certain value, it can map easily.

```go
func (apiServer *APIServer) AddExternalAPI(name string, address any) {
	apiServer.externals[name] = address
}
```

### The Main Function of This Struct

Using gorilla mux package for request routes. The most fascinating feature of gorilla mux is it can filter the url paths like value type filtration(e.g. number or string) with regular expression.

This function handles the incoming http requests and serving the client/frontend website.

```go
func (apiServer *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/api/filter/{city}", makeHTTPHandleFunc(apiServer.FilterCity)).Methods("GET")
	router.HandleFunc("/api/search/{city}", makeHTTPHandleFunc(apiServer.SearchCity)).Methods("GET")
	router.HandleFunc("/api/forecast/{city}", makeHTTPHandleFunc(apiServer.Forecast)).Methods("GET")
	// Ping heath check
	router.HandleFunc("/api/health", makeHTTPHandleFunc(apiServer.Health)).Methods("GET")

	router.PathPrefix("/").Handler(spaHandler("../frontend/dist"))

	fmt.Println("Server starting on port:", apiServer.port[1:])
	log.Fatal(http.ListenAndServe(apiServer.port, router))
}
```

### Wrapper

This wrapper allows to throw a response error of 400. Every function that implements the type of <code>apiFunc</code> can return a custom error message like <code>fmt.Errorf("Error")</code>. But this way makes the http status error code limited to <code>http.StatusBadRequest</code>.

```go
func makeHTTPHandleFunc(f apiFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		EnableCORS(&w)

		if err := f(w, r); err != nil {
			ResponseJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
```

### Response Parameters

Here are the things that requires for http response. The response body is encoded with json format.

```go
func ResponseJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}
```

### Preventing CORS Error

This helps to prevent the browser from throwing error about CORS when accessing the API endpoint with different origin. Let say both backend and frontend host on different server, the browser will prevent the communition between the two. I only used this function in development.

```go
func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true") // Required for credentials
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
```

### Fallback Handler

React Router DOM or any client side routing will now work by just
using a plain file server: http.FileServer(http.Dir("../frontend/dist")),
since there's no actual dist/dashboard file - only index.html knows how to
route it client-side.

```go
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
```