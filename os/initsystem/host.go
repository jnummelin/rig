package initsystem

import "github.com/k0sproject/rig/exec"

// Host interface for init system
type Host interface {
	Execf(string, ...interface{}) error
	ExecOutputf(string, ...interface{}) (string, error)
	ExecOutput(string, ...exec.Option) (string, error)
}
