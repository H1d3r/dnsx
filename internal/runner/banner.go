package runner

import (
	"github.com/projectdiscovery/gologger"
	updateutils "github.com/projectdiscovery/utils/update"
)

const banner = `
      _             __  __
   __| | _ __   ___ \ \/ /
  / _' || '_ \ / __| \  / 
 | (_| || | | |\__ \ /  \ 
  \__,_||_| |_||___//_/\_\
`

// Name
const ToolName = `dnsx`

// version is the current version of dnsx
const version = `1.1.6`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tprojectdiscovery.io\n\n")
}

// GetUpdateCallback returns a callback function that updates dnsx
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("dnsx", version)()
	}
}
