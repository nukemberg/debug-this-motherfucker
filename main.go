package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	dbtm "github.com/avishai-ish-shalom/debug-this-motherfucker/common"
	_ "github.com/avishai-ish-shalom/debug-this-motherfucker/plugins/invisible_net"
	_ "github.com/avishai-ish-shalom/debug-this-motherfucker/plugins/root_resolve"
	_ "github.com/avishai-ish-shalom/debug-this-motherfucker/plugins/shadow_directory"
	_ "github.com/avishai-ish-shalom/debug-this-motherfucker/plugins/immutableFile"
)

func main() {
	kingpin.MustParse(dbtm.App.Parse(os.Args[1:]))
}
