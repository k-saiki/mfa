package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateCommand_Success(t *testing.T) {
	tmpDir := t.TempDir()

	// JBSWY3DPEHPK3PXP is a valid base32 encoded secret
	configContent := `service:
  - name: test-service
    secret: "JBSWY3DPEHPK3PXP"
`
	configPath := filepath.Join(tmpDir, "secrets.yml")
	if err := os.WriteFile(configPath, []byte(configContent), 0600); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	t.Setenv("MFA_CONFIG", configPath)

	cmd := NewCommand()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"gen", "test-service"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := strings.TrimSpace(buf.String())

	// TOTP token should be 6 digits
	matched, _ := regexp.MatchString(`^\d{6}$`, output)
	if !matched {
		t.Errorf("expected 6-digit token, got: %s", output)
	}
}

func TestGenerateCommand_ServiceNotFound(t *testing.T) {
	tmpDir := t.TempDir()

	configContent := `service:
  - name: existing-service
    secret: "JBSWY3DPEHPK3PXP"
`
	configPath := filepath.Join(tmpDir, "secrets.yml")
	if err := os.WriteFile(configPath, []byte(configContent), 0600); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	t.Setenv("MFA_CONFIG", configPath)

	cmd := NewCommand()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"gen", "nonexistent-service"})

	err := cmd.Execute()
	if err == nil {
		t.Error("expected error for nonexistent service")
	}

	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("expected 'not found' in error, got: %v", err)
	}
}

func TestGenerateCommand_InvalidSecret(t *testing.T) {
	tmpDir := t.TempDir()

	configContent := `service:
  - name: invalid-service
    secret: "not-valid-base32!"
`
	configPath := filepath.Join(tmpDir, "secrets.yml")
	if err := os.WriteFile(configPath, []byte(configContent), 0600); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	t.Setenv("MFA_CONFIG", configPath)

	cmd := NewCommand()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"gen", "invalid-service"})

	err := cmd.Execute()
	if err == nil {
		t.Error("expected error for invalid secret")
	}
}
