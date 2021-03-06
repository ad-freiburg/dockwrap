package wrap

import (
	"flag"
	"os"
)

type Rm struct {
	Cmd   *flag.FlagSet
	Force bool
}

func (rm *Rm) InitFlags() {
	rm.Cmd = flag.NewFlagSet("rm", flag.ExitOnError)
	rm.Cmd.BoolVar(&rm.Force, "f", false, "Force removal also kills running container")
}

func (rm *Rm) ParseToArgs(rawArgs []string) []string {
	if err := rm.Cmd.Parse(rawArgs); err != nil {
		// Only returns an error if the Usage was shown
		os.Exit(0)
	}
	args := []string{"rm"}

	if rm.Force {
		args = append(args, "-f")
	}

	if rm.Cmd.NArg() > 0 {
		// add -- to make sure additional arguments are not interpreted as
		// potentially harmful flags. Here this is the container to rm
		args = append(args, "--")
		for _, arg := range rm.Cmd.Args() {
			if !IsHexOnly(arg) {
				arg = PrependUsername(arg)
			}
			args = append(args, arg)
		}
	}
	return args
}
