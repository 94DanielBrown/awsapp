# awsapp

A Go module providing common AWS service utilities and convenience functions for building AWS-backed web APIs.

## Overview

This module provides both high-level convenience functions and low-level utilities for working with AWS services used commonly in Go based Web APIs or at least the ones I make/have made.

## Installation

```bash
go get github.com/94DanielBrown/awsapp
```

## Structure

```
awsapp/
├── awsapp.go           # High-level convenience functions
├── pkg/
│   ├── awsconfig/      # Shared AWS configuration
│   ├── dynamo/         # Low-level DynamoDB operations
│   └── s3/             # Low-level S3 operations
```

## Usage

The root package provides opinionated, high-level functions that handle common initialization patterns:

```go
import (
    "context"
    "github.com/94DanielBrown/awsapp"
)

func main() {
    ctx := context.Background()
    
    // Initialize a DynamoDB table (creates if doesn't exist)
    client, message, err := awsapp.InitDynamo(ctx, "my-table")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message) // "table my-table created successfully"
}
```


For more control, you can use the individual packages directly:

```go
import (
    "context"
    "github.com/94DanielBrown/awsapp/pkg/dynamo"
    "github.com/94DanielBrown/awsapp/pkg/awsconfig"
)

func main() {
    ctx := context.Background()
    
    // Connect to DynamoDB
    client, err := dynamo.Connect()
    if err != nil {
        log.Fatal(err)
    }
    
    // Check if table exists
    exists, err := dynamo.Exists(ctx, client, "my-table")
    if err != nil {
        log.Fatal(err)
    }
    
    // Create table with custom configuration
    if !exists {
        err = dynamo.Create(ctx, client, "my-table")
        // Handle error...
    }
}
```

## Available Functions

### Root Package (`awsapp`)

- `InitDynamo(ctx, tableName)` - Initialize DynamoDB table (creates if doesn't exist)
  - Handles connection, existence check, and creation
  - Includes 5-minute timeout for creation
  - Returns client, status message, and error

### DynamoDB Package (`pkg/dynamo`)

- `Connect()` - Create DynamoDB client using AWS credentials from environment
- `Exists(ctx, client, tableName)` - Check if a table exists
- `Create(ctx, client, tableName)` - Create a new table
- Additional CRUD operations (see package documentation)

### S3 Package (`pkg/s3`)

- S3-specific operations (see package documentation)

### Config Package (`pkg/awsconfig`)

- Shared AWS configuration utilities
- Credential management
- Region configuration

## Configuration

The module uses AWS credentials from environment variables. Ensure you have set:

```bash
export AWS_REGION=us-east-1
export AWS_ACCESS_KEY_ID=your-access-key
export AWS_SECRET_ACCESS_KEY=your-secret-key
```

Or use AWS IAM roles if running on EC2/ECS/Lambda.

## Author

Daniel Brown
