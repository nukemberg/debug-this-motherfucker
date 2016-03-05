package main

import (
	"github.com/avishai-ish-shalom/debug-this-motherfucker/common"
	_ "github.com/avishai-ish-shalom/debug-this-motherfucker/invisible_net"
	_ "github.com/avishai-ish-shalom/debug-this-motherfucker/root_resolve"
	_ "github.com/avishai-ish-shalom/debug-this-motherfucker/shadow_directory"
)

func main() {
	common.Main()
}
