package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed all:dist
var daxUIFS embed.FS

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	uiSub, err := fs.Sub(fs.FS(daxUIFS), "dist")
	if err != nil {
		log.Fatal(err)
	}
	ui := http.FileServer(http.FS(uiSub))
	router := chi.NewMux()
	// router.PathPrefix("/").Methods(http.MethodGet).Handler(http.StripPrefix("/", ui))
	router.Handle("/*", ui)
	// apiRouter := router.PathPrefix("/api/").Subrouter()
	//router.Path("/dux/api/users").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	router.Get("/dux/api/users", func(w http.ResponseWriter, r *http.Request) {
		content := map[string]any{
			"Name":    "Alwin Doss",
			"Age":     32,
			"Address": "Bangalore",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(content)
	})
	addr := fmt.Sprintf("%s", os.Getenv("DUX_ADDR"))
	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, router)
}
