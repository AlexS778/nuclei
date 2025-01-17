package rsync

import (
	lib_rsync "github.com/projectdiscovery/nuclei/v3/pkg/js/libs/rsync"

	"github.com/dop251/goja"
	"github.com/projectdiscovery/nuclei/v3/pkg/js/gojs"
)

var (
	module = gojs.NewGojaModule("nuclei/rsync")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Objects / Classes
			"IsRsyncResponse": gojs.GetClassConstructor[lib_rsync.IsRsyncResponse](&lib_rsync.IsRsyncResponse{}),
			"RsyncClient":     gojs.GetClassConstructor[lib_rsync.RsyncClient](&lib_rsync.RsyncClient{}),
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
