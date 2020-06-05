package tools

type MfaApi interface {
	Init() error
	Run() (MfaResult, error)
	Apply(args []string, response *MfaResult) error
}

type MfaResult interface {
	Result() string
}
