package models

// Conversation represents a chat conversation.
type Conversation struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Message represents a single message in a conversation.
type Message struct {
	ID             string      `json:"id"`
	ConversationID string      `json:"conversation_id"`
	Role           string      `json:"role"`
	Content        string      `json:"content"`
	Attachments    []string    `json:"attachments"`
	Model          string      `json:"model"`
	ModelConfig    interface{} `json:"model_config"`
	PromptVersion  string      `json:"prompt_version"`
	CreatedAt      string      `json:"created_at"`
}

// Prompt represents a prompt template.
type Prompt struct {
	ID          string      `json:"id"`
	Version     string      `json:"version"`
	Author      string      `json:"author"`
	Intent      string      `json:"intent"`
	Description string      `json:"description"`
	PromptText  string      `json:"prompt_text"`
	Settings    interface{} `json:"settings"`
	Examples    interface{} `json:"examples"`
	CreatedAt   string      `json:"created_at"`
}
