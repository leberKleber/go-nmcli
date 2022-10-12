package go_nmcli

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type NMCli struct {
	// should be used to exec custom nmcli commands
	CommandContext func(ctx context.Context, name string, args ...string) Cmd
	logDebug       func(fmt string, args ...interface{})
}

type Option = func(cli *NMCli)

func NewNMCli(opts ...Option) NMCli {
	cli := NMCli{
		logDebug: logrus.Debugf,
		CommandContext: func(ctx context.Context, name string, args ...string) Cmd {
			return exec.CommandContext(ctx, name, args...)
		},
	}
	for i := range opts {
		opts[i](&cli)
	}

	return cli
}

func (cli NMCli) runCommand(ctx context.Context, fields []string, object string, command string) ([][]string, error) {
	nmcliPath := fmt.Sprintf("'%s>%s'", object, command)

	output, err := cli.
		CommandContext(ctx, "nmcli", "-g", strings.Join(fields, ","), object, command).
		Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run %s command: %w", nmcliPath, err)
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(output))

	var recordLines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		cli.logDebug("parse %s response line: %q", nmcliPath, line)

		recordLine := cli.splitOutput(line)
		if len(recordLine) != len(fields) {
			return nil, fmt.Errorf(
				"%s response line conains %d instead of %d options",
				nmcliPath, len(recordLine), len(fields),
			)
		}

		recordLines = append(recordLines, recordLine)
	}

	return recordLines, nil
}

func (cli NMCli) splitOutput(line string) []string {
	escape := `\`
	separator := ":"
	tempEscapedSeparator := "\x00"

	replacedEscape := strings.ReplaceAll(line, escape+separator, tempEscapedSeparator)
	records := strings.Split(replacedEscape, separator)

	for i, record := range records {
		records[i] = strings.ReplaceAll(record, tempEscapedSeparator, separator)
	}

	return records
}
