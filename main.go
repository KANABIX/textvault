package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Session represents a text session with expiration
type Session struct {
	ID        string
	Text      string
	ExpiresAt time.Time
}

var (
	sessions = make(map[string]Session)
	mu       sync.RWMutex
)

func generateID(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	durationStr := r.FormValue("duration")
	if text == "" || durationStr == "" {
		http.Error(w, "Missing text or duration", http.StatusBadRequest)
		return
	}
	duration, err := strconv.Atoi(durationStr)
	if err != nil || duration <= 0 {
		http.Error(w, "Invalid duration", http.StatusBadRequest)
		return
	}
	id := generateID(8)
	expiresAt := time.Now().Add(time.Duration(duration) * time.Minute)
	session := Session{ID: id, Text: text, ExpiresAt: expiresAt}
	mu.Lock()
	sessions[id] = session
	mu.Unlock()
	// Return the link as JSON
	resp := map[string]string{"link": fmt.Sprintf("/paste/%s", id)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func viewSessionHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/paste/")
	mu.RLock()
	session, ok := sessions[id]
	mu.RUnlock()
	if !ok || time.Now().After(session.ExpiresAt) {
		http.Error(w, "Session not found or expired", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles(filepath.Join("static", "view.html"))
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, session)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("static", "index.html"))
}

func cleanupExpiredSessions() {
	for {
		time.Sleep(1 * time.Minute)
		mu.Lock()
		now := time.Now()
		for id, session := range sessions {
			if now.After(session.ExpiresAt) {
				delete(sessions, id)
			}
		}
		mu.Unlock()
	}
}

func main() {
	go cleanupExpiredSessions()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/paste", createSessionHandler)
	http.HandleFunc("/paste/", viewSessionHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
