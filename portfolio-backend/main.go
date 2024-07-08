package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ContactMessage struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

var projects = []Project{
	{ID: 1, Title: "Proyek 1", Description: "Deskripsi proyek 1"},
	{ID: 2, Title: "Proyek 2", Description: "Deskripsi proyek 2"},
	{ID: 3, Title: "Proyek 3", Description: "Deskripsi proyek 3"},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/projects", getProjects)
	mux.HandleFunc("/api/contact", handleContact)

	handler := cors.Default().Handler(mux)

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var message ContactMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Di sini Anda bisa menambahkan logika untuk menyimpan atau mengirim pesan
	log.Printf("Received contact message from %s (%s): %s", message.Name, message.Email, message.Message)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
