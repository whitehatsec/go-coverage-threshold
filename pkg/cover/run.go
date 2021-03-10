package cover

import (
	"os/exec"
)

// Run executes `go test -cover ./...` and returns the raw result.
func Run(fileList string) ([]byte, error) {
	goExe, err := exec.LookPath("go")
	if err != nil {
		return nil, err
	}

	if fileList == "" {
		fileList = "./..."
	}

	// gas lint warns about possible injection here, but we're fine
	cmd := exec.Command("bash", "-c", goExe+" test -coverprofile aggregate.coverprofile "+fileList+" | grep -v 'no test files' || true") // nolint: gas,gosec

	output, err := cmd.CombinedOutput()
	return output, err
}
