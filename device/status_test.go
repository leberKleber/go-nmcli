package device_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leberKleber/go-nmcli/device"
	"github.com/leberKleber/go-nmcli/utils"
	"github.com/stretchr/testify/require"
)

func TestManager_Status(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedCmd := NewMockCmd(ctrl)
	mockedCmd.EXPECT().Output().Return([]byte(`wlp0s20f3:wifi:connected:full:limited:/org/freedesktop/NetworkManager/Devices/3:FRITZ!Box 7530 NT:c29400e9-b18d-4d33-aa63-bed67d06fe6d:/org/freedesktop/NetworkManager/ActiveConnection/34
lo:loopback:connected (externally):unknown:unknown:/org/freedesktop/NetworkManager/Devices/1:lo:71cd44a9-9cd8-4c8b-9875-46ca760647c2:/org/freedesktop/NetworkManager/ActiveConnection/1
p2p-dev-wlp0s20f3:wifi-p2p:disconnected:none:none:/org/freedesktop/NetworkManager/Devices/69:::
enp0s31f6:ethernet:unavailable:none:none:/org/freedesktop/NetworkManager/Devices/2:::`), nil).Times(1)

	m := device.Manager{
		CommandContext: func(ctx context.Context, name string, args ...string) utils.Cmd {
			require.Equal(t, "nmcli", name)
			require.EqualValues(t,
				[]string{"-g", "DEVICE,TYPE,STATE,IP4-CONNECTIVITY,IP6-CONNECTIVITY,DBUS-PATH,CONNECTION,CON-UUID,CON-PATH", "device", "status"},
				args,
			)

			return mockedCmd
		},
	}

	statuss, err := m.Status(context.Background())

	require.NoError(t, err)
	require.EqualValues(t, []device.Status{
		{
			Device:          "wlp0s20f3",
			Type:            "wifi",
			State:           "connected",
			IP4Connectivity: "full",
			IP6Connectivity: "limited",
			DbusPath:        "/org/freedesktop/NetworkManager/Devices/3",
			Connection:      "FRITZ!Box 7530 NT",
			ConUUID:         "c29400e9-b18d-4d33-aa63-bed67d06fe6d",
			ConPath:         "/org/freedesktop/NetworkManager/ActiveConnection/34",
		}, {
			Device:          "lo",
			Type:            "loopback",
			State:           "connected (externally)",
			IP4Connectivity: "unknown",
			IP6Connectivity: "unknown",
			DbusPath:        "/org/freedesktop/NetworkManager/Devices/1",
			Connection:      "lo",
			ConUUID:         "71cd44a9-9cd8-4c8b-9875-46ca760647c2",
			ConPath:         "/org/freedesktop/NetworkManager/ActiveConnection/1",
		}, {
			Device:          "p2p-dev-wlp0s20f3",
			Type:            "wifi-p2p",
			State:           "disconnected",
			IP4Connectivity: "none",
			IP6Connectivity: "none",
			DbusPath:        "/org/freedesktop/NetworkManager/Devices/69",
			Connection:      "",
			ConUUID:         "",
			ConPath:         "",
		}, {
			Device:          "enp0s31f6",
			Type:            "ethernet",
			State:           "unavailable",
			IP4Connectivity: "none",
			IP6Connectivity: "none",
			DbusPath:        "/org/freedesktop/NetworkManager/Devices/2",
			Connection:      "",
			ConUUID:         "",
			ConPath:         "",
		},
	}, statuss)
}
