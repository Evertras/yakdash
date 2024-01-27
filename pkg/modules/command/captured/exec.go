package captured

import (
	"context"
	"fmt"
	"os/exec"
)

type Cmd struct {
	cmd    *exec.Cmd
	stdout *capturedOutput
	stderr *capturedOutput
}

func New(ctx context.Context, command string, args ...string) *Cmd {
	cmd := exec.CommandContext(ctx, command, args...)

	c := &Cmd{
		cmd:    cmd,
		stdout: newCapturedOutput(),
		stderr: newCapturedOutput(),
	}

	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	return c
}

func (c *Cmd) Run() error {
	err := c.cmd.Start()

	if err != nil {
		return fmt.Errorf("cmd.Start: %w", err)
	}

	return c.cmd.Wait()
}

func (r *Cmd) Stop() error {
	if r == nil || r.cmd == nil {
		return fmt.Errorf("cmd is nil")
	}

	err := r.cmd.Process.Kill()

	if err != nil {
		return fmt.Errorf("process.Kill: %w", err)
	}

	return nil
}

func (r *Cmd) Stdout() string {
	output := r.stdout.String()

	return output
}

func (r *Cmd) Stderr() string {
	output := r.stderr.String()

	return output
}

func (r *Cmd) ResetOutput() {
	r.stdout.reset()
	r.stderr.reset()
}
