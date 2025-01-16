# Workflow-go

## Overview

The **workflow-go SDK** simplifies interaction with Temporal workflows, providing easy-to-use methods for managing workflows, activities, and workers.

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

## Getting Started with workflow-go SDK

### Prerequisites

Ensure you have the following installed:
- [Go](https://go.dev/) (version 1.18 or later recommended)
- Docker and Docker Compose

### Installation

Add the SDK to your project using `go get`:
```bash
go get github.com/Zampfi/workflow-sdk-go/sdk
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


## Making a release

1. Pull `main`.
2. Add a git tag with a bumped version as per your change. If the current latest tag is v1.1.50, tag the latest commit on `main` with `v1.1.51`.
   
    ```
    # example: git tag v1.1.51
    git tag <version>
    ```
3. Push the tag
   
   ```
   git push origin v1.1.51
   ```

## Using workflow-go sdk in other repos

1. `go get` the latest version from the go project that you wish to use workflow-go sdk in.

   ```
   go get github.com/Zampfi/workflow-sdk-go/sdk@v1.1.50
   ```
