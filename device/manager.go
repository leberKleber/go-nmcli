//go:generate go run github.com/golang/mock/mockgen -destination=utils_cmd_mock_test.go -package=device_test github.com/leberKleber/go-nmcli/utils Cmd

package device

import (
	"context"

	"github.com/leberKleber/go-nmcli/utils"
)

const nmcliCmd = "nmcli"

type Manager struct {
	CommandContext func(ctx context.Context, name string, args ...string) utils.Cmd
}
