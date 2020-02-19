package control

// CommandsRepository - Data type to store the commands to be sent
type CommandsRepository []StandardCommand

// NewCommandsRepository creates and returns an instance to a CommandRepository
func NewCommandsRepository(numberOfRobots int) *CommandsRepository {
	cmdRepo := make(CommandsRepository, numberOfRobots)
	return &cmdRepo
}
