package layout

// Root is the root configuration for the dashboard.
type Root struct {
	// Screens is a list of screens which will take up the entire
	// display at a time.
	Screens []Node `mapstructure:"screens"`
}

// Style is a set of common styles that can be applied to all panes.
type Style struct {
	// AlignVertical is the alignment of the pane within its parent.
	// Defaults to center.
	AlignVertical string `mapstructure:"align-vertical,omitempty"`

	// AlignHorizontal is the alignment of the pane within its parent.
	// Defaults to center.
	AlignHorizontal string `mapstructure:"align-horizontal,omitempty"`
}

// Node is a single node in the dashboard, which can either be a
// parent node with children or a leaf node with a module.
type Node struct {
	//////////////////////////////////////////////
	// Common attributes

	// Name is a short, human-readable name that may be displayed
	// as part of the frame.  Currently only displays if part
	// of a leaf node.  Optional, and does not have to be unique.
	//
	// Can be specified in a parent node for organization in the
	// configuration file.  May be used more in the future.
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

	// Style is a set of common styles that can be
	// applied to all modules.
	Style Style `mapstructure:"style,omitempty"`

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
