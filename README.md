# myaws CLI

A CLI tool to simplify AWS interactions, built with Go and the Cobra library.

## Overview

`myaws` is a CLI tool that simplifies interactions with AWS services like S3 and ECS. It follows a layered architecture:

- **cmd layer**: Handles command configuration, argument parsing, and calls to the service layer
- **service layer**: Contains the core logic for each command and interacts with the AWS client layer
- **awsclient layer**: Abstracts AWS SDK clients behind Go interfaces for better testability

## Building the CLI

To build the CLI, you need Go 1.19 or later.

```bash
# Clone the repository
git clone https://github.com/username/myaws.git
cd myaws

# Install dependencies
go mod tidy

# Build the CLI
go build -o myaws

# (Optional) Install the CLI
go install
```

## Running Tests

To run the tests for the CLI:

```bash
go test -v ./...
```

This will run all tests in the project, ensuring that commands like `version` work as expected.

## Available Commands

### Version

Displays the current version of the CLI.

```bash
myaws version
```

### List App Config

Lists configuration files from an S3 bucket.

```bash
# Basic usage (table output)
myaws list-app-config --bucket your-bucket-name

# With prefix filter
myaws list-app-config --bucket your-bucket-name --prefix config/

# JSON output
myaws list-app-config --bucket your-bucket-name --output json
```

### Restart App (Proof of Concept)

A stub command for restarting an application running on ECS.

```bash
myaws restart-app --cluster your-cluster --service your-service
```

This command is intentionally left incomplete for demonstration purposes. It will display a message indicating that the implementation is incomplete.

## Architecture

The CLI follows a layered architecture:

1. **cmd layer**: Uses Cobra to define commands, flags, and parse arguments
2. **service layer**: Contains the business logic for each command
3. **awsclient layer**: Provides interfaces and implementations for AWS services

This separation allows for easier testing and maintenance.

## Future Enhancements

- Complete the `restart-app` command to actually restart ECS services
- Add more AWS service integrations
- Add support for profiles and regions
- Implement pagination for listing objects 