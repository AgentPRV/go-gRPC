# gRPC Explorer

This repository demonstrates a basic gRPC implementation in Go. The following steps guide you through the setup and installation process to run the gRPC service and client.

Before Diving into This Repo, Let's See What gRPC Is

## What is gRPC?

gRPC (Google Remote Procedure Call) is a high-performance, open-source framework developed by Google that enables remote procedure calls (RPC) between client and server applications. It allows you to define services and methods using Protocol Buffers (protobuf), a language-agnostic binary serialization format, and then automatically generates client and server code for multiple programming languages.

### Key Features of gRPC:
- **Language-agnostic:** Supports multiple programming languages, allowing for seamless communication across different platforms.
- **Efficient Serialization:** Uses Protocol Buffers, which are faster and more compact than traditional text-based formats like JSON or XML.
- **Streaming:** Supports client-side, server-side, and bidirectional streaming, enabling real-time data exchange.
- **Built on HTTP/2:** gRPC leverages HTTP/2 features like multiplexing, flow control, and header compression to improve performance.
- **Built-in Authentication and Load Balancing:** Provides support for secure communication and integrates with load balancers to manage traffic efficiently.

gRPC is ideal for building microservices, real-time applications, and scenarios where performance, scalability, and cross-language support are critical.


## Prerequisites

Before you begin, ensure you have the following installed:

- Go (latest version recommended)
- Protocol Buffers Compiler (`protoc`)

## Installation Steps

### 1. Install Go Tools

First, install the necessary Go tools for working with Protocol Buffers and gRPC:

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 2. Install protoc
Ensure you have the Protocol Buffers Compiler (protoc) installed. You can follow the installation guide from the official gRPC documentation:

[Protocol Buffers Compiler Installation](https://grpc.io/docs/protoc-installation/)

Note - for mac brew method will work fine.

### 3. Generate Go Code from Proto Files
Once protoc is installed, you can generate the Go code from your .proto files. From root of server run the following command:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       log.proto
```
This command generates both the Go code for your messages and the gRPC client and server interfaces.


### 4. Install missing dependencies
Install dependencies by running - 
```bash
go mod tidy
```


### 5. Running the Project
To run the gRPC server and client, follow these steps:

Compile and run the server:

From server folder - 
```bash
go run .
```

From client folder - 
```bash
go run .
```

This should provide clear guidance to anyone looking to understand or set up your gRPC project.