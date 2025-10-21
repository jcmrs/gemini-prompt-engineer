package integration

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/jcmrs/gemini-prompt-engineer/internal/server"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Set the mock environment variable for all tests in this package.
	os.Setenv("PEA_GEMINI_MOCK", "true")
	os.Exit(m.Run())
}

func TestBasicRun(t *testing.T) {
	// Start the server in a goroutine.
	srv := server.NewServer(":8080")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			t.Errorf("unexpected error from server: %v", err)
		}
	}()
	defer srv.Shutdown(context.Background())

	// Wait for the server to be ready.
	time.Sleep(1 * time.Second)

	// Connect to the WebSocket.
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws/run/demo-run", nil)
	if err != nil {
		t.Fatalf("failed to connect to websocket: %v", err)
	}
	defer conn.Close()

	// Read messages from the WebSocket.
	var receivedTokens []string
	var isFinal bool
	for !isFinal {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var frame map[string]interface{}
		if err := json.Unmarshal(message, &frame); err != nil {
			t.Fatalf("failed to unmarshal json frame: %v", err)
		}

		if frame["type"] == "token" {
			receivedTokens = append(receivedTokens, frame["data"].(string))
		}
		isFinal = frame["is_final"].(bool)
	}

	// Assert that we received the expected messages.
	expectedResponse := "This is a mock response."
	if received := strings.Join(receivedTokens, ""); received != expectedResponse {
		t.Errorf("expected to receive %q, got %q", expectedResponse, received)
	}
}
