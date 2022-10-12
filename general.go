package go_nmcli

import (
	"context"
)

type GeneralPermission struct {
	Permission string
	Value      GeneralPermissionValue
}

type GeneralPermissionValue string

const (
	GeneralPermissionValueAuth GeneralPermissionValue = "auth"
	GeneralPermissionValueYes  GeneralPermissionValue = "yes"
	GeneralPermissionValueNo   GeneralPermissionValue = "no"
)

// GeneralPermissions shows caller permissions for authenticated operations.
func (cli NMCli) GeneralPermissions(ctx context.Context) ([]GeneralPermission, error) {
	cmdFields := []string{"PERMISSION", "VALUE"}

	recordLines, err := cli.runCommand(ctx, cmdFields, "general", "permissions")
	if err != nil {
		return nil, err
	}

	var permissions []GeneralPermission
	for _, line := range recordLines {
		permissions = append(permissions, GeneralPermission{
			Permission: line[0],
			Value:      GeneralPermissionValue(line[1]),
		})
	}

	return permissions, nil
}
