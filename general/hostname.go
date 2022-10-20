package general

import (
	"context"
	"fmt"
)

// Hostname show the persistent system hostname.
func (m Manager) Hostname(ctx context.Context) (string, error) {
	args := []string{"general", "hostname"}

	output, err := m.CommandContext(ctx, nmcliCmd, args...).Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute nmcli with args %+q: %w", args, err)
	}

	return string(output), nil
}

// SetHostname sets the persistent system hostname.
// This execution path needs root privileges.
func (m Manager) SetHostname(ctx context.Context, hostname string) error {
	args := []string{"general", "hostname", hostname}

	err := m.CommandContext(ctx, nmcliCmd, args...).Run()
	if err != nil {
		return fmt.Errorf("failed to execute nmcli with args %+q: %w", args, err)
	}

	return nil
}
