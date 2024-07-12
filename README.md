# CodeChat: Real-time Chat Application

Welcome to CodeChat, a real-time chat application designed to facilitate seamless communication through a terminal-based interface. Built with a focus on simplicity and efficiency, CodeChat leverages WebSockets for real-time messaging and is structured to provide a clear separation of concerns, making it both scalable and easy to maintain.

## Upcoming Integrations (Proposed)

- [ ] Tview, for a better Terminal aspect.
- [ ] Web interface (Optional)

## Current Features

- **Real-time Communication**: Instantly send and receive messages with minimal latency.
- **CLI-based Interface**: Access and interact with the chat directly from your terminal.
- **Scalable Architecture**: Designed with a modular architecture that separates the client, server, and common utilities.

## Getting Started

To get started with CodeChat, clone this repository to your local machine and follow the setup instructions below.

### Prerequisites

- Golang
- WebSocket support

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/<yourusername>/CodeChat.git
   ```

2. Navigate to the cloned directory:

   ```sh
   cd CodeChat
   ```

3. Build the application:

   ```sh
   ./build.sh
   ```

   If you encounter a permission denied error, run `chmod +x build.sh` to make the script executable.

### Running CodeChat

To start the ChatApp server, run:

```sh
./CodeChat server
```

To connect as a client from another terminal, run:

```sh
./CodeChat client
```

## Project Structure

CodeChat is organized into several key components for clarity and ease of development:

- **`cmd/`**: Contains the CLI setup using CobraCLI, including `root.go` for CLI environment setup and `chat.go` for chat command definitions.
- **`internal/`**:
  - **`server/`**: Contains server logic, WebSocket connection management (`server.go`), and client handling (`client.go`).
  - **`client/`**: Manages chat client functionality, WebSocket connections, and terminal UI rendering (`ui.go`).
- **`pkg/`**:
  - **`websocket/`**: Provides utilities for WebSocket operations, like message encoding/decoding.
  - **`common/`**: Defines common data structures, such as message formats.
- **`main.go`**: The application's entry point, initializing and running Cobra commands.

## Contributing

Contributions to CodeChat are welcome!

## License

CodeChat is open-sourced under the MIT License. See the LICENSE file for more details.
