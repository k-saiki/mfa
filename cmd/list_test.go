package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestListCommand_PrintsServices(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "mfa-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	configContent := `service:
  - name: amazon
    secret: "JBSWY3DPEHPK3PXP"
  - name: google
    secret: "HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ"
  - name: github
    secret: "GEZDGNBVGY3TQOJQ"
`
	configPath := filepath.Join(tmpDir, "secrets.yml")
	if err := os.WriteFile(configPath, []byte(configContent), 0600); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	originalEnv := os.Getenv("MFA_CONFIG")
	os.Setenv("MFA_CONFIG", configPath)
	defer os.Setenv("MFA_CONFIG", originalEnv)

	config = Config{}

	cmd := NewCommand()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"list"})

	err = cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	expectedServices := []string{"amazon", "google", "github"}

	for _, svc := range expectedServices {
		if !strings.Contains(output, svc) {
			t.Errorf("expected service '%s' in output, got: %s", svc, output)
		}
	}
}

