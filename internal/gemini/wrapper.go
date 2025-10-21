package gemini

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

// Wrapper defines the interface for interacting with the Gemini CLI.
type Wrapper interface {
	CheckAuth(ctx context.Context) error
	RunChatStreaming(
		ctx context.Context,
		model string,
		input string,
		settings map[string]interface{},
		onToken func(token string, idx int, isFinal bool),
	) (finalResult string, err error)
	Embeddings(ctx context.Context, text string) ([]float32, error) // optional
}

// NewWrapperFromEnv creates a new Gemini wrapper based on the PEA_GEMINI_MOCK environment variable.
func NewWrapperFromEnv() Wrapper {
	if os.Getenv("PEA_GEMINI_MOCK") == "true" {
		return &mockWrapper{}
	}
	return &realWrapper{}
}

// mockWrapper is a mock implementation of the Wrapper interface.
type mockWrapper struct{}

func (w *mockWrapper) CheckAuth(ctx context.Context) error {
	return nil
}

func (w *mockWrapper) RunChatStreaming(
	ctx context.Context,
	model string,
	input string,
	settings map[string]interface{},
	onToken func(token string, idx int, isFinal bool),
) (string, error) {
	tokens := []string{"This", " ", "is", " ", "a", " ", "mock", " ", "response."}
	fullResponse := ""
	for i, token := range tokens {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(50 * time.Millisecond):
			onToken(token, i, false)
			fullResponse += token
		}
	}
	onToken("", len(tokens), true)
	return fullResponse, nil
}

func (w *mockWrapper) Embeddings(ctx context.Context, text string) ([]float32, error) {
	return nil, fmt.Errorf("not implemented")
}

// realWrapper is the real implementation of the Wrapper interface that interacts with the Gemini CLI.
type realWrapper struct{}

func (w *realWrapper) CheckAuth(ctx context.Context) error {
	// [TODO-JULES] Implement the full auth check logic as specified in the instructions.
	// Attempt `gemini whoami --format=json`.
	cmd := exec.CommandContext(ctx, "gemini", "whoami", "--format=json")
	if err := cmd.Run(); err != nil {
		// Fallback to `gemini auth status --format=json`.
		cmd = exec.CommandContext(ctx, "gemini", "auth", "status", "--format=json")
		if err := cmd.Run(); err != nil {
			// Fallback to a safe chat test.
			cmd = exec.CommandContext(ctx, "gemini", "chat", "--model=gemini-2.5-flash", "--format=json")
			stdin, err := cmd.StdinPipe()
			if err != nil {
				return fmt.Errorf("failed to get stdin pipe: %w", err)
			}
			go func() {
				defer stdin.Close()
				io.WriteString(stdin, "ping")
			}()
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to check gemini auth status: %w", err)
			}
		}
	}
	return nil
}

func (w *realWrapper) RunChatStreaming(
	ctx context.Context,
	model string,
	input string,
	settings map[string]interface{},
	onToken func(token string, idx int, isFinal bool),
) (string, error) {
	// [TODO-JULES] Implement the real streaming logic with process group handling and context cancellation.
	// [TODO-JULES] Handle timeouts and retries.
	// [TODO-JULES] Improve streaming parser.
	return "", fmt.Errorf("not implemented")
}

func (w *realWrapper) Embeddings(ctx context.Context, text string) ([]float32, error) {
	return nil, fmt.Errorf("not implemented")
}
