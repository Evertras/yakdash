package clock

import "strings"

type config struct {
	Format   string `mapstructure:"format"`
	Timezone string `mapstructure:"timezone"`
}

func (c config) fillDefaults() config {
	if c.Format == "" {
		c.Format = "15:04:05"
	}

	if c.Timezone == "" {
		c.Timezone = "Local"
	}

	c.Timezone = strings.ReplaceAll(c.Timezone, " ", "_")

	return c
}
