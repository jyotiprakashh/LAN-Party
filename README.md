# 🎉 LAN Party Chat 🕹️

Welcome to **LAN Party Chat**, a fun and simple TCP-based chat application! Gather your friends, connect to the same network, and start chatting away like it's the 90s again!

## 📋 Table of Contents
- [Introduction](#-introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Packages Used](#packages-used)
- [Contributing](#contributing)

## 🌟 Introduction

LAN Party Chat is a terminal-based chat application designed for local area networks. It's built with Go and uses TCP for real-time communication. Whether you're in the same room or just on the same network, you can chat with your friends like it's the good old days of LAN parties!

## 🌟 Features

- **Real-time Chat**: Instant messaging with all connected users.
- **Active User Count**: Always know how many of your friends are online.
- **User-friendly Interface**: Simple and intuitive text-based UI using `tview`.

## 📦 Packages Used

- [**net**](https://pkg.go.dev/net): For network connections and TCP communication.
- [**bufio**](https://pkg.go.dev/bufio): For buffered I/O operations.
- [**tcell**](https://github.com/gdamore/tcell/v2): For handling terminal I/O.
- [**tview**](https://github.com/rivo/tview): For building rich terminal-based UIs.

## 📥 Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/lan-party-chat.git
    cd lan-party-chat
    ```

2. Ensure you have Go installed. You can download it [here](https://golang.org/dl/).

## 🚀 Usage

### Run the Server

First, start the server:
```bash
go run cmd/server/main.go
```
The server will run on localhost:8080 by default
![alt text](https://github.com/jyotiprakashh/LAN-Party/blob/main/screenshots/server.png?raw=true)


### Run the Client

In a separate terminal, start a client:
```bash
go run cmd/client/main.go
```

Your terminal will convert to this ui where it will ask for an username:
![alt text](https://github.com/jyotiprakashh/LAN-Party/blob/main/screenshots/client1.png?raw=true)

If you will be alone, a message will show to invite your friends:
![alt text](https://github.com/jyotiprakashh/LAN-Party/blob/main/screenshots/client2.png?raw=true)

After anyone joins the room notification message wwill be shown that they have joined and the number of active users in the chat will get updated:
![alt text](https://github.com/jyotiprakashh/LAN-Party/blob/main/screenshots/client3.png?raw=true)

If the any user leaves the room, it will show that they have left the chat.
![alt text](https://github.com/jyotiprakashh/LAN-Party/blob/main/screenshots/client4.png?raw=true)

### Using the Makefile
Alternatively, you can use the provided Makefile to build and run the server and client:

Build the Server and Client
```bash
make build
```

Run the Server
```bash
make run-server
```

Run the Client
```bash
make run-client
```


## 🤝 Contributing
Contributions are welcome! Please fork this repository, make your changes, and submit a pull request.
