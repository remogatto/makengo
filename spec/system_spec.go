package main

import (
	"exec"
	. "specify"
	t "../src/testmakengo"
)

func init() {
	Describe("makengo.System", func() {

		It("should pass the given string command to the shell and return stdout", func(e Example) {
			out, err := t.System("echo TEST", exec.DevNull, exec.Pipe)
			e.Value(out).Should(Be("TEST\n"))
			e.Value(err).Should(Be(nil))
		})

		It("should raise errors", func(e Example) {
			_, err := t.System("foobar!!!???", exec.DevNull, exec.Pipe)
			e.Value(err.String() != "").Should(Be(true))
		})

	})

}
