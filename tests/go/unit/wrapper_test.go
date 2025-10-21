package unit

import (
	"context"
	"github.com/jcmrs/gemini-prompt-engineer/internal/gemini"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Set the mock environment variable for all tests in this package.
	os.Setenv("PEA_GEMINI_MOCK", "true")
	os.Exit(m.Run())
}

func TestMockWrapper_RunChatStreaming(t *testing.T) {
	wrapper := gemini.NewWrapperFromEnv()

	var tokens []string
	onToken := func(token string, idx int, isFinal bool) {
		if !isFinal {
			tokens = append(tokens, token)
		}
	}

	finalResult, err := wrapper.RunChatStreaming(context.Background(), "gemini-2.5-flash", "test", nil, onToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if expected := "This is a mock response."; finalResult != expected {
		t.Errorf("expected final result to be %q, got %q", expected, finalResult)
	}

	if expected := "This is a mock response."; strings.Join(tokens, "") != expected {
		t.Errorf("expected concatenated tokens to be %q, got %q", expected, strings.Join(tokens, ""))
	}
}
