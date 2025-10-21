package server

import (
	"encoding/json"
	"github.com/jcmrs/gemini-prompt-engineer/internal/gemini"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// NewServer creates a new HTTP server.
func NewServer(addr string) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/auth/check", authCheckHandler).Methods("GET")
	// [TODO-JULES] Implement the remaining REST endpoints.

	r.HandleFunc("/ws/run/{run_id}", wsRunHandler)

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Include the commit SHA in the response.
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "version": "dev"})
}

func authCheckHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement the real auth check logic.
	wrapper := gemini.NewWrapperFromEnv()
	if err := wrapper.CheckAuth(r.Context()); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"authenticated": false, "message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"authenticated": true, "message": "Authenticated successfully."})
}

func wsRunHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connection established.")

	// Use the Gemini wrapper to stream a mock response.
	wrapper := gemini.NewWrapperFromEnv()
	onToken := func(token string, idx int, isFinal bool) {
		frame := map[string]interface{}{
			"type":        "token",
			"data":        token,
			"chunk_index": idx,
			"is_final":    isFinal,
		}
		bytes, _ := json.Marshal(frame)
		if err := conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
			log.Println("Failed to write message to WebSocket:", err)
		}
	}

	if _, err := wrapper.RunChatStreaming(r.Context(), "gemini-2.5-flash", "test", nil, onToken); err != nil {
		log.Println("Failed to run chat streaming:", err)
	}
}
