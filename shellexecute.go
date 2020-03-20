package su

const (
	// EDIT is the action "edit" for ShellExecute
	EDIT = "edit"
	// EXPLORE is the action "explore" for ShellExecute
	EXPLORE = "explore"
	// OPEN is the action "open" for ShellExecute
	OPEN = "open"
	// PRINT is the action "print" for ShellExecute
	PRINT = "print"
	// PROPERTIES is the action "properties" for ShellExecute
	PROPERTIES = "properties"
	// RUNAS is the action "runas" for ShellExecute
	RUNAS = "runas"
)

func ShellExecute(action, path, param, directory string) (int, error) {
	return shellExecute(action, path, param, directory)
}
