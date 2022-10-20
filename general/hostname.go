package general

import (
	"context"
	"fmt"
)

type HostnameArgs struct {
	Hostname string
}

func (ha HostnameArgs) rawArgs() []string {
	var args []string
	if ha.Hostname != "" {
		args = append(args, ha.Hostname)
	}

	return args
}

// Hostname get or change persistent system hostname.
// With no arguments, this prints currently configured hostname.
// When you pass a hostname, NetworkManager will set it as the new persistent system hostname.
func (m Manager) Hostname(ctx context.Context, args HostnameArgs) (string, error) {
	cmdArgs := []string{"general", "hostname"}
	cmdArgs = append(cmdArgs, args.rawArgs()...)

	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute nmcli with args %+q: %w", args, err)
	}

	return string(output), nil
}
