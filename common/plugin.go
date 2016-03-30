package common

import (
	"fmt"
	"os"

	"github.com/kardianos/osext"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Explanations plugins supported
var (
	App          = kingpin.New("debug-this-motherfucker", "A collection of mind fucking trolling hacks")
	selfDestruct = App.Flag("self-destruct", "Remove this binary after trolling").Bool()
	explain      = App.Flag("explain", "Explain this trolling hack").Bool()
)

func pluginActionWrapper(explanation string, action func(*kingpin.ParseContext) error) func(*kingpin.ParseContext) error {
	return func(ctx *kingpin.ParseContext) error {
		if *explain {
			fmt.Println(explanation)
			return nil
		}

		if *selfDestruct {
			if exec, err := osext.Executable(); err != nil {
				fmt.Printf("Couldn't remove executable: %s\n", err.Error())
			} else {
				os.Remove(exec)
			}
		}
		return action(ctx)
	}
}

// RegisterPlugin registers a plugin
func RegisterPlugin(name string, explanation string, action func(*kingpin.ParseContext) error) *kingpin.CmdClause {
	cmd := App.Command(name, fmt.Sprintf("%s troll", name))
	cmd.Action(pluginActionWrapper(explanation, action))
	return cmd
}
