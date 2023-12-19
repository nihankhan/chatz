
# Chatz

Chatz is a real-time chat application written in Go, leveraging the Gorilla WebSocket library.

## Description

A simple and scalable chat application that allows users to communicate in real-time. It uses WebSocket for efficient and instant message delivery.

## Features

- Real-time WebSocket-based chat.
- User authentication and identification.
- Support for multiple chat rooms.
- Emoji and file sharing.
- ...

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/) (at least version 1.21.5)

### Installation

Clone the repository:

```bash
git clone https://github.com/nihankhan/chatz.git
cd chatz
```

Build and run the project:

```bash
go build
./chatz
```

The server will start on `http://localhost:8080`.

## Usage

1. Open your web browser and navigate to `http://localhost:8080`.
2. Enter your username and start chatting!
3. Explore different chat rooms and features.

## Configuration

Customize configuration parameters in `config.yml`.

## Contributing

We welcome contributions! If you find a bug or have an enhancement in mind, [open an issue](https://github.com/nihankhan/chatz/issues) or create a pull request. Please follow our [contributing guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gorilla WebSocket](https://github.com/gorilla/websocket) - A WebSocket library for Go.
