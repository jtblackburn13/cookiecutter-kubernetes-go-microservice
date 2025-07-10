package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"gopkg.in/yaml.v3"
)

type Handler struct {
	startTime    time.Time
	requestCount int64
	version      string
	buildTime    string
	gitCommit    string
}

type Response struct {
	Message   string    `json:"message"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

type MetricsResponse struct {
	Service       string    `json:"service"`
	Uptime        string    `json:"uptime"`
	RequestsTotal int64     `json:"requests_total"`
	Version       string    `json:"version"`
	Timestamp     time.Time `json:"timestamp"`
}

type VersionResponse struct {
	Service   string `json:"service"`
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
	GitCommit string `json:"git_commit"`
	GoVersion string `json:"go_version"`
}

type ErrorResponse struct {
	Error     string    `json:"error"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

func New() *Handler {
	return &Handler{
		startTime: time.Now(),
		version:   "1.0.0",
		buildTime: time.Now().Format(time.RFC3339),
		gitCommit: "unknown",
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	// API routes
	router.Get("/", h.HandleRoot)
	router.Get("/health", h.HandleHealth)
	router.Get("/ready", h.HandleReady)
	router.Get("/metrics", h.HandleMetrics)
	router.Get("/version", h.HandleVersion)

	// Swagger documentation
	router.Mount("/docs", httpSwagger.Handler(
		httpSwagger.URL("/swagger.yaml"),
	))

	// Serve swagger.yaml file
	router.Get("/swagger.yaml", h.HandleSwaggerYAML)

	// Middleware to count requests
	router.Use(h.RequestCountMiddleware)
}

func (h *Handler) RequestCountMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&h.requestCount, 1)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message:   "Hello from {{cookiecutter.service_name}}!",
		Service:   "{{cookiecutter.service_name}}",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Service:   "{{cookiecutter.service_name}}",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HandleReady(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "ready",
		Service:   "{{cookiecutter.service_name}}",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HandleMetrics(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(h.startTime)

	response := MetricsResponse{
		Service:       "{{cookiecutter.service_name}}",
		Uptime:        uptime.String(),
		RequestsTotal: atomic.LoadInt64(&h.requestCount),
		Version:       h.version,
		Timestamp:     time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HandleVersion(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{
		Service:   "{{cookiecutter.service_name}}",
		Version:   h.version,
		BuildTime: h.buildTime,
		GitCommit: h.gitCommit,
		GoVersion: runtime.Version(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HandleSwaggerYAML(w http.ResponseWriter, r *http.Request) {
	// Read swagger.yaml file
	swaggerPath := filepath.Join("docs", "swagger.yaml")
	swaggerContent, err := ioutil.ReadFile(swaggerPath)
	if err != nil {
		http.Error(w, "Swagger file not found", http.StatusNotFound)
		return
	}

	// Parse and validate YAML
	var swagger interface{}
	if err := yaml.Unmarshal(swaggerContent, &swagger); err != nil {
		http.Error(w, "Invalid swagger YAML", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-yaml")
	w.Write(swaggerContent)
}
