package admin

import (
	"github.com/x-clone/go-ceph/internal/commands"
)

var (
	// ErrStatusNotEmpty is an alias for commands.ErrStatusNotEmpty
	ErrStatusNotEmpty = commands.ErrStatusNotEmpty
	// ErrBodyNotEmpty is an alias for commands.ErrBodyNotEmpty
	ErrBodyNotEmpty = commands.ErrBodyNotEmpty
)

type response = commands.Response

// NotImplementedError is an alias for commands.NotImplementedError.
type NotImplementedError = commands.NotImplementedError

// newResponse returns a response.
func newResponse(b []byte, s string, e error) response {
	return commands.NewResponse(b, s, e)
}
