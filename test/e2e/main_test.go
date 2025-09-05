package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	// cliPath is the absolute path to the compiled test binary.
	cliPath string
)

// TestMain is the entry point for tests in this package.
// It compiles the CLI binary once and makes it available to all tests.
func TestMain(m *testing.M) {
	// Get project root. The test is run from `test/e2e`, so we go up two levels.
	projectRoot, err := filepath.Abs("../../")
	if err != nil {
		fmt.Printf("Failed to get project root: %v\n", err)
		os.Exit(1)
	}

	// Define a unique name for the test binary
	cliPath = filepath.Join(projectRoot, "test/e2e/atelier-cli-test")

	fmt.Println("Building test binary...")
	buildCmd := exec.Command("go", "build", "-o", cliPath, projectRoot)
	if output, err := buildCmd.CombinedOutput(); err != nil {
		fmt.Printf("Failed to build test binary: %v\nOutput:\n%s\n", err, string(output))
		os.Exit(1)
	}
	fmt.Printf("Test binary built at: %s\n", cliPath)

	// Run the tests
	exitCode := m.Run()

	// Cleanup the test binary
	fmt.Println("Cleaning up test binary...")
	os.Remove(cliPath)

	os.Exit(exitCode)
}