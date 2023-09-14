package core

type CommandExecPlugin interface {
	Execute(args map[string]interface{}) string
}
