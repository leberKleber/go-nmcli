package go_nmcli

type testCmd struct {
	output []byte
	err    error
}

func (c testCmd) Output() ([]byte, error) {
	return c.output, c.err
}
