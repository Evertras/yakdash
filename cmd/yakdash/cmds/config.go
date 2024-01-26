package cmds

import "github.com/evertras/yakdash/pkg/layout"

var config struct {
	Layout layout.Root `mapstructure:"layout"`
}
