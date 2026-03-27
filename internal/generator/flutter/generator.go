package flutter

import (
	"fmt"
	"os/exec"
	"strings"
)

type Config struct {
	Name      string
	Org       string // e.g. "com.example"
	OutputDir string
}

func Generate(cfg Config) error {
	// Check if flutter is installed
	if _, err := exec.LookPath("flutter"); err != nil {
		return fmt.Errorf("flutter not found — install: https://docs.flutter.dev/get-started/install")
	}

	// flutter create --org com.example --project-name myapp ./output/myapp
	args := []string{
		"create",
		"--org", cfg.Org,
		"--project-name", strings.ToLower(cfg.Name),
		fmt.Sprintf("%s/%s", cfg.OutputDir, cfg.Name),
	}

	cmd := exec.Command("flutter", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil

	fmt.Printf("Running: flutter %s\n", strings.Join(args, " "))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("flutter create failed: %w", err)
	}

	fmt.Printf("Flutter project created: %s/%s\n", cfg.OutputDir, cfg.Name)
	return nil
}
