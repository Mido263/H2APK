package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCleanOldBuilds(t *testing.T) {
	baseDir := t.TempDir()
	outputDir := filepath.Join(baseDir, "output")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatal(err)
	}

	oldFile := filepath.Join(outputDir, "old.apk")
	newFile := filepath.Join(outputDir, "new.apk")

	if err := os.WriteFile(oldFile, []byte("old"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(newFile, []byte("new"), 0644); err != nil {
		t.Fatal(err)
	}

	oldTime := time.Now().Add(-48 * time.Hour)

	if err := os.Chtimes(oldFile, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}

	cleanOldBuilds(baseDir)

	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Error("old file should have been removed")
	}

	if _, err := os.Stat(newFile); err != nil {
		t.Error("new file should still exist")
	}
}

func TestEnv(t *testing.T) {
	t.Setenv("H2APK_TEST_VALUE", "test-value")

	got := env("H2APK_TEST_VALUE", "default")

	if got != "test-value" {
		t.Errorf("expected test-value, got %s", got)
	}
}

func TestEnvDefault(t *testing.T) {
	t.Setenv("H2APK_EMPTY_VALUE", "")

	got := env("H2APK_EMPTY_VALUE", "default")

	if got != "default" {
		t.Errorf("expected default, got %s", got)
	}
}