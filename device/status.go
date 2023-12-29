package device

import (
	"context"
	"fmt"
	"strings"

	"github.com/leberKleber/go-nmcli/utils"
)

type Status struct {
	Device          string
	Type            string
	State           string
	IP4Connectivity string
	IP6Connectivity string
	DbusPath        string
	Connection      string
	ConUUID         string
	ConPath         string
}

// Status lists the status for all devices.
func (m Manager) Status(ctx context.Context) ([]Status, error) {
	fields := []string{"DEVICE", "TYPE", "STATE", "IP4-CONNECTIVITY", "IP6-CONNECTIVITY", "DBUS-PATH", "CONNECTION", "CON-UUID", "CON-PATH"}

	cmdArgs := []string{"-g", strings.Join(fields, ",")}
	cmdArgs = append(cmdArgs, "device", "status")

	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute nmcli with args %+q: %w", cmdArgs, err)
	}

	parsedOutput, err := utils.ParseCmdOutput(output, len(fields))
	if err != nil {
		return nil, fmt.Errorf("failed to parse nmcli output: %w", err)
	}

	statuss := make([]Status, len(parsedOutput))
	for i, fields := range parsedOutput {
		statuss[i] = Status{
			Device:          fields[0],
			Type:            fields[1],
			State:           fields[2],
			IP4Connectivity: fields[3],
			IP6Connectivity: fields[4],
			DbusPath:        fields[5],
			Connection:      fields[6],
			ConUUID:         fields[7],
			ConPath:         fields[8],
		}
	}

	return statuss, nil
}
