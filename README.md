# Citadel-Go SDK

## Overview

The **Citadel-Go SDK** simplifies interaction with Temporal workflows, providing easy-to-use methods for managing workflows, activities, and workers.

## Setting Up Temporal Locally

Follow these steps to set up Temporal on your local environment:

1. **Clone the Temporal Docker Compose Repository**:
   ```bash
   git clone https://github.com/temporalio/docker-compose
   ```

2. **Navigate to the Cloned Repository**:
   ```bash
   cd docker-compose
   ```

3. **Start the Temporal Server**:
   ```bash
   docker compose up
   ```

   This will start the Temporal server using Docker Compose.

## Getting Started with Citadel-Go SDK

### Prerequisites

Ensure you have the following installed:
- [Go](https://go.dev/) (version 1.18 or later recommended)
- Docker and Docker Compose

### Installation

Add the SDK to your project using `go get`:
```bash
go get github.com/yourusername/citadel-go
```
## Features

- **Seamless Workflow Management**: Start, query, signal, and cancel workflows.
- **Worker Setup**: Easily configure and run Temporal workers.
- **Activity Registration**: Simplified activity management.

## Samples Package

This repository includes a `samples` package that provides examples for various use cases. To use the samples:

1. Start the worker first.
    -- go to the worker/main.go
    -- go run main.go
2. Start the starter to trigger the workflows.
    -- go to the starter/main.go
    -- go run main.go

These examples demonstrate how to integrate and use the SDK in real-world scenarios.

