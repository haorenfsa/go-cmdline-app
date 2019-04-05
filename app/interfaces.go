package app

// CmdSet are a line of what we type in cmdline after program started
type CmdSet []string

// CmdSetRunner represents the ability to run according to a CmdSet
type CmdSetRunner interface {
	RunOnCmdSet(CmdSet)
}

// APP contains all function a cmdline app should implement
type APP interface {
	SetFlags()
	Run()
	CmdSetRunner
}
