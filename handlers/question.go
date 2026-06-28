package handlers

import (
	"encoding/json"
	"net/http"
  "log"
	"question-system/services"
)

type QuestionRequest struct {
	Question string `json:"question"`
}

func SubmitQuestion(w http.ResponseWriter, r *http.Request) {

 w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req QuestionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Question == "" {
		http.Error(w, "Question is required", http.StatusBadRequest)
		return
	}

	err = services.SendQuestionEmail(req.Question)
if err != nil {
    log.Println("Email error:", err)
    http.Error(w, "Failed to send email", http.StatusInternalServerError)
    return
}
	

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question sent successfully",
	})
}