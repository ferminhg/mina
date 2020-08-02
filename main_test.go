package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_E2E(t *testing.T) {
	assertThat := func(assumption string, command string, expectedOutput string) {
		tmpfile, _ := ioutil.TempFile("", "calc-fake-stdout.*")
		defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile
		os.Args = strings.Split(command, " ")

		main()

		output, _ := ioutil.ReadFile(tmpfile.Name())
		actualOutput := string(output)
		assert.Equal(t, expectedOutput, actualOutput, assumption)
	}
	assertThat("should sum 1+1 and return 2", "calc 1 + 1", "sum total: 2\n")
}

func TestMain_ErrorCodes(t *testing.T) {
	assertThat := func(assumption, command, expectedErr, expectedOutput string) {
		exe, _ := os.Executable()

		cmd := exec.Command(exe, "-test.run", "^TestMain_ErrorCodes_Inception$")
		cmd.Env = append(cmd.Env, fmt.Sprintf("ErrorCodes_Args=%s", command))

		output, err := cmd.CombinedOutput()

		e, ok := err.(*exec.ExitError)

		if !ok {
			t.Log("was expecting exit code which did not happen")
			t.FailNow()
		}

		actualOutput := string(output)
		assert.Equal(t, expectedErr, e.Error(), assumption)
		assert.Equal(t, expectedOutput, actualOutput, assumption)
	}

	assertThat("should exit with code 5 if no args provided", "calc", "exit status 5", "error: invalid syntax\n")
}

func TestMain_ErrorCodes_Inception(t *testing.T) {
	args := os.Getenv("ErrorCodes_Args")
	if args != "" {
		os.Args = strings.Split(args, " ")

		main()
	}
}
