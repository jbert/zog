package speccy

import "github.com/jbert/zog"

func InstallKeyboardInputPorts(z *zog.Zog) {
	z.RegisterInputHandler(32766, func() byte {
		return 0xef
	})
}
