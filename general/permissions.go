package general

import (
	"context"
	"fmt"
	"strings"

	"github.com/leberKleber/go-nmcli/utils"
)

type Permission struct {
	Permission string
	Value      PermissionValue
}

type PermissionValue string

const (
	PermissionValueAuth PermissionValue = "auth"
	PermissionValueYes  PermissionValue = "yes"
	PermissionValueNo   PermissionValue = "no"
)

// Permissions shows caller permissions for authenticated operations.
func (m Manager) Permissions(ctx context.Context) ([]Permission, error) {
	fields := []string{"PERMISSION", "VALUE"}
	args := []string{"-g", strings.Join(fields, ","), "general", "permissions"}

	output, err := m.CommandContext(ctx, "nmcli", args...).Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute nmcli with args %+q: %w", args, err)
	}

	lines, err := utils.ParseCmdOutput(output, len(fields))
	if err != nil {
		return nil, fmt.Errorf("failed to parse command output: %w", err)
	}

	var permissions []Permission
	for _, line := range lines {
		permissions = append(permissions, Permission{
			Permission: line[0],
			Value:      PermissionValue(line[1]),
		})
	}

	return permissions, nil
}
