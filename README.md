# PEA: Prompt Engineering Agent

PEA is a local-first, Windows-first Prompt Engineering Agent that helps you craft, test, and manage high-quality prompts for Gemini models. It provides a rich desktop UI, a powerful Go backend, and a privacy-first architecture that keeps your data on your machine.

## Quickstart (Windows 11)

### Prerequisites

1.  **Go:** [Install Go](https://golang.org/doc/install) for Windows.
2.  **Flutter:** [Install Flutter](https://flutter.dev/docs/get-started/install/windows) for Windows desktop.
3.  **Gemini CLI:** [Install and authenticate the Gemini CLI](https://ai.google.dev/tutorials/gemini_cli_quickstart). You must complete the OAuth flow so that the CLI can access your account.

### Running the Application

1.  **Run the Backend Server:**

    ```bash
    # Set the mock environment variable to use the mock Gemini server
    set PEA_GEMINI_MOCK=true

    # Run the server
    go run ./cmd/pea serve --db ./data/pea.db --port 8080
    ```

2.  **Run the Flutter Client:**

    ```bash
    # Navigate to the client directory
    cd client/flutter_app

    # Run the Flutter app
    flutter run -d windows
    ```

### Running Tests

To run the full suite of unit and integration tests, use the following command:

```bash
set PEA_GEMINI_MOCK=true
go test ./...
```

## Privacy

PEA is designed to be privacy-first. Here's what you need to know:

*   **Local-First:** All your data, including conversations, prompts, and attachments, is stored locally in a SQLite database.
*   **Redaction:** By default, PEA redacts sensitive information like email addresses and API keys before storing it.
*   **`--no-store`:** For ephemeral sessions, you can use the `--no-store` flag to prevent any data from being written to the database.
*   **`--no-redact`:** For local debugging, you can use the `--no-redact` flag to disable redaction. **Do not use this in production or with sensitive data.**

## Gemini CLI Authentication

PEA uses the Gemini CLI for all interactions with the Gemini API. This means you don't need to manage API keys. Instead, you authenticate the CLI once using Google's OAuth flow, and PEA will use your existing credentials.

To authenticate the Gemini CLI, run the following command and follow the instructions:

```bash
gemini auth login
```
