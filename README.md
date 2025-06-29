# gocmd

This package provides a simple framework for registering and executing CLI commands.

## Overview

The `cmd` package allows you to define commands with a name, description, usage, and a handler function. You can register multiple commands and execute them based on command-line arguments.

## Usage

### 1. Registering Commands

To add a new command, use the `RegisterCommand` function:

```go
cmd.RegisterCommand(cmd.Command{
    Name: "hello",
    Description: "Prints Hello World",
    Usage: "{run command} hello",
    Handler: func(ctx context.Context) {
        println("Hello World")
    },
})
```

### 2. Executing Commands

To execute the registered commands, call:

```go
cmd.Execute()
```

This will parse the command-line arguments and run the handler for the matching command name.

### 3. Example

```go
package main

import (
    "context"
    "gocmd/cmd"
)

func main() {
    cmd.RegisterCommand(cmd.Command{
        Name: "hello",
        Description: "Prints Hello World",
        Usage: "{run command} hello",
        Handler: func(ctx context.Context) {
            println("Hello World")
        },
    })

    cmd.Execute()
}
```

If you run your program with:

```
go run main.go hello
```

It will print:

```
Hello World
```

If no command or an invalid command is given, it will show usage/help information.

## API

### Command Struct

```go
type Command struct {
    Name        string
    Description string
    Usage       string
    Handler     func(ctx context.Context)
}
```

### Cmd Interface

```go
type Cmd interface {
    RegisterCommand(cmd Command)
    Execute()
}
```

## License

MIT
