# FFmpegCommand

FFmpegCommand is a Go package that provides a simple and flexible way to interact with the `ffmpeg` command line interface. This package is designed to make it easy to build and execute `ffmpeg` commands programmatically, with additional features like custom execution time, output handling, and debug logging.

## Features

- Build `ffmpeg` commands dynamically with easy-to-use methods.
- Set custom execution time limits for commands.
- Retrieve the standard output or error messages from `ffmpeg`.
- Debug mode for detailed logging of command execution.
- Methods to check command execution status and stop it if needed.

## Installation

To use the FFmpegCommand package, first install it using `go get`:

```bash
go get github.com/ZertyCraft/GoFFmpeg
```

Then, import it into your Go project:

```go
import "github.com/ZertyCraft/GoFFmpeg"
```

## Usage

Here's a basic example of how to use the FFmpegCommand package:

```go
package main

import (
    "log"
    "github.com/ZertyCraft/GoFFmpeg"
    "time"
)

func main() {
    ffmpegCmd := ffmpegcommand.NewFFmpegCommand().
        SetReturnOutput(false).
        SetExecutionTime(1 * time.Minute).
        AddArgs("-i input.mp4", "-c:v copy", "output.mp4")

    output, err := ffmpegCmd.Execute()
    if err != nil {
        log.Fatalf("Failed to execute ffmpeg: %v", err)
    }

    if output != "" {
        log.Println("ffmpeg output:", output)
    }
}
```

## Documentation

### Constructor

- `NewFFmpegCommand() *FFmpegCommand`: Initializes a new FFmpeg command.

### Methods

- `AddArgs(args ...string) *FFmpegCommand`: Adds arguments to the ffmpeg command.
- `SetReturnOutput(returnOutput bool) *FFmpegCommand`: Determines if the output of the command should be returned.
- `SetExecutionTime(duration time.Duration) *FFmpegCommand`: Sets the maximum execution time for the command.
- `Execute() (string, error)`: Executes the ffmpeg command and manages the internal context.
- `Stop()`: Stops the ffmpeg command if it's currently running.
- `IsRunning() bool`: Checks if the ffmpeg command is currently running.
- `GetArgs() []string`: Returns the arguments of the ffmpeg command.
- `GetCommand() *exec.Cmd`: Returns the ffmpeg command.
- `GetExecutionTime() time.Duration`: Returns the maximum execution time of the command.
- `GetReturnOutput() bool`: Returns whether the command's output should be returned.
- `GetLastError() string`: Returns the last error message, if any.

### Debug Mode

Enable debug mode to receive detailed logs during command execution:

```go
ffmpegCmd.EnableDebugMode(true)
```
