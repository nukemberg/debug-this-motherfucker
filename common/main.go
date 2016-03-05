package common

import (
	"fmt"
	"github.com/kardianos/osext"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	App           = kingpin.New("debug-this-motherfucker", "A collection of mind fucking trolling hacks")
	self_destruct = App.Flag("self-destruct", "Remove this binary after trolling").Bool()
	explain       = App.Flag("explain", "Explain this trolling hack").Bool()
	explanations  = make(map[string]string)
)

func Main() {
	cmd := kingpin.MustParse(App.Parse(os.Args[1:]))
	if *explain && cmd != "" {
		fmt.Println(explanations[cmd])
		return
	}

	if *self_destruct {
		if exec, err := osext.Executable(); err != nil {
			fmt.Errorf("Couldn't remove executable: %s\n", err.Error())
		} else {
			os.Remove(exec)
		}
	}
}

func IsFileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func RegisterPlugin(name string, explanation string, action func(*kingpin.ParseContext) error) *kingpin.CmdClause {
	cmd := App.Command(name, name+" troll")
	explanations[cmd.FullCommand()] = explanation
	cmd.Action(action)
	return cmd
}

func StringInSlice(slice []string, s string) bool {
	for _, x := range slice {
		if x == s {
			return true
		}
	}
	return false
}
