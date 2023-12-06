package cli

type Executor any
type Args []string

var commands = map[string]Executor{}

// Регистрирует консольную команду cmd с испольнителем executor
func Register(cmd string, executor Executor) {
	commands[cmd] = executor
}
