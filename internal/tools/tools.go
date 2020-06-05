package tools

var (
	Code     string
	FileName string
	Tools    map[string]MfaApi
)

func init() {
	Tools = map[string]MfaApi{
		"duo": &Duo{},
	}
}
