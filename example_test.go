package loggedexec_test

import (
	"fmt"

	"github.com/stapelberg/loggedexec"
)

func ExampleCommand() {
	cmd := loggedexec.Command("ls", "/tmp/nonexistent")
	cmd.Env = []string{"LANG=C"}
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	// Output: Running "ls /tmp/nonexistent": exit status 2
	// See "/tmp/000-ls.invocation.log" for invocation details.
	// See "/tmp/000-ls.stdoutstderr.log" for full stdout/stderr.
	// First stdout/stderr line: "ls: cannot access /tmp/nonexistent: No such file or directory"
}
