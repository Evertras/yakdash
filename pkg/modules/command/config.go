package command

import (
	"fmt"
	"time"
)

type config struct {
	Bash        string `mapstructure:"bash,omitempty"`
	IntervalRaw string `mapstructure:"interval,omitempty"`

	interval time.Duration
}

func (c config) setDefaultsAndParse() (config, error) {
	if c.Bash == "" {
		return c, fmt.Errorf("must supply a command to run")
	}

	if c.IntervalRaw == "" {
		c.IntervalRaw = "10s"
	}

	var err error
	c.interval, err = time.ParseDuration(c.IntervalRaw)

	if err != nil {
		return c, fmt.Errorf("failed to parse time %q: %w", c.IntervalRaw, err)
	}

	return c, nil
}
