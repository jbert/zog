package speccy

import (
	"fmt"

	"github.com/jbert/zog/monitor"
)

func (s *Machine) RegisterCallbacks() {

	monitor.MonitorCallbackTable["startTape"] = StartTapeCallback(s)
	monitor.MonitorCallbackTable["stopTape"] = StopTapeCallback(s)
	monitor.RegisterCommonCallbacks(s.z)
}

func StartTapeCallback(s *Machine) func([]string) error {
	return func(args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("Too many arguments to %s\n", args[0])
		}
		s.tape.BeginTapeRead()
		return nil
	}
}

func StopTapeCallback(s *Machine) func([]string) error {
	return func(args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("Too many arguments to %s\n", args[0])
		}
		s.tape.StopTapeRead()
		return nil
	}
}
