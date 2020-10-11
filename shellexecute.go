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

const (
	HIDE            = 0
	SHOWNORMAL      = 1
	SHOWMINIMIZED   = 2
	MAXIMIZE        = 3
	SHOWNOACTIVATE  = 4
	SHOW            = 5
	MINIMIZE        = 6
	SHOWMINNOACTIVE = 7
	SHOWNA          = 8
	RESTORE         = 9
	SHOWDEFAULT     = 10
	FORCEMINIMIZE   = 11
)

type Param struct {
	Action    string
	Path      string
	Param     string
	Directory string
	Show      int
}

func (i Param) ShellExecute() (int, error) {
	return i.shellExecute()
}

func ShellExecute(action, path, param, directory string) (int, error) {
	return Param{
		Action:    action,
		Path:      path,
		Param:     param,
		Directory: directory,
		Show:      SHOWNORMAL,
	}.shellExecute()
}
