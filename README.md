# aws-ug-cli CLI

A CLI tool to simplify AWS interactions, built with Go and the Cobra library.

## Overview

`aws-ug-cli` is a CLI tool that simplifies interactions with AWS services like S3 and ECS. It follows a layered architecture:

- **cmd layer**: Handles command configuration, argument parsing, and calls to the service layer
- **service layer**: Contains the core logic for each command and interacts with the AWS client layer
- **awsclient layer**: Abstracts AWS SDK clients behind Go interfaces for better testability

## AWS Credentials Configuration

The CLI supports AWS credentials through environment variables:

```bash
export AWS_ACCESS_KEY_ID="your_access_key"
export AWS_SECRET_ACCESS_KEY="your_secret_key"
export AWS_SESSION_TOKEN="your_session_token"  # Optional
export AWS_REGION="your_region"  # e.g., us-east-1
```

The AWS SDK will automatically pick up these environment variables. You can also use other AWS credential providers like:
- AWS credentials file (`~/.aws/credentials`)
- IAM roles for Amazon EC2 or ECS tasks
- AWS SSO

## Building the CLI

To build the CLI, you need Go 1.19 or later.

```bash
# Clone the repository
git clone https://github.com/username/aws-ug-cli.git
cd aws-ug-cli

# Install dependencies
go mod tidy

# Build the CLI
go build -o aws-ug-cli

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
aws-ug-cli version
```

### List Apps

Lists configured apps.

```bash
aws-ug-cli list-apps
```

### Flush cache

Flush cache from a given domain

```bash
aws-ug-cli cache flush --domain <your-domain>
```

### Dump the a db table

```bash
aws-ug-cli db dump --table Products
```

## Architecture

The CLI follows a layered architecture:

1. **cmd layer**: Uses Cobra to define commands, flags, and parse arguments
2. **service layer**: Contains the business logic for each command
3. **awsclient layer**: Provides interfaces and implementations for AWS services

This separation allows for easier testing and maintenance.
