package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tmpDir := t.TempDir()

	configContent := `service:
  - name: test-service
    secret: "JBSWY3DPEHPK3PXP"
  - name: another-service
    secret: "HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ"
`
	configPath := filepath.Join(tmpDir, "secrets.yml")
	if err := os.WriteFile(configPath, []byte(configContent), 0600); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	t.Setenv("MFA_CONFIG", configPath)

	config, _, err := LoadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if len(config.Service) != 2 {
		t.Fatalf("expected 2 services, got %d", len(config.Service))
	}

	if config.Service[0].Name != "test-service" {
		t.Errorf("expected first service name 'test-service', got '%s'", config.Service[0].Name)
	}

	if config.Service[0].Secret != "JBSWY3DPEHPK3PXP" {
		t.Errorf("expected first service secret 'JBSWY3DPEHPK3PXP', got '%s'", config.Service[0].Secret)
	}

	if config.Service[1].Name != "another-service" {
		t.Errorf("expected second service name 'another-service', got '%s'", config.Service[1].Name)
	}
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	t.Setenv("MFA_CONFIG", "/nonexistent/path/secrets.yml")

	_, _, err := LoadConfig()
	if err == nil {
		t.Error("expected error for nonexistent config file")
	}
}
