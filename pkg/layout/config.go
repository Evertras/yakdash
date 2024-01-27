package layout

type Root struct {
	// Screens is a list of screens which will take up the entire
	// display at a time.
	Screens []Node `mapstructure:"screens"`
}

type Node struct {
	// Name is a short, human-readable name that may be displayed
	// as part of the frame.
	Name string `mapstructure:"name,omitempty"`

	//////////////////////////////////////////////
	// Option 1: contains an actual module/display

	// Size is the flexbox-style size.  Defaults to 1.
	// If all children have an equal size, they all
	// will take up an equal ratio.
	Size int `mapstructure:"size,omitempty"`

	// Module is the module name to load.
	Module string `mapstructure:"module,omitempty"`

	// Config is the module-specific configuration.
	Config map[string]interface{} `mapstructure:"config,omitempty"`

	//////////////////////////////////////////////
	// Option 2: contains a sublayout

	// Stack is either 'vertical' or 'horizontal'.
	// A vertical stack goes from top to bottom.
	// A horizontal stack goes from left to right.
	Stack string `mapstructure:"stack,omitempty"`

	// Children is a list of panes (which may be split
	// again by having more children) in this group.
	Children []Node `mapstructure:"children,omitempty"`
}
