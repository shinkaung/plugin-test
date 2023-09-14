package helloworld

type HelloWorldPlugin struct{}

func (h *HelloWorldPlugin) Execute(args map[string]interface{}) string {
	return "Hello, World!"
}
