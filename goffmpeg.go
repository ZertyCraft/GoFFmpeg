package GoFFmpeg

import (
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type FFmpegCommand struct {
	args          []string
	returnOutput  bool
	cmd           *exec.Cmd
	executionTime time.Duration
	logger        *log.Logger
	debugMode     bool
	lastError     string
}

func NewFFmpegCommand() *FFmpegCommand {
	return &FFmpegCommand{
		args:          []string{"ffmpeg"},
		returnOutput:  true,
		executionTime: 30 * time.Second,
		logger:        log.New(os.Stdout, "FFmpegCommand: ", log.LstdFlags),
		debugMode:     false,
	}
}

func (fc *FFmpegCommand) EnableDebugMode(enable bool) *FFmpegCommand {
	fc.debugMode = enable
	return fc
}

// AddArgs ajoute des arguments à la commande FFmpeg
func (fc *FFmpegCommand) AddArgs(args ...string) *FFmpegCommand {
	for _, arg := range args {
		fc.args = append(fc.args, parseArgs(arg)...)
	}
	return fc
}

// SetReturnOutput définit si la sortie de la commande doit être retournée
func (fc *FFmpegCommand) SetReturnOutput(returnOutput bool) *FFmpegCommand {
	fc.returnOutput = returnOutput
	return fc
}

// SetExecutionTime définit le délai d'exécution maximal pour la commande
func (fc *FFmpegCommand) SetExecutionTime(duration time.Duration) *FFmpegCommand {
	fc.executionTime = duration
	return fc
}

// Execute exécute la commande FFmpeg et gère le contexte interne
func (fc *FFmpegCommand) Execute() (string, error) {
	var stderr bytes.Buffer

	ctx, cancel := context.WithTimeout(context.Background(), fc.executionTime)
	defer cancel()

	fc.cmd = exec.CommandContext(ctx, fc.args[0], fc.args[1:]...)
	fc.cmd.Stderr = &stderr

	output, err := fc.cmd.Output()

	if err != nil {
		fc.lastError = stderr.String()
		if fc.debugMode {
			fc.logger.Printf("Erreur lors de l'exécution: %s\n", fc.lastError)
		}
		return "", err
	}

	if fc.returnOutput {
		return string(output), nil
	}
	return "", nil
}

// GetLastError renvoie le dernier message d'erreur
func (fc *FFmpegCommand) GetLastError() string {
	return fc.lastError
}

// Stop arrête la commande FFmpeg en cours d'exécution
func (fc *FFmpegCommand) Stop() {
	if fc.cmd != nil && fc.cmd.Process != nil {
		fc.cmd.Process.Kill()
	}
}

// IsRunning vérifie si la commande FFmpeg est en cours d'exécution
func (fc *FFmpegCommand) IsRunning() bool {
	return fc.cmd != nil && fc.cmd.Process != nil
}

// GetArgs renvoie les arguments de la commande FFmpeg
func (fc *FFmpegCommand) GetArgs() []string {
	return fc.args
}

// GetCommand renvoie la commande FFmpeg
func (fc *FFmpegCommand) GetCommand() *exec.Cmd {
	return fc.cmd
}

// GetExecutionTime renvoie le délai d'exécution maximal de la commande
func (fc *FFmpegCommand) GetExecutionTime() time.Duration {
	return fc.executionTime
}

// GetReturnOutput renvoie si la sortie de la commande doit être retournée
func (fc *FFmpegCommand) GetReturnOutput() bool {
	return fc.returnOutput
}

// parseArgs analyse les arguments et gère les cas spéciaux
func parseArgs(arg string) []string {
	// Implémentez une logique plus robuste pour gérer les cas spéciaux
	return strings.Fields(arg)
}

// func main() {
// 	ffmpegCmd := NewFFmpegCommand().
// 		SetReturnOutput(false).
// 		SetExecutionTime(1*time.Minute).
// 		AddArgs("-i input.mp4", "-c:v copy", "output.mp4")

// 	output, err := ffmpegCmd.Execute()
// 	if err != nil {
// 		log.Fatalf("Échec de l'exécution de ffmpeg: %v", err)
// 	}

// 	if output != "" {
// 		log.Println("Sortie de ffmpeg:", output)
// 	}
// }
