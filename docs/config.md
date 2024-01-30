# Config

_NOTE: none of this is finalized and is subject to change.
Just getting a feel for what it might look like._

Configuration can be provided in YAML, JSON, TOML, HCL, INI.
Examples are in [the examples directory](../examples/).

## Layout

Yakdash works in a similar manner to a tmux session. There are
one or more screens, which include everything that should be
shown at the same time. Then there are panes, which divide up
the screen to show different things.

Panes can be nested within other panes to create more complex layouts.

```yaml
layout:
  screens:
    - name: Host metrics
      # Stack from left to right
      stack: horizontal
      children:
        # Actual pane
        - name: Local time
          module: clock
          config:
            tz: Asia/Tokyo
        # Pane including children stacked top to bottom
        - stack: vertical
          children:
            - name: Disk free
              module: command
              config:
                bash: df -h
            - name: Weather
              module: graph
              config:
                type: line
                source:
                  file: ~/some/timeseries-data.csv
```

This will create a layout like:

```text
       2           1
-------------------------
|              |        |
|              |  Host  |   1 (default)
|              |        |
|    Clock     |--------|
|              |        |
|              |Weather |   1 (default)
|              |        |
-------------------------
```
