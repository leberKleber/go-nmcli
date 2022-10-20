package device_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/leberKleber/go-nmcli/device"
	"github.com/leberKleber/go-nmcli/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManager_WiFiList(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedCmd := NewMockCmd(ctrl)
	mockedCmd.EXPECT().Output().Return([]byte(`AP[1]:FRITZ!Box 7530 NT:465249545A21426F782037353330204E54:3C\:37\:12\:79\:85\:1D:Infra:6:2437 MHz:130 Mbit/s:100:▂▄▆█:WPA2:(none):pair_ccmp group_ccmp psk:wlp58s0:no: :/org/freedesktop/NetworkManager/AccessPoint/68
AP[2]:FRITZ!Box 7530 NT:465249545A21426F782037353330204E54:3C\:37\:12\:79\:85\:1E:Infra:52:5260 MHz:270 Mbit/s:87:▂▄▆█:WPA2:(none):pair_ccmp group_ccmp psk:wlp58s0:yes:*:/org/freedesktop/NetworkManager/AccessPoint/1`), nil).Times(1)

	m := device.Manager{
		CommandContext: func(ctx context.Context, name string, args ...string) utils.Cmd {
			require.Equal(t, "nmcli", name)
			require.EqualValues(t,
				[]string{"-g", "NAME,SSID,SSID-HEX,BSSID,MODE,CHAN,FREQ,RATE,SIGNAL,BARS,SECURITY,WPA-FLAGS,RSN-FLAGS,DEVICE,ACTIVE,IN-USE,DBUS-PATH", "device", "wifi", "list", "ifname", "wlp58s0", "bssid", "3C:37:12:79:85:1E", "--rescan", "auto"},
				args,
			)

			return mockedCmd
		},
	}

	wifis, err := m.WiFiList(context.Background(), device.WiFiListOptions{
		Rescan: device.WiFiListOptionsRescanAuto,
		IfName: "wlp58s0",
		BSSID:  "3C:37:12:79:85:1E",
	})

	require.NoError(t, err)
	require.EqualValues(t, []device.WiFi{
		{"AP[1]", "FRITZ!Box 7530 NT", "465249545A21426F782037353330204E54", "3C:37:12:79:85:1D",
			"Infra", "6", "2437 MHz", "130 Mbit/s", "100", "▂▄▆█", "WPA2",
			"(none)", "pair_ccmp group_ccmp psk", "wlp58s0", "no", " ",
			"/org/freedesktop/NetworkManager/AccessPoint/68",
		},
		{"AP[2]", "FRITZ!Box 7530 NT", "465249545A21426F782037353330204E54", "3C:37:12:79:85:1E",
			"Infra", "52", "5260 MHz", "270 Mbit/s", "87", "▂▄▆█", "WPA2",
			"(none)", "pair_ccmp group_ccmp psk", "wlp58s0", "yes", "*",
			"/org/freedesktop/NetworkManager/AccessPoint/1",
		},
	}, wifis)
}
