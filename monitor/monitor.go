package monitor

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jbert/zog"
)

var MonitorCallbackTable = make(map[string]func([]string) error)

const monitorPrompt = "$: "

func Monitor() {

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", monitorPrompt)
		cmd, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		cmd = cmd[:len(cmd)-1]

		cmd_args := strings.Split(cmd, " ")
		callback, ok := MonitorCallbackTable[cmd_args[0]]
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: no such command \"%s\"\n", cmd_args[0])
			continue
		} else {
			err = callback(cmd_args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v", err)
			}
		}
	}
}

func RegisterCommonCallbacks(z *zog.Zog) {
	MonitorCallbackTable["print"] = Print(z)
}
func Print(z *zog.Zog) func(s []string) error {
	return func(s []string) error {
		fmt.Printf("%04X: %s\n", z.GetRegisters().PC, z.State())
		return nil
	}
}
