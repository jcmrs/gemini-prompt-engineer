package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// 1. Meta frame
	meta := map[string]interface{}{
		"type":  "meta",
		"model": "gemini-2.5-flash",
		"usage": map[string]int{"tokens": 123},
	}
	printJSON(meta)

	// 2. Token frames
	tokens := []string{"This", " ", "is", " ", "a", " ", "mock", " ", "response."}
	for i, token := range tokens {
		frame := map[string]interface{}{
			"type":        "token",
			"data":        token,
			"chunk_index": i,
			"is_final":    false,
		}
		printJSON(frame)
		time.Sleep(50 * time.Millisecond)
	}

	// 3. Progress frames
	for i := 0; i <= 100; i += 10 {
		frame := map[string]interface{}{
			"type":    "progress",
			"percent": i,
		}
		printJSON(frame)
		time.Sleep(20 * time.Millisecond)
	}

	// 4. Final frame
	final := map[string]interface{}{
		"type":    "final",
		"content": "This is a mock response.",
		"metrics": map[string]interface{}{},
	}
	printJSON(final)
}

func printJSON(data map[string]interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("{\"type\":\"error\", \"message\":\"%s\"}\n", err.Error())
		return
	}
	fmt.Println(string(bytes))
}
