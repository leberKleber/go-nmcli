package device

import (
	"context"
	"fmt"
	"github.com/leberKleber/go-nmcli/utils"
	"strings"
)

type WiFiListOptions struct {
	IfName string
	BSSID  string
	Rescan WiFiListOptionsRescan
}

type WiFiListOptionsRescan string

var (
	WiFiListOptionsRescanAuto WiFiListOptionsRescan = "auto"
	WiFiListOptionsRescanYes  WiFiListOptionsRescan = "yes"
	WiFiListOptionsRescanNo   WiFiListOptionsRescan = "no"
)

func (a WiFiListOptions) rawArgs() []string {
	var args []string

	args = appendWhenNotEmpty(args, a.IfName, "ifname")
	args = appendWhenNotEmpty(args, a.BSSID, "bssid")
	args = appendWhenNotEmpty(args, string(a.Rescan), "--rescan")

	return args
}

type WiFi struct {
	Name      string
	SSID      string
	SSIDHEX   string
	BSSID     string
	Mode      string
	Chan      string
	Frequency string
	Rate      string
	Signal    string
	Bars      string
	Security  string
	WPAFlags  string
	RSNFlags  string
	Device    string
	Active    string
	InUse     string
	DBusPath  string
}

// WiFiList List available Wi-Fi access points.
// The IfName and BSSID options can be used to list APs for a particular interface, or with a specific BSSID.
// The Rescan flag tells whether a new Wi-Fi scan should be triggered.
func (m Manager) WiFiList(ctx context.Context, args WiFiListOptions) ([]WiFi, error) {
	fields := []string{"NAME", "SSID", "SSID-HEX", "BSSID", "MODE", "CHAN", "FREQ", "RATE", "SIGNAL", "BARS", "SECURITY", "WPA-FLAGS", "RSN-FLAGS", "DEVICE", "ACTIVE", "IN-USE", "DBUS-PATH"}

	cmdArgs := []string{"-g", strings.Join(fields, ",")}
	cmdArgs = append(cmdArgs, "device", "wifi", "list")
	cmdArgs = append(cmdArgs, args.rawArgs()...)

	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute nmcli with args %+q: %w", cmdArgs, err)
	}

	parsedOutput, err := utils.ParseCmdOutput(output, len(fields))
	if err != nil {
		return nil, fmt.Errorf("failed to parse nmcli output: %w", err)
	}

	var wifis []WiFi
	for _, fields := range parsedOutput {
		wifis = append(wifis, WiFi{
			Name:      fields[0],
			SSID:      fields[1],
			SSIDHEX:   fields[2],
			BSSID:     fields[3],
			Mode:      fields[4],
			Chan:      fields[5],
			Frequency: fields[6],
			Rate:      fields[7],
			Signal:    fields[8],
			Bars:      fields[9],
			Security:  fields[10],
			WPAFlags:  fields[11],
			RSNFlags:  fields[12],
			Device:    fields[13],
			Active:    fields[14],
			InUse:     fields[15],
			DBusPath:  fields[16],
		})
	}

	return wifis, nil
}

func appendWhenNotEmpty(slice []string, toCheck string, preAppend string) []string {
	if toCheck != "" {
		slice = append(slice, preAppend, toCheck)
	}

	return slice
}
