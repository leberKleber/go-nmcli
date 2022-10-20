package utils

type Cmd interface {
	Output() ([]byte, error)
	Run() error
}
