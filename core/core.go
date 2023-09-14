package core

type Plugin interface {
	ExecuteCommand(cmd string, args map[string]interface{}) (string, bool)
}

type Orchestrator struct {
	Plugins map[string]CommandExecPlugin
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		Plugins: make(map[string]CommandExecPlugin),
	}
}

func (o *Orchestrator) RegisterPlugin(cmd string, plugin CommandExecPlugin) {
	o.Plugins[cmd] = plugin
}

func (o *Orchestrator) Execute(cmd string, args map[string]interface{}) {
	if plugin, exists := o.Plugins[cmd]; exists {
		result := plugin.Execute(args)
		println(result)
	} else {
		println("No registered plugin for command:", cmd)
	}
}
