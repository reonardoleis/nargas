package cli

type ArgHandler struct {
	Arg     string
	Handler func()
}

func HandleArgs(args []string, argHandlers []ArgHandler) {
	for _, arg := range args {
		for _, argHandler := range argHandlers {
			if argHandler.Arg == arg {
				argHandler.Handler()
			}
		}
	}
}
