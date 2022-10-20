package go_nmcli

import (
	"context"
	"os/exec"

	"github.com/leberKleber/go-nmcli/general"
	"github.com/leberKleber/go-nmcli/utils"
	"github.com/sirupsen/logrus"
)

type General interface {
	Permissions(ctx context.Context) ([]general.Permission, error)

	Hostname(ctx context.Context) (string, error)
	SetHostname(ctx context.Context, hostname string) error
}

type NMCli struct {
	// should be used to exec custom nmcli commands
	CommandContext func(ctx context.Context, name string, args ...string) utils.Cmd
	logDebug       func(fmt string, args ...interface{})
	General        General
}

type Option = func(cli *NMCli)

func NewNMCli(opts ...Option) NMCli {
	cli := NMCli{
		logDebug: logrus.Debugf,
		CommandContext: func(ctx context.Context, name string, args ...string) utils.Cmd {
			return exec.CommandContext(ctx, name, args...)
		},
	}
	for i := range opts {
		opts[i](&cli)
	}

	cli.General = general.Manager{CommandContext: cli.CommandContext}

	return cli
}
