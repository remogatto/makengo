package makengo

import (
	"os"
	"exec"
	"bytes"
	"io"
	"strings"
)

func System(command string, args ...int) (out string, err os.Error) {

	runParams := [3]int{exec.DevNull, exec.PassThrough, exec.MergeWithStdout}

	if len(args) > 0 && len(args) <= 3 {
		for i := 0; i < len(args); i++ {
			runParams[i] = args[i]
		}
	}

	return run([]string{os.Getenv("SHELL"), "-c", command}, runParams[0], runParams[1], runParams[2])
}

func copy(a []string) []string {
	b := make([]string, len(a))
	for i, s := range a {
		b[i] = s
	}
	return b
}

func run(argv []string, stdin, stdout, stderr int) (out string, err os.Error) {

	if len(argv) < 1 {
		err = os.EINVAL
		goto Error
	}

	var cmd *exec.Cmd

	cmd, err = exec.Run(argv[0], argv, os.Environ(), "", stdin, stdout, stderr)

	if err != nil {
		goto Error
	}

	defer cmd.Close()

	var buf bytes.Buffer

	_, err = io.Copy(&buf, cmd.Stdout)
	out = buf.String()

	if err != nil {
		cmd.Wait(0)
		goto Error
	}

	w, err := cmd.Wait(0)

	if err != nil {
		goto Error
	}

	if !w.Exited() || w.ExitStatus() != 0 {
		err = w
		goto Error
	}

	return

Error:
	err = &runError{copy(argv), err}
	return
}

// A runError represents an error that occurred while running a command.
type runError struct {
	cmd []string
	err os.Error
}

func (e *runError) String() string { return strings.Join(e.cmd, " ") + ": " + e.err.String() }
