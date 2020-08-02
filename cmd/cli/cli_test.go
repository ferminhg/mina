package cli

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type commandStub struct {
	hasExecuted bool
}

func (c *commandStub) run(output io.Writer) {
	c.hasExecuted = true
}

func TestRun(t *testing.T) {
	assertThat := func(assumption string, err error, errored, executedCmd bool) {
		var (
			hasErrored     bool = false
			stdOut, stdErr bytes.Buffer
		)
		stub := &commandStub{}
		c := NewConsole(&stdOut, &stdErr, func(code int) { hasErrored = true })
		c.commandFactory = func(args []string) (cliCommand, error) {
			return stub, err
		}
		c.Run([]string{})
		assert.Equal(t, hasErrored, errored, assumption)
		assert.Equal(t, executedCmd, stub.hasExecuted, assumption)
	}
	assertThat("should not run command when get error", errors.New("some error"), true, false)
	assertThat("should run command when no errors", nil, false, true)
}

func TestAdditionCommand(t *testing.T) {
	assertThat := func(assumption string, v1, v2 int, expectedOutput string) {
		cmd := additionCommand{v1, v2}
		var stdOut bytes.Buffer

		cmd.run(&stdOut)

		actualOutput := stdOut.String()

		assert.Equal(t, expectedOutput, actualOutput, assumption)
	}

	assertThat("should print 'sum total: 5\n' for 3 + 2", 3, 2, "sum total: 5\n")
	assertThat("should print 'sum total: 79\n' for 40 + 39", 40, 39, "sum total: 79\n")
}
