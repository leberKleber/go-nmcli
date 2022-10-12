package go_nmcli

type Cmd interface {
	Output() ([]byte, error)
}
