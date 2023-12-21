package GoFFmpeg

import (
	"reflect"
	"testing"
	"time"
)

func TestFFmpegCommand_Execute(t *testing.T) {
	// Create a new FFmpegCommand instance
	ffmpegCmd := NewFFmpegCommand().
		SetReturnOutput(false).
		SetExecutionTime(1*time.Minute).
		AddArgs("-y", "-i", "tests/video/test_input.mp4", "-c:v", "copy", "tests/video/test_output.mp4").
		EnableDebugMode(true)

	// Execute the command
	output, err := ffmpegCmd.Execute()

	// Check if there was an error
	if err != nil {
		t.Fatalf("Failed to execute ffmpeg command: %v\n%s", err, ffmpegCmd.GetLastError())
	}

	// Check if the output is empty
	if output != "" {
		t.Errorf("Unexpected output from ffmpeg command: %s", output)
	}
}

func TestFFmpegCommand_Stop(t *testing.T) {
	// Create a new FFmpegCommand instance
	ffmpegCmd := NewFFmpegCommand().
		SetReturnOutput(false).
		SetExecutionTime(1*time.Minute).
		AddArgs("-y", "-i", "tests/video/test_input.mp4", "-c:v", "copy", "tests/video/test_output.mp4")

	// Start executing the command in a separate goroutine
	go func() {
		_, _ = ffmpegCmd.Execute()
	}()

	// Stop the command
	ffmpegCmd.Stop()

	// Check if the command was stopped correctly
	if ffmpegCmd.IsRunning() {
		t.Errorf("Failed to stop the command")
	}
}

func TestFFmpegCommand_AddArgs(t *testing.T) {
	// Create a new FFmpegCommand instance
	ffmpegCmd := NewFFmpegCommand().
		SetReturnOutput(false).
		SetExecutionTime(1 * time.Minute)

	// Add arguments to the command
	ffmpegCmd.AddArgs("-i", "tests/video/test_input.mp4", "-c:v", "copy", "tests/video/test_output.mp4")

	// Check if the arguments were added correctly
	args := ffmpegCmd.GetArgs()
	expectedArgs := []string{"ffmpeg", "-i", "tests/video/test_input.mp4", "-c:v", "copy", "tests/video/test_output.mp4"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("Unexpected arguments. Got %v, want %v", args, expectedArgs)
	}
}

func TestFFmpegCommand_SetReturnOutput(t *testing.T) {
	// Create a new FFmpegCommand instance
	ffmpegCmd := NewFFmpegCommand().
		SetExecutionTime(1*time.Minute).
		AddArgs("-y", "-i", "tests/video/test_input.mp4", "-c:v", "copy", "tests/video/test_output.mp4")

	// Set return output to true
	ffmpegCmd.SetReturnOutput(true)

	// Check if the return output flag was set correctly
	if !ffmpegCmd.GetReturnOutput() {
		t.Errorf("Failed to set return output flag correctly")
	}
}

func TestFFmpegCommand_SetExecutionTime(t *testing.T) {
	// Create a new FFmpegCommand instance
	ffmpegCmd := NewFFmpegCommand().
		SetReturnOutput(false).
		AddArgs("-y", "-i", "tests/video/test_input.mp4", "-c:v", "copy", "tests/video/test_output.mp4")

	// Set execution time to 2 minutes
	ffmpegCmd.SetExecutionTime(2 * time.Minute)

	// Check if the execution time was set correctly
	if ffmpegCmd.GetExecutionTime() != 2*time.Minute {
		t.Errorf("Failed to set execution time correctly")
	}
}
